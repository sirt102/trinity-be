package repositories

import (
	"trinity-be/global"
	"trinity-be/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactionByUserAndSubID(userID, subscriptionID uuid.UUID) (*entities.Transaction, error)
	InsertNewTransaction(tx *entities.Transaction) error
	ApplyVoucherToTransaction(transactionID, voucherID uuid.UUID, discountRate float64) error
	UpdatePaySuccessTransaction(id uuid.UUID) error
}

type transactionRepository struct {
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

// GetTransactionByUserAndSubID implements TransactionRepository.
func (tr *transactionRepository) GetTransactionByUserAndSubID(userID, subscriptionID uuid.UUID) (*entities.Transaction, error) {
	var cp entities.Transaction
	err := global.PostgresQLDB.First(&cp, "user_id = (?) AND subscription_id = (?)", userID, subscriptionID).Error
	if err != nil {
		return nil, err
	}

	return &cp, nil
}

func (tr *transactionRepository) InsertNewTransaction(tx *entities.Transaction) error {
	return global.PostgresQLDB.Create(tx).Error
}

func (tr *transactionRepository) UpdatePaySuccessTransaction(id uuid.UUID) error {
	return global.PostgresQLDB.Model(&entities.Transaction{}).Where("transaction_id =?", id).Update("payment_status", "Success").Error
}

// ApplyVoucherToTransaction implements TransactionRepository.
func (tr *transactionRepository) ApplyVoucherToTransaction(transactionID uuid.UUID, voucherID uuid.UUID, discountRate float64) error {
	err := global.PostgresQLDB.Model(&entities.Transaction{}).
		Where("transaction_id =?", transactionID).
		Update("final_amount", gorm.Expr("amount * ?", discountRate)).
		Error

	// Maybe in future we will have more logic right here, so I add if condition checking to prepare.
	if err != nil {
		return err
	}

	return err
}
