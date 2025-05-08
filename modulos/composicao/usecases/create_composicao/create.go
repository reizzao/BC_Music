package create_composicao

// return :: modo de fazer tal coisa
var returnComposicaoFixo FixoComposicao = Objfix_01_Composicao_Fixo

// Function
func CreateComposicao(r RequestComposicao) IComposicao {

	res := IComposicao{
		FixoComposicao:    returnComposicaoFixo,
		RequestComposicao: r,
	}
	return res
}


