package dao

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

// Service is the production implementation of IDAOService.
type Service struct {
	client   *ethclient.Client
	contract *bindings.DAOManager
}

// NewService constructs a DAO service from an RPC URL and contract address.
func NewService(rpcURL, contractAddr string) (IDAOService, error) {
	cli, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("dao dial: %w", err)
	}
	addr := common.HexToAddress(contractAddr)
	c, err := bindings.NewDAOManager(addr, cli)
	if err != nil {
		return nil, fmt.Errorf("dao binding: %w", err)
	}
	return &Service{client: cli, contract: c}, nil
}

func (s *Service) CreateDAO(opts *bind.TransactOpts, creator common.Address, goal uint8, votingPeriodDays uint64) (*types.Transaction, error) {
	return s.contract.CreateDAO(opts, creator, goal, votingPeriodDays)
}

func (s *Service) JoinDAO(opts *bind.TransactOpts, daoId *big.Int, member common.Address, investment *big.Int) (*types.Transaction, error) {
	return s.contract.JoinDAO(opts, daoId, member, investment)
}

func (s *Service) Propose(opts *bind.TransactOpts, daoId *big.Int, data []byte) (*types.Transaction, error) {
	return s.contract.Propose(opts, daoId, data)
}

func (s *Service) Vote(opts *bind.TransactOpts, proposalId *big.Int, support bool) (*types.Transaction, error) {
	return s.contract.Vote(opts, proposalId, support)
}

func (s *Service) FinalizeProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return s.contract.FinalizeProposal(opts, proposalId)
}

func (s *Service) ExecuteProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return s.contract.ExecuteProposal(opts, proposalId)
}

func (s *Service) SetBnplTerms(opts *bind.TransactOpts, daoId *big.Int, numInstallments *big.Int, minDays *big.Int, maxDays *big.Int, lateFeeBps *big.Int, graceDays *big.Int, rescheduleAllowed bool, minDownBps *big.Int) (*types.Transaction, error) {
	return s.contract.SetBnplTerms(opts, daoId, numInstallments, minDays, maxDays, lateFeeBps, graceDays, rescheduleAllowed, minDownBps)
}

func (s *Service) GetBnplTerms(ctx context.Context, daoId *big.Int) (BnplTerms, error) {
	out, err := s.contract.GetBnplTerms(&bind.CallOpts{Context: ctx}, daoId)
	if err != nil {
		return BnplTerms{}, err
	}
	return BnplTerms{
		NumInstallments:        out.NumInstallments,
		AllowedIntervalMinDays: out.AllowedIntervalMinDays,
		AllowedIntervalMaxDays: out.AllowedIntervalMaxDays,
		LateFeeBps:             out.LateFeeBps,
		GracePeriodDays:        out.GracePeriodDays,
		RescheduleAllowed:      out.RescheduleAllowed,
		MinDownPaymentBps:      out.MinDownPaymentBps,
	}, nil
}

func (s *Service) GetTreasuryBalance(ctx context.Context, daoId *big.Int) (*big.Int, error) {
	return s.contract.GetTreasuryBalance(&bind.CallOpts{Context: ctx}, daoId)
}

func (s *Service) CreditTreasury(opts *bind.TransactOpts, daoId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return s.contract.CreditTreasury(opts, daoId, amount)
}

func (s *Service) IsMember(ctx context.Context, daoId *big.Int, member common.Address) (bool, error) {
	fopts := &bind.FilterOpts{Start: 0, End: nil, Context: ctx}
	iter, err := s.contract.FilterMemberJoined(fopts, []*big.Int{daoId}, []common.Address{member})
	if err != nil {
		return false, err
	}
	defer iter.Close()
	return iter.Next(), iter.Error()
}

func (s *Service) WaitForDaoID(ctx context.Context, tx *types.Transaction) (uint64, error) {
	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return 0, err
	}
	for _, log := range receipt.Logs {
		ev, err := s.contract.ParseDaoCreated(*log)
		if err == nil {
			return ev.DaoId.Uint64(), nil
		}
	}
	return 0, fmt.Errorf("DaoCreated event not found in receipt %s", tx.Hash().Hex())
}
