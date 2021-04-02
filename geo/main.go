package main

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

var (
	defaultGeoDB *geoip2.Reader
)

func Initial(dbFile string) error {
	db, err := geoip2.Open(dbFile)
	if err != nil {
		return err
	}

	defaultGeoDB = db
	return nil
}


func main() {
	Initial("geo/GeoLite2-Region.mmdb")

	netIP := net.ParseIP("156.249.25.195")
	record, err := defaultGeoDB.Country(netIP)
	if err != nil {
		log.Println(err)
	}

	log.Println(record)
}