package composicao

import (
	"fmt"
)

var requestSinopse = Sinopse{
	SentimentoCentral: Ruim,
}

var fraseUm = Frase{
	Pergunta: TypeFrase{
		Texto:             "pergunta 1",
		MovimentoMelodico: Sobe,
	},
	Resposta: TypeFrase{
		Texto:             "resposta 1",
		MovimentoMelodico: Desce,
	},
}

func Tester_CreateComposicao() {
	newComposicao := Composicao{
		Sinopse: createSinopse(requestSinopse),
		Frase:   createFrase(fraseUm),
	}

	fmt.Println(newComposicao)
}
