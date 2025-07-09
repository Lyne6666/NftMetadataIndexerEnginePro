// cmd/nftmetadataindexerenginepro/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"nftmetadataindexerenginepro/internal/nftmetadataindexerenginepro"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := nftmetadataindexerenginepro.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
