package mailchimp

import "sync"

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

