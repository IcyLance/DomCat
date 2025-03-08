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
			fmt.Println("Getting domains from Namesilo...")
			page_num := 1
			domList, err := NsList(page_num)
			if err != nil {
				log.Fatalf("Error getting domains: %v", err)
			}
			fmt.Println("Got domains from Namesilo")

			// compiling all info into 1 structure
			// will hold domain info and categorization for easy access
			var domains = make([]Domain, len(domList))
			var domSL = make([]string, len(domList))
			for i, k := range domList {
				domains[i].Details = k
				if i < 25 {
					domSL[i] = k.Domain
				}
			}

			// get categorization for domains
			// place them in the struct with corresponding domain
			fmt.Println("Getting categorization...")
			cat, err := CheckCatBulk(domSL)
			if err != nil {
				log.Fatalf("Error getting categorization: %v", err)
			} else {
				fmt.Println("Finished getting categorization")
			}

			for _, k := range cat {
				fmt.Print(k.Domain, ": ")
				fmt.Println(k.ContentCategories)
			}

			// collect categorization into main variable for storage
			for _, j := range cat {
				for i, k := range domains {
					if j.Domain == k.Details.Domain {
						domains[i].Categories = j.ContentCategories
					}
				}
			}

			fmt.Println("==============================")

			for _, k := range domains {
				fmt.Print(k.Details.Domain, ": ")
				fmt.Println(k.Categories)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
