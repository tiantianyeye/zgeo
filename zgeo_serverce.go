//created by lijinlong use for get country and region info
package zgeo

import (
	"fmt"
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

var (
	loadErr = "WARNING:LoadGetLiteCityDB Path:%s, Err:%s!!!\nMESSAGE:加载Geo数据库错误,使用本地geo需要使用相应的基础镜像getsimple, geo数据库在相应镜像的/app目录下,详细信息参见zgeo README.md\n"
	getErr  = "WARNING:GetCityInfoByGeoIp Err:%s!!! \nMessage:加载Geo数据库错误, 使用本地geo需要使用相应的基础镜像getsimple, geo数据库在相应镜像的/app目录下, 详细信息参见zgeo README.md\n"
)

func LoadGetLiteCityDB(p string) error {
	rd, err := maxminddb.Open(p)
	if err != nil {
		return fmt.Errorf(loadErr, p, err.Error())
	}
	geoDb = rd
	return nil
}

func GetCityInfoByGeoIp(geoIp string) (country string, region string, err error) {
	if geoDb == nil {
		err = fmt.Errorf(getErr, "not init")
		return
	}
	ip := net.ParseIP(geoIp)
	var cityInfo fullCity

	lookErr := geoDb.Lookup(ip, &cityInfo)
	if lookErr != nil {
		err = fmt.Errorf(getErr, lookErr.Error())
		return
	}

	if len(cityInfo.Subdivisions) > 0 {
		region = cityInfo.Subdivisions[0].IsoCode
	}
	country = cityInfo.Country.IsoCode
	return
}
