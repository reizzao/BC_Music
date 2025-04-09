package composicao

type Frase struct {
	Tamanho_Frase Op_Tamanho_Frase
	Pergunta      TypeFrase
	Resposta      TypeFrase
}

type IFrase interface {
	NewFrase() Frase
}