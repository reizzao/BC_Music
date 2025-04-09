package composicao

type Request struct {
	Sinopse Sinopse
	Cenas   Def_Cenas
}

type Sinopse struct {
	SentimentoCentral OpSentimento
}

type Def_Cenas struct {
	Fato_Apresentacao IFrase
	Resposta_do_Fato  IFrase
	Info_Cena         Info_Cena
}

type TypeFrase struct {
	Texto                 string
	Conceito              Op_Conceito
	Gatilho_ParteFrase    Op_Gatilho_ParteFrase
	Silabas_Poeticas      uint
	Abertura_Ultima_Vogal Op_Abertura_UltimaVogal
	ParesNotas            Op_ParesNotas
	Quantidade_Metrica    uint
	Movimento_Melodico    Op_Movimento_Melodico
}

type Info_Cena struct {
	Nome_Cena   Op_Nome_Cena
	Missao_Cena Op_Missao_Cena
}