package datastore

import (
	"database/sql"
	"errors"
	"fiber-tutorial/internal/model"
	"fiber-tutorial/internal/model/translate"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysqldb struct {
	database *gorm.DB
}

func NewMysqldb() (*Mysqldb, error) {
	// get config param
	dsn := "username:password@tcp(localhost:3306)/fiber-tutorial?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return &Mysqldb{
		database: db,
	}, nil
}

func (d *Mysqldb) InitTables() {
	d.database.AutoMigrate(model.User{})
}

func (d *Mysqldb) CreateUser(u *model.User) (*model.User, error) {

	ret, err := d.GetUser(&model.User{Email: u.Email})
	if err != nil {
		return nil, err
	}
	if len(*ret) > 0 {
		//return nil, fmt.Errorf("bu e-posta hesabı daha önce kullanılmış")
		return nil, errors.New(translate.AlreadyExistError)
	}

	res := d.database.Create(u)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected < 1 {
		return nil, errors.New("herhangi bir satir degismedi")
	}

	return u, nil
}

func (d *Mysqldb) GetUser(u *model.User) (*[]model.User, error) {
	var users []model.User

	// return all users if email equal empty string
	query := `SELECT * FROM users WHERE 
				(email = @email or @email = "")`
	res := d.database.Raw(query, sql.Named("email", u.Email)).Scan(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return &users, nil
}
