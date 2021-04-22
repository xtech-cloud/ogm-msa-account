package model

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha512"
	"errors"
	"ogm-msa-account/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	return "ogm_account"
}

type AccountDAO struct {
	conn *Conn
}

func NewAccountDAO(_conn *Conn) *AccountDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &AccountDAO{
		conn: conn,
	}
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

func (this *AccountDAO) GeneratePassword(_password string, _username string) string {
	pwd := doSHA512(_password)
	secretAES := doMD5(_username + config.Schema.Encrypt.Secret)
	passwd := doAES(pwd, secretAES)
	hash, _ := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)
	password := string(hash)
	return password
}

func (this *AccountDAO) VerifyPassword(_password string, _username string, _dbPWD string) error {
	pwd := doSHA512(_password)
	secretAES := doMD5(_username + config.Schema.Encrypt.Secret)
	passwd := doAES(pwd, secretAES)
	return bcrypt.CompareHashAndPassword([]byte(_dbPWD), passwd)
}

func (this *AccountDAO) Exists(_username string) (bool, error) {
	db := this.conn.DB
	var account Account
	result := db.Where("username= ?", _username).Limit(1).Find(&account)
	return "" != account.UUID, result.Error
}

func (this *AccountDAO) Insert(_account *Account) error {
	db := this.conn.DB
	return db.Create(_account).Error
}

func (this *AccountDAO) UpdateProfile(_uuid string, _profile string) error {
	db := this.conn.DB
	return db.Model(&Account{}).Where("uuid = ?", _uuid).Update("profile", _profile).Error
}

func (this *AccountDAO) UpdatePassword(_uuid string, _password string) error {
	db := this.conn.DB
	return db.Model(&Account{}).Where("uuid = ?", _uuid).Update("password", _password).Error
}

func (this *AccountDAO) Find(_uuid string) (Account, error) {
	db := this.conn.DB
	var account Account
	res := db.Where("uuid = ?", _uuid).First(&account)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return Account{}, nil
	}
	return account, res.Error
}

func (this *AccountDAO) WhereUsername(_username string) (Account, error) {
	db := this.conn.DB
	var account Account
	res := db.Where("username= ?", _username).Limit(1).Find(&account)
	return account, res.Error
}

func (this *AccountDAO) List(_offset int64, _count int64) ([]*Account, error) {
	db := this.conn.DB
	var accounts []*Account
	res := db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc").Find(&accounts)
	return accounts, res.Error
}
func (this *AccountDAO) Count() (int64, error) {
	db := this.conn.DB
	count := int64(0)
	res := db.Model(&Account{}).Count(&count)
	return count, res.Error
}
