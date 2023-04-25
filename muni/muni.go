/*
a package that manupulates muni.js
*/
package muni

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ypm-llc/go-msearch-gsi-jp/httpreq"
	"github.com/ypm-llc/go-msearch-gsi-jp/types"
)

// Muni file url
const MuniURL = "https://maps.gsi.go.jp/js/muni.js"
const MuniRecordPrefix = "GSI.MUNI_ARRAY["

// get muni map (city or ward map by city code)
func GetMuniMap() (types.MuniMap, error) {
	req := httpreq.Build(MuniURL, nil)
	body, err := httpreq.ProcessBody(req)
	if err != nil {
		return nil, err
	}

	muniMap, err := parseMuniMap(body)
	if err != nil {
		return nil, err
	}

	return muniMap, nil
}

// converts muni code to address name.
func MuniCodeToAddressName(muniMap types.MuniMap, muniCode int) (string, error) {
	muniRec, ok := muniMap[muniCode]
	if !ok {
		return "", fmt.Errorf("invalid muni code: %d", muniCode)
	}

	addr := fmt.Sprintf("%s%s", muniRec.PrefName, muniRec.CityName)
	return addr, nil
}

// converts address result to address name.
func AddressResultsToAddressName(muniMap types.MuniMap, addrResults *types.AddressResults) (string, error) {
	matcher := regexp.MustCompile(`^0`)
	mc := string(matcher.ReplaceAll([]byte(addrResults.MuniCd), []byte("")))

	muniCode, err := strconv.Atoi(mc)
	if err != nil {
		return "", err
	}
	muniName, err := MuniCodeToAddressName(muniMap, muniCode)
	if err != nil {
		return "", err
	}

	addrName := fmt.Sprintf("%s%s", muniName, addrResults.Lv01Nm)
	return addrName, nil
}

func parseMuniMap(body []byte) (types.MuniMap, error) {
	buf := bytes.NewBuffer(body)
	r := bufio.NewReader(buf)

	muniMap := types.MuniMap{}

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		if bytes.HasPrefix(line, []byte(MuniRecordPrefix)) {
			muniRec, err := parseMuni(line)
			if err != nil {
				return nil, err
			}
			muniMap[muniRec.CityCode] = muniRec
		}
	}

	return muniMap, nil
}

func parseMuni(line []byte) (*types.MuniRecord, error) {
	muniStr := string(line)
	muniStr = strings.Replace(muniStr, MuniRecordPrefix, "", 1)
	muniStr = strings.Replace(muniStr, "] = ", "", 1)
	muniStr = strings.Replace(muniStr, ";", "", 1)
	muniStr = strings.Replace(muniStr, "\"", "", -1)
	muniStr = strings.Replace(muniStr, "'", "", -1)

	items := strings.Split(muniStr, ",")
	if len(items) != 4 {
		return nil, fmt.Errorf("invalid muni record: %s", muniStr)
	}

	prefCode, err := strconv.Atoi(items[0])
	if err != nil {
		return nil, err
	}

	cityCode, err := strconv.Atoi(items[2])
	if err != nil {
		return nil, err
	}

	muniRec := &types.MuniRecord{
		PrefCode: prefCode,
		PrefName: items[1],
		CityCode: cityCode,
		CityName: items[3],
	}

	return muniRec, nil
}
