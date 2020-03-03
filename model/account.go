package model

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha512"
	"errors"
	"omo-msa-account/config"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Embedded gorm.Model `gorm:"embedded"`
	UUID     string     `gorm:"column:uuid;type:char(32);not null;unique"`
	Username string     `gorm:"column:username;type:varchar(32);not null;unique"`
	Password string     `gorm:"column:password;type:char(60);not null"`
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

func doMD5(_pwd string) []byte {
	h := md5.New()
	h.Write([]byte(_pwd))
	return h.Sum(nil)
}

func doSHA512(_pwd string) []byte {
	hash := sha512.New()
	hash.Write([]byte(_pwd))
	return hash.Sum(nil)
}

func doAES(_orig []byte, _secret []byte) []byte {
	block, _ := aes.NewCipher(_secret)
	blocksize := block.BlockSize()

	padding := blocksize - len(_orig)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	orig := append(_orig, padtext...)

	blockmode := cipher.NewCBCEncrypter(block, _secret[:blocksize])
	crypted := make([]byte, len(orig))
	blockmode.CryptBlocks(crypted, orig)
	return crypted
}

func (AccountDAO) GeneratePassword(_password string, _username string) string {
	pwd := doSHA512(_password)
	secretAES := doMD5(_username + config.Schema.Encrypt.Secret)
	passwd := doAES(pwd, secretAES)
	hash, _ := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)
	password := string(hash)
	return password
}

func (AccountDAO) VerifyPassword(_password string, _username string, _dbPWD string) error {
	pwd := doSHA512(_password)
	secretAES := doMD5(_username + config.Schema.Encrypt.Secret)
	passwd := doAES(pwd, secretAES)
	return bcrypt.CompareHashAndPassword([]byte(_dbPWD), passwd)
}

func (AccountDAO) Exists(_username string) (bool, error) {
	db, err := openSqlDB()
	if nil != err {
		return false, err
	}
	defer closeSqlDB(db)

	var account Account
	result := db.Where("username= ?", _username).First(&account)
	if result.RecordNotFound() {
		return false, nil
	}

	return "" != account.UUID, result.Error
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
