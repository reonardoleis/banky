package domain

type Account struct {
	ID             uint
	InitialBalance int64
	Limit          int64
}

type AccountRepository interface {
	LoadAccounts() (map[uint]*AccountInformation, error)
}
