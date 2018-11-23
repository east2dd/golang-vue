package models

import (
	"os"
	"strings"

	u "github.com/xyingsoft/golang-vue/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
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

	println(account.Email)
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
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
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
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

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {
	account := &Account{}
	rows, err := db.Query(`SELECT id, email, password FROM accounts WHERE id = ?`, u)

	for rows.Next() {
		err = rows.Scan(&account.ID, &account.Email, &account.Password)
	}

	checkErr(err)

	if account.Email == "" { //User not found!
		return nil
	}

	account.Password = ""
	return account
}
