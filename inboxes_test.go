package disposable_test

import (
	"fmt"
	"testing"

	"github.com/buffreak/disposable-email"
	"github.com/stretchr/testify/assert"
)

func TestRandomEmailAddress(t *testing.T) {
	ib := disposable.Inboxes{}
	ib.RandomEmailAddress()
	t.Logf("email address is: %s", ib.Email)
}

func TestSetEmailAddress(t *testing.T) {
	ib := disposable.Inboxes{}
	ib.SetEmailAddress("apatarswag880")
	t.Logf("email address is: %s", ib.Email)
}

func TestGetEmails(t *testing.T) {
	ib := disposable.Inboxes{}
	ib.Email = "jane2344@blondmail.com"
	emails, err := ib.GetEmails()
	assert.Nil(t, err)
	for _, email := range emails {
		t.Logf("Email ID: %s\n", email.ID)
		t.Logf("Email From: %s\n", email.From)
		t.Logf("Email Timestamp: %d\n", email.Timestamp)
		// t.Logf("Email Body: %+v\n", email.Body)
		fmt.Println()
	}
}
