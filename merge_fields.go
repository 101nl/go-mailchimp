package mailchimp

import (
	"errors"
	"fmt"
	"io/ioutil"
)

var allowedMergeFieldTypes = []string{
	"text", "number", "address",
	"phone", "date", "url",
	"imageurl", "radio", "dropdown",
	"birthday", "zip",
}

func (c *Client) CreateMergeField(listID, tag, name, fieldType string, required bool) (err error) {
	allowedType := false
	for _, t := range allowedMergeFieldTypes {
		if t == fieldType {
			allowedType = true
			break
		}
	}

	// Return err if fieldType is not allowed
	if !allowedType {
		return errors.New("Given fieldtype is not allowed")
	}

	params := map[string]interface{}{
		"tag": tag,
		"name": name,
		"type": fieldType,
		"required": required,
	}

	response, err := c.do("POST", fmt.Sprintf("/lists/%s/merge-fields", listID), &params)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode/100 == 2 {
		return nil
	} else {
		errorResponse, err := extractError(data)
		if err != nil {
			return err
		}

		return errorResponse
	}
}
