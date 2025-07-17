package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dominik-matic/dddns/apiserver/internal/db"
	"github.com/dominik-matic/dddns/apiserver/pkg/models"
)

func NewUpdateHandler(authToken string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handling shit")
		if ok, msg := validateMethod(request); !ok {
			http.Error(writer, msg, http.StatusMethodNotAllowed)
			return
		}
		if ok, msg := validateAuthorization(request, authToken); !ok {
			http.Error(writer, msg, http.StatusUnauthorized)
			return
		}

		var data models.RequestData
		err := json.NewDecoder(request.Body).Decode(&data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		if ok, msg := validateRequestData(&data); !ok {
			http.Error(writer, msg, http.StatusBadRequest)
			return
		}

		prepareData(&data)

		if data.Action == "update" {
			err = db.InsertOrUpdate(data)
		} else {
			err = db.Delete(data)
		}
		if err != nil {
			http.Error(writer, "Db error", http.StatusInternalServerError)
			log.Printf("Db error: %v", err)
			return
		}

		fmt.Fprintln(writer, "Update successful")
	}
}
func prepareData(data *models.RequestData) {
	data.Name = strings.ToLower(data.Name)
	data.Type = strings.ToUpper(data.Type)
	data.Value = strings.ToLower(data.Value)
	data.Action = strings.ToLower(data.Action)

	if data.Type == "" {
		data.Type = "A"
	}
}
