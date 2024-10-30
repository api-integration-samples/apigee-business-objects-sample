package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Health struct {
	Status string `json:"status"`
}

type SalesOrdersResponse struct {
	Orders []SalesOrder `json:"orders"`
}

type SalesOrder struct {
	SalesOrder                 string
	SalesOrderType             string
	SalesOrganization          string
	DistributionChannel        string
	OrganizationDivision       string
	SalesGroup                 string
	SalesOffice                string
	SalesDistrict              string
	SoldToParty                string
	CreationDate               string
	CreatedByUser              string
	LastChangeDate             string
	SalesOrderDate             string
	TotalNetAmount             float32
	TransactionCurrency        string
	PricingDate                string
	RequestedDeliveryDate      string
	ShippingCondition          string
	OverallTotalDeliveryStatus string
}

var orderCol []SalesOrder = []SalesOrder{
	{
		SalesOrder:                 "9000000152",
		SalesOrderType:             "OR",
		SalesOrganization:          "US1100",
		DistributionChannel:        "01",
		OrganizationDivision:       "01",
		SalesGroup:                 "",
		SalesOffice:                "",
		SalesDistrict:              "7",
		SoldToParty:                "1003766",
		CreationDate:               "1/2/2021",
		CreatedByUser:              "LARRY",
		LastChangeDate:             "",
		SalesOrderDate:             "12/30/2020",
		TotalNetAmount:             245.83,
		TransactionCurrency:        "USD",
		PricingDate:                "1/2/2021",
		RequestedDeliveryDate:      "3/31/2021",
		OverallTotalDeliveryStatus: "DELIVERED",
	},
	{
		SalesOrder:                 "9000000158",
		SalesOrderType:             "OR",
		SalesOrganization:          "US1100",
		DistributionChannel:        "01",
		OrganizationDivision:       "01",
		SalesGroup:                 "",
		SalesOffice:                "",
		SalesDistrict:              "",
		SoldToParty:                "1003765",
		CreationDate:               "2/11/2021",
		CreatedByUser:              "LARRY",
		LastChangeDate:             "",
		SalesOrderDate:             "2/7/2021",
		TotalNetAmount:             901.36,
		TransactionCurrency:        "USD",
		PricingDate:                "2/11/2021",
		RequestedDeliveryDate:      "7/1/2021",
		OverallTotalDeliveryStatus: "DELIVERED",
	},
	{
		SalesOrder:                 "9000000173",
		SalesOrderType:             "OR",
		SalesOrganization:          "US1100",
		DistributionChannel:        "01",
		OrganizationDivision:       "",
		SalesGroup:                 "",
		SalesOffice:                "",
		SalesDistrict:              "",
		SoldToParty:                "1003765",
		CreationDate:               "12/29/2020",
		CreatedByUser:              "LARRY",
		LastChangeDate:             "",
		SalesOrderDate:             "12/29/2020",
		TotalNetAmount:             1180.86,
		TransactionCurrency:        "USD",
		PricingDate:                "12/29/2020",
		RequestedDeliveryDate:      "4/16/2021",
		OverallTotalDeliveryStatus: "DELIVERED",
	},
	{
		SalesOrder:            "9000000253",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1003764",
		CreationDate:          "1/20/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "1/19/2021",
		TotalNetAmount:        175.75,
		TransactionCurrency:   "USD",
		PricingDate:           "1/20/2021",
		RequestedDeliveryDate: "4/21/2021",
	},
	{
		SalesOrder:            "9000000348",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1003764",
		CreationDate:          "3/8/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "3/7/2021",
		TotalNetAmount:        1576.55,
		TransactionCurrency:   "USD",
		PricingDate:           "3/8/2021",
		RequestedDeliveryDate: "8/8/2021",
	},
	{
		SalesOrder:            "9000000363",
		SalesOrderType:        "ZSMP",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "2",
		SoldToParty:           "1003767",
		CreationDate:          "2/23/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "2/23/2021",
		TotalNetAmount:        1112.74,
		TransactionCurrency:   "USD",
		PricingDate:           "2/23/2021",
		RequestedDeliveryDate: "5/7/2021",
	},
	{
		SalesOrder:            "9000000364",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "2",
		SoldToParty:           "1003767",
		CreationDate:          "1/6/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "1/5/2021",
		TotalNetAmount:        1051.28,
		TransactionCurrency:   "USD",
		PricingDate:           "1/6/2021",
		RequestedDeliveryDate: "7/3/2021",
	},
	{
		SalesOrder:            "9000000043",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "7",
		SoldToParty:           "1003768",
		CreationDate:          "12/17/2020",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "12/17/2020",
		TotalNetAmount:        2260.6,
		TransactionCurrency:   "USD",
		PricingDate:           "12/17/2020",
		RequestedDeliveryDate: "2/14/2021",
	},
	{
		SalesOrder:            "9000000232",
		SalesOrderType:        "ZOR",
		SalesOrganization:     "2000",
		DistributionChannel:   "01",
		OrganizationDivision:  "",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1003768",
		CreationDate:          "3/18/2021",
		CreatedByUser:         "SERGEY",
		LastChangeDate:        "",
		SalesOrderDate:        "3/15/2021",
		TotalNetAmount:        1356.08,
		TransactionCurrency:   "EUR",
		PricingDate:           "3/18/2021",
		RequestedDeliveryDate: "6/10/2021",
	},
	{
		SalesOrder:            "9000000237",
		SalesOrderType:        "ZOR",
		SalesOrganization:     "2000",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1003768",
		CreationDate:          "2/8/2021",
		CreatedByUser:         "SERGEY",
		LastChangeDate:        "",
		SalesOrderDate:        "2/5/2021",
		TotalNetAmount:        81.86,
		TransactionCurrency:   "EUR",
		PricingDate:           "2/8/2021",
		RequestedDeliveryDate: "6/15/2021",
	},
}

