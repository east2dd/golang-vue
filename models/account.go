package models

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	u "github.com/xyingsoft/golang-vue/utils"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID       uint
	Email    string
	Password string
	Token    string
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	rows, err := db.Query(`SELECT id, email FROM accounts WHERE email = ?`, account.Email)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&temp.ID, &temp.Email)
		checkErr(err)
	}

	if temp.Email != "" {
		return u.Message(false, "Duplicated Email"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	res, err := db.Exec(`INSERT INTO accounts(email, password) VALUES( ?, ? )`, account.Email, account.Password)
	checkErr(err)

	if err == nil {
		id, err := res.LastInsertId()
		checkErr(err)

		account.ID = uint(id)
	}

	if account.ID <= 0 {
		return u.Message(false, "Connection Error")
	}

	account.Token = GenerateUniqToken(32)
	account.UpdateToken()
	account.Password = "" //delete password

	response := u.Message(true, "Account Created")
	response["data"] = account

	return response
}

func Login(email, password string) map[string]interface{} {
	account := &Account{}

	rows, err := db.Query(`SELECT id, email, password FROM accounts WHERE email = ?`, email)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&account.ID, &account.Email, &account.Password)
		checkErr(err)
	}

	if account.Email == "" {
		return u.Message(false, "User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Wrong credentials")
	}

	//Worked! Logged In
	account.Password = ""
	account.Token = GenerateUniqToken(32) //Store the token in the response
	account.UpdateToken()

	if err != nil {
		panic(err)
	}

	resp := u.Message(true, "Logged In")
	resp["data"] = account

	return resp
}

func GetUser(u uint) *Account {
	account := &Account{}
	rows, err := db.Query(`SELECT id, email FROM accounts WHERE id = ?`, u)

	for rows.Next() {
		err = rows.Scan(&account.ID, &account.Email, &account.Password)
	}

	checkErr(err)

	if account.Email == "" {
		return nil
	}

	account.Password = ""
	return account
}

func GenerateUniqToken(length uint) string {
	for {
		tokenString, err := GenerateRandomString(32)

		if err != nil {
			panic(err)
		}

		account := GetUserByToken(tokenString)
		if account == nil {
			return tokenString
		}
	}

	return ""
}

func (account *Account) UpdateToken() bool {
	res, err := db.Exec(`UPDATE accounts SET token = ? WHERE id = ?`, account.Token, account.ID)
	checkErr(err)

	var count int64
	count, err = res.RowsAffected()
	checkErr(err)

	if count > 0 {
		return true
	} else {
		return false
	}
}

func GetUserByToken(token string) *Account {
	account := &Account{}
	rows, err := db.Query(`SELECT id, email FROM accounts WHERE token = ?`, token)

	for rows.Next() {
		err = rows.Scan(&account.ID, &account.Email)
	}

	checkErr(err)

	if account.Email == "" { //User not found!
		return nil
	}

	return account
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
