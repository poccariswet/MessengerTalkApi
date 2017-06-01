package talk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const talkApiUrl = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"

var talkApiId = os.Getenv("TALKAPIID")

type TalkJson struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Results []talkApiResult `json:"results"`
}

type talkApiResult struct {
	Perplexity float64 `json:"perplexity"`
	Reply      string  `json:"reply"`
}

func post(url string, params url.Values, out interface{}) error {
	resp, err := http.PostForm(url, params)
	// fmt.Println(resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, out)
	//fmt.Println(out)
	if err != nil {
		return err
	}
	if out == nil {
		return err
	}

	return nil
}

func TalkApi(text string) string {

	params := url.Values{
		"apikey": {talkApiId},
		"query":  {text},
	}
	json := TalkJson{}

	post(talkApiUrl, params, &json)
	if json.Message == "empty reply" {
		fmt.Println("おいおい")
		return "よくわかりません"
	} else {
		return json.Results[0].Reply
	}
}
