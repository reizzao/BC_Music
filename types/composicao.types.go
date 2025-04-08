package types

type Composicao struct {
	Global  Global_Composicao
	Sinopse Sinopse
	Cenas Def_Cenas

}

type Def_Cenas struct{
	Fato_Apresentacao  Op_Frase
	Resposta_do_Fato  Op_Frase
	Info_Cena Info_Cena

}

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
	Pergunta TypeFrase
	Resposta TypeFrase
}

type TypeFrase struct {
	Texto                 string
	Conceito              Op_Conceito
	Gatilho_ParteFrase    Op_Gatilho_ParteFrase
	Silabas_Poeticas      Op_Numeracao
	Abertura_Ultima_Vogal Op_Abertura_UltimaVogal
	ParesNotas            Op_ParesNotas
	Quantidade_Metrica    uint
	Movimento_Melodico    Op_Movimento_Melodico
}

type Info_Cena struct {
	Nome_Cena Op_Nome_Cena
	Missao_Cena Op_Missao_Cena
}

type Op_Movimento_Melodico string

const (
	Linear = "Linear"
	Sobe   = "Sobe"
	Desce  = "Desce"
)

type Op_ParesNotas string

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
	NaoBom = "tristeza, saudade, "
)

type Op_Tamanho_Frase string
const(
	Frase_Curta_Definida = "Frase_Curta_Definida"
	Frase_Longa = "Frase_Longa"
)

type Op_Nome_Cena string
const(
	Cena_A = "Cena_A"
	Cena_B = "Cena_B"
	Cena_Ponte_Refrao = "Cena_Ponte_Refrao"
	Cena_Refrao = "Cena_Refrao"
)

type Op_CenaRefrao string
const(
	Cena_Refrao_IN = "Cena_Refrao_IN"
	Cena_Refrao_MeioAcalma = "Cena_Refrao_MeioAcalma"
	Cena_Refrao_Fecha = "Cena_Refrao_Fecha"
)

type Op_Missao_Cena string
const(
	A_Cena = "A_Apresenta_Fato e Responder a essa Apresentação"

)

type Op_Gatilho_ParteFrase string

const (
	GatilhoPergunta = "Fato, Pergunta, Afirmacao,Eu to, eu vou, Atitude Minha reacao ao problema"
	GatilhoResposta = "O que, Do que,"
)

type Op_Conceito string

const (
	Conceito_Pergunta = "frase incompleta conclusão"
	Conceito_Resposta = "conclusão da pergunta incompleta, repita a pergunta para dar a resposta, "
)

type Dois uint = 2
type Cinco uint = 5

type Def_Frase_Por_Cena = Dois
type Def_Silabas_Poeticas_Master = Cinco


type Op_Abertura_UltimaVogal string

const (
	AberturaAberta  = "AberturaAberta [ a, e, i ]"
	AberturaFechada = "AberturaFechada [ o, u, m ]"
)
