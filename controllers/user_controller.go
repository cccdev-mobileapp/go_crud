package controllers

import (
	"encoding/json"
	"go-crud/config"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {

	rows, err := config.GetDB().Query("SELECT name, is_verified, role FROM user")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	defer rows.Close()

	var userList []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Name, &user.IsVerified, &user.Role)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		userList = append(userList, user)
	}

	json.NewEncoder(writer).Encode(userList)
}

func AddNewUser(writer http.ResponseWriter, request *http.Request) {
	var newUser models.User

	err := json.NewDecoder(request.Body).Decode(&newUser)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the new user into the database
	id, insertErr := models.InsertUser(newUser)
	if insertErr != nil {
		http.Error(writer, insertErr.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	type UserResponse struct {
		Message string      `json:"message"`
		User    models.User `json:"user"`
		Id      int         `json:"id"`
	}

	userResponse := UserResponse{
		Message: "User Created Successfully!",
		User:    newUser,
		Id:      id,
	}

	json.NewEncoder(writer).Encode(userResponse)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(writer, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteUserByID(id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	type DeleteUserResponse struct {
		Message string `json:"message"`
		ID      int    `json:"id"`
	}

	deleteUserResponse := DeleteUserResponse{Message: "User Deleted Successfully!", ID: id}

	json.NewEncoder(writer).Encode(deleteUserResponse)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(writer, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User

	err = json.NewDecoder(request.Body).Decode(&updatedUser)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	updateError := models.UpdateUserByID(id, updatedUser)

	if updateError != nil {
		http.Error(writer, updateError.Error(), http.StatusBadRequest)
		return
	}

	type UpdateUserResonse struct {
		Message string
		User    models.User
	}

	updateUserResonse := UpdateUserResonse{Message: "User Updated Successfully!", User: updatedUser}

	json.NewEncoder(writer).Encode(updateUserResonse)
}
