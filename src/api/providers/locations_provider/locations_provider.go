package locations_provider

import (
	"Golang-all-type-of-tests/src/api/domain/locations"
	"Golang-all-type-of-tests/src/utils/errors"
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	response := rest.Get(fmt.Sprintf(urlGetCountry, countryId))

	if response == nil || response.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}

	if response.StatusCode < 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error response when getting country %s", countryId),
			}
		}
	}
	var result locations.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying unmarshal country data for %s", countryId),
		}
	}
	return &result, nil
}
