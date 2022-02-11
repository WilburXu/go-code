package main

import (
	"github.com/mholt/archiver/v3"
	"log"
)

func main() {

	err := archiver.Unarchive("./e3f245acb83f036f9b5aae65d9f7d873.tar.gz", "./")
	if err != nil {
		log.Printf("%+v", err)
	}

}
