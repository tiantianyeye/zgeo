//created by lijinlong use for city and privacy service
package zgeo

import (
	"github.com/oschwald/maxminddb-golang"
)

var geoDb *maxminddb.Reader

func LoadGetLiteCityDB(p string) error {
	rd, err := maxminddb.Open(p)
	if err != nil {
		return err
	}
	geoDb = rd
	return nil
}
