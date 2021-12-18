package doucovna

import (
	"encoding/json"
	"log"
)

type Subject struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewSubject(name string) Subject {
	return Subject{Id: -1, Name: name}
}

func (d *DbConnector) GetSubjectsJson() (string, error) {
	jsonSubjects, err := json.Marshal(d.GetSubjects())
	return string(jsonSubjects), err
}

func (d *DbConnector) GetSubjects() []Subject {
	row, err := d.db.Query("SELECT rowid,* FROM subjects")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var subjects []Subject
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int64
		var name string

		row.Scan(&id, &name)
		subjects = append(subjects, Subject{
			Id:   id,
			Name: name,
		})
	}

	return subjects
}

func (d *DbConnector) GetSubjectById(id int64) Subject {
	row, err := d.db.Query("SELECT rowid,* FROM subjects WHERE rowid=?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var subjects []Subject
	for row.Next() {
		var id int64
		var name string

		row.Scan(&id, &name)
		subjects = append(subjects, Subject{
			Id:   id,
			Name: name,
		})
	}

	if len(subjects) == 0 {
		return Subject{}
	}

	return subjects[0]
}
