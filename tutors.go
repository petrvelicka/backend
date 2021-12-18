package main

import (
	"doucovna/doucovna"
	"fmt"
	"github.com/drexedam/gravatar"
	"html/template"
	"net/http"
)

type OffersTutorsData struct {
	Subjects []doucovna.Subject
	Tutors   []OffersData
}

type OffersData struct {
	ProfilePicture string
	TutorId        int64
	Name           string
	Subject        string
	Description    string
}

func genProfilePicture(email string) string {
	return gravatar.New(email).Size(256).AvatarURL()
}

func getOffersData() []OffersData {
	offers := db.GetOffers()
	var offersData []OffersData

	for _, offer := range offers {
		tutor := db.GetTutorById(offer.Teacher)
		subject := db.GetSubjectById(offer.Subject)
		fmt.Println(tutor.Email)
		offersData = append(offersData, OffersData{
			ProfilePicture: genProfilePicture(tutor.Email),
			TutorId:        tutor.Id,
			Name:           tutor.Username,
			Subject:        subject.Name,
			Description:    offer.Description,
		})
	}

	return offersData
}

func tutorsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/tutors.html"))
	data := OffersTutorsData{Subjects: db.GetSubjects(), Tutors: getOffersData()}
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
