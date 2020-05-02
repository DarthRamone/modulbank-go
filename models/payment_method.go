package models

type PaymentMethod string

const (
	//Предоплата 100%
	FULL_PREPAYMENT = "full_prepayment"
	//Предоплата
	PREPAYMENT = "prepayment"
	//Аванс
	ADVANCE = "advance"
	//Полный расчёт
	FULL_PAYMENT = "full_payment"
	//Частичный расчет и кредит
	PARTIAL_PAYMENT = "partial_payment"
	//Передача в кредит
	CREDIT = "credit"
	//Оплата кредита
	CREDIT_PAYMENT = "credit_payment"
)
