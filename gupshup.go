package go_gupshup

import (
	"net/url"
	"net/http"
) 

type Gupshup struct {
	apiURL string
	apiParams map[string]string
}

func EnterpriseInitialize(opts map[string]string) *Gupshup {

	apiUrl := "http://enterprise.smsgupshup.com/GatewayAPI/rest"
	var apiParams map[string]string
	apiParams["v"] =  "1.1"
	apiParams["auth_scheme"] = "PLAIN"

	if val, ok := opts["api_url"]; ok {
		apiUrl = val 
	}

	if val ,ok := opts["userId"]; ok {
		apiParams["userId"] = val
	} 

	if val, ok := opts["password"]; ok {
		apiParams["password"] = val
	} 

	if val, ok := opts["token"]; ok {
		apiParams["auth_scheme"] = "TOKEN"
		apiParams["token"] = val
		delete(apiParams,"password")
	}

	return &Gupshup{apiUrl, apiParams}
}

func callApi(gupshup *Gupshup) {
	var params url.Values
	for k, v := range gupshup.apiParams {
		gupshup.apiParams[k] = v
	}

	for k, v := range gupshup.apiParams {
		params.Add(k,v)
	}

	_, err := http.PostForm(gupshup.apiURL, params)

	if err != nil {

	}
}

func sendMessage(gupshup *Gupshup) string {
	var msg string 
	var number string 
	if val, ok := gupshup.apiParams["msg"]; ok {
		msg = val
	}

	if val, ok := gupshup.apiParams["send_to"]; ok {
		number = val
	}

	if len(number) < 12 {
		return "Phone number is too short"
	}

	if len(number) > 12 {
		return "Phone number is too long"
	}	

	if len(msg) > 724 {
		return "Message should be less than 725 characters long"
	}

	if _, ok := gupshup.apiParams["msg_type"]; !ok {
		gupshup.apiParams["msg_type"] = "TEXT"
	}

	callApi(gupshup)
	return ""
}

func SendFlashMessage(gupshup *Gupshup) {
	gupshup.apiParams["msg_type"] = "FLASH"
	sendMessage(gupshup)
}

func SendTextMessage(gupshup *Gupshup) {
	gupshup.apiParams["msg_type"] = "TEXT"
	sendMessage(gupshup)
}

func SendVCard(gupshup *Gupshup) {
	gupshup.apiParams["msg_type"] = "VCARD"
	sendMessage(gupshup)
}

func SendUnicodeMessage(gupshup *Gupshup) {
	gupshup.apiParams["msg_type"] = "UNICODE_TEXT"
	sendMessage(gupshup)
}

func GroupPost(gupshup *Gupshup) string {

	if _,ok := gupshup.apiParams["group_name"]; !ok {
		return "Invalid group name"
	}

	if _,ok := gupshup.apiParams["msg"]; !ok {
		return "Invalid message" 
	}
 
	if _,ok := gupshup.apiParams["msg_type"]; !ok {
		return "Invalid message type"
	}

	gupshup.apiParams["method"] = "post_group"
	callApi(gupshup)
	return ""
}
