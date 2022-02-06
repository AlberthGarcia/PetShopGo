package handlers

import (
	"encoding/json"
	"net/http"
	"package/models"
	"strconv"

	"github.com/gorilla/mux"
)

//HANDLER TO GET ALL CATEGORIES
func GetCategories(rw http.ResponseWriter, r *http.Request) {
	if categories, err := models.GetCategories(); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendResponse(rw, categories, "Get categories succesful")
	}
}

//HANDLER TO GET A CATEGORIES
func GetCategoryById(rw http.ResponseWriter, r *http.Request) {
	if category, err := getIdByUrl(r); err != nil {
		models.SendNotFound(rw)
	} else if category.IdCategory == 0 {
		models.SendResponse(rw, category, "Category do not exists")
	} else {
		models.SendResponse(rw, category, "Category found")
	}
}

//FUNC TO CREATE A CATEGORY
func CreateCategory(rw http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&category); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		category.Save()
		models.SendResponse(rw, category, "Created succefull")
	}
}

//hanlder to delete a category
func DeleteCategory(rw http.ResponseWriter, r *http.Request) {

	if category, err := getIdByUrl(r); err != nil {
		models.SendNotFound(rw)
	} else {
		category.DeleteCategoryById()
		models.SendResponse(rw, category, "category deleted")
	}
}

//hanlder to update a category
func UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	var categoryId int64

	if cate, err := getIdByUrl(r); err != nil {
		models.SendNotFound(rw)
	} else {
		categoryId = cate.IdCategory
	}

	decode := json.NewDecoder(r.Body)
	category := models.Category{}
	if err := decode.Decode(&category); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		category.IdCategory = categoryId
		category.Save()
		models.SendResponse(rw, category, "Updated succesfull")
	}
}

//func to get the ID through the URL
func getIdByUrl(r *http.Request) (models.Category, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if category, err := models.GetCategoryById(int64(userId)); err != nil {
		return category, err
	} else {
		return category, nil
	}
}
