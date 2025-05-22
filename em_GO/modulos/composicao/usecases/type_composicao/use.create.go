package type_composicao

// return :: modo de fazer tal coisa -> todo: deve ser trocada por repositorio

var returnComposicaoFixo FixoComposicao = Objfix_01_Composicao_Fixo

// Function
func CreateComposicao(r RequestComposicao) IComposicao {

	res := IComposicao{
		FixoComposicao:    returnComposicaoFixo,
		RequestComposicao: r,
	}
	return res
}
