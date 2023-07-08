package initializer

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	godotenv.Load()

}
