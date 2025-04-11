package service

import "github.com/gaabrielbrocco/go-gateway/internal/domain"

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    domain.AccountService
}
