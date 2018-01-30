package mailchimp

import (
	"sync"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

type BatchOperation struct {
	Requests []MailchimpRequest
	mutex    *sync.Mutex
}

func NewBatchOperation() BatchOperation {
	return BatchOperation{
		Requests: make([]MailchimpRequest, 0),
		mutex:    &sync.Mutex{},
	}
}

func (b *BatchOperation) AddRequest(req MailchimpRequest) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.Requests = append(b.Requests, req)
}

// Execute BatchOperation
func (c *Client) ExecuteBatchOperation(b BatchOperation) (batchResponse BatchResponse, err error) {

	jsonRequests := make([]mailchimpJSONRequest, 0)
	for _, r := range b.Requests {
		var jsonReq mailchimpJSONRequest
		jsonReq, err = r.ToMailchimpJSONRequest()
		if err != nil {
			break
		}
		jsonRequests = append(jsonRequests, jsonReq)
	}

	if err != nil {
		return
	}

	data := make(map[string]interface{})
	data["operations"] = jsonRequests

	// Do Request
	resp, err := c.do("POST", "/batches", data)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Unmarshal into BatchResponse
	batchResponse = BatchResponse{}
	err = json.Unmarshal(responseBody, &batchResponse)
	if err != nil {
		return
	}

	return batchResponse, nil
}

func (c *Client) GetBatchOperationStatus(b BatchResponse) (batchResponse BatchResponse, err error) {
	return c.GetBatchOperationStatusByID(b.ID)
}

func (c *Client) GetBatchOperationStatusByID(id string) (batchResponse BatchResponse, err error) {
	resp, err := c.do("GET", fmt.Sprintf("/batches/%s", id), nil)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Print(string(body))

	err = json.Unmarshal(body, &batchResponse)
	if err != nil {
		return
	}

	return
}