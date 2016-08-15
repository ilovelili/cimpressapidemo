package document

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	Auth "CimpressApiSampleApp/Auth"
)

// PreviewRequest Document request
type PreviewRequest struct {
	Sku                  string
	InstructionSourceURL string
	Width                string
}

// PreviewResponse Document response
type PreviewResponse struct {
	PreviewUrls []string `json:"PreviewUrls"`
}

const documentPreviewAPI = "https://api.cimpress.io/sandbox/vcs/printapi/v1/documents/previews"

// PreviewDocument preview document
func PreviewDocument(request PreviewRequest, tokenres Auth.Response) (res PreviewResponse, err error) {
	req, _ := http.NewRequest("GET", documentPreviewAPI, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenres.IDToken)

	q := req.URL.Query()
	q.Add("sku", request.Sku)
	q.Add("instructionSourceURL", request.InstructionSourceURL)
	q.Add("Width", request.Width)
	req.URL.RawQuery = q.Encode()

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
