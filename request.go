package mailchimp

import "encoding/json"

type MailchimpRequest struct {
	Method string      `json:"method"`
	Path   string      `json:"path"`
	Body   interface{} `json:"body"`
}

type mailchimpJSONRequest struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Body   string `json:"body"`
}

func (m MailchimpRequest) ToMailchimpJSONRequest() (jsonReq mailchimpJSONRequest, err error) {
	body, err := json.Marshal(m.Body)
	if err != nil {
		return
	}

	jsonReq.Method = m.Method
	jsonReq.Path = m.Path
	jsonReq.Body = string(body)

	return
}
