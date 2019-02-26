/*
 * ApproveAPISwagger
 *
 * The simple API to request a user's approval on anything via email + sms.
 *
 * API version: 1.0.1
 * Contact: dev@approveapi.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package approveapi

type PromptAnswer struct {
	// The unix timestamp when the user answered the prompt.
	Time float32 `json:"time"`
	// The user's answer to whether or not they approve this prompt.
	Result bool `json:"result"`
	Metadata *AnswerMetadata `json:"metadata,omitempty"`
}
