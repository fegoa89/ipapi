package ipapi

import (
	_ "fmt"
	"gopkg.in/h2non/gock.v1"
	"reflect"
	"testing"
)

var apiResponse = map[string]interface{}{
	"ip":                   "178.13.214.11",
	"city":                 "Frankfurt am Main",
	"region":               "Hesse",
	"region_code":          "HE",
	"country":              "DE",
	"country_name":         "Germany",
	"continent_code":       "EU",
	"in_eu":                true,
	"postal":               "60313",
	"latitude":             50.1153,
	"longitude":            8.6823,
	"timezone":             "Europe/Berlin",
	"utc_offset":           "+0100",
	"country_calling_code": "+49",
	"currency":             "EUR",
	"languages":            "de",
	"asn":                  "AS3209",
	"org":                  "Vodafone GmbH",
}

func TestClientLocation(t *testing.T) {
	defer gock.Off()

	gock.New("https://ipapi.co").
		Get("/json").
		Reply(200).
		JSON(apiResponse)

	response, err := ClientLocation()

	if err != nil {
		t.Errorf("Method raised an unexpected error %s", err)
	}

	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		t.Errorf("Returned response is not a pointer")
	}

	if reflect.TypeOf(response).String() != "*ipapi.ApiResponse" {
		t.Errorf("Returned response object does not have the expected type")
	}
}

func TestFindLocation(t *testing.T) {
	defer gock.Off()

	gock.New("https://ipapi.co").
		Get("/178.13.214.11/json").
		Reply(200).
		JSON(apiResponse)

	response, err := FindLocation("178.13.214.11")

	if err != nil {
		t.Errorf("Method raised an unexpected error %s", err)
	}

	if reflect.ValueOf(response).Kind() != reflect.Ptr {
		t.Errorf("Returned response is not a pointer")
	}

	if reflect.TypeOf(response).String() != "*ipapi.ApiResponse" {
		t.Errorf("Returned response object does not have the expected type")
	}
}
