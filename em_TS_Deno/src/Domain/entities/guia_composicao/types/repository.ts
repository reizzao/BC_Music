import { IGuiaComposicaoModel } from "@guia_composicao";

interface IGuiaComposicaoRepository {
  collection?: IGuiaComposicaoModel[]
  save(model: IGuiaComposicaoModel): Promise<IGuiaComposicaoModel>
  
}

export type { IGuiaComposicaoRepository, }