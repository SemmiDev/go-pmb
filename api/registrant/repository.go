package registrant

type QueryResult struct {
	Result interface{}
	Error  error
}

type Repository struct {
	Cmd   CmdRepository
	Query QueryRepository
}

type CmdRepository interface {
	Save(registrant *Registrant) <-chan error
	UpdateStatus(id string, status PaymentStatus) <-chan error
}

type QueryRepository interface {
	GetByID(id string) <-chan QueryResult
	GetByUsername(username string) <-chan QueryResult
	GetByUsernameAndPassword(username, password string) <-chan QueryResult
}
