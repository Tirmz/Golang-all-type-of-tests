package locations_provider

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetCountryRestClientError(t *testing.T) {

	//Execution
	country, err := GetCountry("AR")

	//Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country AR", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	//Execution
	country, err := GetCountry("AR")

	//Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	//Execution
	country, err := GetCountry("AR")

	//Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country AR", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	//Execution
	country, err := GetCountry("AR")

	//Validation
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	//Execution
	country, err := GetCountry("AR")

	//Validation
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))

}
