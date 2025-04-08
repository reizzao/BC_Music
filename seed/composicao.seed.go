package seed

import (
	"github.com/musicalrzj/types"
)

var RequestSinopse = types.Sinopse{
	SentimentoCentral: types.NaoBom,
}

var FraseUm = types.Frase{
	Pergunta: types.TypeFrase{
		Texto:              "pergunta 1",
		ParesNotas:         types.UmDois,
		Quantidade_Metrica: 2,
		Movimento_Melodico: types.Sobe,
	},
	Resposta: types.TypeFrase{
		Texto:              "resposta 1",
		ParesNotas:         types.UmDois,
		Quantidade_Metrica: 2,
		Movimento_Melodico: types.Desce,
	},
}
