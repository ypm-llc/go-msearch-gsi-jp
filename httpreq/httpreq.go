package httpreq

import (
	"io/ioutil"
	"net/http"

	"github.com/ddo/rq"
)

// build http request
func Build(endpoint string, query map[string][]string) *rq.Rq {
	req := rq.Get(endpoint)
	req.Query = query
	return req
}

// process body by request
func ProcessBody(req *rq.Rq) ([]byte, error) {
	res, err := doRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func doRequest(req *rq.Rq) (*http.Response, error) {
	hr, err := req.ParseRequest()
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(hr)
}
