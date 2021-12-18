package doucovna

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
)

type Tutor struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Bio      string `json:"bio,omitempty"`
	password string
}

func hashPassword(email string, password string) string {
	hashedPwd := fmt.Sprintf("%x", sha1.Sum([]byte(email+password)))
	return hashedPwd
}

func NewTutor(username string, email string, bio string, password string) Tutor {
	hashedPassword := hashPassword(email, password)

	return Tutor{
		Id:       -1,
		Username: username,
		Email:    email,
		Bio:      bio,
		password: hashedPassword,
	}
}

func (d *DbConnector) GetTutors() []Tutor {
	row, err := d.db.Query("SELECT rowid,* FROM tutors")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var tutors []Tutor
	for row.Next() {
		var id int64
		var username string
		var email string
		var password string
		var bio string

		row.Scan(&id, &email, &username, &password, &bio)
		tutors = append(tutors, Tutor{
			Id:       id,
			Username: username,
			password: password,
			Email:    email,
			Bio:      bio,
		})
	}

	return tutors
}

func (d *DbConnector) GetTutorsJson() (string, error) {
	jsonTutors, err := json.Marshal(d.GetTutors())
	return string(jsonTutors), err
}

func (d *DbConnector) GetTutorById(id int64) Tutor {
	row, err := d.db.Query("SELECT rowid,* FROM tutors WHERE rowid=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var tutors []Tutor
	for row.Next() {
		var id int64
		var username string
		var email string
		var password string
		var bio string

		row.Scan(&id, &email, &username, &password, &bio)
		tutors = append(tutors, Tutor{
			Id:       id,
			Username: username,
			password: password,
			Email:    email,
			Bio:      bio,
		})
	}

	if len(tutors) == 0 {
		return Tutor{}
	}

	return tutors[0]
}

func (d *DbConnector) GetTutorByEmail(email string) Tutor {
	row, err := d.db.Query("SELECT rowid,* FROM tutors WHERE email=?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var tutors []Tutor
	for row.Next() {
		var Id int64
		var Username string
		var Email string
		var password string
		var Bio string

		row.Scan(&Id, &Email, &Username, &password, &Bio)
		tutors = append(tutors, Tutor{
			Id:       Id,
			Username: Username,
			password: password,
			Email:    Email,
			Bio:      Bio,
		})
	}

	if len(tutors) == 0 {
		return Tutor{}
	}

	return tutors[0]
}

func (d *DbConnector) CheckPassword(email string, password string) bool {
	tutor := d.GetTutorByEmail(email)

	if tutor.password == hashPassword(email, password) {
		return true
	}

	return false
}
