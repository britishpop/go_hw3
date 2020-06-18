package card

func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5921": "Рестораны",
		"5411": "Супермаркеты",
		"5912": "Аптеки",
	}
	notCategory := "Категория не указана"
	value, ok := mcc[code]

	if ok == true {
		return value
	}
	return notCategory
}
