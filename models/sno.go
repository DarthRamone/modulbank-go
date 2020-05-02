package models

//Система налогообложения.
type SNO string

const (
	//Общая СН;
	OSN = "osn"
	//Упрощенная СН (доходы);
	USN_INCOME = "usn_income"
	//Упрощенная СН (доходы минус расходы);
	USN_INCOME_OUTCOME = "usn_income_outcome"
	//Единый налог на вмененный доход;
	ENVD = "envd"
	//Единый сельскохозяйственный налог;
	ESN = "esn"
	//Патентная СН
	PATENT = "patent"
)
