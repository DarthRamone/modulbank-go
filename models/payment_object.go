package models

//Предмет расчета.
type PaymentObject string

const (
	//Товар
	COMMODITY = "commodity"
	//Подакцизный товар
	EXCISE = "excise"
	//Работа
	JOB = "job"
	//Услуга
	SERVICE = "service"
	//Ставка азартной игры
	GAMBLING_BET = "gambling_bet"
	//Выигрыш азартной игры
	GAMBLING_PRIZE = "gambling_prize"
	//Лотерейный билет
	LOTTERY = "lottery"
	//Выигрыш лотереи
	LOTTERY_PRIZE = "lottery_prize"
	//Предоставление результатов интеллектуальной деятельности
	INTELLECTUAL_ACTIVITY = "intellectual_activity"
	//Платеж
	PAYMENT = "payment"
	//Агентское вознаграждение
	AGENT_COMISSION = "agent_commission"
	//Составной предмет расчета
	COMPOSITE = "composite"
	//Другое
	ANOTHER = "another"
)
