# V1Refund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [***V1RefundExtra**](v1RefundExtra.md) | 支付渠道退款元参数 | [optional] [default to null]
**Amount** | **float32** | 退款金额 | [default to null]
**ChargeId** | **string** | Charge 对象 id | [default to null]
**ChargeMerchantTradeId** | **string** | 商户系统订单号 | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | 退款创建时间 | [optional] [default to null]
**Description** | **string** | 退款说明 | [default to null]
**FailureCode** | **string** | 支付渠道失败错误码 | [default to null]
**FailureMsg** | **string** | 支付渠道失败原因描述 | [default to null]
**IsSuccess** | **bool** | 退款是否成功 | [default to null]
**Metadata** | **map[string]string** | 元数据，原样返回 | [optional] [default to null]
**RefundId** | **string** | Refund 对象 ID | [default to null]
**RefundNo** | **string** | 退款单号 | [default to null]
**Status** | **string** | 退款状态 | [default to null]
**SuccessAt** | [**time.Time**](time.Time.md) | 退款成功时间 | [optional] [default to null]
**TransactionNo** | **string** | 交易号 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


