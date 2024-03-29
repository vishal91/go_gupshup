package go_gupshup

import (
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
) 

type Gupshup struct {
	apiURL string
	apiParams map[string]string
}

func EnterpriseInitialize(opts map[string]string) *Gupshup {
	apiUrl := "http://enterprise.smsgupshup.com/GatewayAPI/rest"
	opts["v"] =  "1.1"
	opts["auth_scheme"] = "PLAIN"
	opts["method"] = "sendMessage"
	if val, ok := opts["api_url"]; ok {
		apiUrl = val 
	}

	if val ,ok := opts["userid"]; ok {
		opts["userid"] = val
	} 

	if val, ok := opts["password"]; ok {
		opts["password"] = val
	} 

	if val, ok := opts["token"]; ok {
		opts["auth_scheme"] = "TOKEN"
		opts["token"] = val
		delete(opts,"password")
	}

	return &Gupshup{apiUrl, opts}
}

func callApi(gupshup *Gupshup) error {
	params := url.Values{}

	for k, v := range gupshup.apiParams {
		params.Add(k,v)
	}
	r, err := http.PostForm(gupshup.apiURL, params)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	log.Print("Response for gupshup",string(body))
	return err
}

func (gupshup *Gupshup) sendMessage() (error, string) {
	var msg string 
	var number string 
	if val, ok := gupshup.apiParams["msg"]; ok {
		msg = val
	}

	if val, ok := gupshup.apiParams["send_to"]; ok {
		number = val
	}

	if len(number) < 12 {
		return nil,"Phone number is too short"
	}

	if len(number) > 12 {
		return nil,"Phone number is too long"
	}	

	if len(msg) > 724 {
		return nil,"Message should be less than 725 characters long"
	}

	if _, ok := gupshup.apiParams["msg_type"]; !ok {
		gupshup.apiParams["msg_type"] = "TEXT"
	}

	err := callApi(gupshup)
	return err, "Gupshup api call failing"
}

func (gupshup *Gupshup) SendFlashMessage() (error, string) {
	gupshup.apiParams["msg_type"] = "FLASH"
	return gupshup.sendMessage()
}

func (gupshup *Gupshup) SendTextMessage() (error, string) {
	gupshup.apiParams["msg_type"] = "TEXT"
	return gupshup.sendMessage()
}

func (gupshup *Gupshup) SendVCard() (error, string) {
	gupshup.apiParams["msg_type"] = "VCARD"
	return gupshup.sendMessage()
}

func (gupshup *Gupshup) SendUnicodeMessage() (error, string) {
	gupshup.apiParams["msg_type"] = "UNICODE_TEXT"
	return gupshup.sendMessage()
}

func (gupshup *Gupshup) GroupPost() (error, string) {

	if _,ok := gupshup.apiParams["group_name"]; !ok {
		return nil, "Invalid group name"
	}

	if _,ok := gupshup.apiParams["msg"]; !ok {
		return nil, "Invalid message" 
	}
 
	if _,ok := gupshup.apiParams["msg_type"]; !ok {
		return nil, "Invalid message type"
	}

	gupshup.apiParams["method"] = "post_group"
	err := callApi(gupshup)
	return err,""
}
