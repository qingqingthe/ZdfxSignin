package util

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func init() {
	if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); ok {
		log.SetLevel(log.DebugLevel)
	}
}
