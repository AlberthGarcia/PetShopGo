package handlers

import (
	"net/http"
	"package/models"
)

func GetCategories(rw http.ResponseWriter, r *http.Request) {
	if categories, err := models.GetCategories(); err != nil {
		panic(err)
	} else {
		models.SendResponse(rw, categories)
	}

}
