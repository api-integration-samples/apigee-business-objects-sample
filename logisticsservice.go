package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type SalesOrderLocation struct {
	SalesOrder             string
	Stage                  string
	LocationType           string
	LocationAddress        string
	LocationLat            string
	LocationLong           string
	LocationAdditionalInfo string
}

type SalesOrderLocationResponse struct {
	Locations []SalesOrderLocation `json:"locations"`
}

var locCol []SalesOrderLocation = []SalesOrderLocation{
	{
		SalesOrder:             "9000000152",
		Stage:                  "RECEIVED",
		LocationType:           "SHIPPING",
		LocationAddress:        "Tucholskystr. 2, 10117 Berlin, Germany",
		LocationLat:            "52.520",
		LocationLong:           "13.405",
		LocationAdditionalInfo: "",
	},
}

func locationHandler(w http.ResponseWriter, r *http.Request) {

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
			var resp SalesOrderLocation

			for i := 0; i < len(locCol); i++ {
				if locCol[i].SalesOrder == filterValue {
					resp = locCol[i]

					j, _ := json.Marshal(resp)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusCreated)
					w.Write(j)

					break
				}
			}
		} else {
			resp := SalesOrderLocationResponse{Locations: locCol}

			j, _ := json.Marshal(resp)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write(j)
		}
	case "POST":
		d := json.NewDecoder(r.Body)
		p := &SalesOrderLocation{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		locCol = append(locCol, *p)

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
