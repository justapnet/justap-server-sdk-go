# \DefaultApi

All URIs are relative to *http://127.0.0.1:21011*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessUserServiceCreateUser**](DefaultApi.md#BusinessUserServiceCreateUser) | **Post** /v1/business_users | 创建 Business User 对象
[**BusinessUserServiceDeleteUser**](DefaultApi.md#BusinessUserServiceDeleteUser) | **Delete** /v1/business_users/{id} | 删除 Business User 对象
[**BusinessUserServiceListAllUsers**](DefaultApi.md#BusinessUserServiceListAllUsers) | **Get** /v1/business_users | 查询 Business User 对象列表
[**BusinessUserServiceRetrieveUser**](DefaultApi.md#BusinessUserServiceRetrieveUser) | **Get** /v1/business_users/{id} | 查询 Business User 对象
[**BusinessUserServiceSearchUsers**](DefaultApi.md#BusinessUserServiceSearchUsers) | **Get** /v1/business_users/search | 查询 Business User 对象列表
[**BusinessUserServiceUpdateUser**](DefaultApi.md#BusinessUserServiceUpdateUser) | **Put** /v1/business_users/{user.id} | 更新 Business User 对象
[**BusinessUserServiceUpdateUser2**](DefaultApi.md#BusinessUserServiceUpdateUser2) | **Patch** /v1/business_users/{user.id} | 更新 Business User 对象
[**ChargeServiceCharges**](DefaultApi.md#ChargeServiceCharges) | **Post** /transaction/v1/charges | 创建 Charge 对象
[**ChargeServiceCharges2**](DefaultApi.md#ChargeServiceCharges2) | **Post** /v1/charges | 创建 Charge 对象
[**ChargeServiceQueryCharge**](DefaultApi.md#ChargeServiceQueryCharge) | **Get** /transaction/v1/charges/{charge_id} | 查询 Charge 对象
[**ChargeServiceQueryCharge2**](DefaultApi.md#ChargeServiceQueryCharge2) | **Get** /v1/charges/{charge_id} | 查询 Charge 对象
[**ChargeServiceQueryCharge3**](DefaultApi.md#ChargeServiceQueryCharge3) | **Get** /v1/charges/merchant_trade_id/{merchant_trade_id} | 查询 Charge 对象
[**ChargeServiceQueryChargeList**](DefaultApi.md#ChargeServiceQueryChargeList) | **Get** /transaction/v1/charges | 查询 Charge 对象列表
[**ChargeServiceQueryChargeList2**](DefaultApi.md#ChargeServiceQueryChargeList2) | **Get** /v1/charges | 查询 Charge 对象列表
[**ChargeServiceReverseCharge**](DefaultApi.md#ChargeServiceReverseCharge) | **Post** /transaction/v1/charges/{charge_id}/reverse | 撤销 Charge 对象
[**ChargeServiceReverseCharge2**](DefaultApi.md#ChargeServiceReverseCharge2) | **Post** /v1/charges/{charge_id}/reverse | 撤销 Charge 对象
[**RefundServiceQueryRefund**](DefaultApi.md#RefundServiceQueryRefund) | **Get** /transaction/v1/charges/{charge_id}/refunds/{refund_id} | 查询 Refund 对象
[**RefundServiceQueryRefund2**](DefaultApi.md#RefundServiceQueryRefund2) | **Get** /v1/refunds/{refund_id} | 查询 Refund 对象
[**RefundServiceQueryRefundList**](DefaultApi.md#RefundServiceQueryRefundList) | **Get** /transaction/v1/charges/{charge_id}/refunds | 查询 Refund 对象列表
[**RefundServiceQueryRefundList2**](DefaultApi.md#RefundServiceQueryRefundList2) | **Get** /v1/refunds | 查询 Refund 对象列表
[**RefundServiceRefunds**](DefaultApi.md#RefundServiceRefunds) | **Post** /transaction/v1/refunds | 创建 Refund 对象
[**RefundServiceRefunds2**](DefaultApi.md#RefundServiceRefunds2) | **Post** /v1/refunds | 创建 Refund 对象
[**RoyaltyServiceCreateRoyalty**](DefaultApi.md#RoyaltyServiceCreateRoyalty) | **Post** /v1/royalties | 创建 Royalty 对象
[**RoyaltyServiceListAllRoyalties**](DefaultApi.md#RoyaltyServiceListAllRoyalties) | **Get** /v1/royalties | 查询 Royalty 对象列表
[**RoyaltyServiceRetrieveRoyalty**](DefaultApi.md#RoyaltyServiceRetrieveRoyalty) | **Get** /v1/royalties/{id} | 查询 Royalty 对象
[**SettlementServiceCreateSettlementAccount**](DefaultApi.md#SettlementServiceCreateSettlementAccount) | **Post** /v1/settlement_accounts | 创建结算账户
[**SettlementServiceDeleteSettlementAccount**](DefaultApi.md#SettlementServiceDeleteSettlementAccount) | **Delete** /v1/settlement_accounts/{id} | 删除结算账户
[**SettlementServiceListAllSettlementAccounts**](DefaultApi.md#SettlementServiceListAllSettlementAccounts) | **Get** /v1/settlement_accounts | 查询结算账户列表
[**SettlementServiceRetrieveSettlementAccount**](DefaultApi.md#SettlementServiceRetrieveSettlementAccount) | **Get** /v1/settlement_accounts/{id} | 查询结算账户
[**SettlementServiceSearchSettlementAccounts**](DefaultApi.md#SettlementServiceSearchSettlementAccounts) | **Get** /v1/settlement_accounts/search | 查询结算账户列表
[**SettlementServiceUpdateSettlementAccount**](DefaultApi.md#SettlementServiceUpdateSettlementAccount) | **Put** /v1/settlement_accounts/{settlementAccount.id} | 更新结算账户
[**SettlementServiceUpdateSettlementAccount2**](DefaultApi.md#SettlementServiceUpdateSettlementAccount2) | **Patch** /v1/settlement_accounts/{settlementAccount.id} | 更新结算账户


# **BusinessUserServiceCreateUser**
> V1UserResponse BusinessUserServiceCreateUser(ctx, body)
创建 Business User 对象

创建 Business User 对象。商业用户是本系统中的一种账户类型，在交易完成之后可以对该类型的账户进行分账等操作。

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
删除 Business User 对象

删除 Business User 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***DefaultApiBusinessUserServiceDeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiBusinessUserServiceDeleteUserOpts struct

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
查询 Business User 对象列表

查询 Business User 对象列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiBusinessUserServiceListAllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiBusinessUserServiceListAllUsersOpts struct

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
查询 Business User 对象

查询 Business User 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***DefaultApiBusinessUserServiceRetrieveUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiBusinessUserServiceRetrieveUserOpts struct

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
查询 Business User 对象列表

查询 Business User 对象列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiBusinessUserServiceSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiBusinessUserServiceSearchUsersOpts struct

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
> V1UserResponse BusinessUserServiceUpdateUser(ctx, userId, body, optional)
更新 Business User 对象

更新 Business User 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **body** | [**V1BusinessUser**](V1BusinessUser.md)|  | 
 **optional** | ***DefaultApiBusinessUserServiceUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiBusinessUserServiceUpdateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **optional.String**|  | 

### Return type

[**V1UserResponse**](v1UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BusinessUserServiceUpdateUser2**
> V1UserResponse BusinessUserServiceUpdateUser2(ctx, userId, body, optional)
更新 Business User 对象

更新 Business User 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 
  **body** | [**V1BusinessUser**](V1BusinessUser.md)|  | 
 **optional** | ***DefaultApiBusinessUserServiceUpdateUser2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiBusinessUserServiceUpdateUser2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **optional.String**|  | 

### Return type

[**V1UserResponse**](v1UserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceCharges**
> V1ChargeResponse ChargeServiceCharges(ctx, body)
创建 Charge 对象

发起一次支付请求时需要创建一个新的 charge 对象，获取一个可用的支付凭据用于客户端向第三方渠道发起支付请求。如果使用测试模式的 API Key，则不会发生真实交易。当支付成功后，会发送 Webhooks 通知。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateChargeRequest**](V1CreateChargeRequest.md)|  | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceCharges2**
> V1ChargeResponse ChargeServiceCharges2(ctx, body)
创建 Charge 对象

发起一次支付请求时需要创建一个新的 charge 对象，获取一个可用的支付凭据用于客户端向第三方渠道发起支付请求。如果使用测试模式的 API Key，则不会发生真实交易。当支付成功后，会发送 Webhooks 通知。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateChargeRequest**](V1CreateChargeRequest.md)|  | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceQueryCharge**
> V1ChargeResponse ChargeServiceQueryCharge(ctx, chargeId, optional)
查询 Charge 对象

你可以在后台异步通知之前，通过查询接口确认支付状态。通过charge对象的id查询一个已创建的charge对象。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| [REQUIRED] Charge 对象 id | 
 **optional** | ***DefaultApiChargeServiceQueryChargeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiChargeServiceQueryChargeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**| [REQUIRED] 应用 id | 
 **merchantTradeId** | **optional.String**| [OPTIONAL] 商户订单号 | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceQueryCharge2**
> V1ChargeResponse ChargeServiceQueryCharge2(ctx, chargeId, optional)
查询 Charge 对象

你可以在后台异步通知之前，通过查询接口确认支付状态。通过charge对象的id查询一个已创建的charge对象。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| [REQUIRED] Charge 对象 id | 
 **optional** | ***DefaultApiChargeServiceQueryCharge2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiChargeServiceQueryCharge2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**| [REQUIRED] 应用 id | 
 **merchantTradeId** | **optional.String**| [OPTIONAL] 商户订单号 | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceQueryCharge3**
> V1ChargeResponse ChargeServiceQueryCharge3(ctx, merchantTradeId, optional)
查询 Charge 对象

你可以在后台异步通知之前，通过查询接口确认支付状态。通过charge对象的id查询一个已创建的charge对象。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **merchantTradeId** | **string**| [OPTIONAL] 商户订单号 | 
 **optional** | ***DefaultApiChargeServiceQueryCharge3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiChargeServiceQueryCharge3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeId** | **optional.String**| [REQUIRED] Charge 对象 id | 
 **appId** | **optional.String**| [REQUIRED] 应用 id | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceQueryChargeList**
> V1ChargeListResponse ChargeServiceQueryChargeList(ctx, optional)
查询 Charge 对象列表

返回之前创建过 charge 对象的一个列表。列表是按创建时间进行排序，总是将最新的 charge 对象显示在最前。如果不设置 created 参数，默认查询近一个月的数据；设置了 created 参数，会按照对应的时间段查询。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiChargeServiceQueryChargeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiChargeServiceQueryChargeListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **optional.String**| [REQUIRED] 应用 id | 
 **limit** | **optional.Int32**| [OPTIONAL] 限制有多少对象可以被返回，限制范围是从 1~100 项，默认是 10 项 | [default to 10]
 **startingAfter** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的第一项从何处开始。假设你的一次请求返回列表的最后一项的 id 是 obj_end，你可以使用 starting_after &#x3D; obj_end 去获取下一页 | 
 **endingBefore** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的最末项在何处结束。假设你的一次请求返回列表的第一项的 id 是 obj_start，你可以使用 ending_before &#x3D; obj_start 去获取上一页 | 
 **merchantTradeId** | **optional.String**| [OPTIONAL] 客户系统订单号 | 
 **createdLt** | **optional.Int64**| 大于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdLte** | **optional.Int64**| 大于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGt** | **optional.Int64**| 小于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGte** | **optional.Int64**| 小于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **channel** | **optional.String**| [OPTIONAL] 渠道名称   - BALANCE: 余额支付  - AlipayQR: 支付宝扫码支付  - AlipayScan: 支付宝条码支付  - AlipayApp: 支付宝 App 支付  - AlipayWap: 支付宝手机网站支付  - AlipayPage: 支付宝电脑网站支付  - AlipayFace: 支付宝刷脸支付  - AlipayLite: 支付宝小程序支付  - AlipayJSAPI: 支付宝 JSAPI 支付  - WechatpayApp: 微信 App 支付  - WechatpayJSAPI: 微信 JSAPI 支付  - WechatpayH5: 微信 H5 支付  - WechatpayNative: 微信 Native 支付  - WechatpayLite: 微信小程序支付  - WechatpayFace: 刷脸支付  - WechatpayScan: 微信付款码支付  - UnionPayQr: 银联二维码支付（云闪付扫码） | [default to CHANNEL_INVALID_UNSPECIFIED]
 **paid** | **optional.Bool**| [OPTIONAL] 是否已付款 | [default to false]
 **refunded** | **optional.Bool**| [OPTIONAL] 是否存在退款信息，无论退款是否成功。 | [default to false]
 **reversed** | **optional.Bool**| [OPTIONAL] 是否已撤销 | [default to false]
 **closed** | **optional.Bool**| [OPTIONAL] 是否已关闭 | [default to false]
 **expired** | **optional.Bool**| [OPTIONAL] 是否已过期 | [default to false]

### Return type

[**V1ChargeListResponse**](v1ChargeListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceQueryChargeList2**
> V1ChargeListResponse ChargeServiceQueryChargeList2(ctx, optional)
查询 Charge 对象列表

返回之前创建过 charge 对象的一个列表。列表是按创建时间进行排序，总是将最新的 charge 对象显示在最前。如果不设置 created 参数，默认查询近一个月的数据；设置了 created 参数，会按照对应的时间段查询。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiChargeServiceQueryChargeList2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiChargeServiceQueryChargeList2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **optional.String**| [REQUIRED] 应用 id | 
 **limit** | **optional.Int32**| [OPTIONAL] 限制有多少对象可以被返回，限制范围是从 1~100 项，默认是 10 项 | [default to 10]
 **startingAfter** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的第一项从何处开始。假设你的一次请求返回列表的最后一项的 id 是 obj_end，你可以使用 starting_after &#x3D; obj_end 去获取下一页 | 
 **endingBefore** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的最末项在何处结束。假设你的一次请求返回列表的第一项的 id 是 obj_start，你可以使用 ending_before &#x3D; obj_start 去获取上一页 | 
 **merchantTradeId** | **optional.String**| [OPTIONAL] 客户系统订单号 | 
 **createdLt** | **optional.Int64**| 大于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdLte** | **optional.Int64**| 大于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGt** | **optional.Int64**| 小于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGte** | **optional.Int64**| 小于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **channel** | **optional.String**| [OPTIONAL] 渠道名称   - BALANCE: 余额支付  - AlipayQR: 支付宝扫码支付  - AlipayScan: 支付宝条码支付  - AlipayApp: 支付宝 App 支付  - AlipayWap: 支付宝手机网站支付  - AlipayPage: 支付宝电脑网站支付  - AlipayFace: 支付宝刷脸支付  - AlipayLite: 支付宝小程序支付  - AlipayJSAPI: 支付宝 JSAPI 支付  - WechatpayApp: 微信 App 支付  - WechatpayJSAPI: 微信 JSAPI 支付  - WechatpayH5: 微信 H5 支付  - WechatpayNative: 微信 Native 支付  - WechatpayLite: 微信小程序支付  - WechatpayFace: 刷脸支付  - WechatpayScan: 微信付款码支付  - UnionPayQr: 银联二维码支付（云闪付扫码） | [default to CHANNEL_INVALID_UNSPECIFIED]
 **paid** | **optional.Bool**| [OPTIONAL] 是否已付款 | [default to false]
 **refunded** | **optional.Bool**| [OPTIONAL] 是否存在退款信息，无论退款是否成功。 | [default to false]
 **reversed** | **optional.Bool**| [OPTIONAL] 是否已撤销 | [default to false]
 **closed** | **optional.Bool**| [OPTIONAL] 是否已关闭 | [default to false]
 **expired** | **optional.Bool**| [OPTIONAL] 是否已过期 | [default to false]

### Return type

[**V1ChargeListResponse**](v1ChargeListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceReverseCharge**
> V1ChargeResponse ChargeServiceReverseCharge(ctx, chargeId)
撤销 Charge 对象

针对已经创建的 Charge，你可以调用撤销接口进行交易的关闭。接口支持对于未成功付款的订单进行撤销，则关闭交易。调用此接口后用户后期不能支付成功。  注：撤销订单在不同收单机构会有不同的行为。对于成功付款的订单请使用 退款 接口进行退款处理。只有针对未支付的订单，我们建议你调用撤销接口。  - 微信支付：如果此订单用户支付失败，微信支付系统会将此订单关闭；如果用户支付成功，微信支付系统会将此订单资金退还给用户。(7天以内的交易单可调用撤销) - 支付宝：如果此订单用户支付失败，支付宝系统会将此订单关闭；如果用户支付成功，支付宝系统会将此订单资金退还给用户。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| Charge 对象 id | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChargeServiceReverseCharge2**
> V1ChargeResponse ChargeServiceReverseCharge2(ctx, chargeId)
撤销 Charge 对象

针对已经创建的 Charge，你可以调用撤销接口进行交易的关闭。接口支持对于未成功付款的订单进行撤销，则关闭交易。调用此接口后用户后期不能支付成功。  注：撤销订单在不同收单机构会有不同的行为。对于成功付款的订单请使用 退款 接口进行退款处理。只有针对未支付的订单，我们建议你调用撤销接口。  - 微信支付：如果此订单用户支付失败，微信支付系统会将此订单关闭；如果用户支付成功，微信支付系统会将此订单资金退还给用户。(7天以内的交易单可调用撤销) - 支付宝：如果此订单用户支付失败，支付宝系统会将此订单关闭；如果用户支付成功，支付宝系统会将此订单资金退还给用户。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| Charge 对象 id | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundServiceQueryRefund**
> V1RefundResponse RefundServiceQueryRefund(ctx, chargeId, refundId, optional)
查询 Refund 对象

可以通过 charge 对象的查询接口查询某一个 charge 对象的退款列表，也可以通过 refund 对象的 id 查询一个已创建的 refund 对象。可以在 Webhooks 通知之前，通过查询接口确认退款状态。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| [REQUIRED] 支付 Charge Id | 
  **refundId** | **string**| [REQUIRED] Refund 对象 id | 
 **optional** | ***DefaultApiRefundServiceQueryRefundOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRefundServiceQueryRefundOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appId** | **optional.String**| [REQUIRED] 应用 id | 

### Return type

[**V1RefundResponse**](v1RefundResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundServiceQueryRefund2**
> V1RefundResponse RefundServiceQueryRefund2(ctx, refundId, optional)
查询 Refund 对象

可以通过 charge 对象的查询接口查询某一个 charge 对象的退款列表，也可以通过 refund 对象的 id 查询一个已创建的 refund 对象。可以在 Webhooks 通知之前，通过查询接口确认退款状态。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **refundId** | **string**| [REQUIRED] Refund 对象 id | 
 **optional** | ***DefaultApiRefundServiceQueryRefund2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRefundServiceQueryRefund2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeId** | **optional.String**| [REQUIRED] 支付 Charge Id | 
 **appId** | **optional.String**| [REQUIRED] 应用 id | 

### Return type

[**V1RefundResponse**](v1RefundResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundServiceQueryRefundList**
> V1RefundListResponse RefundServiceQueryRefundList(ctx, chargeId, optional)
查询 Refund 对象列表

返回之前创建 charge 对象的一个 refund 对象列表。列表是按创建时间进行排序，总是将最新的 refund 对象显示在最前。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| [REQUIRED] 支付 Charge Id | 
 **optional** | ***DefaultApiRefundServiceQueryRefundListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRefundServiceQueryRefundListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**| [REQUIRED] 应用 id | 
 **limit** | **optional.Int32**| [OPTIONAL] 限制有多少对象可以被返回，限制范围是从 1~100 项，默认是 10 项 | [default to 10]
 **startingAfter** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的第一项从何处开始。假设你的一次请求返回列表的最后一项的 id 是 obj_end，你可以使用 starting_after &#x3D; obj_end 去获取下一页 | 
 **endingBefore** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的最末项在何处结束。假设你的一次请求返回列表的第一项的 id 是 obj_start，你可以使用 ending_before &#x3D; obj_start 去获取上一页 | 

### Return type

[**V1RefundListResponse**](v1RefundListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundServiceQueryRefundList2**
> V1RefundListResponse RefundServiceQueryRefundList2(ctx, optional)
查询 Refund 对象列表

返回之前创建 charge 对象的一个 refund 对象列表。列表是按创建时间进行排序，总是将最新的 refund 对象显示在最前。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiRefundServiceQueryRefundList2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRefundServiceQueryRefundList2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chargeId** | **optional.String**| [REQUIRED] 支付 Charge Id | 
 **appId** | **optional.String**| [REQUIRED] 应用 id | 
 **limit** | **optional.Int32**| [OPTIONAL] 限制有多少对象可以被返回，限制范围是从 1~100 项，默认是 10 项 | [default to 10]
 **startingAfter** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的第一项从何处开始。假设你的一次请求返回列表的最后一项的 id 是 obj_end，你可以使用 starting_after &#x3D; obj_end 去获取下一页 | 
 **endingBefore** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的最末项在何处结束。假设你的一次请求返回列表的第一项的 id 是 obj_start，你可以使用 ending_before &#x3D; obj_start 去获取上一页 | 

### Return type

[**V1RefundListResponse**](v1RefundListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundServiceRefunds**
> V1RefundResponse RefundServiceRefunds(ctx, body)
创建 Refund 对象

通过发起一次退款请求创建一个新的 refund 对象，只能对已经发生交易并且没有全额退款的 charge 对象发起退款。当进行全额退款之前，可以进行多次退款，直至全额退款。当一次退款成功后，会发送 Webhooks 通知。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateRefundRequest**](V1CreateRefundRequest.md)|  | 

### Return type

[**V1RefundResponse**](v1RefundResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundServiceRefunds2**
> V1RefundResponse RefundServiceRefunds2(ctx, body)
创建 Refund 对象

通过发起一次退款请求创建一个新的 refund 对象，只能对已经发生交易并且没有全额退款的 charge 对象发起退款。当进行全额退款之前，可以进行多次退款，直至全额退款。当一次退款成功后，会发送 Webhooks 通知。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateRefundRequest**](V1CreateRefundRequest.md)|  | 

### Return type

[**V1RefundResponse**](v1RefundResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoyaltyServiceCreateRoyalty**
> V1RoyaltyResponse RoyaltyServiceCreateRoyalty(ctx, body)
创建 Royalty 对象

对一个 Charge 对象进行分账，分账的金额和分账接收方由 Royalty 对象指定。Royalty 创建仅代表本系统成功接收分账申请，尚未提交到支付机构清分，更不代表分账立即成功，相关结果信息请调用查询接口确认

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

# **RoyaltyServiceListAllRoyalties**
> V1ListAllRoyaltiesResponse RoyaltyServiceListAllRoyalties(ctx, optional)
查询 Royalty 对象列表

查询 Royalty 对象的列表信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiRoyaltyServiceListAllRoyaltiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRoyaltyServiceListAllRoyaltiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| [OPTIONAL] 限制有多少对象可以被返回，限制范围是从 1~100 项，默认是 10 项 | [default to 10]
 **startingAfter** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的第一项从何处开始。假设你的一次请求返回列表的最后一项的 id 是 obj_end，你可以使用 starting_after &#x3D; obj_end 去获取下一页 | 
 **endingBefore** | **optional.String**| [OPTIONAL] 在分页时使用的指针，决定了列表的最末项在何处结束。假设你的一次请求返回列表的第一项的 id 是 obj_start，你可以使用 ending_before &#x3D; obj_start 去获取上一页 | 
 **merchantTradeId** | **optional.String**| [OPTIONAL] 客户系统订单号 | 
 **createdLt** | **optional.Int64**| 大于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdLte** | **optional.Int64**| 大于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGt** | **optional.Int64**| 小于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **createdGte** | **optional.Int64**| 小于或等于 charge 对象的创建时间，用 Unix 时间戳表示 | [default to 0]
 **appId** | **optional.String**|  | 
 **settleAccountId** | **optional.String**|  | 
 **royaltySettlementId** | **optional.String**|  | 

### Return type

[**V1ListAllRoyaltiesResponse**](v1ListAllRoyaltiesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoyaltyServiceRetrieveRoyalty**
> V1RoyaltyResponse RoyaltyServiceRetrieveRoyalty(ctx, id)
查询 Royalty 对象

查询 Royalty 对象的信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**V1RoyaltyResponse**](v1RoyaltyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettlementServiceCreateSettlementAccount**
> V1SettlementAccountResponse SettlementServiceCreateSettlementAccount(ctx, body)
创建结算账户

为用户创建一个结算账户。添加结算账户信息后方可对该用进行分账、余额充值、转账等操作。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateSettlementAccountRequest**](V1CreateSettlementAccountRequest.md)|  | 

### Return type

[**V1SettlementAccountResponse**](v1SettlementAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettlementServiceDeleteSettlementAccount**
> V1DeleteSettlementAccountResponse SettlementServiceDeleteSettlementAccount(ctx, id, optional)
删除结算账户

删除结算账户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***DefaultApiSettlementServiceDeleteSettlementAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSettlementServiceDeleteSettlementAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 

### Return type

[**V1DeleteSettlementAccountResponse**](v1DeleteSettlementAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettlementServiceListAllSettlementAccounts**
> V1SettlementAccountListResponse SettlementServiceListAllSettlementAccounts(ctx, optional)
查询结算账户列表

查询结算账户列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiSettlementServiceListAllSettlementAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSettlementServiceListAllSettlementAccountsOpts struct

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
 **customerId** | **optional.String**| [OPTIONAL] 客户 ID | 
 **businessUserId** | **optional.String**| [OPTIONAL] 商户用户 ID | 

### Return type

[**V1SettlementAccountListResponse**](v1SettlementAccountListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettlementServiceRetrieveSettlementAccount**
> V1SettlementAccountResponse SettlementServiceRetrieveSettlementAccount(ctx, id, optional)
查询结算账户

查询结算账户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***DefaultApiSettlementServiceRetrieveSettlementAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSettlementServiceRetrieveSettlementAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 
 **object** | **optional.String**| 对象类型 | [default to SettlementAccount]
 **dataId** | **optional.String**| 分账接收方的唯一标识 | [default to 0]
 **dataAppId** | **optional.String**| 分账接收方所在的应用 ID | [default to 0]
 **dataBusinessUserId** | **optional.String**| 分账接收方的用户 ID | [default to 0]
 **dataCustomerId** | **optional.String**| 分账接收方的用户 ID | [default to 0]
 **dataChannel** | **optional.String**| 分账接收方的账户类型 | [default to CHANNEL_UNKNOWN]
 **dataRecipientWechatpayAccount** | **optional.String**| openid 或者商户号，由类型决定. 微信支付分账接收方账户，OPENID或者商户号 | 
 **dataRecipientWechatpayName** | **optional.String**| 微信支付分账接收方姓名或名称 | 
 **dataRecipientWechatpayForceCheck** | **optional.Bool**| 是否强制校验收款人姓名 | [default to false]
 **dataRecipientWechatpayType** | **optional.String**| 微信支付分账接收方类型 | [default to TYPE_UNSET]
 **dataRecipientWechatpayAccountType** | **optional.String**| 微信支付分账接收方账户类型 | [default to ACCOUNT_TYPE_UNSET]
 **dataRecipientWechatpayAppId** | **optional.String**| 微信支付分账接收方 openid 所对应的服务商公众号 ID | 
 **dataRecipientWechatpaySubAppId** | **optional.String**| 微信支付分账接收方 openid 所对应的商户公众号 ID | 
 **dataRecipientPaymentAlipayAccount** | **optional.String**| 支付宝账号，账号ID或者登录邮箱 | 
 **dataRecipientPaymentAlipayName** | **optional.String**| 支付宝账号真实姓名 | 
 **dataRecipientPaymentAlipayType** | **optional.String**| 支付宝账号类型 | [default to TYPE_UNSET]
 **dataRecipientPaymentAlipayAccountType** | **optional.String**| 支付宝账号类型 | [default to ACCOUNT_TYPE_UNSET]
 **dataRecipientBankAccount** | **optional.String**| 银行卡号 | 
 **dataRecipientBankName** | **optional.String**| 银行卡开户名 | 
 **dataRecipientBankType** | **optional.String**| 银行卡类型 | 
 **dataRecipientBankBankName** | **optional.String**| 银行卡开户行编码 | 
 **dataRecipientBankBankBranch** | **optional.String**| 银行卡开户支行 | 
 **dataRecipientBankBankProvince** | **optional.String**| 银行卡开户省份 | 
 **dataRecipientBankBankCity** | **optional.String**| 银行卡开户城市 | 
 **dataRecipientYsepayMerchantDivisionMerUsercode** | **optional.String**| 银盛商户号 | 
 **dataCreated** | **optional.Int64**| 分账接收方的创建时间 | [default to 0]
 **dataUpdated** | **optional.Int64**| 分账接收方的更新时间 | [default to 0]
 **dataObject** | **optional.String**| 对象类型 | [default to Recipient]

### Return type

[**V1SettlementAccountResponse**](v1SettlementAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettlementServiceSearchSettlementAccounts**
> V1SettlementAccountListResponse SettlementServiceSearchSettlementAccounts(ctx, optional)
查询结算账户列表

查询结算账户列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiSettlementServiceSearchSettlementAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSettlementServiceSearchSettlementAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.String**|  | 
 **appId** | **optional.String**|  | 

### Return type

[**V1SettlementAccountListResponse**](v1SettlementAccountListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettlementServiceUpdateSettlementAccount**
> V1SettlementAccountResponse SettlementServiceUpdateSettlementAccount(ctx, settlementAccountId, body, optional)
更新结算账户

更新结算账户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settlementAccountId** | **string**|  | 
  **body** | [**V1UpdateAndPatchRequestBody**](V1UpdateAndPatchRequestBody.md)|  | 
 **optional** | ***DefaultApiSettlementServiceUpdateSettlementAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSettlementServiceUpdateSettlementAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **optional.String**|  | 

### Return type

[**V1SettlementAccountResponse**](v1SettlementAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettlementServiceUpdateSettlementAccount2**
> V1SettlementAccountResponse SettlementServiceUpdateSettlementAccount2(ctx, settlementAccountId, body, optional)
更新结算账户

更新结算账户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settlementAccountId** | **string**|  | 
  **body** | [**V1UpdateAndPatchRequestBody**](V1UpdateAndPatchRequestBody.md)|  | 
 **optional** | ***DefaultApiSettlementServiceUpdateSettlementAccount2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSettlementServiceUpdateSettlementAccount2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **optional.String**|  | 

### Return type

[**V1SettlementAccountResponse**](v1SettlementAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

