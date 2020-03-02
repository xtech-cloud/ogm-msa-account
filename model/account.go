package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Account struct {
	Embedded gorm.Model `gorm:"embedded"`
	UUID     string     `gorm:"column:uuid;type:char(32);not null;unique"`
	Username string     `gorm:"column:username;type:varchar(32);not null;unique"`
	Password string     `gorm:"column:password;type:char(32);not null"`
	Profile  string     `gorm:"column:profile;type:text"`
}

var ErrAccountExits = errors.New("account exists")

func (Account) TableName() string {
	return "msa_account"
}

type AccountDAO struct {
}

func NewAccountDAO() *AccountDAO {
	return &AccountDAO{}
}

func (AccountDAO) StrengthenPassword(_password string, _salt string) string {
	return ToUUID(_password + _salt + saltSuffix)
}

func (AccountDAO) Exists(_username string) (bool, error) {
	db, err := openSqlDB()
	if nil != err {
		return false, err
	}
	defer closeSqlDB(db)

	var account Account
	err = db.Where("username= ?", _username).First(&account).Error
	if nil != err {
		return false, err
	}
	return "" != account.UUID, nil
}

func (AccountDAO) Insert(_account Account) error {
	db, err := openSqlDB()
	if nil != err {
		return err
	}
	defer closeSqlDB(db)

	return db.Create(&_account).Error
}

func (AccountDAO) UpdateProfile(_uuid string, _profile string) error {
	db, err := openSqlDB()
	if nil != err {
		return err
	}
	defer closeSqlDB(db)

	return db.Model(&Account{}).Where("uuid = ?", _uuid).Update("profile", _profile).Error
}

func (AccountDAO) UpdatePassword(_uuid string, _password string) error {
	db, err := openSqlDB()
	if nil != err {
		return err
	}
	defer closeSqlDB(db)

	return db.Model(&Account{}).Where("uuid = ?", _uuid).Update("password", _password).Error
}

/*

func (AccountDAO) List() ([]Account, error) {
	var accounts []Account
	err := db.Find(&accounts).Error
	return accounts, err
}
*/

func (AccountDAO) Find(_uuid string) (Account, error) {
	var account Account
	db, err := openSqlDB()
	if nil != err {
		return account, err
	}
	defer closeSqlDB(db)

	res := db.Where("uuid = ?", _uuid).First(&account)
	if res.RecordNotFound() {
		return Account{}, nil
	}
	return account, err
}

func (AccountDAO) WhereUsername(_username string) (Account, error) {
	var account Account
	db, err := openSqlDB()
	if nil != err {
		return account, err
	}
	defer closeSqlDB(db)

	res := db.Where("username= ?", _username).First(&account)
	if res.RecordNotFound() {
		return Account{}, nil
	}
	return account, res.Error
}
