# \ApproveApi

All URIs are relative to *https://approve.sh*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrompt**](ApproveApi.md#CreatePrompt) | **Post** /prompt | Sending a prompt
[**GetPrompt**](ApproveApi.md#GetPrompt) | **Get** /prompt/{id} | Retrieve a prompt
[**GetPromptStatus**](ApproveApi.md#GetPromptStatus) | **Get** /prompt/{id}/status | Check prompt status


# **CreatePrompt**
> Prompt CreatePrompt(ctx, createPromptRequest)
Sending a prompt

Creates a prompt and pushes it to the user (sends via email, sms, or other supported protocols).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPromptRequest** | [**CreatePromptRequest**](CreatePromptRequest.md)|  | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrompt**
> Prompt GetPrompt(ctx, id, optional)
Retrieve a prompt

Retrieve the prompt object with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier for a pending or completed prompt. This is returned when you create a prompt. | 
 **optional** | ***GetPromptOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPromptOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **longPoll** | **optional.Bool**| If true, the request waits (long-polls) until the user responds to the prompt or more than 10 minutes pass. Defaults to false. | 

### Return type

[**Prompt**](Prompt.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPromptStatus**
> PromptStatus GetPromptStatus(ctx, id)
Check prompt status

Returns whether a prompt has been completed by the user. This request does not require authentication, and so can be used client-side without sharing API credentials.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The prompt identifier. | 

### Return type

[**PromptStatus**](PromptStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

