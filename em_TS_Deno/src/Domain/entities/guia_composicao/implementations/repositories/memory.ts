import { IGuiaComposicaoModel, IRequestCreateGuiaComposicaoDTO, IGuiaComposicaoRepository, } from "@guia_composicao";


class MemoryGuiaComposicaoRepository implements IGuiaComposicaoRepository {
  public readonly collection: IGuiaComposicaoModel[] = []

  async save(model: IGuiaComposicaoModel): Promise<IGuiaComposicaoModel> {
    this.collection.push(model)
    return await model
  }

}

export { MemoryGuiaComposicaoRepository }