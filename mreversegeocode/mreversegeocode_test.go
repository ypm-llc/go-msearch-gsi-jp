package mreversegeocode_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ypm-llc/go-msearch-gsi-jp/mreversegeocode"
	"github.com/ypm-llc/go-msearch-gsi-jp/types"
)

func TestReverseGeocode(t *testing.T) {
	type reverseGeocodeTest struct {
		name        string
		lat         float64
		lon         float64
		response    string
		expected    *types.Address
		expectedErr string
	}

	tests := []reverseGeocodeTest{
		{
			name:     "success",
			lon:      140.756941807778,
			lat:      41.797539945,
			response: `{"results":{"muniCd":"01202","lv01Nm":"五稜郭町"}}`,
			expected: &types.Address{
				Results: &types.AddressResults{
					MuniCd: "01202",
					Lv01Nm: "五稜郭町",
				},
			},
		},
		{
			name:     "no result",
			lat:      0,
			lon:      0,
			response: `[]`,
			expected: &types.Address{},
		},
	}

	for _, tc := range tests {
		addr, err := mreversegeocode.LatLonToAddress(tc.lat, tc.lon)

		if !reflect.DeepEqual(addr, tc.expected) {
			t.Errorf("test name=%v, expected: %#v, but got: %#v", tc.name, tc.expected.Results, addr.Results)
		}

		if (err != nil && tc.expectedErr == "") ||
			(err != nil && !strings.Contains(err.Error(), tc.expectedErr)) {
			t.Errorf("test name=%v, expected err: %#v, but got: %#v", tc.name, tc.expectedErr, err)
		}
	}

}
