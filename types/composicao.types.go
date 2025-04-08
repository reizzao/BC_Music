package types

type Composicao struct {
	Global  Global_Composicao
	Sinopse Sinopse
	Cenas   Def_Cenas
}

type Def_Cenas struct {
	Fato_Apresentacao Op_Frase
	Resposta_do_Fato  Op_Frase
	Info_Cena         Info_Cena
}

// todo: tem que ser opcoes ou uma frase ou outra em objeto.
type Op_Frase Frase

const (
	Frase_Media = Frase
	Frase_Curta = Frase
)

type Global_Composicao struct {
	Frase_Por_Cena          Def_Frase_Por_Cena
	Silabas_Poeticas_Master Def_Silabas_Poeticas_Master
}

type Frase struct {
	Tamanho_Frase Op_Tamanho_Frase
	Pergunta      TypeFrase
	Resposta      TypeFrase
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

type Sinopse struct {
	SentimentoCentral OpSentimento
}
