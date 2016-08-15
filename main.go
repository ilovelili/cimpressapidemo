package main

import (
	"bufio"
	"fmt"
	"os"

	Auth "CimpressApiSampleApp/Auth"
	DeliveryOptions "CimpressApiSampleApp/DeliveryOptions"
	Document "CimpressApiSampleApp/Document"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func continueStep() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	// get authed
	fmt.Println("Step1: Get Authed")
	tokenreq := Auth.Request{ClientID: "4GtkxJhz0U1bdggHMdaySAy05IV4MEDV", UserName: "route666@live.cn", Password: "Aa7059970599", Connection: "default", Scope: "openid email app_metadata"}
	tokenres, tokenerr := Auth.DoAuth(tokenreq)
	panicOnError(tokenerr)
	continueStep()

	// create document
	fmt.Println("Step2: Create Document")

	const imageURL = "https://www.google.co.jp/logos/doodles/2016/2016-doodle-fruit-games-day-11-5698592858701824-scta.png"
	// fmt.Scanln(&imageURL)
	documentreq := Document.CreateRequest{Images: []Document.Image{Document.Image{ImageURL: imageURL, MultipagePdf: false}}, Sku: "VIP-45696"}
	createdocumentres, createdocumenterror := Document.CreateDocument(documentreq, tokenres)
	panicOnError(createdocumenterror)
	fmt.Println("InstructionSourceURL: ", createdocumentres.InstructionSourceURL)
	continueStep()

	fmt.Println("Step3: Preview Document")
	// fmt.Scanln(&imageURL)
	documentpreviewreq := Document.PreviewRequest{Sku: "VIP-45696", InstructionSourceURL: createdocumentres.InstructionSourceURL, Width: "640"}
	previewdocumentres, previewdocumenterror := Document.PreviewDocument(documentpreviewreq, tokenres)
	panicOnError(previewdocumenterror)
	fmt.Println("PreviewUrls: ", previewdocumentres.PreviewUrls)
	continueStep()

	fmt.Println("Step4: Get Delivery Options")
	deliveryoptionsreq := DeliveryOptions.GetDeliveryOptionRequest{Items: []DeliveryOptions.Item{DeliveryOptions.Item{Quantity: 100, Sku: "VIP-45696"}}, DestinationAddress: DeliveryOptions.DestinationAddress{AddressLine1: "Harumi Triton Tower F25", AddressLine2: "Harumi 1-1-1 chuo-ku", City: "Tokyo", CountryCode: "us", County: "us", StateOrRegion: "tokyo", PostalCode: "200101"}}
	deliveryoptionsres, deliveryoptionserr := DeliveryOptions.GetDeliveryOptions(deliveryoptionsreq, tokenres)
	panicOnError(deliveryoptionserr)
	fmt.Println("DeliveryOptions: ", deliveryoptionsres.DeliveryOptions)
	continueStep()
}
