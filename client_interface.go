package mailchimp

import (
	"net/url"
	"net/http"
)

// ClientInterface defines exported methods
type ClientInterface interface {
	// Exported methods
	CheckSubscription(listID string, email string) (*MemberResponse, error)
	Subscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error)
	UpdateSubscription(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error)
	RemoveSubscription(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error)
	SetBaseURL(baseURL *url.URL)
	GetBaseURL() *url.URL
	CreateMergeField(listID, tag, name, fieldType string, required bool) (error)
	GetUpdateSubscriptionRequest(listID, email string, mergeFields map[string]interface{}) (MailchimpRequest)
	ExecuteBatchOperation(b BatchOperation) (*http.Response, error)
}
