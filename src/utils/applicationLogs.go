package utils

import (
	"log"
	"os"
)

const (
	LOG_FLAGS = log.Ldate | log.Ltime | log.Lshortfile

	INFO_LOG_PREFIX  = "INFO: "
	ERROR_LOG_PREFIX = "ERROR: "
	PANIC_LOG_PREFIX = "PANIC: "
)

var (
	InfoLogger  *log.Logger = log.New(os.Stdout, INFO_LOG_PREFIX, LOG_FLAGS)
	ErrorLogger *log.Logger = log.New(os.Stdout, ERROR_LOG_PREFIX, LOG_FLAGS)
	PanicLogger *log.Logger = log.New(os.Stdout, PANIC_LOG_PREFIX, LOG_FLAGS)
)
