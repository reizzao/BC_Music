package guia_composicao

type Guia_Composicao struct {
	Guia_Frase Guia_Frase
	// Frase_Por_Cena                    Def_Frase_Por_Cena
	// Silabas_Poeticas_NotasPorCompasso Def_Silabas_Poeticas
}

// type Def_Frase_Por_Cena = global.Numeracao
// type Def_Silabas_Poeticas = global.Numeracao

type Guia_Frase struct {
	GPergunta PropsPartFrase
	GResposta PropsPartFrase
}

type PropsPartFrase struct {
	AberturaSom   OpAberturaSom
	Interpretacao OpInterpretacao
}

type OpAberturaSom string

const (
	Fechado = "Fechado"
	Aberto  = "Aberto"
)

type OpInterpretacao string

const (
	Perguntando = "Perguntando"
	Explicando  = "Explicando"
)
