// Package middleware provides HTTP middleware for the Optimus protocol server.
package middleware

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ─── Context keys ────────────────────────────────────────────────────────────

type ctxKey int

const (
	// CtxPrivyUserID is the context key for the authenticated Privy user id.
	CtxPrivyUserID ctxKey = iota
	// CtxWalletAddress is the context key for the wallet address extracted from claims.
	CtxWalletAddress
)

// UserIDFromCtx extracts the Privy user id from the request context.
func UserIDFromCtx(ctx context.Context) string {
	v, _ := ctx.Value(CtxPrivyUserID).(string)
	return v
}

// WalletFromCtx extracts the wallet address from the request context.
func WalletFromCtx(ctx context.Context) string {
	v, _ := ctx.Value(CtxWalletAddress).(string)
	return v
}

// ─── JWKS types ──────────────────────────────────────────────────────────────

type jwksResponse struct {
	Keys []jwkKey `json:"keys"`
}

type jwkKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
	Alg string `json:"alg"`
	Use string `json:"use"`
}

// ─── JWKS cache ──────────────────────────────────────────────────────────────

type jwksCache struct {
	mu       sync.RWMutex
	keys     map[string]*rsa.PublicKey
	fetched  time.Time
	jwksURL  string
	cacheTTL time.Duration
}

func newJWKSCache(jwksURL string) *jwksCache {
	return &jwksCache{
		keys:     make(map[string]*rsa.PublicKey),
		jwksURL:  jwksURL,
		cacheTTL: 1 * time.Hour,
	}
}

func (c *jwksCache) getKey(kid string) (*rsa.PublicKey, error) {
	c.mu.RLock()
	if key, ok := c.keys[kid]; ok && time.Since(c.fetched) < c.cacheTTL {
		c.mu.RUnlock()
		return key, nil
	}
	c.mu.RUnlock()

	// Fetch fresh keys.
	if err := c.refresh(); err != nil {
		return nil, fmt.Errorf("jwks refresh: %w", err)
	}

	c.mu.RLock()
	defer c.mu.RUnlock()
	key, ok := c.keys[kid]
	if !ok {
		return nil, fmt.Errorf("kid %q not found in JWKS", kid)
	}
	return key, nil
}

func (c *jwksCache) refresh() error {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(c.jwksURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("jwks fetch returned %d", resp.StatusCode)
	}

	var jwks jwksResponse
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return err
	}

	keys := make(map[string]*rsa.PublicKey, len(jwks.Keys))
	for _, k := range jwks.Keys {
		if k.Kty != "RSA" {
			continue
		}
		pub, err := parseRSAPublicKey(k)
		if err != nil {
			continue
		}
		keys[k.Kid] = pub
	}

	c.mu.Lock()
	c.keys = keys
	c.fetched = time.Now()
	c.mu.Unlock()
	return nil
}

func parseRSAPublicKey(k jwkKey) (*rsa.PublicKey, error) {
	nb, err := base64.RawURLEncoding.DecodeString(k.N)
	if err != nil {
		return nil, err
	}
	eb, err := base64.RawURLEncoding.DecodeString(k.E)
	if err != nil {
		return nil, err
	}
	n := new(big.Int).SetBytes(nb)
	e := 0
	for _, b := range eb {
		e = e<<8 + int(b)
	}
	return &rsa.PublicKey{N: n, E: e}, nil
}

// ─── Minimal JWT parsing ─────────────────────────────────────────────────────
// We avoid a heavy third-party JWT library; Privy tokens are RS256 JWTs with
// a standard structure.

type jwtHeader struct {
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	Typ string `json:"typ"`
}

