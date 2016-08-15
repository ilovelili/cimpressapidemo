package main

import (
	"fmt"

	Auth "./CimpressApiSampleApp/Auth"
)

func main() {
	res, err := Auth.DoAuth(Auth.Request{ClientID: "4GtkxJhz0U1bdggHMdaySAy05IV4MEDV", UserName: "route666@live.cn", Password: "Aa7059970599", Connection: "default", Scope: "openid email app_metadata"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
