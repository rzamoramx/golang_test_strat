package database

import (
	"errors"
	"golang_test_strat/domain/models"
	"log"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	db *sql.DB
}

func NewSqlite() any {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}

	prepareDb(db)

	return &SqliteDB{db: db}
}

// Retrieve first occurence of user by email and phone
// returns error when ocurred or User empty when is not located
func (class *SqliteDB) RetrieveUserByUniqueFields(emailVal string, phoneVal string, userVal string) (models.User, error) {
	row := class.db.QueryRow("SELECT user, email, password, phone FROM users WHERE email = ? or phone = ? or user = ?", emailVal, phoneVal, userVal)
	userObj, err := class.dataTo(row)
	if err != nil {
		return userObj, err
	}

	log.Printf("user retrieved from SqliteDB.RetrieveUserByUniqueFields(): %+v\n", userObj)

	return userObj, nil
}

// Retrieve user by one field
// fieldName is the name of the field to filter
// returns error when ocurred or User empty when is not located
func (class *SqliteDB) RetrieveUser(fieldName string, value string) (models.User, error) {
	row := class.db.QueryRow("SELECT user, email, password, phone FROM users WHERE "+fieldName+" = ?", value)
	userObj, err := class.dataTo(row)
	if err != nil {
		return userObj, err
	}

	log.Printf("user retrieved from SqliteDB.RetrieveUser(): %+v\n", userObj)

	return userObj, nil
}

// Saves user
// returns error or nil when all is ok
func (class *SqliteDB) SaveUser(user models.User) error {
	sts := `
	INSERT INTO users(user, email, password, phone) VALUES('` +
		user.User + `', '` + user.Email + `', '` + user.Password + `', '` + user.Phone + `');
	`
	res, err := class.db.Exec(sts)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("cannot insert")
	}

	return nil
}

// Transform row to models.User
func (class *SqliteDB) dataTo(row *sql.Row) (models.User, error) {
	var user, password, email, phone string

	err := row.Scan(&user, &email, &password, &phone)
	if err == sql.ErrNoRows {
		return models.User{}, nil
	} else if err != nil {
		log.Fatal(err)
		return models.User{}, err
	}

	return models.User{User: user, Email: email, Password: password, Phone: phone}, nil
}

// Initialize database
func prepareDb(db *sql.DB) error {
	sts := `
	DROP TABLE IF EXISTS users;
	CREATE TABLE users(id INTEGER PRIMARY KEY AUTOINCREMENT, user TEXT, email TEXT, phone TEXT, password TEXT);
	`
	_, err := db.Exec(sts)
	if err != nil {
		return err
	}

	log.Println("table users created")
	return nil
}
