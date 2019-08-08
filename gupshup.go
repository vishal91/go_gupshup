package go_gupshup

import (
	"net/url"
	"net/http"
) 

var apiUrl string
var apiParams map[string]string

func EnterpriseInitialize(opts map[string]string) {

	apiUrl = "http://enterprise.smsgupshup.com/GatewayAPI/rest"
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
}

func callApi(opts map[string]string) {
	var params url.Values
	for k, v := range opts {
		apiParams[k] = v
	}

	for k, v := range apiParams {
		params.Add(k,v)
	}

	_, err := http.PostForm(apiUrl, params)

	if err != nil {

	}
}

func sendMessage(opts map[string]string) string {
	var msg string 
	var number string 
	if val, ok := opts["msg"]; ok {
		msg = val
	}

	if val, ok := opts["send_to"]; ok {
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

	if _, ok := opts["msg_type"]; !ok {
		opts["msg_type"] = "TEXT"
	}

	callApi(opts)
	return ""
}

func SendFlashMessage(opts map[string]string) {
	opts["msg_type"] = "FLASH"
	sendMessage(opts)
}

func SendTextMessage(opts map[string]string) {
	opts["msg_type"] = "TEXT"
	sendMessage(opts)
}

func SendVCard(opts map[string]string) {
	opts["msg_type"] = "VCARD"
	sendMessage(opts)
}

func SendUnicodeMessage(opts map[string]string) {
	opts["msg_type"] = "UNICODE_TEXT"
	sendMessage(opts)
}

func GroupPost(opts map[string]string) string {

	if _,ok := opts["group_name"]; !ok {
		return "Invalid group name"
	}

	if _,ok := opts["msg"]; !ok {
		return "Invalid message" 
	}

	if _,ok := opts["msg_type"]; !ok {
		return "Invalid message type"
	}

	opts["method"] = "post_group"
	callApi(opts)
	return ""
}
