package type_composicao

type IComposicao struct {
	FixoComposicao    FixoComposicao
	RequestComposicao RequestComposicao
}

type FixoComposicao struct {
	Guia_Tema                      string
	Guia_SentimentoCentral_do_Tema string
	Guia_Sentimento_Master         string
}

type RequestComposicao struct {
	Tema                       string
	Sentimento_Central_do_Tema string
	Sentimento_Master          string
}

// type Composicao struct {
// 	Partes                IPartes
// }


// type GuiaeResposta struct{
// 	Fixo_GuiaResposta Fixo_GuiaResposta
// 	Request string

// }

// type Fixo_GuiaResposta struct{
// 	ComoFazer string
// }

// type IPartes struct {
// 	Parte_A        PartesModel
// 	Parte_B        PartesModel
// 	Parte_R_01_IN  PartesModel
// 	Parte_R_01_OUT PartesModel
// 	Parte_C        PartesModel
// }

// type PartesModel struct {
// 	Fixo    PartesFixo
// 	Request PartesRequest
// }

// type PartesFixo struct {
// 	Identificador      string
// 	MissaoDaParte      string
// 	TecnicaDaParte     string
// 	Metrica            string
// 	VelocidadeMelodica string
// 	VocalSom           string
// 	Exemplos           []string
// 	Tutoriais          []string
// }

// type PartesRequest struct {
// 	Frases IFrases
// }

// type IFrases struct {
// 	Frase_01_Apresenta IFrasesModel
// 	Frase_02_Conclui   IFrasesModel
// }

// type IFrasesModel struct {
// 	MissaoFrase string
// 	Inicio      string
// 	Meio        string
// 	Fim         string
// }
