/*
msearch is a package for searching address by query string.
*/
package msearch

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ypm-llc/go-msearch-gsi-jp/httpreq"
	"github.com/ypm-llc/go-msearch-gsi-jp/types"
)

// base url for msearch api
const BaseURL = "https://msearch.gsi.go.jp"

// search address by query
func SearchAddress(q string) ([]*types.Feature, error) {
	url := fmt.Sprintf("%v/address-search/AddressSearch", BaseURL)
	query := map[string][]string{"q": {q}}
	req := httpreq.Build(url, query)

	body, err := httpreq.ProcessBody(req)
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

// search address that contains query string
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
