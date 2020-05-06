package modulbank_go

type PaymentMethod string

const (
	//Предоплата 100%
	PaymentMethodFullPrepayment = "full_prepayment"
	//Предоплата
	PaymentMethodPrepayment = "prepayment"
	//Аванс
	PaymentMethodAdvance = "advance"
	//Полный расчёт
	PaymentMethodFullPayment = "full_payment"
	//Частичный расчет и кредит
	PaymentMethodPartialPayment = "partial_payment"
	//Передача в кредит
	PaymentMethodCredit = "credit"
	//Оплата кредита
	PaymentMethodCreditPayment = "credit_payment"
)
