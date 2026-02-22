package bnpl

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/optimus-boc-protocol/bindings"
)

// Service is the production implementation of IBNPLService.
type Service struct {
	client   bind.ContractBackend
	contract *bindings.BNPLManager
}

// NewService constructs a BNPL service from an RPC URL and contract address.
func NewService(rpcURL, contractAddr string) (IBNPLService, error) {
	cli, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("bnpl dial: %w", err)
	}
	addr := common.HexToAddress(contractAddr)
	c, err := bindings.NewBNPLManager(addr, cli)
	if err != nil {
		return nil, fmt.Errorf("bnpl binding: %w", err)
	}
	return &Service{client: cli, contract: c}, nil
}

func (s *Service) CreateBNPL(opts *bind.TransactOpts, daoId *big.Int, recipient common.Address, totalAmount *big.Int, startTimestamp *big.Int, intervalSeconds *big.Int, metadata []byte) (*types.Transaction, error) {
	return s.contract.CreateBNPL(opts, daoId, recipient, totalAmount, startTimestamp, intervalSeconds, metadata)
}

func (s *Service) MakePayment(opts *bind.TransactOpts, arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return s.contract.MakePayment(opts, arrangementId, installmentNumber)
}

func (s *Service) ActivateBNPL(opts *bind.TransactOpts, arrangementId *big.Int) (*types.Transaction, error) {
	return s.contract.ActivateBNPL(opts, arrangementId)
}

func (s *Service) ApplyLateFee(opts *bind.TransactOpts, arrangementId *big.Int, installmentNumber uint8) (*types.Transaction, error) {
	return s.contract.ApplyLateFee(opts, arrangementId, installmentNumber)
}

func (s *Service) Reschedule(opts *bind.TransactOpts, arrangementId *big.Int, newStartTimestamp *big.Int, newIntervalSeconds *big.Int) (*types.Transaction, error) {
	return s.contract.Reschedule(opts, arrangementId, newStartTimestamp, newIntervalSeconds)
}

func (s *Service) GetArrangement(ctx context.Context, id *big.Int) (BNPLArrangement, error) {
	out, err := s.contract.GetArrangement(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return BNPLArrangement{}, err
	}
	return BNPLArrangement{
		ID:                 out.Id,
		DaoId:              out.DaoId,
		Payer:              out.Payer,
		Recipient:          out.Recipient,
		TotalAmount:        out.TotalAmount,
		NumInstallments:    out.NumInstallments,
		InstallmentAmounts: out.InstallmentAmounts,
		StartTimestamp:     out.StartTimestamp,
		IntervalSeconds:    out.IntervalSeconds,
		LateFeeBps:         out.LateFeeBps,
		Status:             out.Status,
	}, nil
}

func (s *Service) WaitForArrangementID(ctx context.Context, tx *types.Transaction) (uint64, error) {
	backend, ok := s.client.(bind.DeployBackend)
	if !ok {
		return 0, fmt.Errorf("client does not support WaitMined")
	}
	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return 0, err
	}
	for _, log := range receipt.Logs {
		ev, err := s.contract.ParseBNPLCreated(*log)
		if err == nil {
			return ev.ArrangementId.Uint64(), nil
		}
	}
	return 0, fmt.Errorf("BNPLCreated event not found in receipt %s", tx.Hash().Hex())
}
