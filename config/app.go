package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/subosito/gotenv"
)

const (
	Production  = "production"
	Development = "development"
)

var ApplicationName = "GoVue"
var Env string
var initer sync.Once

func Init() {
	initer.Do(func() {
		gotenv.Load()
		var env = os.Getenv("APP_ENV")
		if env == "" {
			env = Development
		}
		Env = env

		fmt.Println("Initializing App: ", ApplicationName)
		fmt.Println("Initializing ....  env:", env)

		initVars()
	})
}
