//////////////////////////////////////////////////////////////////////
// google.go
//////////////////////////////////////////////////////////////////////
package google

import (
)


const (
)


type Location struct {
    LocationCode int
    LocationName string
    LocationCodeParent *int
    CountryIsoCode string
    LocationType string
}


//////////////////////////////////////////////////////////////////////
// Get location for Australia.
//////////////////////////////////////////////////////////////////////
func GetLocationAustralia() *Location {
    return &Location{
        LocationCode: 2036,
        LocationName: "Australia",
        LocationCodeParent: nil,
        CountryIsoCode: "AU",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for Canada.
//////////////////////////////////////////////////////////////////////
func GetLocationCanada() *Location {
    return &Location{
        LocationCode: 2124,
        LocationName: "Canada",
        LocationCodeParent: nil,
        CountryIsoCode: "CA",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for Germany.
//////////////////////////////////////////////////////////////////////
func GetLocationGermany() *Location {
    return &Location{
        LocationCode: 2276,
        LocationName: "Germany",
        LocationCodeParent: nil,
        CountryIsoCode: "DE",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for India.
//////////////////////////////////////////////////////////////////////
func GetLocationIndia() *Location {
    return &Location{
        LocationCode: 2356,
        LocationName: "India",
        LocationCodeParent: nil,
        CountryIsoCode: "IN",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for Ireland.
//////////////////////////////////////////////////////////////////////
func GetLocationIreland() *Location {
    return &Location{
        LocationCode: 2372,
        LocationName: "Ireland",
        LocationCodeParent: nil,
        CountryIsoCode: "IE",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for Italy.
//////////////////////////////////////////////////////////////////////
func GetLocationItaly() *Location {
    return &Location{
        LocationCode: 2380,
        LocationName: "Italy",
        LocationCodeParent: nil,
        CountryIsoCode: "IT",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for Japan.
//////////////////////////////////////////////////////////////////////
func GetLocationJapan() *Location {
    return &Location{
        LocationCode: 2392,
        LocationName: "Japan",
        LocationCodeParent: nil,
        CountryIsoCode: "JP",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for New Zealand.
//////////////////////////////////////////////////////////////////////
func GetLocationNewZealand() *Location {
    return &Location{
        LocationCode: 2554,
        LocationName: "New Zealand",
        LocationCodeParent: nil,
        CountryIsoCode: "NZ",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for Singapore.
//////////////////////////////////////////////////////////////////////
func GetLocationSingapore() *Location {
    return &Location{
        LocationCode: 2702,
        LocationName: "Singapore",
        LocationCodeParent: nil,
        CountryIsoCode: "SG",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for South Africa.
//////////////////////////////////////////////////////////////////////
func GetLocationSouthAfrica() *Location {
    return &Location{
        LocationCode: 2710,
        LocationName: "South Africa",
        LocationCodeParent: nil,
        CountryIsoCode: "ZA",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for South Korea.
//////////////////////////////////////////////////////////////////////
func GetLocationSouthKorea() *Location {
    return &Location{
        LocationCode: 2410,
        LocationName: "South Korea",
        LocationCodeParent: nil,
        CountryIsoCode: "KR",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for Spain.
//////////////////////////////////////////////////////////////////////
func GetLocationSpain() *Location {
    return &Location{
        LocationCode: 2724,
        LocationName: "Spain",
        LocationCodeParent: nil,
        CountryIsoCode: "ES",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for United Kingdom.
//////////////////////////////////////////////////////////////////////
func GetLocationUnitedKingdom() *Location {
    return &Location{
        LocationCode: 2826,
        LocationName: "United Kingdom",
        LocationCodeParent: nil,
        CountryIsoCode: "GB",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location for United States.
//////////////////////////////////////////////////////////////////////
func GetLocationUnitedStates() *Location {
    return &Location{
        LocationCode: 2840,
        LocationName: "United States",
        LocationCodeParent: nil,
        CountryIsoCode: "US",
        LocationType: "Country",
    }
}


//////////////////////////////////////////////////////////////////////
// Get location code for Australia.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeAustralia() int {
    if o := GetLocationAustralia(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for Canada.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeCanada() int {
    if o := GetLocationCanada(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for Germany.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeGermany() int {
    if o := GetLocationGermany(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for India.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeIndia() int {
    if o := GetLocationIndia(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for Ireland.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeIreland() int {
    if o := GetLocationIreland(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for Italy.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeItaly() int {
    if o := GetLocationItaly(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for Japan.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeJapan() int {
    if o := GetLocationJapan(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for New Zealand.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeNewZealand() int {
    if o := GetLocationNewZealand(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for Singapore.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeSingapore() int {
    if o := GetLocationSingapore(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for South Africa.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeSouthAfrica() int {
    if o := GetLocationSouthAfrica(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for South Korea.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeSouthKorea() int {
    if o := GetLocationSouthKorea(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for Spain.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeSpain() int {
    if o := GetLocationSpain(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for United Kingdom.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeUnitedKingdom() int {
    if o := GetLocationUnitedKingdom(); o != nil {
        return o.LocationCode
    }
    return 0
}


//////////////////////////////////////////////////////////////////////
// Get location code for United States.
//////////////////////////////////////////////////////////////////////
func GetLocationCodeUnitedStates() int {
    if o := GetLocationUnitedStates(); o != nil {
        return o.LocationCode
    }
    return 0
}
