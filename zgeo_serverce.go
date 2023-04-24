//created by lijinlong use for get country and region info
package zgeo

import (
	"net"

	"github.com/oschwald/maxminddb-golang"
)

var geoDb *maxminddb.Reader

type fullCity struct {
	//City struct {
	//	GeoNameID uint              `maxminddb:"geoname_id"`
	//	Names     map[string]string `maxminddb:"names"`
	//} `maxminddb:"city"`
	//Continent struct {
	//	Code      string            `maxminddb:"code"`
	//	GeoNameID uint              `maxminddb:"geoname_id"`
	//	Names     map[string]string `maxminddb:"names"`
	//} `maxminddb:"continent"`
	Country struct {
		GeoNameID         uint              `maxminddb:"geoname_id"`
		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
		IsoCode           string            `maxminddb:"iso_code"`
		Names             map[string]string `maxminddb:"names"`
	} `maxminddb:"country"`
	//Location struct {
	//	AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
	//	Latitude       float64 `maxminddb:"latitude"`
	//	Longitude      float64 `maxminddb:"longitude"`
	//	MetroCode      uint    `maxminddb:"metro_code"`
	//	TimeZone       string  `maxminddb:"time_zone"`
	//} `maxminddb:"location"`
	//Postal struct {
	//	Code string `maxminddb:"code"`
	//} `maxminddb:"postal"`
	//RegisteredCountry struct {
	//	GeoNameID         uint              `maxminddb:"geoname_id"`
	//	IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
	//	IsoCode           string            `maxminddb:"iso_code"`
	//	Names             map[string]string `maxminddb:"names"`
	//} `maxminddb:"registered_country"`
	//RepresentedCountry struct {
	//	GeoNameID         uint              `maxminddb:"geoname_id"`
	//	IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
	//	IsoCode           string            `maxminddb:"iso_code"`
	//	Names             map[string]string `maxminddb:"names"`
	//	Type              string            `maxminddb:"type"`
	//} `maxminddb:"represented_country"`
	Subdivisions []struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		IsoCode   string            `maxminddb:"iso_code"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
	//Traits struct {
	//	IsAnonymousProxy    bool `maxminddb:"is_anonymous_proxy"`
	//	IsSatelliteProvider bool `maxminddb:"is_satellite_provider"`
	//} `maxminddb:"traits"`
}

func LoadGetLiteCityDB(p string) error {
	rd, err := maxminddb.Open(p)
	if err != nil {
		return err
	}
	geoDb = rd
	return nil
}

func GetCityInfoByGeoIp(geoIp string) (country string, region string, err error) {
	ip := net.ParseIP(geoIp)
	var cityInfo fullCity

	err = geoDb.Lookup(ip, &cityInfo)
	if err != nil {
		return
	}

	if len(cityInfo.Subdivisions) > 0 {
		region = cityInfo.Subdivisions[0].IsoCode
	}
	country = cityInfo.Country.IsoCode
	return
}
