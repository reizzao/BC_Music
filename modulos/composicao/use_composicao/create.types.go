package use_composicao

import (
	"github.com/musicalrzj/modulos/composicao"
)

type Composicao struct {
	Sinopse        Sinopse
	Cenas_Estrofes Props_Cenas_Estrofes

}

type Sinopse struct {
	SentimentoCentral composicao.OpSentimento
}

type Frase struct {
	Pergunta string
	Resposta string
}

type Props_Cenas_Estrofes struct {
	A_Apresentacao Frase
	A_Resolucao    Frase
	B_Apresentacao Frase
	B_Resolucao    Frase
	R_Apresentacao Frase
	R_Resolucao    Frase
	E_Apresentacao Frase
	E_Resolucao    Frase
}
