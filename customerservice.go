package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CustomerResponse struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	BusinessPartner         string
	Customer                string
	Supplier                string
	DistributionChannel     string
	OrganizationDivision    string
	BusinessPartnerCategory string
	BusinessPartnerFullName string
}

var customerCol []Customer = []Customer{
	{
		BusinessPartner:         "1003766",
		Customer:                "OR",
		Supplier:                "US1100",
		DistributionChannel:     "01",
		OrganizationDivision:    "01",
		BusinessPartnerCategory: "",
		BusinessPartnerFullName: "",
	},
	{
		BusinessPartner:         "1003765",
		Customer:                "OR",
		Supplier:                "US1100",
		DistributionChannel:     "01",
		OrganizationDivision:    "01",
		BusinessPartnerCategory: "",
		BusinessPartnerFullName: "",
	},
	{
		BusinessPartner:         "1003764",
		Customer:                "OR",
		Supplier:                "US1100",
		DistributionChannel:     "01",
		OrganizationDivision:    "01",
		BusinessPartnerCategory: "",
		BusinessPartnerFullName: "",
	},
	{
		BusinessPartner:         "1003767",
		Customer:                "OR",
		Supplier:                "US1100",
		DistributionChannel:     "01",
		OrganizationDivision:    "01",
		BusinessPartnerCategory: "",
		BusinessPartnerFullName: "",
	},
	{
		BusinessPartner:         "1003768",
		Customer:                "OR",
		Supplier:                "US1100",
		DistributionChannel:     "01",
		OrganizationDivision:    "01",
		BusinessPartnerCategory: "",
		BusinessPartnerFullName: "",
	},
}

func customerHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		filterProp := "none"
		filterValue := "none"

		for k, v := range r.URL.Query() {
			if k == "$filter" {
				log.Println("request filter: " + v[0])
				filterProp = v[0][0:strings.Index(v[0], "eq")]
				log.Println("request filter prop: " + filterProp)
				filterValue = v[0][strings.Index(v[0], "'")+1 : strings.LastIndex(v[0], "'")]
				log.Println("request filter prop value: " + filterValue)
			}
		}

		if filterProp != "none" {
			var resp = Customer{}

			for i := 0; i < len(customerCol); i++ {
				if customerCol[i].BusinessPartner == filterValue {
					resp = customerCol[i]
					j, _ := json.Marshal(resp)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusCreated)
					w.Write(j)

					break
				}
			}
		} else {
			resp := CustomerResponse{Customers: customerCol}
			j, _ := json.Marshal(resp)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write(j)
		}
	case "POST":
		d := json.NewDecoder(r.Body)
		p := &Customer{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		customerCol = append(customerCol, *p)

		j, _ := json.Marshal(p)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
