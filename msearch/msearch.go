package msearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ddo/rq"
	"github.com/ypm-llc/go-msearch-gsi-jp/types"
)

func buildRequest(endpoint string, query map[string][]string) *rq.Rq {
	req := rq.Get(endpoint)
	req.Query = query
	return req
}

func doRequest(req *rq.Rq) (*http.Response, error) {
	hr, err := req.ParseRequest()
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(hr)
}

func processBodyByRequest(req *rq.Rq) ([]byte, error) {
	res, err := doRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func SearchAddress(q string) ([]*types.Feature, error) {
	req := buildRequest("https://msearch.gsi.go.jp/address-search/AddressSearch", map[string][]string{"q": {q}})

	body, err := processBodyByRequest(req)
	if err != nil {
		return nil, err
	}

	resData := []*types.Feature{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return nil, err
	}

	return resData, nil
}

func ContainsSearchAddress(q string) ([]*types.Feature, error) {
	features, err := SearchAddress(q)
	if err != nil {
		return nil, err
	}

	res := []*types.Feature{}

	for _, feature := range features {
		fmt.Printf("title: %v, q: %v", feature.Properties.Title, q)
		if strings.Contains(feature.Properties.Title, q) {
			res = append(res, feature)
		}
	}
	return res, nil
}

func ReverseGeocode(lat float64, lon float64) (*types.Address, error) {
	req := buildRequest("https://mreversegeocoder.gsi.go.jp/reverse-geocoder/LonLatToAddress", map[string][]string{"lat": {fmt.Sprintf("%v", lat)}, "lon": {fmt.Sprintf("%v", lon)}})

	body, err := processBodyByRequest(req)
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
