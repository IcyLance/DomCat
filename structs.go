package main

import "github.com/cloudflare/cloudflare-go"

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
	Details    Info
	Categories []cloudflare.ContentCategories
}
