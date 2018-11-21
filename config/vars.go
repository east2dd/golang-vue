package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type VarKey int

var ErrMissingVar = errors.New("Missing var.")

const (
	VAR_EMPTY VarKey = iota
	VAR_APP_URL
	VAR_APP_ENV
	VAR_DATABASE_URL
)

var vars map[VarKey]string

func initVars() {
	vars = map[VarKey]string{
		VAR_APP_URL: os.Getenv("APP_URL"),
		VAR_APP_ENV: os.Getenv("APP_ENV"),
		VAR_DATABASE_URL: fmt.Sprintf("%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USERNAME"),
			//os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE")),
	}
}

func GetVar(key VarKey) string {
	value, ok := vars[key]
	if !ok {
		fmt.Println(int(key))
		panic(ErrMissingVar)
	}
	return value
}

func (x VarKey) String() string {
	return GetVar(x)
}

func (x VarKey) Int() int {
	str := GetVar(x)
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func (x VarKey) Bool() bool {
	str := GetVar(x)
	if str == "" {
		return false
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return b
}
