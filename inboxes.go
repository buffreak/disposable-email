package disposable

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-faker/faker/v4"
)

type Inboxes struct {
	Email string
	Data  []EmailDetail
}

type InboxesResonse struct {
	Messages []struct {
		ID        string `json:"uid"`
		From      string `json:"s"`
		Timestamp uint64 `json:"r"`
		Body      string
	} `json:"msgs"`
}

var inboxesExtension = []string{"blondmail.com", "chapsmail.com", "clowmail.com", "dropjar.com", "fivermail.com", "getairmail.com", "getmule.com", "getnada.com", "gimpmail.com", "givmail.com", "guysmail.com", "inboxbear.com", "replyloop.com", "robot-mail.com", "spicysoda.com", "tafmail.com", "temptami.com", "tupmail.com", "vomoto.com"}

func (ib *Inboxes) RandomEmailAddress() error {
	var firstName = faker.FirstName()
	ib.Email = strings.ToLower(fmt.Sprintf("%s%d@%s", firstName, rand.Intn(1000), inboxesExtension[rand.Intn(len(inboxesExtension))]))
	return nil
}

func (ib *Inboxes) SetEmailAddress(username string) error {
	ib.Email = strings.ToLower(fmt.Sprintf("%s@%s", username, inboxesExtension[rand.Intn(len(inboxesExtension))]))
	return nil
}

func (ib *Inboxes) GetEmails() ([]EmailDetail, error) {
	defer HandlePanic("Inboxes base.GetEmails")

	response, err := Request("GET", "https://inboxes.com/api/v2/inbox/"+ib.Email, nil, http.Header{
		"user-agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0"},
		"content-type": {"application/json"},
	})
	if err != nil {
		return nil, errors.New("failed do http response from inboxes get emails")
	}

	var inboxes InboxesResonse
	err = json.NewDecoder(response.Body).Decode(&inboxes)
	if err != nil {
		return nil, errors.New("failed do decode response from inboxes get emails")
	}

	if len(inboxes.Messages) > 0 {
		for key, inbox := range inboxes.Messages {

			ib.Data = append(ib.Data, EmailDetail{
				ID:        inbox.ID,
				From:      inbox.From,
				Timestamp: inbox.Timestamp,
			})

			response, err := Request("GET", "https://inboxes.com/api/v2/message/"+inbox.ID, nil, http.Header{
				"user-agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0"},
				"content-type": {"application/json"},
			})
			if err == nil {
				var body struct {
					Email string `json:"html"`
				}
				err = json.NewDecoder(response.Body).Decode(&body)
				if err == nil {
					ib.Data[key].Body = body.Email
				}
			}
		}
		return ib.Data, nil
	}
	return nil, errors.New("no email found from inboxes")
}

func (ib *Inboxes) GetEmailAddress() string {
	return ib.Email
}
