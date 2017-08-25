package util

import (
	"log"
	"os"
)

func GetenvOrDie(k string) string {
	s := os.Getenv(k)
	if s == "" {
		log.Fatalln("missing req'd env var", k)
	}

	return s
}
