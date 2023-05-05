# V1ExtraAlipayScan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancePaymentType** | **string** | 预授权类型 | [optional] [default to null]
**AuthCode** | **string** | 用户的条码 | [optional] [default to null]
**AuthConfirmMode** | **string** | 授权确认方式 | [optional] [default to null]
**AuthNo** | **string** | 授权号 | [optional] [default to null]
**BuyerId** | **string** | 买家的支付宝用户id | [optional] [default to null]
**BuyerLogonId** | **string** | [ONLY IN RESPONSE] 买家支付宝账号 | [optional] [default to null]
**BuyerPayAmount** | **float64** | [ONLY IN RESPONSE] 付款金额 | [optional] [default to null]
**BuyerUserId** | **string** | [ONLY IN RESPONSE] 买家在支付宝的用户id | [optional] [default to null]
**DiscountAmount** | **float64** | [ONLY IN RESPONSE] 商家优惠金额 | [optional] [default to null]
**DiscountGoodsDetail** | **string** | [ONLY IN RESPONSE] 商家优惠商品明细 | [optional] [default to null]
**DiscountableAmount** | **float64** | 可打折金额 | [optional] [default to null]
**ExtendParams** | [***V1ExtraAlipayExtendParams**](v1ExtraAlipayExtendParams.md) | 业务扩展参数 | [optional] [default to null]
**FundBillList** | [***V1ExtraAlipayFundBillList**](v1ExtraAlipayFundBillList.md) | [ONLY IN RESPONSE] 支付金额信息 | [optional] [default to null]
**GmtPayment** | **string** | [ONLY IN RESPONSE] 支付时间 | [optional] [default to null]
**GoodsDetail** | [**[]V1ExtraAlipayGoodsDetail**](v1ExtraAlipayGoodsDetail.md) | 商品明细列表 | [optional] [default to null]
**InvoiceAmount** | **float64** | [ONLY IN RESPONSE] 开票金额 | [optional] [default to null]
**IsAsyncPay** | [***V1ExtraAlipayPayParams**](v1ExtraAlipayPayParams.md) | 是否异步支付 | [optional] [default to null]
**OperatorId** | **string** | 商户操作员编号 | [optional] [default to null]
**PayParams** | **string** | [ONLY IN RESPONSE] 支付宝返回的支付参数 | [optional] [default to null]
**PointAmount** | **float64** | [ONLY IN RESPONSE] 集分宝金额 | [optional] [default to null]
**ProductCode** | **string** | 销售产品码 | [optional] [default to null]
**QueryOptions** | **string** | 商户授权查询类型 | [optional] [default to null]
**ReceiptAmount** | **float64** | [ONLY IN RESPONSE] 实收金额 | [optional] [default to null]
**RequestOrgPid** | **string** | 请求方机构id | [optional] [default to null]
**Scene** | **string** | 支付场景 | [optional] [default to null]
**StoreId** | **string** | 商户门店编号 | [optional] [default to null]
**StoreName** | **string** | [ONLY IN RESPONSE] 商户门店名称 | [optional] [default to null]
**TerminalId** | **string** | 商户机具终端编号 | [optional] [default to null]
**TotalAmount** | **float64** | [ONLY IN RESPONSE] 订单金额 | [optional] [default to null]
**VoucherDetailList** | [***V1ExtraAlipayVoucherDetailList**](v1ExtraAlipayVoucherDetailList.md) | [ONLY IN RESPONSE] 商家优惠明细列表 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


