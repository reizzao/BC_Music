import { IRequestCreateGuiaComposicaoDTO } from "@guia_composicao";

interface IGuiaComposicaoModel {
  ID: string
  Request: IRequestCreateGuiaComposicaoDTO
}

class GuiaComposicaoModel implements IGuiaComposicaoModel {
  public readonly ID: string

  constructor(
    public readonly Request: IRequestCreateGuiaComposicaoDTO,
    ID?: string
  ) {
    this.ID = ID || "todo: makeID() - chamar aqui"
  }

}

export { GuiaComposicaoModel, }
export type { IGuiaComposicaoModel, }