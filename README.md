# Go API client for justap
- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Document

[https://www.justap.cn/docs](https://www.justap.cn/docs)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./justap"
```

## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:21011/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**TradeServiceCharges**](docs/DefaultApi.md#tradeservicecharges) | **Post** /transaction/v1/charges | 创建 Charge 对象
*DefaultApi* | [**TradeServiceQueryCharge**](docs/DefaultApi.md#tradeservicequerycharge) | **Get** /transaction/v1/charges/{charge_id} | 查询 Charge 对象
*DefaultApi* | [**TradeServiceQueryChargeList**](docs/DefaultApi.md#tradeservicequerychargelist) | **Get** /transaction/v1/charges | 查询 Charge 对象列表
*DefaultApi* | [**TradeServiceQueryRefund**](docs/DefaultApi.md#tradeservicequeryrefund) | **Get** /transaction/v1/charges/{charge_id}/refunds/{refund_id} | 查询 Refund 对象
*DefaultApi* | [**TradeServiceQueryRefundList**](docs/DefaultApi.md#tradeservicequeryrefundlist) | **Get** /transaction/v1/charges/{charge_id}/refunds | 查询 Refund 对象列表
*DefaultApi* | [**TradeServiceRefunds**](docs/DefaultApi.md#tradeservicerefunds) | **Post** /transaction/v1/refunds | 创建 Refund 对象
*DefaultApi* | [**TradeServiceReverseCharge**](docs/DefaultApi.md#tradeservicereversecharge) | **Post** /transaction/v1/charges/{charge_id}/reverse | 撤销 Charge 对象


## Documentation For Models

 - [ExtraAlipayAgreementSignParamsAccessParams](docs/ExtraAlipayAgreementSignParamsAccessParams.md)
 - [ExtraAlipayAgreementSignParamsPeriodRuleParams](docs/ExtraAlipayAgreementSignParamsPeriodRuleParams.md)
 - [ExtraAlipayAgreementSignParamsSubMerchant](docs/ExtraAlipayAgreementSignParamsSubMerchant.md)
 - [ExtraAlipayInvoiceInfoKeyInfo](docs/ExtraAlipayInvoiceInfoKeyInfo.md)
 - [ExtraAlipayRoyaltyInfoRoyaltyDetailInfos](docs/ExtraAlipayRoyaltyInfoRoyaltyDetailInfos.md)
 - [ExtraAlipaySettleInfoSettleDetailInfos](docs/ExtraAlipaySettleInfoSettleDetailInfos.md)
 - [ExtraWechatpayDetailGoodsDetail](docs/ExtraWechatpayDetailGoodsDetail.md)
 - [ExtraWechatpaySceneInfoH5Info](docs/ExtraWechatpaySceneInfoH5Info.md)
 - [ExtraWechatpaySceneInfoStoreInfo](docs/ExtraWechatpaySceneInfoStoreInfo.md)
 - [OpenApiRoyaltyDetailInfoPojoTradeFundBillItem](docs/OpenApiRoyaltyDetailInfoPojoTradeFundBillItem.md)
 - [QueryChargeListRequestCreated](docs/QueryChargeListRequestCreated.md)
 - [RefundExtraAlipayOpenApiRoyaltyDetailInfoPojo](docs/RefundExtraAlipayOpenApiRoyaltyDetailInfoPojo.md)
 - [RefundExtraWechatPayAccount](docs/RefundExtraWechatPayAccount.md)
 - [RefundExtraWechatPayGoodsDetailItem](docs/RefundExtraWechatPayGoodsDetailItem.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [V1AlipayCallbackResponse](docs/V1AlipayCallbackResponse.md)
 - [V1AlipayNotifyResponse](docs/V1AlipayNotifyResponse.md)
 - [V1Channel](docs/V1Channel.md)
 - [V1Charge](docs/V1Charge.md)
 - [V1ChargeExtra](docs/V1ChargeExtra.md)
 - [V1ChargeListResponse](docs/V1ChargeListResponse.md)
 - [V1ChargeResponse](docs/V1ChargeResponse.md)
 - [V1CreateChargeRequest](docs/V1CreateChargeRequest.md)
 - [V1CreateChargeRequestExtra](docs/V1CreateChargeRequestExtra.md)
 - [V1CreateRefundRequest](docs/V1CreateRefundRequest.md)
 - [V1ExtraAlipayAgreementParams](docs/V1ExtraAlipayAgreementParams.md)
 - [V1ExtraAlipayAgreementSignParams](docs/V1ExtraAlipayAgreementSignParams.md)
 - [V1ExtraAlipayApp](docs/V1ExtraAlipayApp.md)
 - [V1ExtraAlipayBusinessParams](docs/V1ExtraAlipayBusinessParams.md)
 - [V1ExtraAlipayExtUserInfo](docs/V1ExtraAlipayExtUserInfo.md)
 - [V1ExtraAlipayExtendParams](docs/V1ExtraAlipayExtendParams.md)
 - [V1ExtraAlipayFace](docs/V1ExtraAlipayFace.md)
 - [V1ExtraAlipayFundBillList](docs/V1ExtraAlipayFundBillList.md)
 - [V1ExtraAlipayGoodsDetail](docs/V1ExtraAlipayGoodsDetail.md)
 - [V1ExtraAlipayInvoiceInfo](docs/V1ExtraAlipayInvoiceInfo.md)
 - [V1ExtraAlipayLite](docs/V1ExtraAlipayLite.md)
 - [V1ExtraAlipayLogisticsDetail](docs/V1ExtraAlipayLogisticsDetail.md)
 - [V1ExtraAlipayPage](docs/V1ExtraAlipayPage.md)
 - [V1ExtraAlipayPayParams](docs/V1ExtraAlipayPayParams.md)
 - [V1ExtraAlipayQr](docs/V1ExtraAlipayQr.md)
 - [V1ExtraAlipayReceiverAddressInfo](docs/V1ExtraAlipayReceiverAddressInfo.md)
 - [V1ExtraAlipayRoyaltyInfo](docs/V1ExtraAlipayRoyaltyInfo.md)
 - [V1ExtraAlipayScan](docs/V1ExtraAlipayScan.md)
 - [V1ExtraAlipaySettleInfo](docs/V1ExtraAlipaySettleInfo.md)
 - [V1ExtraAlipaySubMerchant](docs/V1ExtraAlipaySubMerchant.md)
 - [V1ExtraAlipayVoucherDetailList](docs/V1ExtraAlipayVoucherDetailList.md)
 - [V1ExtraAlipayWap](docs/V1ExtraAlipayWap.md)
 - [V1ExtraWechatpayApp](docs/V1ExtraWechatpayApp.md)
 - [V1ExtraWechatpayAppConfig](docs/V1ExtraWechatpayAppConfig.md)
 - [V1ExtraWechatpayAppletConfig](docs/V1ExtraWechatpayAppletConfig.md)
 - [V1ExtraWechatpayDetail](docs/V1ExtraWechatpayDetail.md)
 - [V1ExtraWechatpayH5](docs/V1ExtraWechatpayH5.md)
 - [V1ExtraWechatpayJsapi](docs/V1ExtraWechatpayJsapi.md)
 - [V1ExtraWechatpayJsapiConfig](docs/V1ExtraWechatpayJsapiConfig.md)
 - [V1ExtraWechatpayLite](docs/V1ExtraWechatpayLite.md)
 - [V1ExtraWechatpayNative](docs/V1ExtraWechatpayNative.md)
 - [V1ExtraWechatpayPayer](docs/V1ExtraWechatpayPayer.md)
 - [V1ExtraWechatpayScan](docs/V1ExtraWechatpayScan.md)
 - [V1ExtraWechatpaySceneInfo](docs/V1ExtraWechatpaySceneInfo.md)
 - [V1ExtraWechatpaySettleInfo](docs/V1ExtraWechatpaySettleInfo.md)
 - [V1Refund](docs/V1Refund.md)
 - [V1RefundExtra](docs/V1RefundExtra.md)
 - [V1RefundExtraAlipay](docs/V1RefundExtraAlipay.md)
 - [V1RefundExtraWechatPay](docs/V1RefundExtraWechatPay.md)
 - [V1RefundListResponse](docs/V1RefundListResponse.md)
 - [V1RefundResponse](docs/V1RefundResponse.md)
 - [V1WechatpayCallbackResponse](docs/V1WechatpayCallbackResponse.md)
 - [V1WechatpayNotifyResponse](docs/V1WechatpayNotifyResponse.md)


## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

support@justap.net

