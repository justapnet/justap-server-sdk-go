# V1ExtraAlipayQr

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerId** | **string** | 买家的支付宝唯一用户号（2088开头的16位纯数字） | [default to null]
**DiscountableAmount** | **string** | 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】，【不可打折金额】则该值默认为【订单总金额】-【不可打折金额】 | [optional] [default to null]
**GoodsDetail** | [**[]V1ExtraAlipayGoodsDetail**](v1ExtraAlipayGoodsDetail.md) | 商品明细列表 | [optional] [default to null]
**OperatorId** | **string** | 商户操作员编号 | [default to null]
**ProductCode** | **string** | 销售产品码，商家和支付宝签约的产品码，为固定值QUICK_MSECURITY_PAY | [default to null]
**QrCode** | **string** | [ONLY IN RESPONSE] 二维码 | [default to null]
**QrCodeTimeoutExpress** | **string** | 支付场景。 条码支付，取值：bar_code； 声波支付，取值：wave_code | [default to null]
**QrLink** | **string** | [ONLY IN RESPONSE] 二维码图片的URL地址 | [default to null]
**QueryOptions** | **string** | 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。注：若为空，则默认为15d。 | [optional] [default to null]
**StoreId** | **string** | 商户门店编号 | [default to null]
**TerminalId** | **string** | 商户机具终端编号 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


