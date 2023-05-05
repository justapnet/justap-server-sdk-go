# V1CreateChargeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | [REQUIRED] 订单金额，单位元， 如 0.01 | [optional] [default to null]
**AppId** | **string** | [REQUIRED] 应用 id | [optional] [default to null]
**Body** | **string** | [REQUIRED] 服务明细 | [optional] [default to null]
**CallbackUrl** | **string** | [OPTIONAL] 回调地址，如不传则使用 APP 设置中的回调地址。若都为空，则无法跳回原页面 | [optional] [default to null]
**Channel** | [***Tradev1Channel**](tradev1Channel.md) | [REQUIRED] 渠道名称 | [optional] [default to null]
**ClientIp** | **string** | [REQUIRED] 客户端机器 IP | [optional] [default to null]
**Currency** | **string** | 货币单位。国内收单机构仅支持 CNY | [optional] [default to null]
**Description** | **string** | [OPTIONAL] 交易描述 | [optional] [default to null]
**Extra** | [***V1CreateChargeRequestExtra**](v1CreateChargeRequestExtra.md) | [OPTIONAL] 各支付渠道元数据 | [optional] [default to null]
**MerchantTradeId** | **string** | [REQUIRED] 客户系统的交易单号（订单号），必须在应用下唯一。长度不超过30个字符 | [optional] [default to null]
**Metadata** | **map[string]string** | [OPTIONAL] 订单元数据，原样返回 | [optional] [default to null]
**NotificationArea** | **string** | [OPTIONAL] 接受通知服务器所在区域，为确保消息能够送达，请选择服务器所在国家的国家码。如不填默认为 CN | [optional] [default to null]
**NotifyUrl** | **string** | [OPTIONAL] 通知地址，如不传则使用 APP 设置中的通知地址。若都为空，则不发送通知 | [optional] [default to null]
**Subject** | **string** | [REQUIRED] 物品或服务名称（交易标题） | [optional] [default to null]
**Ttl** | **int32** | [OPTIONAL] 订单超时时间，单位秒 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


