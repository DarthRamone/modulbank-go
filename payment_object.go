package modulbank_go

//Предмет расчета.
type PaymentObject string

const (
	//Товар
	PaymentObjectCommodity = "commodity"
	//Подакцизный товар
	PaymentObjectExcise = "excise"
	//Работа
	PaymentObjectJob = "job"
	//Услуга
	PaymentObjectService = "service"
	//Ставка азартной игры
	PaymentObjectGamblingBet = "gambling_bet"
	//Выигрыш азартной игры
	PaymentObjectGamblingPrize = "gambling_prize"
	//Лотерейный билет
	PaymentObjectLottery = "lottery"
	//Выигрыш лотереи
	PaymentObjectLotteryPrize = "lottery_prize"
	//Предоставление результатов интеллектуальной деятельности
	PaymentObjectIntellectualActivity = "intellectual_activity"
	//Платеж
	PaymentObjectPayment = "payment"
	//Агентское вознаграждение
	PaymentObjectAgentComission = "agent_commission"
	//Составной предмет расчета
	PaymentObjectComposite = "composite"
	//Другое
	PaymentObjectAnother = "another"
)
