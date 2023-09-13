package main

type Pass struct {
	FormatVersion       int         `json:"formatVersion"`
	PassTypeIdentifier  string      `json:"passTypeIdentifier"`
	SerialNumber        string      `json:"serialNumber"`
	TeamIdentifier      string      `json:"teamIdentifier"`
	WebServiceURL       string      `json:"webServiceURL"`
	AuthenticationToken string      `json:"authenticationToken"`
	RelevantDate        string      `json:"relevantDate"`
	Locations           []Location  `json:"locations"`
	Barcodes            []Barcode   `json:"barcodes"`
	OrganizationName    string      `json:"organizationName"`
	Description         string      `json:"description"`
	ForegroundColor     string      `json:"foregroundColor"`
	BackgroundColor     string      `json:"backgroundColor"`
	LabelColor          string      `json:"labelColor"`
	EventTicket         EventTicket `json:"eventTicket"`
	ExpirationDate      string      `json:"expirationDate"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Barcode struct {
	AltText         string `json:"altText"`
	Message         string `json:"message"`
	Format          string `json:"format"`
	MessageEncoding string `json:"messageEncoding"`
}
type PrimaryFields struct {
	Key           string `json:"key"`
	Label         string `json:"label"`
	Value         string `json:"value"`
	TextAlignment string `json:"textAlignment"`
}
type SecondaryFields struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Value string `json:"value"`
}
type AuxiliaryFields struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Label string `json:"label"`
	Row   int    `json:"row"`
}
type EventTicket struct {
	HeaderFields    []any             `json:"headerFields"`
	PrimaryFields   []PrimaryFields   `json:"primaryFields"`
	SecondaryFields []SecondaryFields `json:"secondaryFields"`
	AuxiliaryFields []AuxiliaryFields `json:"auxiliaryFields"`
	BackFields      []any             `json:"backFields"`
}
