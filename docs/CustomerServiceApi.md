# \CustomerServiceApi

All URIs are relative to *http://127.0.0.1:21011*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerServiceCreateCustomer**](CustomerServiceApi.md#CustomerServiceCreateCustomer) | **Post** /v1/customers | 
[**CustomerServiceDeleteCustomer**](CustomerServiceApi.md#CustomerServiceDeleteCustomer) | **Delete** /v1/customers/{id} | 
[**CustomerServiceListAllCustomers**](CustomerServiceApi.md#CustomerServiceListAllCustomers) | **Get** /v1/customers | 
[**CustomerServiceRetrieveCustomer**](CustomerServiceApi.md#CustomerServiceRetrieveCustomer) | **Get** /v1/customers/{id} | 
[**CustomerServiceSearchCustomers**](CustomerServiceApi.md#CustomerServiceSearchCustomers) | **Get** /v1/customers/search | 
[**CustomerServiceUpdateCustomer**](CustomerServiceApi.md#CustomerServiceUpdateCustomer) | **Post** /v1/customers/{id} | 


# **CustomerServiceCreateCustomer**
> V1CustomerResponse CustomerServiceCreateCustomer(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateCustomerRequest**](V1CreateCustomerRequest.md)|  | 

### Return type

[**V1CustomerResponse**](v1CustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerServiceDeleteCustomer**
> V1DeleteCustomerResponse CustomerServiceDeleteCustomer(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomerServiceApiCustomerServiceDeleteCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerServiceApiCustomerServiceDeleteCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 

### Return type

[**V1DeleteCustomerResponse**](v1DeleteCustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerServiceListAllCustomers**
> V1CustomerListResponse CustomerServiceListAllCustomers(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerServiceApiCustomerServiceListAllCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerServiceApiCustomerServiceListAllCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **optional.String**|  | 
 **limit** | **optional.Int32**| [OPTIONAL] 限制有多少对象可以被返回，限制范围是从 1~100 项，默认是 10 项 | [default to 10]
 **startingAfter** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的第一项从何处开始。假设你的一次请求返回列表的最后一项的 id 是 obj_end，你可以使用 starting_after &#x3D; obj_end 去获取下一页 | 
 **endingBefore** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的最末项在何处结束。假设你的一次请求返回列表的第一项的 id 是 obj_start，你可以使用 ending_before &#x3D; obj_start 去获取上一页 | 
 **createdLt** | **optional.Int64**| 大于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdLte** | **optional.Int64**| 大于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGt** | **optional.Int64**| 小于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGte** | **optional.Int64**| 小于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **disabled** | **optional.Bool**| [OPTIONAL] 是否禁用，默认为 false | 

### Return type

[**V1CustomerListResponse**](v1CustomerListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerServiceRetrieveCustomer**
> V1CustomerResponse CustomerServiceRetrieveCustomer(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomerServiceApiCustomerServiceRetrieveCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerServiceApiCustomerServiceRetrieveCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 

### Return type

[**V1CustomerResponse**](v1CustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerServiceSearchCustomers**
> V1CustomerListResponse CustomerServiceSearchCustomers(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerServiceApiCustomerServiceSearchCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerServiceApiCustomerServiceSearchCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **optional.String**|  | 
 **limit** | **optional.Int32**| [OPTIONAL] 限制有多少对象可以被返回，限制范围是从 1~100 项，默认是 10 项 | [default to 10]
 **createdLt** | **optional.Int64**| 大于 BusinessUser 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdLte** | **optional.Int64**| 大于或等于 BusinessUser 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGt** | **optional.Int64**| 小于 BusinessUser 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGte** | **optional.Int64**| 小于或等于 BusinessUser 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **email** | **optional.String**| [OPTIONAL] BusinessUser 对象的邮箱地址。支持模糊匹配 | 
 **name** | **optional.String**| [OPTIONAL] BusinessUser 对象的用户名。支持模糊匹配 | 
 **phone** | **optional.String**| [OPTIONAL] BusinessUser 对象的手机号码 | 

### Return type

[**V1CustomerListResponse**](v1CustomerListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerServiceUpdateCustomer**
> V1CustomerResponse CustomerServiceUpdateCustomer(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CustomerServiceApiCustomerServiceUpdateCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerServiceApiCustomerServiceUpdateCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 
 **address** | **optional.String**|  | 
 **currency** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **phone** | **optional.String**|  | 
 **avatar** | **optional.String**|  | 
 **disabled** | **optional.Bool**|  | 
 **gender** | **optional.String**|  - GENDER_UNKNOWN: 未设置  - MALE: 男  - FE_MALE: 女  - PRIVACY: 保密  - ThirdGender: 第三性别 | [default to GENDER_UNKNOWN]
 **parentCustomerId** | **optional.String**|  | 
 **outCustomerId** | **optional.String**|  | 

### Return type

[**V1CustomerResponse**](v1CustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

