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

1- Abaixa Laringe, Quase fale B, Cara de Surpreso, Suspira antes de cantar, 
2- Deixa queixo descer
3- Franzir Superior da Boca
4- canta em Staccato_Arrastado gatilho:: Há Entendi!,
5- Apoio : Barriga contraída aos poucos como se levantasse as pernas jogando corpo pra traz.,

OQueReproduzimos: somente a melodia notas, não dá pra contar acorde, mais que uma nota por vês.

Movimentos : mova a melodia em escala subindo ou descendo

Enfatize : as tônicas nelas se trocam os acordes, abra mais o Queixo nessas trocas, por padrão vem em 2° depois do suave,

Temos: 3 ambientes, a Tonica_Forte Maior e Tonica_Fraca NaoMaior, e Suave Maior ou Menor

Tutoriais: by: PedroHenriqueSiqueira, https://youtu.be/kh5ncfGRouw?si=Kqwz0HnF576pzhnt

3- Fim: falsete bravo bicudo grave ou agudo, enfatizando harmonicos no fim da vogal final.
Tutoriais: "https://www.instagram.com/reel/DJKAFPgxNX3/?igsh=MWFzdHdrZDRiMnhhdA==", 


4- Aspira pro Pulmão e Costelas

	Abertura_Boca:      "Pequena",
	Tonal:              "Tonica",
	Vogal_Boca:         "Â , Todas com acento Circunflexo",
	Som:                "Grave",
	Interpretar_Fingir: "Estar Perguntando ?",
}

var Obv_VozMeio = VozProps{
	Estimulo:           "Acesso: Impulso > Rá te peguei! , Contrair Palato Mole - Fundo Boca",
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

Afinacao: "Respeite intervalos Maiores_Espacosos e NaoMaiores_Curtos, já inicie nas Escalas de altura 4 ."
}
