package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func NsList(page_num int) ([]Details, error) {
	NS_API_KEY := os.Getenv("NS_API_KEY")
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var domains []Details

	client := &http.Client{}

	url := fmt.Sprintf("https://www.namesilo.com/public/api/listAuctions?version=1&type=json&key=%s&page=%d&pageSize=25&buyNow=1", NS_API_KEY, page_num)

	// create request to namesilo
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// set headers in request
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle 429
	if resp.StatusCode == 429 {
		err := errors.New("429 response")
		return nil, err
	}

	// decode response into a useable format
	var data NSResp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	// gather all domins into 1 easy to read variable
	domains = append(domains, data.Reply.Body...)
	if len(domains) == 0 {
		return domains, errors.New("no domains found")
	}
	// time.Sleep(1 * time.Second)
	// }
	return domains, nil
}

func CheckCat(domain string) (CheckCatReturn, error) {
	// create instance of API using token
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return CheckCatReturn{}, err
	}

	//create blank coontext for API calls
	ctx := context.Background()

	// Get account ID for later call
	var paramsA cloudflare.AccountsListParams
	accounts, _, err := api.Accounts(ctx, paramsA)
	if err != nil {
		return CheckCatReturn{}, err
	}

	// initialize parameters for api function call
	var paramsD cloudflare.GetDomainDetailsParameters
	paramsD.Domain = domain
	paramsD.AccountID = accounts[0].ID

	// call to cloudflare
	// returns more data than just categorization
	info, err := api.IntelligenceDomainDetails(ctx, paramsD)
	if err != nil {
		return CheckCatReturn{}, err
	}

	var out CheckCatReturn

	var categories []string
	for _, category := range info.ContentCategories {
		categories = append(categories, category.Name)
	}
	out = CheckCatReturn{
		Domain:     info.Domain,
		Categories: categories,
	}

	return out, nil
}

func CheckCatBulk(domains []string) ([]CheckCatReturn, error) {
	// create instance of API using token
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return nil, err
	}

	//create blank coontext for API calls
	ctx := context.Background()

	// Get account ID for later call
	var paramsA cloudflare.AccountsListParams
	accounts, _, err := api.Accounts(ctx, paramsA)
	if err != nil {
		return nil, err
	}

	// initialize parameters for api function call
	var paramsD cloudflare.GetBulkDomainDetailsParameters
	paramsD.Domains = domains
	paramsD.AccountID = accounts[0].ID

	// call to cloudflare
	// returns more data than just categorization
	info, err := api.IntelligenceBulkDomainDetails(ctx, paramsD)
	if err != nil {
		return nil, err
	}

	var out []CheckCatReturn

	// putting the categorizies in a more accessible form
	// also dont need some of the info so compacting down a bit
	for _, k := range info {
		var categories []string
		for _, category := range k.ContentCategories {
			categories = append(categories, category.Name)
		}
		out = append(out, CheckCatReturn{
			Domain:     k.Domain,
			Categories: categories,
		})
	}

	return out, nil
}
