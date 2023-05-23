package pgdatabase

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func CreateUsersToDatabase(email, name, password string) (bool, uint32, error) {
	var id uint32
	sqlQuery := `INSERT INTO users(email, name, password) VALUES($1, $2, $3) RETURNING id;`
	if err := DB.QueryRow(sqlQuery, email, name, password).Scan(&id); err != nil {
		log.Println("CreateUsersToDatabase QueryRow " + err.Error())
		return false, id, err
	}
	return true, id, nil
}
func GetUsersByIdFromDatabase(id uint32) (Users, error) {
	var users Users
	var usersNull UsersNull
	sqlQuery := `SELECT id, email, name, password, session_id FROM users WHERE id = $1`
	if err := DB.QueryRow(sqlQuery, id).Scan(&users.ID, &users.Email, &usersNull.Name, &users.Password, &usersNull.SessionId); err != nil {
		log.Println("GetUsersFromDatabase Scan ", err.Error())
		return users, err
	}
	users.Name = usersNull.Name.String
	users.SessionId = usersNull.SessionId.String

	return users, nil
}
func GetUsersByEmailFromDatabase(email string) (Users, error) {
	var users Users
	var usersNull UsersNull
	sqlQuery := `SELECT id, email, name, password, session_id FROM users WHERE email = $1`
	if err := DB.QueryRow(sqlQuery, email).Scan(&users.ID, &users.Email, &usersNull.Name, &users.Password, &usersNull.SessionId); err != nil {
		log.Println("GetUsersFromDatabase Scan ", err.Error())
		return users, err
	}
	users.Name = usersNull.Name.String
	users.SessionId = usersNull.SessionId.String

	return users, nil
}
func VerifyUsersEmailToDatabase(email string) bool {
	var id string
	SqlQuery := `SELECT id FROM users WHERE email = $1`
	if err := DB.QueryRow(SqlQuery, email).Scan(&id); err != nil && err == sql.ErrNoRows {
		log.Println("VerifyUsersEmailToDatabase QueryRow ", err.Error())
		return false
	}
	return true
}
func PatchSessionIDUsingEmail(email string, session_id uuid.UUID) (bool, error) {
	sqlQuery := `UPDATE users SET session_id = $1 WHERE email = $2`
	if _, errExec := DB.Exec(sqlQuery, session_id, email); errExec != nil {
		log.Println("PatchSessionIDUsingEmail Exec ", errExec.Error())
		return false, errExec
	}
	return true, nil
}
func VerifyUsersSessionIDToDatabase(session_id string) bool {
	var id string
	SqlQuery := `SELECT id FROM users WHERE session_id = $1`
	if err := DB.QueryRow(SqlQuery, session_id).Scan(&id); err != nil && err == sql.ErrNoRows {
		log.Println("VerifyUsersSessionIDToDatabase QueryRow ", err.Error())
		return false
	}
	return true
}
