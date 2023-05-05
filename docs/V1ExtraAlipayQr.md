# V1ExtraAlipayQr

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerId** | **string** | 买家的支付宝唯一用户号（2088开头的16位纯数字） | [optional] [default to null]
**DiscountableAmount** | **string** | 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】，【不可打折金额】则该值默认为【订单总金额】-【不可打折金额】 | [optional] [default to null]
**GoodsDetail** | [**[]V1ExtraAlipayGoodsDetail**](v1ExtraAlipayGoodsDetail.md) | 商品明细列表 | [optional] [default to null]
**OperatorId** | **string** | 商户操作员编号 | [optional] [default to null]
**ProductCode** | **string** | 销售产品码，商家和支付宝签约的产品码，为固定值QUICK_MSECURITY_PAY | [optional] [default to null]
**QrCode** | **string** | [ONLY IN RESPONSE] 二维码 | [optional] [default to null]
**QrCodeTimeoutExpress** | **string** | 支付场景。 条码支付，取值：bar_code； 声波支付，取值：wave_code | [optional] [default to null]
**QrLink** | **string** | [ONLY IN RESPONSE] 二维码图片的URL地址 | [optional] [default to null]
**StoreId** | **string** | 商户门店编号 | [optional] [default to null]
**TerminalId** | **string** | 商户机具终端编号 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


