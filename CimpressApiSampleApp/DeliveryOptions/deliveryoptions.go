package deliveryoptions

import (
	Auth "CimpressApiSampleApp/Auth"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// DestinationAddress Destination Address
type DestinationAddress struct {
	AddressLine2  string `json:"AddressLine2"`
	City          string `json:"City"`
	CountryCode   string `json:"CountryCode"`
	AddressLine1  string `json:"AddressLine1"`
	County        string `json:"County"`
	StateOrRegion string `json:"StateOrRegion"`
	PostalCode    string `json:"PostalCode"`
}

// Item item
type Item struct {
	Sku      string `json:"Sku"`
	Quantity int    `json:"Quantity"`
}

// GetDeliveryOptionRequest Delivery Option request
type GetDeliveryOptionRequest struct {
	DestinationAddress DestinationAddress `json:"DestinationAddress"`
	Items              []Item             `json:"Items"`
}

// Price price
type Price struct {
	Amount   int    `json:"Amount"`
	Currency string `json:"Currency"`
}

// DeliveryOption Delivery Option
type DeliveryOption struct {
	DeliveryArriveByDate string `json:"DeliveryArriveByDate"`
	BusinessDays         int    `json:"BusinessDays"`
	Carrier              string `json:"Carrier"`
	Price                Price  `json:"Price"`
	DeliveryOptionID     string `json:"DeliveryOptionId"`
	CarrierService       string `json:"CarrierService"`
}

// GetDeliveryOptionResponse Delivery Option response
type GetDeliveryOptionResponse struct {
	DeliveryOptions []DeliveryOption `json:"DeliveryOptions"`
}

const deliveryOptionsAPI = "https://api.cimpress.io/sandbox/vcs/printapi/v1/delivery-options"

// GetDeliveryOptions get delivery options
func GetDeliveryOptions(request GetDeliveryOptionRequest, tokenres Auth.Response) (res GetDeliveryOptionResponse, err error) {
	params, _ := json.Marshal(request)
	paramstr := string(params)
	req, _ := http.NewRequest("POST", deliveryOptionsAPI, strings.NewReader(paramstr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenres.IDToken)

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	httpres, httperr := client.Do(req)
	if httperr != nil {
		err = httperr
	} else if httpres.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(httpres.Body)
		json.Unmarshal(bodyBytes, &res)
	}

	return
}
