package types

type Composicao struct {
	Sinopse Sinopse
	Frase   Frase
}

type Frase struct {
	Pergunta TypeFrase
	Resposta TypeFrase
}

type TypeFrase struct {
	Texto              string
	ParesNotas         OpParesNotas
	Quantidade_Metrica uint
	MovimentoMelodico  OpMovimentoMelodico
}

type OpMovimentoMelodico string

const (
	Linear = "Linear"
	Sobe   = "Sobe"
	Desce  = "Desce"
)

type OpParesNotas string

const (
	UmUmUm = "UmUmUm"
	UmDois = "UmDois"
)

type Sinopse struct {
	SentimentoCentral OpSentimento
}

type OpSentimento string

const (
	Bom  = "alegre, esperança,  "
	Ruim = "tristeza, saudade, "
)
