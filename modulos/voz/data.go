package voz

type Voz struct {
	Comeco         VozProps
	Meio_Transicao VozProps
	Final          VozProps

	Importantes Importantes_Voz
}

type VozProps struct {
	Estimulo           string
	Abertura_Boca      string
	Tonal              string
	Vogal_Boca         string
	Som                string
	Interpretar_Fingir string
}

type Importantes_Voz struct {
	Projetar      string
	Postura_Peito string
	Gogo          string
	Aspirar       string
 Afinacao      string
}

var Obv_VozComeco = VozProps{
	Estimulo: "
1- Deixa queixo descer , 
2- Contrai Meio Superior do Labio Superior, 
3- Aspira pro Pulmão e Costelas
4- Solta SoproVoz sem apagar a Vela",

	Abertura_Boca:      "Pequena",
	Tonal:              "Tonica",
	Vogal_Boca:         "Â , Todas com acento Circunflexo",
	Som:                "Grave",
	Interpretar_Fingir: "Estar Perguntando ?",
}

var Obv_VozMeio = VozProps{
	Estimulo:           "Contrair Palato Mole - Fundo Boca",
	Abertura_Boca:      "Grande",
	Tonal:              "Atona",
	Vogal_Boca:         "I, Todas com acento Circunflexo",
	Som:                "Agudo",
	Interpretar_Fingir: "Estar Perguntando - Deixa no Ar",
}

var Obv_VozFinal = VozProps{
	Estimulo:           "Como RIR dar Risada, Contrair Palato Duro - Superior Boca",
	Abertura_Boca:      "Media",
	Tonal:              "Tonica",
	Vogal_Boca:         "Ô, Deixa o beiço inferior cair",
	Som:                "Agudo",
	Interpretar_Fingir: "Estar Respondendo, Afirmando !",
}

var Obv_Importantes_Voz = Importantes_Voz{
	Projetar: "como assoprar vela sem sair Ar e Sem apagar a vela.",

	Postura_Peito: "peito pra frente > PEITO DO SUPERMAN, barriga reta, como puxar pernas pra traz",

	Gogo: "baixa a Laringe Gogó.. pensa em falar letra B",

	Aspirar: "aspire já com a boca aberta, direto pro pulmão, 1- Aspire, 2- Segure, 3 solte aos poucos.",

Afinacao: "Respeite intervalos Maiores_Espacosos e NaoMaiores_Curtos ."
}
