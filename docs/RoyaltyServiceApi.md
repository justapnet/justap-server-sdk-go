# \RoyaltyServiceApi

All URIs are relative to *http://127.0.0.1:21011*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoyaltyServiceCreateRoyalty**](RoyaltyServiceApi.md#RoyaltyServiceCreateRoyalty) | **Post** /v1/royalties | 创建 Royalty 对象


# **RoyaltyServiceCreateRoyalty**
> V1RoyaltyResponse RoyaltyServiceCreateRoyalty(ctx, body)
创建 Royalty 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateRoyaltyRequest**](V1CreateRoyaltyRequest.md)|  | 

### Return type

[**V1RoyaltyResponse**](v1RoyaltyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

