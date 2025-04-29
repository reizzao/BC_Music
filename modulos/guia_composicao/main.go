package guia_composicao

import "github.com/musicalrzj/global"

type Guia_Composicao struct {
	Frase_Por_Cena                    Def_Frase_Por_Cena
	Silabas_Poeticas_NotasPorCompasso Def_Silabas_Poeticas
}

type Def_Frase_Por_Cena = global.Numeracao
type Def_Silabas_Poeticas = global.Numeracao
