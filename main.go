package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "domCat",
		Usage: "Placeholder",
		Action: func(*cli.Context) error {
			// Loading keys from .env
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}

			// // get list of expired domains from Namesilo
			page_num := 1
			domList, err := NsList(page_num)
			if err != nil {
				log.Fatalf("Error getting domains: %v", err)
			}

			// compiling all info into 1 structure
			// will hold domain info and categorization for easy access
			var domSL = make([]string, len(domList))
			for i, k := range domList {
				domSL[i] = k.Domain
			}

			// get categorization for domains
			// place them in the struct with corresponding domain
			cats, err := CheckCatBulk(domSL)
			if err != nil {
				log.Fatalf("Error getting categorization: %v", err)
			}

			var domains []Domain

			// collect information into main variable for easy access
			for _, cat := range cats {
				for _, dom := range domList {
					if len(cat.Categories) != 0 {
						if cat.Domain == dom.Domain {
							domains = append(domains, Domain{
								Details:    dom,
								Categories: cat.Categories,
							})
						}
					}
				}
			}

			for i, k := range domains {
				fmt.Print(i, ": ", k.Details.Domain, " - ")
				for _, l := range k.Categories {
					fmt.Print(l, ", ")
				}
				fmt.Println()
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
