# V1CreateRefundRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | [OPTIONAL] 退款金额大于 0, 单位为对应币种的最小货币单位，例如：人民币为分（如退款金额为 1 元，此处请填 100）。必须小于等于可退款金额，默认为全额退款。 | [optional] [default to null]
**AppId** | **string** | [REQUIRED] 应用 id | [optional] [default to null]
**ChargeId** | **string** | [REQUIRED] 支付 Charge Id | [optional] [default to null]
**Description** | **string** | [REQUIRED] 退款原因，最多 255 个 Unicode 字符。 | [optional] [default to null]
**Extra** | [***ProtobufAny**](protobufAny.md) | [OPTIONAL] 退款 extra 参数。 | [optional] [default to null]
**MerchantRefundId** | **string** | [REQUIRED] 商户系统的退款单号，必须保证唯一。由于 charge 支持多次退款，对于失败重试动作确保使用相同的订单号，以避免重复退款造成损失。 | [optional] [default to null]
**Metadata** | **map[string]string** | [OPTIONAL] 参考元数据。 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


