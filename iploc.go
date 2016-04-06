// Geolocalize an IP address through ip-api.com
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

// Data from ip-api.com
type GeoIP struct {
    Country     string  `json:"country"`
    CountryCode string  `json:"countryCode"`
    City        string  `json:"city"`
    Lat         float64 `json:"lat"`
    Lon         float64 `json:"lon"`
}

// Geoloc an IP address, return city + country
func Geolocalize(addr string) (geo GeoIP, err error) {
    resp, err := http.Get("http://ip-api.com/json/" + addr)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    err = json.NewDecoder(resp.Body).Decode(&geo)
    if err != nil {
        fmt.Println(err)
        return
    }
    return geo, nil
}

func main() {
    if len(os.Args)<2 {
        fmt.Printf("use: %s IP\n", os.Args[0])
        return
    }
    geo, err := Geolocalize(os.Args[1])
    if err!=nil {
        fmt.Println(err)
        return
    }
    fmt.Printf(`
       City: %s
    Country: %s (%s)
        Lat: %.2f
        Lon: %.2f

`, geo.City, geo.Country, geo.CountryCode, geo.Lat, geo.Lon)
}

