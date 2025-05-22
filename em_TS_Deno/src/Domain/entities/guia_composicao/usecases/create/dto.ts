
interface IRequestCreateGuiaComposicaoDTO{
  Regras: IRegrasGuiaComposicao
}

interface IRegrasGuiaComposicao{
  Tema: string[]
  Evitar: string[]
  UsarDiversos: string[]
  Use_Sete_Sentidos: string[]
  MetricaPadraoNotasPrincipaisPorCompasso: number

}

export type{ IRequestCreateGuiaComposicaoDTO, }