package store

import (
	"accounts-service/model"
	"errors"
	"sync"
)

type AccountStore struct {
	mu       sync.RWMutex
	accounts map[string]*model.Account
}

func NewAccountStore() *AccountStore {
	return &AccountStore{
		accounts: make(map[string]*model.Account),
	}
}

func (s *AccountStore) Create(account *model.Account) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.accounts[account.ID] = account
}

func (s *AccountStore) Get(id string) (*model.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	acc, exists := s.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}
	return acc, nil
}

func (s *AccountStore) Debit(id string, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	acc, exists := s.accounts[id]
	if !exists {
		return errors.New("account not found")
	}
	if acc.Balance < amount {
		return errors.New("insufficient funds")
	}
	acc.Balance -= amount
	return nil
}

func (s *AccountStore) Credit(id string, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	acc, exists := s.accounts[id]
	if !exists {
		return errors.New("account not found")
	}
	acc.Balance += amount
	return nil
}
