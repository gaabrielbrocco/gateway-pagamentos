package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindByID(ID string) (*Account, error)
	UpdateBalance(accout *Account) error
}
