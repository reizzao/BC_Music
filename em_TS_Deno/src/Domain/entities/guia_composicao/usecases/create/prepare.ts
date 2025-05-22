import { IRequestCreateGuiaComposicaoDTO, IGuiaComposicaoRepository, GuiaComposicaoModel } from "@guia_composicao";

class CreateGuiaComposicaoPrepare {

  constructor(
    private readonly repo: IGuiaComposicaoRepository,
  ) { }

  async execute(requestDTO: IRequestCreateGuiaComposicaoDTO) {
    const model = new GuiaComposicaoModel(requestDTO)
    const saved = await this.repo.save(model)

    return await saved
  }
  
}
  export { CreateGuiaComposicaoPrepare }