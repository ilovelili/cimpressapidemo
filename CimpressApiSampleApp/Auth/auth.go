package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const tokenAPI = "https://cimpress.auth0.com/oauth/ro"

// Request Request
type Request struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
	ClientID   string `json:"client_id"`
	Connection string `json:"connection"`
	Scope      string `json:"scope"`
}

// Response TokenResponse
type Response struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
}

// DoAuth do auth and get the token response
func DoAuth(request Request) (res Response, err error) {
	params, _ := json.Marshal(request)
	paramstr := string(params)
	req, _ := http.NewRequest("POST", tokenAPI, strings.NewReader(paramstr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

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
