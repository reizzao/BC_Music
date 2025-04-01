package composicao

type Sinopse struct {
	SentimentoCentral OpSentimento
	
}

type OpSentimento string

const (
	Bom  = "alegre, esperança,  "
	Ruim = "saudade, "
)

func createSinopse (s Sinopse) Sinopse {
	return s
}