import { IGuiaComposicaoModel, IRequestCreateGuiaComposicaoDTO, IGuiaComposicaoRepository, } from "@guia_composicao";


class JsonGuiaComposicaoRepository implements IGuiaComposicaoRepository {
  public readonly collection: IGuiaComposicaoModel[] = []

  async save(model: IGuiaComposicaoModel): Promise<IGuiaComposicaoModel> {
    this.collection.push(model)
    return await model
  }

}

export { JsonGuiaComposicaoRepository }