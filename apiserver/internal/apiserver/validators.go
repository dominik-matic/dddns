package apiserver

import (
	"net/http"
	"strings"

	"github.com/dominik-matic/dddns/apiserver/pkg/models"
)

func validateMethod(request *http.Request) (bool, string) {
	if request.Method != http.MethodPost {
		return false, "Only POST allowed"
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

func validateRequestData(data *models.RequestData) (bool, string) {
	action := strings.ToLower(data.Action)
	if action != "update" && action != "delete" {
		return false, "Invalid action"
	}
	if data.Name == "" {
		return false, "Missing field: name"
	}
	if action == "update" && data.Value == "" {
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
