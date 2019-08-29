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

	if val, ok := opts["api_url"]; ok {
		apiUrl = val 
	}

	if val ,ok := opts["userId"]; ok {
		opts["userId"] = val
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
		gupshup.apiParams[k] = v
	}

	for k, v := range gupshup.apiParams {
		params.Add(k,v)
	}
	r, err := http.PostForm(gupshup.apiURL, params)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	// read the payload, in this case, Jhon's info
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	log.Print("Response for gupshup",body)
	return err
}

func (gupshup *Gupshup) callApi2() error {
	params := url.Values{}
	for k, v := range gupshup.apiParams {
		gupshup.apiParams[k] = v
	}

	for k, v := range gupshup.apiParams {
		params.Add(k,v)
	}

	_, err := http.PostForm(gupshup.apiURL, params)

	return err
}

func sendMessage(gupshup *Gupshup) (error, string) {
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

func SendFlashMessage(gupshup *Gupshup) (error, string) {
	gupshup.apiParams["msg_type"] = "FLASH"
	return sendMessage(gupshup)
}

func SendTextMessage(gupshup *Gupshup) (error, string) {
	gupshup.apiParams["msg_type"] = "TEXT"
	return sendMessage(gupshup)
}

func SendVCard(gupshup *Gupshup) (error, string) {
	gupshup.apiParams["msg_type"] = "VCARD"
	return sendMessage(gupshup)
}

func SendUnicodeMessage(gupshup *Gupshup) (error, string) {
	gupshup.apiParams["msg_type"] = "UNICODE_TEXT"
	return sendMessage(gupshup)
}

func GroupPost(gupshup *Gupshup) (error, string) {

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
