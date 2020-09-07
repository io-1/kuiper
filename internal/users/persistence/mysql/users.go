package mysql

import "github.com/io-1/kuiper/internal/users/persistence"

func (p *MysqlPersistence) CreateUser(user persistence.User) int64 {
	rowsAffected := p.db.Create(&user).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) GetUser(username string) (bool, persistence.User) {
	var user persistence.User
	recordNotFound := p.db.Where("username=?", username).First(&user).RecordNotFound()
	return recordNotFound, user
}

func (p *MysqlPersistence) UpdateUser(user persistence.User) int64 {
	rowsAffected := p.db.Model(&user).Where("username= ?", user.Username).Updates(persistence.User{Password: user.Password, Name: user.Name, Email: user.Email}).RowsAffected
	return rowsAffected
}

func (p *MysqlPersistence) DeleteUser(user persistence.User) int64 {
	rowsAffected := p.db.Delete(&user).RowsAffected
	return rowsAffected
}
