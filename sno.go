package modulbank_go

//Система налогообложения.
type SNO string

const (
	//Общая СН;
	SNOOsn = "osn"
	//Упрощенная СН (доходы);
	SNOUsnIncome = "usn_income"
	//Упрощенная СН (доходы минус расходы);
	SNOUsnIncomeOutcome = "usn_income_outcome"
	//Единый налог на вмененный доход;
	SNOEnvd = "envd"
	//Единый сельскохозяйственный налог;
	SNOEsn = "esn"
	//Патентная СН
	SNOPatent = "patent"
)
