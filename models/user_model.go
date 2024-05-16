package models

import "go-crud/config"

type User struct {
	Name       string `json:"name"`
	IsVerified bool   `json:"isVerified"`
	Role       int    `json:"roleType"`
}

func InsertUser(user User) (int, error) {
	result, err := config.GetDB().Exec("INSERT INTO user (name, is_verified, role) VALUES (?, ?, ?)", user.Name, user.IsVerified, user.Role)
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func UpdateUserByID(id int, user User) error {
	_, err := config.GetDB().Exec("UPDATE user SET name=?, is_verified=?, role=? WHERE id=?", user.Name, user.IsVerified, user.Role, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByID(id int) error {
	_, err := config.GetDB().Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