type privyClaims struct {
	Sub string `json:"sub"` // Privy user id  (did:privy:…)
	Iss string `json:"iss"` // issuer – should be "privy.io"
	Aud string `json:"aud"` // audience – should be the Privy App ID
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`

	// Privy-specific optional claims
	SID            string `json:"sid,omitempty"`              // session id
	WalletAddress  string `json:"wallet_address,omitempty"`   // embedded wallet
	LinkedAccounts []struct {
		Type    string `json:"type"`
		Address string `json:"address,omitempty"`
	} `json:"linked_accounts,omitempty"`
}

func decodeJWTParts(token string) (*jwtHeader, *privyClaims, []byte, []byte, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, nil, nil, fmt.Errorf("invalid JWT: expected 3 parts")
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("decode header: %w", err)
	}
	var hdr jwtHeader
	if err := json.Unmarshal(headerBytes, &hdr); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("unmarshal header: %w", err)
	}

	claimsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("decode claims: %w", err)
	}
	var claims privyClaims
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("unmarshal claims: %w", err)
	}

	sigBytes, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("decode signature: %w", err)
	}

	// The signed payload is the raw base64url header.payload (not decoded).
	signed := []byte(parts[0] + "." + parts[1])
	return &hdr, &claims, signed, sigBytes, nil
}

// ─── Middleware constructor ──────────────────────────────────────────────────

// PrivyAuth returns a chi-compatible middleware that validates Privy JWTs.
//
// It:
//  1. Extracts the Bearer token from the Authorization header.
//  2. Decodes the JWT header + claims without verification first to get the `kid`.
//  3. Fetches the matching RSA public key from the cached JWKS.
//  4. Verifies the RS256 signature.
//  5. Validates `exp`, `iss`, and `aud` claims.
//  6. Stores the Privy user ID and wallet address in the request context.
//
// If JWKS URL is empty the middleware is a no-op passthrough (dev mode).
func PrivyAuth(appID, jwksURL string) func(http.Handler) http.Handler {
	if jwksURL == "" {
		// No JWKS configured – skip auth (useful for local dev).
		return func(next http.Handler) http.Handler { return next }
	}

	cache := newJWKSCache(jwksURL)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, `{"error":"missing or invalid Authorization header"}`, http.StatusUnauthorized)
				return
			}
			token := strings.TrimPrefix(authHeader, "Bearer ")

			hdr, claims, signedPayload, signature, err := decodeJWTParts(token)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error":"bad token: %s"}`, err), http.StatusUnauthorized)
				return
			}

			if hdr.Alg != "RS256" {
				http.Error(w, `{"error":"unsupported JWT algorithm"}`, http.StatusUnauthorized)
				return
			}

			// Fetch public key by kid.
			pubKey, err := cache.getKey(hdr.Kid)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error":"unknown signing key: %s"}`, err), http.StatusUnauthorized)
				return
			}

			// Verify RS256 signature.
			if err := verifyRS256(pubKey, signedPayload, signature); err != nil {
				http.Error(w, `{"error":"invalid token signature"}`, http.StatusUnauthorized)
				return
			}

			// Validate standard claims.
			now := time.Now().Unix()
			if claims.Exp < now {
				http.Error(w, `{"error":"token expired"}`, http.StatusUnauthorized)
				return
			}
			if claims.Iss != "privy.io" {
				http.Error(w, `{"error":"invalid token issuer"}`, http.StatusUnauthorized)
				return
			}
			if appID != "" && claims.Aud != appID {
				http.Error(w, `{"error":"invalid token audience"}`, http.StatusUnauthorized)
				return
			}

			// Extract wallet address – try dedicated claim first, then linked accounts.
			wallet := claims.WalletAddress
			if wallet == "" {
				for _, la := range claims.LinkedAccounts {
					if la.Type == "wallet" && la.Address != "" {
						wallet = la.Address
						break
					}
				}
			}

			// Store in context.
			ctx := context.WithValue(r.Context(), CtxPrivyUserID, claims.Sub)
			ctx = context.WithValue(ctx, CtxWalletAddress, wallet)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ─── RS256 verification ──────────────────────────────────────────────────────

func verifyRS256(pub *rsa.PublicKey, signed, sig []byte) error {
	h := sha256.Sum256(signed)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, h[:], sig)
}
