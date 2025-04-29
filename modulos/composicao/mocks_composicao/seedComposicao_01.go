package mocks_composicao

import (
	"github.com/musicalrzj/modulos/composicao"
	u "github.com/musicalrzj/modulos/composicao/use_composicao"
)

var seedComposicao_01 = u.Composicao{
	Sinopse: u.Sinopse{
		SentimentoCentral: composicao.Bom,
	},
	Cenas_Estrofes: u.Props_Cenas_Estrofes{
		A_Apresentacao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
		A_Resolucao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
		B_Apresentacao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
		B_Resolucao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
		R_Apresentacao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
		R_Resolucao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
		E_Apresentacao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
		E_Resolucao: u.Frase{
			Pergunta: "pergunta",
			Resposta: "resposta",
		},
	},
}
