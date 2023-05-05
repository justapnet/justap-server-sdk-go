# \CheckoutServiceApi

All URIs are relative to *http://127.0.0.1:21011*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckoutServiceCreateUnionQrCheckout**](CheckoutServiceApi.md#CheckoutServiceCreateUnionQrCheckout) | **Post** /v1/checkout/union_qr | 通过聚合收款码创建订单


# **CheckoutServiceCreateUnionQrCheckout**
> V1UnionQrRequest CheckoutServiceCreateUnionQrCheckout(ctx, body)
通过聚合收款码创建订单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1UnionQrRequest**](V1UnionQrRequest.md)|  | 

### Return type

[**V1UnionQrRequest**](v1UnionQrRequest.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

