# V1Charge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | 订单金额 | [default to 0.0]
**AmountFee** | **float32** | 下单金额 | [default to null]
**AmountRefund** | **float32** | 订单退款总金额 | [default to null]
**AmountRoyalty** | **float32** | 分账金额 | [default to null]
**AmountSettle** | **float64** | 结算金额，不一定有，视支付通道情况返回 | [default to null]
**AppId** | **string** | 应用ID | [default to null]
**Body** | **string** | 订单描述信息 | [default to null]
**Channel** | [***V1Channel**](v1Channel.md) |  | [default to null]
**ChargeId** | **string** | Charge 对象 id | [default to null]
**ClientIp** | **string** | 顾客IP | [default to null]
**Closed** | **bool** | 是否关闭 | [default to false]
**ClosedAt** | [**time.Time**](time.Time.md) | 关闭时间 | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Charge 对象创建时间 | [optional] [default to null]
**Credential** | [***map[string]interface{}**](map.md) |  | [optional] [default to null]
**Currency** | **string** | 货币单位，当前仅支持 CNY | [default to null]
**Description** | **string** | 描述信息 | [default to null]
**Extra** | [***V1ChargeExtra**](v1ChargeExtra.md) |  | [optional] [default to null]
**FailureCode** | **string** | 收单机构错误码 | [default to null]
**FailureMsg** | **string** | 收单机构错误描述信息 | [default to null]
**LiveMode** | **bool** | 表明是否是沙箱环境 | [default to false]
**MerchantTradeId** | **string** | 商户系统订单号，APP下需唯一 | [default to null]
**Metadata** | **map[string]string** | 订单元数据，原样返回 | [optional] [default to null]
**Paid** | **bool** | 表明是否已支付 | [default to false]
**PaidAt** | [**time.Time**](time.Time.md) | 支付时间 | [optional] [default to null]
**Refunded** | **bool** | 表明是否包含退款，含退款失败的 | [default to false]
**Refunds** | [**[]V1Refund**](v1Refund.md) | Refund 对象列表 | [optional] [default to null]
**Reversed** | **bool** | 表明是否已经撤销 | [default to false]
**ReversedAt** | [**time.Time**](time.Time.md) | 冲正时间 | [optional] [default to null]
**Subject** | **string** | 订单描述主题 | [default to null]
**TimeExpire** | [**time.Time**](time.Time.md) | 订单过期时间 | [optional] [default to null]
**TransactionNo** | **string** | Charge 的支付单号 | [default to null]
**Ttl** | **int32** | 订单生存时间，单位秒 | [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


