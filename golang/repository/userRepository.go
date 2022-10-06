package repository

import (
	"log"
	"sanbertutor/config"
	"sanbertutor/models"
)

func AddUserRepository(user models.User) (string, error) {
	db := config.ConnectDB()
	defer db.Close()

	var guid string

	sqlStatement := `INSERT INTO user_table (guid_user, email, password, firstname, lastname) VALUES ($1, $2, $3, $4, $5) RETURNING guid_user`

	err := db.QueryRowx(sqlStatement, user.Guid_user, user.Email, user.Password, user.Firstname, user.Lastname).Scan(&guid)
	if err != nil {
		log.Printf("Error create user to database with err: %s", err)
		return guid, err
	}

	return guid, nil
}

func GetUserLoginByEmailRepository(email string) (models.UserRelational, error) {
	db := config.ConnectDB()
	defer db.Close()

	var user models.UserRelational

	sqlStatement := `SELECT guid_user, email, firstname, lastname, password FROM user_table WHERE email=$1`
	err := db.QueryRowx(sqlStatement, email).Scan(&user.Guid_user, &user.Email, &user.Firstname, &user.Lastname, &user.Password)
	if err != nil {
		log.Printf("User with email %s not found with err: %s", email, err)
		return user, err
	}

	return user, nil
}
