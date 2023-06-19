package configs

import (
	"os"
)

func GetSuperAccountID() string {
	return os.Getenv("SUPER_USER_ID")
}
