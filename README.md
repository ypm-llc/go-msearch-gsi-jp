# go-msearch-gsi-jp

A client library for APIs that presented by Geospatial Information Authority of Japan

## List of Supported APIs

* Geocoding API - `GET https://msearch.gsi.go.jp/address-search/AddressSearch?q=XXXX`
* Reverse Geocoding API - `GET https://mreversegeocoder.gsi.go.jp/reverse-geocoder/LonLatToAddress?lat=XXXX&lon=XXXX`

## Synopsis

### Geocoding

Example 1: Search address that relates to query "新宿駅" (Shinjuku station) and show their titles

    import (
        "fmt"
        "github.com/ypm-llc/go-msearch-gsi-jp/msearch"
    )
    query := "新宿駅"
    geolist, err := msearch.SearchAddress(query)
    if err != nil {
        panic(fmt.Sprintf("Error: %v", err))
    }
    for _, geo := range geolist {
        fmt.Printf("address: %v", geo.Properties.Title)
    }

Example 2. Search address that contains "五反田" (Gotanda) in their title and show these

    import (
        "fmt"
        "github.com/ypm-llc/go-msearch-gsi-jp/msearch"
    )
    query := "五反田"
    geolist, err := msearch.ContainsSearchAddress(query)
    if err != nil {
        panic(fmt.Sprintf("Error: %v", err))
    }
    for _, geo := range geolist {
        fmt.Printf("address: %v", geo.Properties.Title)
    }

### Reverse Geocoding

Example. Fetch address that matches by specified coordinates and show these name

    import (
        "fmt"
        "github.com/ypm-llc/go-msearch-gsi-jp/msearch"
    )
    latitude := 140.756941807778
    longitude := 41.797539945
    addr, err := msearch.ReverseGeocoding(latitude, longitude)
    if err != nil {
        panic(fmt.Sprintf("Error: %v", err))
    }
    fmt.Printf("address: %v", addr.Results.Lv01Nm) // "五稜郭町" (Goryoukaku-cho)

## See Also

* [Geospatial Information Authority of Japan](https://www.gsi.go.jp/)
* [https://memo.appri.me/programming/gsi-geocoding-api](https://memo.appri.me/programming/gsi-geocoding-api)