package persistence

type Persistence interface {
	CreateUser(user User) int64
	GetUser(username string) (bool, User)
	UpdateUser(user User) int64
	DeleteUser(user User) int64
}
