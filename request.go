package mailchimp

type MailchimpRequest struct {
	Method string
	Path   string
	Body   interface{}
}
