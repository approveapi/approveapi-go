/*
 * ApproveAPISwagger
 *
 * The simple API to request a user's approval on anything via email + sms.
 *
 * API version: 1.0.0
 * Contact: dev@approveapi.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package approveapi

type CreatePromptRequest struct {
	// The body of the approval request to show the user.
	Body string `json:"body"`
	// The title of an approval request. Defaults to an empty string.
	Title *string `json:"title,omitempty"`
	// The reject action text. Defaults to 'Reject'.
	RejectText *string `json:"reject_text,omitempty"`
	// The number of seconds until this request can no longer be answered.
	ExpiresIn *float32 `json:"expires_in,omitempty"`
	// If true, the request waits (long-polls) until the user responds to the prompt or more than 10 minutes pass. Defaults to false.
	LongPoll *bool `json:"long_poll,omitempty"`
	// The user to send the approval request to. Can be either an email address or a phone number.
	User string `json:"user"`
	// The approve action text. Defaults to 'Approve'.
	ApproveText *string `json:"approve_text,omitempty"`
	Metadata *PromptMetadata `json:"metadata,omitempty"`
}
