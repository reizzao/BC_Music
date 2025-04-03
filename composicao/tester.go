package composicao

var requestSinopse = Sinopse{
	SentimentoCentral: Ruim,
}

var fraseUm = Frase{
	Pergunta: TypeFrase{
		Texto:              "pergunta 1",
		ParesNotas:         UmDois,
		Quantidade_Metrica: 2,
		MovimentoMelodico:  Sobe,
	},
	Resposta: TypeFrase{
		Texto:              "resposta 1",
		ParesNotas:         UmDois,
		Quantidade_Metrica: 2,
		MovimentoMelodico:  Desce,
	},
	}

func Tester_CreateComposicao() Composicao{
		newComposicao := Composicao{
		Sinopse: createSinopse(requestSinopse),
		Frase:   createFrase(fraseUm),
	}

	return newComposicao
}
