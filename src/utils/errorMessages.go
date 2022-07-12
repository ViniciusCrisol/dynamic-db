package utils

import (
	"strconv"
)

const INTERNAL_SERVER_ERR_STATUS = 500

var ErrMessages = map[string]string{
	"route-not-found":         "404 - route not found",
	"internal-server-err":     "500 - internal server error",
	"domain-already-exists":   "400 - error, domain already exists",
	"domain-does-not-exists":  "404 - error, domain does not exists",
	"cluster-already-exists":  "400 - error, cluster already exists",
	"cluster-does-not-exists": "404 - error, cluster does not exists",
}

func GetMessageAndHTTPStatusFromErr(err error) (string, int) {
	errContent := err.Error()
	errMessage := ErrMessages[errContent]

	if errMessage == "" {
		return ErrMessages["internal-server-err"], INTERNAL_SERVER_ERR_STATUS
	}

	HTTPStatus, getHTTPStatusErr := getHTTPStatusFromErrMessage(errMessage)
	if getHTTPStatusErr != nil {
		return ErrMessages["internal-server-err"], INTERNAL_SERVER_ERR_STATUS
	}
	return errMessage, HTTPStatus
}

func getHTTPStatusFromErrMessage(errMessage string) (int, error) {
	HTTPStatusPrefix := errMessage[0:3]
	return strconv.Atoi(HTTPStatusPrefix)
}
