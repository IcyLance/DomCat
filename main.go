package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/joho/godotenv"
)

type Info struct {
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
		Code   int    `json:"code"`
		Detail string `json:"detail"`
		Body   []Info `json:"body"`
	} `json:"reply"`
}

type Domain struct {
	Categories []cloudflare.ContentCategories
	Details    Info
}

func handle429(num int) {
	num = num % 3
	backoff := time.Duration(1<<num) * time.Second
	jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
	waitTime := backoff + jitter
	time.Sleep(waitTime)
}

func NsList() []Info {
	NS_API_KEY := os.Getenv("NS_API_KEY")
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	page_num := 0
	handle_times := 0
	var domains []Info

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
			handle_times++
			handle429(handle_times)
			page_num--
			continue
		}

		var data NSResp
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&data)
		if err != nil {
			log.Fatalf("Error decoding response: %v\n Response: %v", err, resp.Status)
		}

		if len(data.Reply.Body) == 0 {
			break
		}

		domains = append(domains, data.Reply.Body...)

		time.Sleep(1 * time.Second)
	}

	fmt.Println(len(domains), page_num)
	return domains
}

func CheckCat(domains []string) []cloudflare.DomainDetails {
	// create instance of API using token
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatalf("Error making cloudflare API object: %v", err)
	}

	ctx := context.Background()

	// Get account ID for later call
	var paramsA cloudflare.AccountsListParams
	accounts, _, err := api.Accounts(ctx, paramsA)
	if err != nil {
		log.Fatalf("Error getting accounts: %v", err)
	}

	// initialize parameters for api function call
	var paramsD cloudflare.GetBulkDomainDetailsParameters
	paramsD.Domains = domains
	paramsD.AccountID = accounts[0].ID

	// call to cloudflare
	// returns more data than just categorization
	info, err := api.IntelligenceBulkDomainDetails(ctx, paramsD)
	if err != nil {
		log.Fatalf("Error getting domain details: %v", err)
	}

	return info
}

func main() {
	// Loading keys from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get list of expired domains from Namesilo
	fmt.Println("Getting domains from Namesilo")
	domainList := NsList()
	if len(domainList) == 0 {
		log.Fatal("Error: No domains found")
	}

	//compiling all info into 1 structure
	// will hold domain info and categorization for easy access
	var domains = make([]Domain, len(domainList))
	for i, k := range domainList {
		domains[i].Details = k
	}

	//get categorization for domains
	//place them in the struct with corresponding domain
	// for i, k := range domains {
	// 	cat := CheckCat(k.Details.Domain)
	// 	domains[i].Categories = cat
	// }

	// for _, val := range cat {
	// 	fmt.Println(val)
	// 	time.Sleep(5 * time.Second)
	// }

	// var input string
	// fmt.Scanln(&input)
	// cat := CheckCat(input)
	// fmt.Println(cat)
}
