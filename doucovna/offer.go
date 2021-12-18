package doucovna

import (
	"encoding/json"
	"log"
)

type Offer struct {
	Id          int64  `json:"id,omitempty"`
	Teacher     int64  `json:"teacher,omitempty"`
	Subject     int64  `json:"subject,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewOffer(teacher int64, subject int64, description string) Offer {
	return Offer{}
	//Offer{
	//	Id:          Id,
	//	Teacher:     Teacher,
	//	Subject:     Subject,
	//	Description: Description,
	//}
}

func (d *DbConnector) GetOffersJson() (string, error) {
	jsonOffers, err := json.Marshal(d.GetOffers())
	return string(jsonOffers), err
}

func (d *DbConnector) GetOffers() []Offer {
	row, err := d.db.Query("SELECT rowid,* FROM offers")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var offers []Offer
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int64
		var teacher int64
		var subject int64
		var description string

		row.Scan(&id, &teacher, &subject, &description)
		offers = append(offers, Offer{
			Id:          id,
			Teacher:     teacher,
			Subject:     subject,
			Description: description,
		})
	}

	return offers
}
