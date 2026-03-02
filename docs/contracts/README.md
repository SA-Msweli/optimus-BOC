# Optimus Protocol — Smart Contracts

All contracts are Solidity ^0.8.19, deployed on **Ethereum Sepolia** (chain ID 11155111).  
They use OpenZeppelin `AccessControl` for role-based security and `ReentrancyGuard` where ETH flows occur.

## Deployed Addresses

| Contract       | Address                                      |
|----------------|----------------------------------------------|
| DIDRegistry    | `0x0E9D8959bCD99e7AFD7C693e51781058A998b756`  |
| BNPLManager    | `0x4d99Dc2e504c15496319339E822C4a8EAfe3e2ba`  |
| LoanManager    | `0xbB0D4067488edf4a007822407e2486412dC8D39D`  |
| DAOManager     | `0x561289A9B8439E3fb288a33b3c39C78E0923Cd2b`  |
| TokenVault     | `0x4C704D51fc47cfe582F8c5477de3AE398B344907`  |

**Deployer:** `0x08DEB6b37c3659D045a7Fb93C742f33D1f9B3780`

## Contract Relationship Diagram

```
┌──────────────┐     reads BnplTerms      ┌──────────────┐
│ BNPLManager  │ ──────────────────────►  │  DAOManager   │
│              │     creditTreasury()     │               │
│              │ ──────────────────────►  │               │
└──────────────┘                          └───────┬───────┘
                                                  │ setTokenVault()
                                                  ▼
┌──────────────┐                          ┌──────────────┐
│ LoanManager  │                          │  TokenVault   │
│              │  (independent)           │  (ERC-20 vault│
└──────────────┘                          └──────────────┘

┌──────────────┐
│ DIDRegistry  │  (independent identity layer)
└──────────────┘
```

## Individual Contract Documentation

- [DIDRegistry](./DIDRegistry.md) — On-chain identity registry with Privy linking and risk scoring
- [BNPLManager](./BNPLManager.md) — Buy-Now-Pay-Later installment arrangement management
- [LoanManager](./LoanManager.md) — Traditional loan lifecycle management
- [DAOManager](./DAOManager.md) — DAO governance, membership, BNPL policy, and treasury
- [TokenVault](./TokenVault.md) — ERC-20 token custody vault
