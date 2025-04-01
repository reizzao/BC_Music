package composicao

type Sinopse struct {
	SentimentoCentral OpSentimento
	
}

type OpSentimento string

const (
	Bom  = "alegre, esperança,  "
	Ruim = "tristeza, saudade, "
)

func createSinopse (s Sinopse) Sinopse {
	return s
}