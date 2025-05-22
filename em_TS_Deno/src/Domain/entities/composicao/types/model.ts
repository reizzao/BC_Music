import { IRequestCreateComposicaoDTO } from "@composicao";


class ComposicaoModel {
  public readonly ID: string

  constructor(
    public readonly Request: IRequestCreateComposicaoDTO,
    public readonly Guia: IGuiaComposicaoDTO,
    ID?: string
  ) {
    this.ID = ID || "todo: makeID() - chamar aqui"
  }

}

export { ComposicaoModel, }