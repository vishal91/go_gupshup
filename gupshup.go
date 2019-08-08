package go_gupshup

import (
	"net/url"
	"net/http"
) 

var apiUrl string
var apiParams url.Values

func EnterpriseInitialize(opts map[string]string) {

	apiUrl = "http://enterprise.smsgupshup.com/GatewayAPI/rest"
	apiParams.Add("v", "1.1")
	apiParams.Add("auth_scheme", "PLAIN")

	if val, ok := opts["api_url"]; ok {
		apiUrl = val 
	}

	if val ,ok := opts["userId"]; ok {
		apiParams.Add("userId",val)
	} 

	if val, ok := opts["password"]; ok {
		apiParams.Add("password",val)
	} 

	if val, ok := opts["token"]; ok {
		apiParams.Add("auth_scheme", "TOKEN")
		apiParams.Add("token", val)
		apiParams.Del("password")
	}
}

func callApi(opts url.Values) {
	for k, v := range opts {
		apiParams.Add(k, v[0])
	}

	_, err := http.PostForm(apiUrl, apiParams)

	if err != nil {

	}
}

func sendMessage(opts url.Values) string {
	var msg string 
	var number string 
	if val, ok := opts["msg"]; ok {
		msg = val[0]
	}

	if val, ok := opts["send_to"]; ok {
		number = val[0]
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
		opts.Add("msg_type", "TEXT")
	}

	callApi(opts)
	return ""
}

func SendFlashMessage(opts url.Values) {
	opts.Add("msg_type", "FLASH")
	sendMessage(opts)
}

func SendTextMessage(opts url.Values) {
	opts.Add("msg_type", "TEXT")
	sendMessage(opts)
}

func SendVCard(opts url.Values) {
	opts.Add("msg_type", "VCARD")
	sendMessage(opts)
}

func SendUnicodeMessage(opts url.Values) {
	opts.Add("msg_type", "UNICODE_TEXT")
	sendMessage(opts)
}

func GroupPost(opts url.Values) string {

	if _,ok := opts["group_name"]; !ok {
		return "Invalid group name"
	}

	if _,ok := opts["msg"]; !ok {
		return "Invalid message" 
	}

	if _,ok := opts["msg_type"]; !ok {
		return "Invalid message type"
	}

	opts.Add("method", "post_group")
	callApi(opts)
	return ""
}