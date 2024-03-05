package service

import "fiber-tutorial/internal/datastore"

type Services struct {
	UserService UserService
}

func InitServices(mysql_db *datastore.Mysqldb) Services {
	return Services{
		UserService: *NewUserService(mysql_db),
	}
}
