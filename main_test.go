package bands

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	APP_ID         = "barcelona"
	VALID_ARTIST   = "Skrillex"
	INVALID_ARTIST = "donotexists"
)

var artists = []string{
	"mbid_65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab",
	"Skrillex",
	"fbid_6885814958",
}

var client = New(APP_ID)

func TestAcceptKey(t *testing.T) {
	assert.Equal(t, client.API_KEY, APP_ID, "app id is not the same")
}

func TestGetArtist(t *testing.T) {
	for _, artistName := range artists {
		artist, err := client.GetArtist(artistName)

		assert.Nil(t, err)
		assert.NotEmpty(t, artist, "artist return is incorrect")
	}
}

func TestGetArtistEvents(t *testing.T) {
	events, err := client.GetArtistEvents(VALID_ARTIST)

	assert.Nil(t, err)
	assert.True(t, len(events) > 0, "events should return more than 1")
	event := events[0]
	assert.NotNil(t, event.ID)
}

func TestEventsParsing(t *testing.T) {
	events, err := client.GetArtistEvents(VALID_ARTIST)
	t.Log(events)
	assert.Nil(t, err)
	payload, err := json.Marshal(events)
	assert.Nil(t, err)

	err = json.Unmarshal(payload, &events)
	assert.Nil(t, err)
}

func TestGetArtistNotFound(t *testing.T) {
	_, err := client.GetArtistEvents(INVALID_ARTIST)

	assert.Error(t, err, "does not return Unknown Artist error object")
	assert.Equal(t, err, errors.New("Unknown Artist"))
}

func TestNotFound(t *testing.T) {
	_, err := client.GetArtist("test/test")

	assert.Error(t, err, "does not return 404 error")
	assert.Equal(t, err, errors.New("status code 404"))
}
