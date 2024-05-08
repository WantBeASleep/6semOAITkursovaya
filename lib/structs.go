package lib

type User struct {
	Login    string
	Email    string
	Password string
	Role     int
}

type ExternalResource struct {
	Link  string
	RType string
}
