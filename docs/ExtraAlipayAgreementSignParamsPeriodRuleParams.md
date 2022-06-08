# ExtraAlipayAgreementSignParamsPeriodRuleParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecuteTime** | **string** | 扣款执行时间execute_time是周期扣款产品必填，枚举值为TIMING和FIXED | [default to null]
**Period** | **string** | 周期数period是周期扣款产品必填。与另一参数period_type组合使用确定扣款周期，例如period_type为DAY，period&#x3D;90，则扣款周期为90天 | [default to null]
**PeriodType** | **string** | 周期类型period_type是周期扣款产品必填，枚举值为DAY和MONTH | [default to null]
**SingleAmount** | **string** | 单笔扣款金额single_amount是周期扣款产品必填，单位为元 | [default to null]
**TotalAmount** | **string** | 扣款总金额total_amount是周期扣款产品必填，单位为元 | [default to null]
**TotalPayments** | **string** | 扣款总笔数total_payments是周期扣款产品必填，单位为笔 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


