package composicao

type Frase struct {
	Pergunta TypeFrase
	Resposta TypeFrase

}

type TypeFrase struct{
	Texto string
	ParesNotas OpParesNotas
	Quantidade_Metrica uint
	MovimentoMelodico OpMovimentoMelodico
}

type OpMovimentoMelodico string
const (
	Linear = "Linear"
	Sobe = "Sobe"
	Desce = "Desce"
)

type OpParesNotas string
const (
	UmUmUm = "UmUmUm"
	UmDois = "UmDois"

)



func createFrase (frase Frase) Frase {
	return frase
}