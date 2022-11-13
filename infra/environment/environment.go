package environment

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Netflix/go-env"
	_ "github.com/lib/pq"
)

// var secretCache, _ = secretcache.New()

type DB struct {
	DbUser string `json:"username"`
	DBPass string `json:"password"`
	DBhost string `json:"host"`
	DBPort int64  `json:"port"`
}

type Environment struct {
	App struct {
		App_version string `env:"APPVERSION"`
		App_name    string `env:"APPNAME"`
	} `json:"App"`

	Env string `env:"ENV"`
	DB  DB
}

// LoadOrDie responsible for loading the environment variables in the cloud
func LoadOrDie() *Environment {
	env := os.Getenv("ENV")

	if env == "DEV" {
		return LoadFromLocal()
	}
	return LoadFromSecret()
}

func LoadFromLocal() *Environment {

	environment := new(Environment)
	_, err := env.UnmarshalFromEnviron(environment)
	if err != nil {
		log.Fatal("Error on Load environment variables ", err)
	}

	return environment
}

func LoadFromSecret() *Environment {

	environment := Environment{}

	fmt.Println("################################################################")
	mostrar, _ := json.MarshalIndent(environment, "", "")
	fmt.Println(string(mostrar))
	fmt.Println("################################################################")

	return &environment

}
