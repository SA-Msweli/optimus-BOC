// Package shared – Confidential HTTP helpers for making privacy-preserving
// authenticated requests to the Optimus backend API.  Secrets (backendApiKey)
// are resolved inside the Vault DON enclave via template syntax and never
// appear in workflow or node memory.
package shared

import (
	"fmt"

	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/confidentialhttp"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

// BackendReq carries the URL and optional body for a confidential HTTP call.
type BackendReq struct {
	URL    string
	Method string
	Body   string
	Owner  string // vault DON owner – empty for simulation
}

// BackendResp holds the plaintext response from the backend.
type BackendResp struct {
	Body       string
	StatusCode int
}

// ConfidentialRequest makes an authenticated HTTP call to the Optimus backend
// via the confidential-http capability running inside the CRE TEE enclave.
// The backendApiKey secret is injected via Vault DON template syntax.
//
// Usage inside a handler (which receives cre.Runtime):
//
//	resp, err := shared.ConfidentialRequest(
//	    shared.BackendReq{URL: cfg.BackendURL+"/api/...", Method: "POST", Body: jsonStr, Owner: cfg.Owner},
//	    runtime,
//	)
func ConfidentialRequest(req BackendReq, runtime cre.Runtime) (BackendResp, error) {
	client := confidentialhttp.Client{}

	httpReq := &confidentialhttp.HTTPRequest{
		Url:    req.URL,
		Method: req.Method,
		MultiHeaders: map[string]*confidentialhttp.HeaderValues{
			"Authorization": {Values: []string{"Bearer {{.backendApiKey}}"}},
		},
	}

	if req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH" {
		httpReq.MultiHeaders["Content-Type"] = &confidentialhttp.HeaderValues{
			Values: []string{"application/json"},
		}
		httpReq.Body = &confidentialhttp.HTTPRequest_BodyString{BodyString: req.Body}
	}

	secrets := []*confidentialhttp.SecretIdentifier{
		{Key: "backendApiKey", Owner: &req.Owner},
	}

	resp, err := client.SendRequest(runtime, &confidentialhttp.ConfidentialHTTPRequest{
		Request:         httpReq,
		VaultDonSecrets: secrets,
	}).Await()
	if err != nil {
		return BackendResp{}, fmt.Errorf("confidential %s %s failed: %w", req.Method, req.URL, err)
	}

	return BackendResp{
		Body:       string(resp.Body),
		StatusCode: int(resp.StatusCode),
	}, nil
}
