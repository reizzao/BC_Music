package composicao

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

type OpSentimento string

const (
	Bom    = "alegre, esperança,  "
	NaoBom = "tristeza, saudade, "
)

type Op_Tamanho_Frase string

const (
	Frase_Curta_Definida = "Frase_Curta_Definida"
	Frase_Longa          = "Frase_Longa"
)

type Op_Nome_Cena string

const (
	Cena_A            = "Cena_A - Introdução"
	Cena_B            = "Cena_B - Envolvimento"
	Cena_Ponte_Refrao = "Cena_Ponte_Refrao - Final da Historia"
	Cena_Refrao       = "Cena_Refrao"
)

type Op_CenaRefrao string

const (
	Cena_Refrao_IN         = "Cena_Refrao_IN"
	Cena_Refrao_MeioAcalma = "Cena_Refrao_MeioAcalma"
	Cena_Refrao_Fecha      = "Cena_Refrao_Fecha"
)

type Op_Missao_Cena string

const (
	A_Cena = "Ibtrofucao ao Tema, A_Apresenta_Fato e Responder a essa Apresentação"
)

type Op_Gatilho_ParteFrase string

const (
	GatilhoPergunta = "Fato, Movimento a fazer ou feito, O que esta Acontecendo Agora, Pergunta, Afirmacao,Eu to, eu vou, Atitude Minha reacao ao problema"
	GatilhoResposta = "O que, Do que,"
)

type Op_Conceito string

const (
	Conceito_Pergunta = "frase incompleta conclusão,  comeca com Tonica, "
	Conceito_Resposta = "conclusão da pergunta incompleta, repita a pergunta para dar a resposta, termina em tonica, "
)

type Op_Abertura_UltimaVogal string

const (
	AberturaAberta  = "AberturaAberta [ a, e, i ]"
	AberturaFechada = "AberturaFechada [ o, u, m ]"
)

