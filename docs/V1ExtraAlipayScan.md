# V1ExtraAlipayScan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancePaymentType** | **string** | 预授权类型 | [default to null]
**AgreementParams** | [***V1ExtraAlipayAgreementParams**](v1ExtraAlipayAgreementParams.md) | 协议参数 | [optional] [default to null]
**AuthCode** | **string** | 用户的条码 | [default to null]
**AuthConfirmMode** | **string** | 授权确认方式 | [default to null]
**AuthNo** | **string** | 授权号 | [default to null]
**BuyerId** | **string** | 买家的支付宝用户id | [default to null]
**BuyerLogonId** | **string** | [ONLY IN RESPONSE] 买家支付宝账号 | [default to null]
**BuyerPayAmount** | **float64** | [ONLY IN RESPONSE] 付款金额 | [default to null]
**BuyerUserId** | **string** | [ONLY IN RESPONSE] 买家在支付宝的用户id | [default to null]
**DiscountAmount** | **float64** | [ONLY IN RESPONSE] 商家优惠金额 | [default to null]
**DiscountGoodsDetail** | **string** | [ONLY IN RESPONSE] 商家优惠商品明细 | [default to null]
**DiscountableAmount** | **float64** | 可打折金额 | [default to null]
**ExtendParams** | [***V1ExtraAlipayExtendParams**](v1ExtraAlipayExtendParams.md) | 业务扩展参数 | [optional] [default to null]
**FundBillList** | [***V1ExtraAlipayFundBillList**](v1ExtraAlipayFundBillList.md) | [ONLY IN RESPONSE] 支付金额信息 | [optional] [default to null]
**GmtPayment** | **string** | [ONLY IN RESPONSE] 支付时间 | [default to null]
**GoodsDetail** | [**[]V1ExtraAlipayGoodsDetail**](v1ExtraAlipayGoodsDetail.md) | 商品明细列表 | [optional] [default to null]
**InvoiceAmount** | **float64** | [ONLY IN RESPONSE] 开票金额 | [default to null]
**IsAsyncPay** | [***V1ExtraAlipayPayParams**](v1ExtraAlipayPayParams.md) | 是否异步支付 | [optional] [default to null]
**MdiscountAmount** | **float64** | [ONLY IN RESPONSE] 平台优惠金额 | [default to null]
**OperatorId** | **string** | 商户操作员编号 | [default to null]
**PayParams** | **string** | [ONLY IN RESPONSE] 支付宝返回的支付参数 | [default to null]
**PointAmount** | **float64** | [ONLY IN RESPONSE] 集分宝金额 | [default to null]
**ProductCode** | **string** | 销售产品码 | [default to null]
**QueryOptions** | **string** | 商户授权查询类型 | [default to null]
**ReceiptAmount** | **float64** | [ONLY IN RESPONSE] 实收金额 | [default to null]
**RequestOrgPid** | **string** | 请求方机构id | [default to null]
**Scene** | **string** | 支付场景 | [default to null]
**SettleInfo** | [***V1ExtraAlipaySettleInfo**](v1ExtraAlipaySettleInfo.md) | 结算信息 | [optional] [default to null]
**StoreId** | **string** | 商户门店编号 | [default to null]
**StoreName** | **string** | [ONLY IN RESPONSE] 商户门店名称 | [default to null]
**SubMerchant** | [***V1ExtraAlipaySubMerchant**](v1ExtraAlipaySubMerchant.md) | 子商户信息 | [optional] [default to null]
**TerminalId** | **string** | 商户机具终端编号 | [default to null]
**TotalAmount** | **float64** | [ONLY IN RESPONSE] 订单金额 | [default to null]
**UndiscountableAmount** | **float64** | 不可打折金额 | [default to null]
**VoucherDetailList** | [***V1ExtraAlipayVoucherDetailList**](v1ExtraAlipayVoucherDetailList.md) | [ONLY IN RESPONSE] 商家优惠明细列表 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


