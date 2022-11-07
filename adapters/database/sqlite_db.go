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

func NewDummyDD() any {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}

	prepareDb(db)

	return &SqliteDB{db: db}
}

func (class *SqliteDB) Retrieve(email string) (models.User, error) {
	stm, err := class.db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stm.Close()

	var user string
	var password string
	var phone string

	err = stm.QueryRow(email).Scan(&user, &password, &phone)
	if err == sql.ErrNoRows {
		return models.User{}, nil
	} else if err != nil {
		log.Fatal(err)
		return models.User{}, err
	}

	userObj := models.User{User: user, Email: email, Password: password, Phone: phone}
	log.Printf("user retrieved from sqlite: %+v\n", userObj)

	return userObj, nil
}

func (class *SqliteDB) Save(user models.User) error {
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

func prepareDb(db *sql.DB) error {
	sts := `
	DROP TABLE IF EXISTS users;
	CREATE TABLE users(id INTEGER PRIMARY KEY, user TEXT, email TEXT, phone TEXT, password TEXT);
	`
	_, err := db.Exec(sts)
	if err != nil {
		return err
	}

	log.Println("table users created")
	return nil
}
