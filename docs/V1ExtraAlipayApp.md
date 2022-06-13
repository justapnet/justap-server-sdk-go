# V1ExtraAlipayApp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementSignParams** | [***V1ExtraAlipayAgreementSignParams**](v1ExtraAlipayAgreementSignParams.md) | 签约参数。如果希望在sdk中支付并签约，需要在这里传入签约信息。周期扣款场景 product_code 为 CYCLE_PAY_AUTH 时必填。 | [optional] [default to null]
**CreditAgreementId** | **string** | [ONLY IN RESPONSE] 信用支付协议号 | [default to null]
**CreditBizOrderId** | **string** | [ONLY IN RESPONSE] 信用支付业务订单号 | [default to null]
**CreditPayMode** | **string** | [ONLY IN RESPONSE] 信用支付模式 | [default to null]
**DisablePayChannels** | **string** | 禁用渠道 | [default to null]
**EnablePayChannels** | **string** | 可用渠道 | [default to null]
**ExtUserInfo** | [***V1ExtraAlipayExtUserInfo**](v1ExtraAlipayExtUserInfo.md) | 外部指定买家 | [optional] [default to null]
**ExtendParams** | [***V1ExtraAlipayExtendParams**](v1ExtraAlipayExtendParams.md) | 业务扩展参数 | [optional] [default to null]
**GoodsDetail** | [**[]V1ExtraAlipayGoodsDetail**](v1ExtraAlipayGoodsDetail.md) | 商品明细列表 | [optional] [default to null]
**GoodsType** | **string** | 商品类型 | [default to null]
**MerchantTradeId** | **string** | [ONLY IN RESPONSE] 商户订单号 | [default to null]
**PayParam** | **string** | [ONLY IN RESPONSE] App 用于拉起支付的请求字符串 | [default to null]
**ProductCode** | **string** | 销售产品码，商家和支付宝签约的产品码 | [default to null]
**SellerId** | **string** | [ONLY IN RESPONSE] 支付宝卖家支付宝用户ID | [default to null]
**StoreId** | **string** | 商户门店编号 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


