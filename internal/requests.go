package internal

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const liveEndpoint = "https://metaname.net/api/1.1/"
const testEndpoint = "https://test.metaname.net/api/1.1"

type ApiResponse struct {
	httpCode    int
	messageBody string
}

/**
take metaname payloads and send http requests
handle network errors
bulk operation functions
return api errors
**/

func MetanameRequest(test bool, jsonPayload string) ApiResponse {
	url := ""
	if !test {
		url = liveEndpoint
	} else {
		url = testEndpoint
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonPayload))
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	responseData, err := io.ReadAll(resp.Body)
	responseString := string(responseData)
	return ApiResponse{httpCode: resp.StatusCode, messageBody: responseString}
}
