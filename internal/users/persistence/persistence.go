//go:generate mockgen -source persistence.go -destination=mock/mockpersistence.go -package=mock
package persistence

type Persistence interface {
	CreateUser(user User) int64
	GetUser(id string) (bool, User)
	GetUserByUsername(username string) (bool, User)
	UpdateUser(user User) int64
	DeleteUser(user User) int64
}
