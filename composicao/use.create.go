package composicao

type Composicao struct {
	Sinopse Sinopse
	Frase   Frase
}

func CreateComposicao(c Composicao) Composicao {
	newComposicao := Composicao{
		Sinopse: createSinopse(c.Sinopse),
		Frase:   createFrase(c.Frase),
	}

	return newComposicao
}
