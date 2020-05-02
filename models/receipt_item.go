package models

type (
	ReceiptItem struct {
		/*
			Наименование товара.
			Максимум 128 символов. */
		Name string `json:"name"`

		/*
			Количество.
			Максимум 5 символов до запятой и 3 символа после запятой.*/
		Quantity uint32 `json:"quantity"`

		/*
			Цена за единицу товара (рубли, в формате 10.99).
			Максимум 8 символов до запятой и 2 символа после запятой. */
		Price Money `json:"price"`

		/*
			Скидка на позицию (рубли, в формате 5.99).
			Максимум 8 символов до запятой и 2 символа после запятой.
			Используется только в работе через Модулькассу! */
		DiscountSum Money `json:"discount_sum,omitempty"`

		//[{\"name\":\"Test\",\"quantity\":1,\"price\":100.00,\"sno\":\"usn_income\",\"payment_object\":\"service\",\"payment_method\":\"full_prepayment\",\"vat\":\"vat10\"}]
		//[{"name":"Test","quantity":1,"price":100.00,"sno":"usn_income","payment_object":"service","payment_method":"full_prepayment","vat":"vat10"}]
		/*
			Система налогообложения. */
		SNO SNO `json:"sno"`

		/*
			Предмет расчета */
		PaymentObject PaymentObject `json:"payment_object"`

		/*
			Метод платежа */
		PaymentMethod PaymentMethod `json:"payment_method"`

		/*
			Ставка НДС */
		VAT VAT `json:"vat"`
	}
)
