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

type PromptStatus struct {
	// Whether the prompt has been answered or not.
	IsAnswered bool `json:"is_answered"`
	// Whether the prompt can still be answered.
	IsExpired bool `json:"is_expired"`
}
