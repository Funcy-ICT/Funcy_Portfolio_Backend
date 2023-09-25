package configs

import (
	"os"
)

type Mode uint

const (
	Unset Mode = iota
	Production
	Local
)

func (m Mode) String() string {
	switch m {
	case Unset:
		return "Unset"
	case Production:
		return "Production"
	case Local:
		return "Local"
	default:
		return "Unset"
	}
}

func GetMode() Mode {
	if m := os.Getenv("MODE"); m == "" {
		return Unset
	} else {
		switch m {
		case "production":
			return Production
		case "local":
			return Local
		default:
			return Unset
		}
	}
}
