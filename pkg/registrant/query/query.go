package query

type RegistrantQueryResult struct {
	Result interface{}
	Error  error
}

type RegistrantQuery interface {
	GetByID(id string) <-chan RegistrantQueryResult
	GetByUsername(username string) <-chan RegistrantQueryResult
	GetByUsernameAndPassword(username, password string) <-chan RegistrantQueryResult
}
