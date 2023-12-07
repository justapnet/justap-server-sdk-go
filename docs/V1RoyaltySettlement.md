# V1RoyaltySettlement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | 结算金额 | [default to null]
**AmountCanceled** | **float32** | 结算取消金额 | [default to null]
**AmountFailed** | **float32** | 结算失败金额 | [default to null]
**AmountSucceeded** | **float32** | 结算成功金额 | [default to null]
**AppId** | **string** | 付款方 App ID | [default to null]
**Count** | **int64** | 分账总笔数 | [default to 0]
**CountCanceled** | **int64** | 分账取消笔数 | [default to 0]
**CountFailed** | **int64** | 分账失败笔数 | [default to 0]
**CountSucceeded** | **int64** | 分账成功笔数 | [default to 0]
**Fee** | **float32** | 手续费 | [default to null]
**Id** | **string** | 分账结算单 ID | [default to null]
**Livemode** | **bool** |  | [optional] [default to null]
**Metadata** | **map[string]string** | 元数据 | [default to null]
**Object** | **string** | 对象类型 | [default to null]
**OperationUrl** | **string** | 操作链接 | [default to null]
**Source** | [***V1RoyaltySettlementSource**](v1RoyaltySettlementSource.md) | 分账来源 | [default to null]
**Status** | [***RoyaltySettlementRoyaltySettlementStatus**](RoyaltySettlementRoyaltySettlementStatus.md) | 结算状态 | [default to null]
**TimeFinished** | **int64** | 分账完成时间 | [default to 0]
**Transactions** | [**[]V1RoyaltySettlementTransaction**](v1RoyaltySettlementTransaction.md) | 分账处理流水列表 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


