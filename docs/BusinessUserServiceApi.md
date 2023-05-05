# \BusinessUserServiceApi

All URIs are relative to *http://127.0.0.1:21011*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessUserServiceCreateUser**](BusinessUserServiceApi.md#BusinessUserServiceCreateUser) | **Post** /v1/business_users | 创建 BusinessUser 对象
[**BusinessUserServiceDeleteUser**](BusinessUserServiceApi.md#BusinessUserServiceDeleteUser) | **Delete** /v1/business_users/{id} | 删除 BusinessUser 对象
[**BusinessUserServiceListAllUsers**](BusinessUserServiceApi.md#BusinessUserServiceListAllUsers) | **Get** /v1/business_users | 查询 BusinessUser 对象列表
[**BusinessUserServiceRetrieveUser**](BusinessUserServiceApi.md#BusinessUserServiceRetrieveUser) | **Get** /v1/business_users/{id} | 查询 BusinessUser 对象
[**BusinessUserServiceSearchUsers**](BusinessUserServiceApi.md#BusinessUserServiceSearchUsers) | **Get** /v1/business_users/search | 搜索 BusinessUser 对象
[**BusinessUserServiceUpdateUser**](BusinessUserServiceApi.md#BusinessUserServiceUpdateUser) | **Post** /v1/business_users/{id} | 更新 BusinessUser 对象


# **BusinessUserServiceCreateUser**
> V1UserResponse BusinessUserServiceCreateUser(ctx, body)
创建 BusinessUser 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateUserRequest**](V1CreateUserRequest.md)|  | 

### Return type

[**V1UserResponse**](v1UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BusinessUserServiceDeleteUser**
> V1DeleteUserResponse BusinessUserServiceDeleteUser(ctx, id, optional)
删除 BusinessUser 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***BusinessUserServiceApiBusinessUserServiceDeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUserServiceApiBusinessUserServiceDeleteUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 

### Return type

[**V1DeleteUserResponse**](v1DeleteUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BusinessUserServiceListAllUsers**
> V1UserListResponse BusinessUserServiceListAllUsers(ctx, optional)
查询 BusinessUser 对象列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessUserServiceApiBusinessUserServiceListAllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUserServiceApiBusinessUserServiceListAllUsersOpts struct

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

[**V1UserListResponse**](v1UserListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BusinessUserServiceRetrieveUser**
> V1UserResponse BusinessUserServiceRetrieveUser(ctx, id, optional)
查询 BusinessUser 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***BusinessUserServiceApiBusinessUserServiceRetrieveUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUserServiceApiBusinessUserServiceRetrieveUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 

### Return type

[**V1UserResponse**](v1UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BusinessUserServiceSearchUsers**
> V1UserListResponse BusinessUserServiceSearchUsers(ctx, optional)
搜索 BusinessUser 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BusinessUserServiceApiBusinessUserServiceSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUserServiceApiBusinessUserServiceSearchUsersOpts struct

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

[**V1UserListResponse**](v1UserListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BusinessUserServiceUpdateUser**
> V1UserResponse BusinessUserServiceUpdateUser(ctx, id, optional)
更新 BusinessUser 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***BusinessUserServiceApiBusinessUserServiceUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessUserServiceApiBusinessUserServiceUpdateUserOpts struct

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
 **parentUserId** | **optional.String**|  | 

### Return type

[**V1UserResponse**](v1UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

