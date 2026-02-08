package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repository *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repository}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
