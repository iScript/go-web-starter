package model

import "fmt"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func UserAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	fmt.Println(err)

	if err != nil {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
