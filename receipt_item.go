package modulbank_go

type (
	ReceiptItem struct {
		/*
			Наименование товара.
			Максимум 128 символов. */
		Name string `json:"name,omitempty"`

		/*
			Количество.
			Максимум 5 символов до запятой и 3 символа после запятой.*/
		Quantity uint32 `json:"quantity,omitempty"`

		/*
			Цена за единицу товара (рубли, в формате 10.99).
			Максимум 8 символов до запятой и 2 символа после запятой. */
		Price Money `json:"price,omitempty"`

		/*
			Скидка на позицию (рубли, в формате 5.99).
			Максимум 8 символов до запятой и 2 символа после запятой.
			Используется только в работе через Модулькассу! */
		DiscountSum Money `json:"discount_sum,omitempty"`

		/*
			Система налогообложения. */
		SNO SNO `json:"sno,omitempty"`

		/*
			Предмет расчета */
		PaymentObject PaymentObject `json:"payment_object,omitempty"`

		/*
			Метод платежа */
		PaymentMethod PaymentMethod `json:"payment_method,omitempty"`

		/*
			Ставка НДС */
		VAT VAT `json:"vat,omitempty"`
	}
)
