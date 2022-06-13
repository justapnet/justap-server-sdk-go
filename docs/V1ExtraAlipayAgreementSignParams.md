# V1ExtraAlipayAgreementSignParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessParams** | [***ExtraAlipayAgreementSignParamsAccessParams**](ExtraAlipayAgreementSignParamsAccessParams.md) | 请按当前接入的方式进行填充，且输入值必须为文档中的参数取值范围 | [optional] [default to null]
**ExternalAgreementNo** | **string** | 外部协议号 | [default to null]
**ExternalLogonId** | **string** | 外部用户唯一标识 | [default to null]
**PeriodRuleParams** | [***ExtraAlipayAgreementSignParamsPeriodRuleParams**](ExtraAlipayAgreementSignParamsPeriodRuleParams.md) | 周期管控规则参数period_rule_params，在签约周期扣款产品（如CYCLE_PAY_AUTH_P）时必传，在签约其他产品时无需传入。 周期扣款产品，会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制。 | [optional] [default to null]
**PersonalProductCode** | **string** | 个人签约产品码 | [default to null]
**SignNotifyUrl** | **string** | 签约回调地址 | [default to null]
**SignScene** | **string** | 签约场景 | [default to null]
**SubMerchant** | [***ExtraAlipayAgreementSignParamsSubMerchant**](ExtraAlipayAgreementSignParamsSubMerchant.md) | 此参数用于传递子商户信息，无特殊需求时不用关注。目前商户代扣、海外代扣、淘旅行信用住产品支持传入该参数（在销售方案中“是否允许自定义子商户信息”需要选是）。 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


