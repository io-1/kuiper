package persistence

type Persistence interface {
	CreateUser(user User) (int64, User)
	GetUser(username string) (bool, User)
	UpdateUser(user User) int64
	DeleteUser(user User) int64
}
