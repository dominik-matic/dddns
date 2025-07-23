package apiserver

import (
	"net/http"
	"strings"

	"github.com/dominik-matic/dddns/apiserver/pkg/models"
)

func validateMethod(request *http.Request) (bool, string) {
	if request.Method != http.MethodPost && request.Method != http.MethodDelete {
		return false, "Only POST and DELETE allowed"
	}
	return true, ""
}

func validateAuthorization(request *http.Request, expectedToken string) (bool, string) {
	authHeader := request.Header.Get("Authorization")
	if !checkAuth(authHeader, expectedToken) {
		return false, "Unauthorized"
	}
	return true, ""
}

func validateRequestData(data *models.RequestData, method string) (bool, string) {
	if data.Name == "" {
		return false, "Missing field: name"
	}
	if method == http.MethodPost && data.Value == "" {
		return false, "Missing field: value"
	}
	return true, ""
}

func checkAuth(header, expectedToken string) bool {
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return false
	}
	token := strings.TrimPrefix(header, "Bearer ")
	return token == expectedToken
}
