package mysql

import "github.com/io-1/kuiper/internal/users/persistence"

func (p *MysqlPersistence) CreateUser(user persistence.User) int64 {
	rowsAffected := p.db.Create(&user).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) GetUser(id string) (bool, persistence.User) {
	var user persistence.User
	recordNotFound := p.db.Where("id=?", id).First(&user).RecordNotFound()
	return recordNotFound, user
}

func (p *MysqlPersistence) GetUserByUsername(username string) (bool, persistence.User) {
	var user persistence.User
	recordNotFound := p.db.Where("username=?", username).First(&user).RecordNotFound()
	return recordNotFound, user
}

func (p *MysqlPersistence) UpdateUser(user persistence.User) int64 {
	rowsAffected := p.db.Model(&user).Where("id=?", user.ID).Updates(persistence.User{Username: user.Username, Name: user.Name, Email: user.Email}).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) DeleteUser(user persistence.User) int64 {
	rowsAffected := p.db.Delete(&user).RowsAffected
	return rowsAffected
}
