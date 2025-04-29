package mocks_composicao

import (
	"github.com/musicalrzj/global"
	"github.com/musicalrzj/modulos/composicao/use_composicao"
)

func tester_Use_CreateComposicao() {
	global.Console(use_composicao.CreateComposicao(seedComposicao_01))
}
