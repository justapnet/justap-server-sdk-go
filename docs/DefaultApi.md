# \DefaultApi

All URIs are relative to *http://127.0.0.1:21011*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TradeServiceCharges**](DefaultApi.md#TradeServiceCharges) | **Post** /transaction/v1/charges | 创建 Charge 对象
[**TradeServiceQueryCharge**](DefaultApi.md#TradeServiceQueryCharge) | **Get** /transaction/v1/charges/{charge_id} | 查询 Charge 对象
[**TradeServiceQueryChargeList**](DefaultApi.md#TradeServiceQueryChargeList) | **Get** /transaction/v1/charges | 查询 Charge 对象列表
[**TradeServiceQueryRefund**](DefaultApi.md#TradeServiceQueryRefund) | **Get** /transaction/v1/charges/{charge_id}/refunds/{refund_id} | 查询 Refund 对象
[**TradeServiceQueryRefundList**](DefaultApi.md#TradeServiceQueryRefundList) | **Get** /transaction/v1/charges/{charge_id}/refunds | 查询 Refund 对象列表
[**TradeServiceRefunds**](DefaultApi.md#TradeServiceRefunds) | **Post** /transaction/v1/refunds | 创建 Refund 对象
[**TradeServiceReverseCharge**](DefaultApi.md#TradeServiceReverseCharge) | **Post** /transaction/v1/charges/{charge_id}/reverse | 撤销 Charge 对象


# **TradeServiceCharges**
> V1ChargeResponse TradeServiceCharges(ctx, body)
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

# **TradeServiceQueryCharge**
> V1ChargeResponse TradeServiceQueryCharge(ctx, chargeId, optional)
查询 Charge 对象

你可以在后台异步通知之前，通过查询接口确认支付状态。通过charge对象的id查询一个已创建的charge对象。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| [REQUIRED] Charge 对象 id | 
 **optional** | ***DefaultApiTradeServiceQueryChargeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTradeServiceQueryChargeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**| [REQUIRED] 应用 id | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeServiceQueryChargeList**
> V1ChargeListResponse TradeServiceQueryChargeList(ctx, optional)
查询 Charge 对象列表

返回之前创建过 charge 对象的一个列表。列表是按创建时间进行排序，总是将最新的 charge 对象显示在最前。如果不设置 created 参数，默认查询近一个月的数据；设置了 created 参数，会按照对应的时间段查询。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiTradeServiceQueryChargeListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTradeServiceQueryChargeListOpts struct

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
 **channel** | **optional.String**| [OPTIONAL] 渠道名称   - BALANCE: 余额  - AlipayQR: 支付宝扫码支付  - AlipayScan: 支付宝条码支付  - AlipayApp: 支付宝 App 支付  - AlipayWap: 支付宝手机网站支付  - AlipayPage: 支付宝电脑网站支付  - AlipayFace: 支付宝刷脸支付  - AlipayLite: 支付宝小程序支付  - WechatpayApp: 微信 App 支付  - WechatpayJSAPI: 微信 JSAPI 支付  - WechatpayH5: 微信 H5 支付  - WechatpayNative: 微信 Native 支付  - WechatpayLite: 微信小程序支付  - WechatpayFace: 刷脸支付  - WechatpayScan: 微信付款码支付 | [default to CHANNEL_INVALID_UNSPECIFIED]
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

# **TradeServiceQueryRefund**
> V1RefundResponse TradeServiceQueryRefund(ctx, chargeId, refundId, optional)
查询 Refund 对象

可以通过 charge 对象的查询接口查询某一个 charge 对象的退款列表，也可以通过 refund 对象的 id 查询一个已创建的 refund 对象。可以在 Webhooks 通知之前，通过查询接口确认退款状态。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| [REQUIRED] 支付 Charge Id | 
  **refundId** | **string**| [REQUIRED] Refund 对象 id | 
 **optional** | ***DefaultApiTradeServiceQueryRefundOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTradeServiceQueryRefundOpts struct

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

# **TradeServiceQueryRefundList**
> V1RefundListResponse TradeServiceQueryRefundList(ctx, chargeId, optional)
查询 Refund 对象列表

返回之前创建 charge 对象的一个 refund 对象列表。列表是按创建时间进行排序，总是将最新的 refund 对象显示在最前。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| [REQUIRED] 支付 Charge Id | 
 **optional** | ***DefaultApiTradeServiceQueryRefundListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTradeServiceQueryRefundListOpts struct

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

# **TradeServiceRefunds**
> V1RefundResponse TradeServiceRefunds(ctx, body)
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

# **TradeServiceReverseCharge**
> V1ChargeResponse TradeServiceReverseCharge(ctx, chargeId, optional)
撤销 Charge 对象

针对已经创建的 Charge，你可以调用撤销接口进行交易的关闭。接口支持对于未成功付款的订单进行撤销，则关闭交易。调用此接口后用户后期不能支付成功。  注：撤销订单在不同收单机构会有不同的行为。对于成功付款的订单请使用 退款 接口进行退款处理。只有针对未支付的订单，我们建议你调用撤销接口。  - 微信支付：如果此订单用户支付失败，微信支付系统会将此订单关闭；如果用户支付成功，微信支付系统会将此订单资金退还给用户。(7天以内的交易单可调用撤销) - 支付宝：如果此订单用户支付失败，支付宝系统会将此订单关闭；如果用户支付成功，支付宝系统会将此订单资金退还给用户。

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargeId** | **string**| Charge 对象 id | 
 **optional** | ***DefaultApiTradeServiceReverseChargeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTradeServiceReverseChargeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **optional.String**| [REQUIRED] 应用 id | 

### Return type

[**V1ChargeResponse**](v1ChargeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

