//go:build !std
// +build !std

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/extism/go-pdk"
)

var API_KEY string = "YOUR_API_KEY"

type input struct {
	Prompt      string `json:"prompt"`
	TargetModel string `json:"targetModel"`
}

type optimizationResponse struct {
	Code   int `json:"code"`
	Status int `json:"status"`
	Result struct {
		PromptOptimized string `json:"promptOptimized"`
	} `json:"result"`
}

//export run
func run() int32 {
	in := input{}
	err := pdk.InputJSON(&in)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	payload := []byte(fmt.Sprintf(`{
		"data": {
			"prompt": "%s",
			"targetModel": "%s"
		}
	}`, in.Prompt, in.TargetModel))

	url := "https://api.promptperfect.jina.ai/optimize"

	req := pdk.NewHTTPRequest(pdk.MethodPost, url)
	req.SetHeader("content-type", "application/json")
	req.SetHeader("x-api-key", "token "+API_KEY)
	req.SetBody(bytes.NewBuffer(payload).Bytes())
	res := req.Send()

	if res.Status() != 200 {
		pdk.SetError(fmt.Errorf("failed to optimize prompt: %v", res.Status()))
		return 1
	}

	var response optimizationResponse
	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	pdk.OutputString(response.Result.PromptOptimized)
	return 0
}

func main() {}
