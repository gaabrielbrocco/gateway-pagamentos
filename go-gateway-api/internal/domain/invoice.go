package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardholderName string
}

// construtor de uma nova fatura (Invoice)
// onde garante as regras de negócio antes de criar uma fatura
func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	// irá buscar os ultimos 4 numeros
	// card.Number[12:<aqui>] -> 1234
	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

// retorna em branco para que a fatura fique como pending e possa ser enviada para o apache kafka
func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != StatusPending {
		return errors.New("invalid status")
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
