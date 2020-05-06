package modulbank_go

type VAT string

const (
	//Без НДС
	VATNone = "none"
	//НДС по ставке 0%
	VAT0 = "vat0"
	//НДС чека по ставке 10%
	VAT10 = "vat10"
	//НДС чека по ставке 20%
	VAT20 = "vat20"
	//НДС чека по расчетной ставке 10%
	VAT110 = "vat110"
	//НДС чека по расчетной ставке 20%
	VAT120 = "vat120"
)
