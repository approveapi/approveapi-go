# CreatePromptRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The body of the approval request to show the user. | 
**Title** | **string** | The title of an approval request. Defaults to an empty string. | [optional] 
**RejectText** | **string** | The reject action text. If not specified the reject button will NOT be rendered, and the user will only see an approve action button. | [optional] 
**ExpiresIn** | **float32** | The number of seconds until this request can no longer be answered. | [optional] 
**LongPoll** | **bool** | If true, the request waits (long-polls) until the user responds to the prompt or more than 10 minutes pass. Defaults to false. | [optional] 
**User** | **string** | The user to send the approval request to. Can be either an email address or a phone number. | 
**ApproveText** | **string** | The approve action text. Defaults to &#39;Approve&#39;. | [optional] 
**ApproveRedirectUrl** | **string** | An HTTPS URL to redirect the user to if the prompt is approved. This URL is kept secret until the user is redirected to it. | [optional] 
**RejectRedirectUrl** | **string** | An HTTPS URL to redirect the user to if the prompt is rejected. This URL is kept secret until the user is redirected to it. | [optional] 
**Metadata** | [**PromptMetadata**](PromptMetadata.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


