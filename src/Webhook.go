package main

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	//"net/http"
	//"strings"
	// "log"
	//"bytes"
	"io/ioutil"
	"net/url"
)

func main() {
	c := hipchat.NewHttpClient{AuthToken: "3e3ca2b2bcee915cd46c3cd133727f"}
	uri := "api.hipchat.com/v2/room/stoopkids/webhook"

	payload := url.Values{"url": {"http://synchrotronics.net"}, "event": {"room_message"}, "name": {"drd"}}

	req, neg := http.NewRequest("POSTFORM", uri, payload)
	if neg != nil {
		fmt.Println("Error : ", neg)
	}
	req.Header.Set("Authorization", "Bearer 3e3ca2b2bcee915cd46c3cd133727f")
	//req.Header.Set("Content-Type", "application/json")
	resp, neg := client.Do(req)
	defer resp.Body.Close()
	body, err0 := ioutil.ReadAll(resp.Body)
	if err0 != nil {
		fmt.Println("Error : ", err0)
	}
	var web *webhook
	if err1 := json.Unmarshal(body, &web); err1 != nil {
		fmt.Println("Error : ", err1)
	}

}

type webhook struct {
	id    string
	name  string
	links struct {
		self string
	}
}
