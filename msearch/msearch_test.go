package msearch_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ypm-llc/go-msearch-gsi-jp/msearch"
	"github.com/ypm-llc/go-msearch-gsi-jp/types"
)

func TestSearchAddress(t *testing.T) {
	type searchAddressTest struct {
		name        string
		query       string
		response    string
		expected    []*types.Feature
		expectedErr string
	}

	tests := []searchAddressTest{
		{
			name:     "success",
			query:    "函館市電",
			response: `[{"geometry":{"coordinates":[140.728943,41.768688],"type":"Point"},"type":"Feature","properties":{"addressCode":"","title":"北海道函館市"}}]`,
			expected: []*types.Feature{
				{
					Type: "Feature",
					Properties: &types.Properties{
						AddressCode: "",
						Title:       "北海道函館市",
						DataSource:  "",
					},
					Geometry: &types.Geometry{
						Coordinates: []float64{140.728943, 41.768688},
						Type:        "Point",
					},
				},
			},
		},
		{
			name:     "no result",
			query:    "",
			response: `[]`,
			expected: []*types.Feature{},
		},
	}

	for _, tc := range tests {
		res, err := msearch.SearchAddress(tc.query)

		if !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("test name=%v, expected: %#v, but got: %#v", tc.name, tc.expected, res)
		}

		if (err != nil && tc.expectedErr == "") ||
			(err != nil && !strings.Contains(err.Error(), tc.expectedErr)) {
			t.Errorf("test name=%v, expected err: %#v, but got: %#v", tc.name, tc.expectedErr, err)
		}
	}
}

func TestContainsSearchAddress(t *testing.T) {
	type containsSearchAddressTest struct {
		name        string
		query       string
		response    string
		expected    []*types.Feature
		expectedErr string
	}

	tests := []containsSearchAddressTest{
		{
			name:     "success",
			query:    "函館工業",
			response: `[{"geometry":{"coordinates":[140.7683382,41.79252925],"type":"Point"},"type":"Feature","properties":{"addressCode":"1202","title":"道立北海道函館工業高等学校","dataSource":"3"}},{"geometry":{"coordinates":[140.8028406,41.78336056],"type":"Point"},"type":"Feature","properties":{"addressCode":"1202","title":"函館工業高等専門学校","dataSource":"3"}}]`,
			expected: []*types.Feature{
				{
					Type: "Feature",
					Properties: &types.Properties{
						AddressCode: "1202",
						Title:       "道立北海道函館工業高等学校",
						DataSource:  "3",
					},
					Geometry: &types.Geometry{
						Coordinates: []float64{140.7683382, 41.79252925},
						Type:        "Point",
					},
				},
				{
					Type: "Feature",
					Properties: &types.Properties{
						AddressCode: "1202",
						Title:       "函館工業高等専門学校",
						DataSource:  "3",
					},
					Geometry: &types.Geometry{
						Coordinates: []float64{140.8028406, 41.78336056},
						Type:        "Point",
					},
				},
			},
		},
		{
			name:     "no result",
			query:    "",
			response: `[]`,
			expected: []*types.Feature{},
		},
	}

	for _, tc := range tests {
		res, err := msearch.ContainsSearchAddress(tc.query)

		if !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("test name=%v, expected: %#v, but got: %#v", tc.name, tc.expected, res)
		}

		if (err != nil && tc.expectedErr == "") ||
			(err != nil && !strings.Contains(err.Error(), tc.expectedErr)) {
			t.Errorf("test name=%v, expected err: %#v, but got: %#v", tc.name, tc.expectedErr, err)
		}
	}

}
