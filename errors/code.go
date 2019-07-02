package errors

const (
	CodeTypeInternalErr        uint32 = 1
	CodeTypeEncodingErr        uint32 = 2
	CodeTypeUnauthorized       uint32 = 3
	CodeTypeUnknownRequest     uint32 = 4
	CodeTypeUnknownAddress     uint32 = 5
	CodeTypeBaseUnknownAddress uint32 = 5 //
	CodeTypeBadNonce           uint32 = 6

	CodeTypeBaseInvalidInput  uint32 = 20
	CodeTypeBaseInvalidOutput uint32 = 21

	CodeLowGasPriceErr        uint32 = 101
	CodeHighGasLimitErr       uint32 = 102
	CodeLowPriceTxCapErr      uint32 = 103
)
