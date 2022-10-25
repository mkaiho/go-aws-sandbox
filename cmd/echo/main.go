package main

import (
	"os"
	"strings"

	"github.com/mkaiho/go-aws-sandbox/logging"
)

func init() {
	logging.InitLoggerWithZap()
}

func main() {
	logging.GetLogger().Info(strings.Join(os.Args[1:], " "))
}
