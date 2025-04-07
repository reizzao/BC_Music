package use

import (
	"github.com/musicalrzj/types"
	"github.com/musicalrzj/useprepare"
)




func CreateComposicao(c types.Composicao) types.Composicao {
	newComposicao := types.Composicao{
		Sinopse: useprepare.CreateSinopse(c.Sinopse),
		Frase:   useprepare.CreateFrase(c.Frase),
	}

	return newComposicao
}





