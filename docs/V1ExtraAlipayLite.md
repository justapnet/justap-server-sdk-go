# V1ExtraAlipayLite

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | 商品描述 | [optional] [default to null]
**BusinessParams** | [***V1ExtraAlipayBusinessParams**](v1ExtraAlipayBusinessParams.md) | 业务扩展参数 | [optional] [default to null]
**BuyerId** | **string** | 买家的支付宝唯一用户号（2088开头的16位纯数字） | [optional] [default to null]
**DiscountableAmount** | **float64** | 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】，【不可打折金额】则该值默认为【订单总金额】-【不可打折金额】 | [optional] [default to null]
**ExtendParams** | [***V1ExtraAlipayExtendParams**](v1ExtraAlipayExtendParams.md) | 业务扩展参数 | [optional] [default to null]
**LogisticsDetail** | [***V1ExtraAlipayLogisticsDetail**](v1ExtraAlipayLogisticsDetail.md) | 物流信息 | [optional] [default to null]
**OperatorId** | **string** | 商户操作员编号 | [optional] [default to null]
**ProductCode** | **string** | 销售产品码，商家和支付宝签约的产品码，为固定值 FACE_TO_FACE_PAYMENT | [optional] [default to null]
**ReceiverAddressInfo** | [***V1ExtraAlipayReceiverAddressInfo**](v1ExtraAlipayReceiverAddressInfo.md) | 收货信息 | [optional] [default to null]
**SellerId** | **string** | 卖家支付宝用户号 | [optional] [default to null]
**SettleInfo** | [***V1ExtraAlipaySettleInfo**](v1ExtraAlipaySettleInfo.md) | 结算信息 | [optional] [default to null]
**StoreId** | **string** | 商户门店编号 | [optional] [default to null]
**TerminalId** | **string** | 商户机具终端编号 | [optional] [default to null]
**TimeExpire** | **string** | 绝对超时时间，格式为yyyy-MM-dd HH:mm:ss | [optional] [default to null]
**TimeoutExpress** | **string** | 订单有效时间，该时间段内订单可以进行支付，结束后订单将关闭，天数为0表示永久有效 | [optional] [default to null]
**TradeNo** | **string** | [ONLY IN RESPONSE] 支付宝交易号 | [optional] [default to null]
**UndiscountableAmount** | **float64** | 不可打折金额. 不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】,【可打折金额】，则该值默认为【订单总金额】-【可打折金额】 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


