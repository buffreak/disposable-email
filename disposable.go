package disposable

type Disposable interface {
	RandomEmailAddress() error
	SetEmailAddress(username string) error
	GetEmailAddress() string
	GetEmails() ([]EmailDetail, error)
}

type EmailDetail struct {
	ID        string
	Body      string
	Timestamp uint64
	From      string
}
