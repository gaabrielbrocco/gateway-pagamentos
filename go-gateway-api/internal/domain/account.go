package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

// RWMutex - bloquear a escrita do amount enquanto está sendo adicionado
type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		APIKey:    generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}

func (a *Account) AddBalance(amount float64) {
	// nesse momento ninguem consegue alterar o valor do Balance, porque está em lock (precisa aguardar terminar a transação)
	a.mu.Lock()

	//defer - aguarda tudo terminar para rodar
	defer a.mu.Unlock()

	a.Balance += amount
	a.UpdatedAt = time.Now()
}
