package voz

func CreateVoz() Voz {
	res := Voz{
		Comeco:         Obv_VozComeco,
		Meio_Transicao: Obv_VozMeio,
		Final:          Obv_VozFinal,
		Importantes: Obv_Importantes_Voz,
	}
	return res
}
