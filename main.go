package main

import (
	"ShiftAlt/mockit/configs"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var jsonMockResponse []byte
var xmlMockResponse []byte

func mockAPIWithJSONResponse(ctx *fasthttp.RequestCtx) {
	fmt.Println("mockit json!")
	ctx.Response.Header.Set("Content-Type", "application/json")
	time.Sleep(time.Duration(configs.LatencyForApiWithJSONResponseinMilliSeconds) * time.Millisecond)
	ctx.Write(jsonMockResponse)
}
func mockAPIWithXMLResponse(ctx *fasthttp.RequestCtx) {
	fmt.Println("mockit xml!")
	ctx.Response.Header.Set("Content-Type", "application/xml")
	time.Sleep(time.Duration(configs.LatencyForApiWithXMLResponseinMilliSeconds) * time.Millisecond)
	ctx.Write(xmlMockResponse)
}

func main() {

	jsonMessage, e := ioutil.ReadFile(configs.JSONResponseFileRelativePath)
	if e != nil {
		fmt.Printf("Error reading mock response file: %v\n", e)
		os.Exit(1)
	}
	xmlMessage, e := ioutil.ReadFile(configs.XMLResponseFileRelativePath)
	if e != nil {
		fmt.Printf("Error reading mock response file: %v\n", e)
		os.Exit(1)
	}

	jsonMockResponse = jsonMessage
	xmlMockResponse = xmlMessage

	fmt.Println("starting mock server...")

	router := fasthttprouter.New()
	router.POST(configs.ApiPathWithXMLResponse, mockAPIWithXMLResponse)
	router.POST(configs.ApiPathWithJSONResponse, mockAPIWithJSONResponse)

	log.Fatal(fasthttp.ListenAndServe(":8010", router.Handler))

}
