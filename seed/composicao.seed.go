package seed

import (
	"github.com/musicalrzj/types"
)

var RequestSinopse = types.Sinopse{
	SentimentoCentral: types.Ruim,
}

var FraseUm = types.Frase{
	Pergunta: types.TypeFrase{
		Texto:              "pergunta 1",
		ParesNotas:         types.UmDois,
		Quantidade_Metrica: 2,
		MovimentoMelodico:  types.Sobe,
	},
	Resposta: types.TypeFrase{
		Texto:              "resposta 1",
		ParesNotas:         types.UmDois,
		Quantidade_Metrica: 2,
		MovimentoMelodico:  types.Desce,
	},
	}