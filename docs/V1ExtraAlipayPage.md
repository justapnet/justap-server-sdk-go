# V1ExtraAlipayPage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementSignParams** | [***V1ExtraAlipayAgreementSignParams**](v1ExtraAlipayAgreementSignParams.md) | 签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。周期扣款场景 product_code 为 CYCLE_PAY_AUTH 时必填。 | [optional] [default to null]
**BusinessParams** | **float32** | 业务扩展参数 | [default to null]
**DisablePayChannels** | **string** | 禁用渠道 | [default to null]
**EnablePayChannels** | **string** | 可用渠道 | [default to null]
**ExtUserInfo** | [***V1ExtraAlipayExtUserInfo**](v1ExtraAlipayExtUserInfo.md) | 支付宝用户信息 | [optional] [default to null]
**ExtendParams** | [***V1ExtraAlipayExtendParams**](v1ExtraAlipayExtendParams.md) | 业务扩展参数 | [optional] [default to null]
**GoodsDetail** | [**[]V1ExtraAlipayGoodsDetail**](v1ExtraAlipayGoodsDetail.md) | 商品明细列表 | [optional] [default to null]
**GoodsType** | **string** | 商品类型 | [default to null]
**IntegrationType** | **float32** | 支付宝用户ID | [default to null]
**InvoiceInfo** | [***V1ExtraAlipayInvoiceInfo**](v1ExtraAlipayInvoiceInfo.md) | 发票信息 | [optional] [default to null]
**MerchantTradeId** | **string** | [ONLY IN RESPONSE] 商户订单号 | [default to null]
**PayUrl** | **string** | [ONLY IN RESPONSE] 支付链接 | [default to null]
**PromoParams** | **string** | 优惠参数 | [default to null]
**QrPayMode** | **string** | 扫码支付模式 | [default to null]
**QrcodeWidth** | **float32** | 二维码宽度 | [default to null]
**RequestFromUrl** | **float32** | 请求来源地址 | [default to null]
**RoyaltyInfo** | [***V1ExtraAlipayRoyaltyInfo**](v1ExtraAlipayRoyaltyInfo.md) | 分账类型卖家的分账类型，目前只支持传入ROYALTY（普通分账类型）。 | [optional] [default to null]
**SellerId** | **string** | [ONLY IN RESPONSE] 收款支付宝用户ID | [default to null]
**SettleInfo** | [***V1ExtraAlipaySettleInfo**](v1ExtraAlipaySettleInfo.md) | 结算信息 | [optional] [default to null]
**StoreId** | **string** | 商户门店编号 | [default to null]
**SubMerchant** | [***V1ExtraAlipaySubMerchant**](v1ExtraAlipaySubMerchant.md) | 二级商户信息 | [optional] [default to null]
**TimeExpire** | **int32** | 订单失效时间 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


