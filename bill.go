package modulbank_go

type (
	GetBillRequest struct {
		Id            string `json:"id"`
		Merchant      string `json:"merchant"`
		UnixTimestamp int64  `json:"unix_timestamp,omitempty"`
		Salt          string `json:"salt,omitempty"`
		Signature     string `json:"signature,omitempty"`
	}

	CreateBillRequest struct {
		/*
			Идентификатор магазина, который выдается в личном кабинете на этапе интеграции
			Обязательный параметр.
			Строка (максимум 128 символов).
			Допускаются только печатные ASCII­символы. */
		Merchant string `json:"merchant"`

		/*
			Сумма платежа.
			Обязательный параметр.*/
		Amount Money `json:"amount"`

		/*
			Идентификатор заказа, который будет отображаться покупателю.
			Необязательный параметр
			Строка (максимум 50 символов)*/
		CustomOrderId string `json:"custom_order_id"`

		/*
			Описание платежа.
			Обязательный параметр.
			Строка (максимум 250 символов).*/
		Description string `json:"description,omitempty"`

		/*
			Флаг отправки письма с уведомлением о выставленном счете.
			Необязательный параметр.
			Если указать этот  флаг (и client_email), то клиенту будет отправлено письмо, в котором будет ссылка на оплату.*/
		SendLetter IntBool `json:"send_letter,omitempty"`

		/*
			Флаг тестового режима, в котором можно совершать произвольное количество транзакций с использованием тестовых карт.
			Необязательный параметр.
			По умолчанию реальные платежи.*/
		Testing IntBool `json:"testing,omitempty"`

		/*
			Срок актуальности счета в секундах. По истечению счет будет недоступен к оплате.
			Необязательный параметр.
			Целое число
			По умолчанию срок актуальности равен одной неделе*/
		Lifetime uint64 `json:"lifetime,omitempty,omitempty"`

		/*
			Е-mail получателя чека.
			Необязательный параметр
			Если в ЛК включена удаленная регистрация чеков через онлайн-кассу, на этот адрес отправится чек.
			Строка (максимум 64 символа).*/
		ReceiptContact string `json:"receipt_contact,omitempty"`

		/*
			Позиции чека.
			Обязательный параметр, если в ЛК включена удаленная регистрация чеков через онлайн-кассу.*/
		ReceiptItems ReceiptItems `json:"receipt_items"`

		/*
			Текущее время.
			Обязательный параметр.
			Дата и время.
			Формат: UNIX Time.*/
		UnixTimestamp int64 `json:"unix_timestamp,omitempty"`

		/*
			Случайная величина.
			Необязательный параметр.
			Строка (максимум 32 символа)
			Допускаются только печатные
			ASCII­ символы.*/
		Salt string `json:"salt,omitempty"`

		/*
			Криптографическая подпись.
			Обязательный параметр.
			Строка (40 символов в нижнем регистре).*/
		Signature string `json:"signature,omitempty"`
	}

	Bill struct {
		Id       string `json:"id"`
		IsActive bool   `json:"is_active"`
		Paid     int    `json:"paid"`
		Expired  int    `json:"expired"`
		Url      string `json:"url"`
	}

	billResponse struct {
		Bill   Bill   `json:"bill"`
		Status string `json:"status"`
	}
)