func SetResponseDefaults(o *SalesOrdersResponse) {
	for i := 0; i < len(o.Orders); i++ {
		SetOrderDefaults(&o.Orders[i], i)
	}
}

func SetSingleResponseDefaults(o *SalesOrder) {
	SetOrderDefaults(o, 0)
}

func SetOrderDefaults(o *SalesOrder, i int) {

	if o.SalesOrderType == "" {
		o.SalesOrderType = "OR"
	}
	if o.SalesOrganization == "" {
		o.SalesOrganization = "1710"
	}
	if o.DistributionChannel == "" {
		o.DistributionChannel = "10"
	}

	if o.SalesOrderDate == "" {
		newDate := time.Now()
		o.SalesOrderDate = newDate.Format("1/2/2006")
		o.CreationDate = newDate.Format("1/2/2006")
	}

	if o.RequestedDeliveryDate == "" {
		tempDate, _ := time.Parse("1/2/2006", o.SalesOrderDate)
		o.RequestedDeliveryDate = tempDate.Add(time.Hour * 24 * 4).Format("1/2/2006")
	}

	if o.CreatedByUser == "" {
		o.CreatedByUser = "WHO_KNOWS"
	}

	// if o.TotalNetAmount == "" {
	// 	o.TotalNetAmount = 0.00
	// }

	if i == 0 && o.OverallTotalDeliveryStatus == "" {
		o.OverallTotalDeliveryStatus = "SCHEDULED"
	} else if o.OverallTotalDeliveryStatus == "" {
		o.OverallTotalDeliveryStatus = "FULFILLMENT"
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {

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
			resp := SalesOrder{}

			for i := 0; i < len(orderCol); i++ {
				if orderCol[i].SalesOrder == filterValue {
					resp = orderCol[i]
					SetSingleResponseDefaults(&resp)
					j, _ := json.Marshal(resp)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusCreated)
					w.Write(j)

					break
				}
			}
		} else {
			resp := SalesOrdersResponse{Orders: orderCol}
			SetResponseDefaults(&resp)
			j, _ := json.Marshal(resp)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write(j)
		}
	case "POST":
		d := json.NewDecoder(r.Body)
		p := &SalesOrder{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if p.SalesOrderDate == "" {
			newDate := time.Now()
			p.SalesOrderDate = newDate.Format("1/2/2006")
			p.CreationDate = newDate.Format("1/2/2006")
		}
		if p.OverallTotalDeliveryStatus == "" {
			p.OverallTotalDeliveryStatus = "SCHEDULED"
		}
		orderCol = append(orderCol, *p)

		SetSingleResponseDefaults(p)

		// Do new sort
		sort.Slice(orderCol, func(i, j int) bool {
			myDate1, _ := time.Parse("1/2/2006", orderCol[i].SalesOrderDate)
			myDate2, _ := time.Parse("1/2/2006", orderCol[j].SalesOrderDate)
			return myDate1.After(myDate2)
		})

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

func healthHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		resp := Health{Status: "Ok"}
		j, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}

func main() {

	// Do initial sorting of orders
	sort.Slice(orderCol, func(i, j int) bool {
		myDate1, _ := time.Parse("1/2/2006", orderCol[i].SalesOrderDate)
		myDate2, _ := time.Parse("1/2/2006", orderCol[j].SalesOrderDate)
		return myDate1.After(myDate2)
	})

	http.HandleFunc("/orders", orderHandler)
	http.HandleFunc("/logistics", locationHandler)
	http.HandleFunc("/customers", customerHandler)
	http.HandleFunc("/", healthHandler)

	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}
