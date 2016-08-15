package document

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	Auth "CimpressApiSampleApp/Auth"
)

// Image image entity
type Image struct {
	ImageURL     string `json:"ImageUrl"`
	MultipagePdf bool   `json:"MultipagePdf"`
}

// CreateRequest Document request
type CreateRequest struct {
	Images []Image `json:"Images"`
	Sku    string  `json:"Sku"`
}

// CreateResponse Document response
type CreateResponse struct {
	DocumentID           string `json:"DocumentId"`
	InstructionVersion   string `json:"InstructionVersion"`
	InstructionSourceURL string `json:"InstructionSourceUrl"`
}

const documentCreateAPI = "https://api.cimpress.io/sandbox/vcs/printapi/v1/documents/creators/url"

// CreateDocument create document
func CreateDocument(request CreateRequest, tokenres Auth.Response) (res CreateResponse, err error) {
	params, _ := json.Marshal(request)
	paramstr := string(params)
	req, _ := http.NewRequest("POST", documentCreateAPI, strings.NewReader(paramstr))
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
