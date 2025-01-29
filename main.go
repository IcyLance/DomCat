package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type DomainInfo struct {
	ID               int     `json:"id"`
	LeaderUserID     int     `json:"leaderUserId"`
	OwnerUserID      int     `json:"ownerUserId"`
	DomainID         int     `json:"domainId"`
	Domain           string  `json:"domain"`
	StatusID         int     `json:"statusId"`
	TypeID           int     `json:"typeId"`
	OpeningBid       float64 `json:"openingBid"`
	CurrentBid       float64 `json:"currentBid"`
	MaxBid           float64 `json:"maxBid"`
	HasBids          bool    `json:"hasBids"`
	BidsQuantity     int     `json:"bidsQuantity"`
	BidID            any     `json:"bidId"`
	DomainCreatedOn  string  `json:"domainCreatedOn"`
	AuctionEndsOn    string  `json:"auctionEndsOn"`
	AuctionEndsOnUtc string  `json:"auctionEndsOnUtc"`
	URL              string  `json:"url"`
}

type NSResp struct {
	Request struct {
		Operation string `json:"operation"`
		IP        string `json:"ip"`
	} `json:"request"`
	Reply struct {
		Code   int          `json:"code"`
		Detail string       `json:"detail"`
		Body   []DomainInfo `json:"body"`
	} `json:"reply"`
}

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func nsList() []DomainInfo {
	fmt.Println("Getting domains from Namesilo")

	NS_API_KEY := os.Getenv("NS_API_KEY")
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	page_num := 0
	var domains []DomainInfo

	client := &http.Client{}

	for {
		page_num++

		url := fmt.Sprintf("https://www.namesilo.com/public/api/listAuctions?version=1&type=json&key=%s&page=%d&pageSize=500&buyNow=1", NS_API_KEY, page_num)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 429 {
			time.Sleep(5 * time.Second)
			resp, err = client.Do(req)
			if err != nil {
				log.Fatalf("Error making request: %v", err)
			}
		}

		var data NSResp
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&data)
		if err != nil {
			log.Fatalf("Error decodoing response: %v", err)
		}

		if len(data.Reply.Body) == 0 {
			break
		}

		domains = append(domains, data.Reply.Body...)

		time.Sleep(1 * time.Second)
	}

	fmt.Println(len(domains))
	return domains
}

func main() {
	loadEnvVars()
	nsList()
}
