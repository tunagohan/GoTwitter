package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"bufio"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

type ApiConf struct {
	ConsumerKey       string `json:"consumer_key"`
	ConsumerSecret    string `json:"consumer_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {
	var apiConf ApiConf
	{
		apiConfPath := flag.String("conf", "config.json", "API Config File")
		flag.Parse()
		data, err_file := ioutil.ReadFile(*apiConfPath)
		check(err_file)
		err_json := json.Unmarshal(data, &apiConf)
		check(err_json)
	}

	anaconda.SetConsumerKey(apiConf.ConsumerKey)
	anaconda.SetConsumerSecret(apiConf.ConsumerSecret)
	api := anaconda.NewTwitterApi(apiConf.AccessToken, apiConf.AccessTokenSecret)

	text := StrStdin()
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("送信完了: ", tweet.Text)

}
