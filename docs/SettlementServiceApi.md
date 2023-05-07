# \SettlementServiceApi

All URIs are relative to *http://127.0.0.1:21011*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettlementServiceCreateSettlementAccount**](SettlementServiceApi.md#SettlementServiceCreateSettlementAccount) | **Post** /v1/settlement_accounts | 创建 SettlementAccount 对象
[**SettlementServiceDeleteSettlementAccount**](SettlementServiceApi.md#SettlementServiceDeleteSettlementAccount) | **Delete** /v1/settlement_accounts/{id} | 删除 SettlementAccount 对象
[**SettlementServiceListAllSettlementAccounts**](SettlementServiceApi.md#SettlementServiceListAllSettlementAccounts) | **Get** /v1/settlement_accounts | 查询 SettlementAccount 对象列表
[**SettlementServiceRetrieveSettlementAccount**](SettlementServiceApi.md#SettlementServiceRetrieveSettlementAccount) | **Get** /v1/settlement_accounts/{id} | 查询 SettlementAccount 对象
[**SettlementServiceSearchSettlementAccounts**](SettlementServiceApi.md#SettlementServiceSearchSettlementAccounts) | **Get** /v1/settlement_accounts/search | 搜索 SettlementAccount 对象
[**SettlementServiceUpdateSettlementAccount**](SettlementServiceApi.md#SettlementServiceUpdateSettlementAccount) | **Post** /v1/settlement_accounts/{id} | 更新 SettlementAccount 对象


# **SettlementServiceCreateSettlementAccount**
> V1SettlementAccountResponse SettlementServiceCreateSettlementAccount(ctx, optional)
创建 SettlementAccount 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettlementServiceApiSettlementServiceCreateSettlementAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettlementServiceApiSettlementServiceCreateSettlementAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **optional.String**|  | 
 **userId** | **optional.String**|  | 
 **customerId** | **optional.String**|  | 
 **channel** | **optional.String**|  - ALIPAY: 分账到支付宝  - WECHANTPAY: 分账到微信支付  - BANK: 分账到银行卡  - BALANCE: 分账到 justap 账户余额 | [default to CHANNEL_UNKNOWN]
 **recipientWechatpayChannelRecipientAccount** | **optional.String**| openid 或者商户号，由类型决定  微信支付分账接收方账户，OPENID或者商户号 | 
 **recipientWechatpayChannelRecipientName** | **optional.String**| 微信支付分账接收方姓名或名称 | 
 **recipientWechatpayChannelRecipientForceCheck** | **optional.Bool**| 是否强制校验收款人姓名 | [default to false]
 **recipientWechatpayChannelRecipientType** | **optional.String**| 微信支付分账接收方类型 | [default to TYPE_UNSET]
 **recipientWechatpayChannelRecipientAccountType** | **optional.String**| 微信支付分账接收方账户类型   - ACCOUNT_TYPE_UNSET: 未设置  - MERCHANT_ID: 分账到微信商户号  - OPENID: 分账到个人微信号（父公众号的openid，或服务商的openid））  - SUB_OPENID: 分账到个人微信号，子账号的  - LOGIN_NAME: 分账到微信登录号 | [default to ACCOUNT_TYPE_UNSET]
 **recipientWechatpayChannelRecipientAppId** | **optional.String**| 微信支付分账接收方 openid 所对应的公众号 ID | 
 **recipientAlipayChannelRecipientAccount** | **optional.String**| 支付宝账号，账号ID或者登录邮箱 | 
 **recipientAlipayChannelRecipientName** | **optional.String**| 支付宝账号真实姓名 | 
 **recipientAlipayChannelRecipientType** | **optional.String**| 支付宝账号类型 | [default to TYPE_UNSET]
 **recipientAlipayChannelRecipientAccountType** | **optional.String**| 支付宝账号类型   - ACCOUNT_TYPE_UNSET: 未设置  - MERCHANT_ID: 分账到微信商户号  - OPENID: 分账到个人微信号（父公众号的openid，或服务商的openid））  - SUB_OPENID: 分账到个人微信号，子账号的  - LOGIN_NAME: 分账到微信登录号 | [default to ACCOUNT_TYPE_UNSET]
 **recipientBankChannelRecipientAccount** | **optional.String**| 银行卡号 | 
 **recipientBankChannelRecipientName** | **optional.String**| 银行卡开户名 | 
 **recipientBankChannelRecipientType** | **optional.String**| 银行卡类型 | 
 **recipientBankChannelRecipientBankName** | **optional.String**| 银行卡开户行编码 | 
 **recipientBankChannelRecipientBankBranch** | **optional.String**| 银行卡开户支行 | 
 **recipientBankChannelRecipientBankProvince** | **optional.String**| 银行卡开户省份 | 
 **recipientBankChannelRecipientBankCity** | **optional.String**| 银行卡开户城市 | 

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
删除 SettlementAccount 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SettlementServiceApiSettlementServiceDeleteSettlementAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettlementServiceApiSettlementServiceDeleteSettlementAccountOpts struct

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
查询 SettlementAccount 对象列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettlementServiceApiSettlementServiceListAllSettlementAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettlementServiceApiSettlementServiceListAllSettlementAccountsOpts struct

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
 **userId** | **optional.String**| [OPTIONAL] 商户用户 ID | 

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
查询 SettlementAccount 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SettlementServiceApiSettlementServiceRetrieveSettlementAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettlementServiceApiSettlementServiceRetrieveSettlementAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**|  | 
 **object** | **optional.String**| 对象类型 | [default to SettlementAccount]
 **dataId** | **optional.String**| 分账接收方的唯一标识 | [default to 0]
 **dataAppId** | **optional.String**| 分账接收方所在的应用 ID | [default to 0]
 **dataUserId** | **optional.String**| 分账接收方的用户 ID | [default to 0]
 **dataChannel** | **optional.String**| 分账接收方的账户类型   - ALIPAY: 分账到支付宝  - WECHANTPAY: 分账到微信支付  - BANK: 分账到银行卡  - BALANCE: 分账到 justap 账户余额 | [default to CHANNEL_UNKNOWN]
 **dataRecipientWechatpayChannelRecipientAccount** | **optional.String**| openid 或者商户号，由类型决定  微信支付分账接收方账户，OPENID或者商户号 | 
 **dataRecipientWechatpayChannelRecipientName** | **optional.String**| 微信支付分账接收方姓名或名称 | 
 **dataRecipientWechatpayChannelRecipientForceCheck** | **optional.Bool**| 是否强制校验收款人姓名 | [default to false]
 **dataRecipientWechatpayChannelRecipientType** | **optional.String**| 微信支付分账接收方类型 | [default to TYPE_UNSET]
 **dataRecipientWechatpayChannelRecipientAccountType** | **optional.String**| 微信支付分账接收方账户类型   - ACCOUNT_TYPE_UNSET: 未设置  - MERCHANT_ID: 分账到微信商户号  - OPENID: 分账到个人微信号（父公众号的openid，或服务商的openid））  - SUB_OPENID: 分账到个人微信号，子账号的  - LOGIN_NAME: 分账到微信登录号 | [default to ACCOUNT_TYPE_UNSET]
 **dataRecipientWechatpayChannelRecipientAppId** | **optional.String**| 微信支付分账接收方 openid 所对应的公众号 ID | 
 **dataRecipientAlipayChannelRecipientAccount** | **optional.String**| 支付宝账号，账号ID或者登录邮箱 | 
 **dataRecipientAlipayChannelRecipientName** | **optional.String**| 支付宝账号真实姓名 | 
 **dataRecipientAlipayChannelRecipientType** | **optional.String**| 支付宝账号类型 | [default to TYPE_UNSET]
 **dataRecipientAlipayChannelRecipientAccountType** | **optional.String**| 支付宝账号类型   - ACCOUNT_TYPE_UNSET: 未设置  - MERCHANT_ID: 分账到微信商户号  - OPENID: 分账到个人微信号（父公众号的openid，或服务商的openid））  - SUB_OPENID: 分账到个人微信号，子账号的  - LOGIN_NAME: 分账到微信登录号 | [default to ACCOUNT_TYPE_UNSET]
 **dataRecipientBankChannelRecipientAccount** | **optional.String**| 银行卡号 | 
 **dataRecipientBankChannelRecipientName** | **optional.String**| 银行卡开户名 | 
 **dataRecipientBankChannelRecipientType** | **optional.String**| 银行卡类型 | 
 **dataRecipientBankChannelRecipientBankName** | **optional.String**| 银行卡开户行编码 | 
 **dataRecipientBankChannelRecipientBankBranch** | **optional.String**| 银行卡开户支行 | 
 **dataRecipientBankChannelRecipientBankProvince** | **optional.String**| 银行卡开户省份 | 
 **dataRecipientBankChannelRecipientBankCity** | **optional.String**| 银行卡开户城市 | 
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
搜索 SettlementAccount 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettlementServiceApiSettlementServiceSearchSettlementAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettlementServiceApiSettlementServiceSearchSettlementAccountsOpts struct

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
> V1SettlementAccountResponse SettlementServiceUpdateSettlementAccount(ctx, id, optional)
更新 SettlementAccount 对象

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SettlementServiceApiSettlementServiceUpdateSettlementAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettlementServiceApiSettlementServiceUpdateSettlementAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerId** | **optional.String**|  | 
 **userId** | **optional.String**|  | 
 **channel** | **optional.String**|  - ALIPAY: 分账到支付宝  - WECHANTPAY: 分账到微信支付  - BANK: 分账到银行卡  - BALANCE: 分账到 justap 账户余额 | [default to CHANNEL_UNKNOWN]
 **recipientWechatpayChannelRecipientAccount** | **optional.String**| openid 或者商户号，由类型决定  微信支付分账接收方账户，OPENID或者商户号 | 
 **recipientWechatpayChannelRecipientName** | **optional.String**| 微信支付分账接收方姓名或名称 | 
 **recipientWechatpayChannelRecipientForceCheck** | **optional.Bool**| 是否强制校验收款人姓名 | [default to false]
 **recipientWechatpayChannelRecipientType** | **optional.String**| 微信支付分账接收方类型 | [default to TYPE_UNSET]
 **recipientWechatpayChannelRecipientAccountType** | **optional.String**| 微信支付分账接收方账户类型   - ACCOUNT_TYPE_UNSET: 未设置  - MERCHANT_ID: 分账到微信商户号  - OPENID: 分账到个人微信号（父公众号的openid，或服务商的openid））  - SUB_OPENID: 分账到个人微信号，子账号的  - LOGIN_NAME: 分账到微信登录号 | [default to ACCOUNT_TYPE_UNSET]
 **recipientWechatpayChannelRecipientAppId** | **optional.String**| 微信支付分账接收方 openid 所对应的公众号 ID | 
 **recipientAlipayChannelRecipientAccount** | **optional.String**| 支付宝账号，账号ID或者登录邮箱 | 
 **recipientAlipayChannelRecipientName** | **optional.String**| 支付宝账号真实姓名 | 
 **recipientAlipayChannelRecipientType** | **optional.String**| 支付宝账号类型 | [default to TYPE_UNSET]
 **recipientAlipayChannelRecipientAccountType** | **optional.String**| 支付宝账号类型   - ACCOUNT_TYPE_UNSET: 未设置  - MERCHANT_ID: 分账到微信商户号  - OPENID: 分账到个人微信号（父公众号的openid，或服务商的openid））  - SUB_OPENID: 分账到个人微信号，子账号的  - LOGIN_NAME: 分账到微信登录号 | [default to ACCOUNT_TYPE_UNSET]
 **recipientBankChannelRecipientAccount** | **optional.String**| 银行卡号 | 
 **recipientBankChannelRecipientName** | **optional.String**| 银行卡开户名 | 
 **recipientBankChannelRecipientType** | **optional.String**| 银行卡类型 | 
 **recipientBankChannelRecipientBankName** | **optional.String**| 银行卡开户行编码 | 
 **recipientBankChannelRecipientBankBranch** | **optional.String**| 银行卡开户支行 | 
 **recipientBankChannelRecipientBankProvince** | **optional.String**| 银行卡开户省份 | 
 **recipientBankChannelRecipientBankCity** | **optional.String**| 银行卡开户城市 | 

### Return type

[**V1SettlementAccountResponse**](v1SettlementAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

