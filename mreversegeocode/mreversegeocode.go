/*
mreversegeocode is a package for reverse geocoding by latitude and longitude.
*/
package mreversegeocode

import (
	"encoding/json"
	"fmt"

	"github.com/ypm-llc/go-msearch-gsi-jp/httpreq"
	"github.com/ypm-llc/go-msearch-gsi-jp/types"
)

// base url for mreversegeocode api
const BaseURL = "https://mreversegeocoder.gsi.go.jp"

// reverse geocoding by latitude and longitude
func LatLonToAddress(lat float64, lon float64) (*types.Address, error) {
	url := fmt.Sprintf("%v/reverse-geocoder/LonLatToAddress", BaseURL)
	query := map[string][]string{"lat": {fmt.Sprintf("%v", lat)}, "lon": {fmt.Sprintf("%v", lon)}}
	req := httpreq.Build(url, query)

	body, err := httpreq.ProcessBody(req)
	if err != nil {
		return nil, err
	}

	resData := &types.Address{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return nil, err
	}

	return resData, nil
}
