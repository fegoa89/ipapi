//ipapi package contains a simple wrapper that finds the location of an IP address using ipapi.co.
package ipapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const baseURL string = "https://ipapi.co"

type ApiResponse struct {
	Ip                 string  `json:"ip"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float32 `json:"latitude"`
	Longitude          float32 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	Languages          string  `json:"languages"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

func ClientLocation() (*ApiResponse, error) {
	basicUrl := []string{fmt.Sprintf(baseURL + "/json")}
	requestUrl := strings.Join(basicUrl, appendIpapiKey())
	url := fmt.Sprintf(requestUrl)
	response, err := performLookupRequest(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func FindLocation(ip string) (*ApiResponse, error) {
	basicUrl := []string{fmt.Sprintf(baseURL+"/%s/json", ip)}
	requestUrl := strings.Join(basicUrl, appendIpapiKey())
	url := fmt.Sprintf(requestUrl)
	response, err := performLookupRequest(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func performLookupRequest(url string) (*ApiResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := performRequest(req)
	if err != nil {
		return nil, err
	}
	var data ApiResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func performRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func appendIpapiKey() string {
	ipapiKey, presentIpapiKey := os.LookupEnv("IPAPI_KEY")
	if presentIpapiKey {
		return "?key=" + url.PathEscape(ipapiKey)
	}

	return ""
}
