package composicao

type Frase struct {
	Pergunta TypeFrase
	Resposta TypeFrase

}

type TypeFrase struct{
	Texto string
	MovimentoMelodico OpMovimentoMelodico
}

type OpMovimentoMelodico string
const (
	Linear = "Linear"
	Sobe = "Sobe"
	Desce = "Desce"
)



func createFrase (frase Frase) Frase {
	return frase
}