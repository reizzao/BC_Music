# Composicao_Tecnicas_Padrao

Composicao_Dados_A_Serem_Coletados :
  DirecoesTematicas :
    Tema_Titulo : [ Acao_1 || Adjetivo , ]
    Logica_Sentimental - Triste ? NaoTriste
    Emocao_Central : Sentimento

  Sinopse :
    O_Grande_Fato_Que_Esta_Acontecendo
    Relatar_O_Que_Eu_Quero_Sobre_O_Grande_Fato_que_Esta_Acontecendo
    O_Que_Houve_Pra_Tudo_Estar_Assim
    O_Que_Precisa_Pra_Melhorar_O_Cenario
    A_Grande_Solucao_Pro_Futuro
    Coro_ResumoDaHistoria_TudoIssoPraDizer_GrandeSolucao_Pro_Futuro

  Vogais :
    Estilo_Vogal_Disponiveis :
      Estilo_Vogal_Aberta = [ A, E ]
      Estilo_Vogal_Fechada = [ I, O, U ]

  Config_VogalFinal :
      vA1 = "Estilo_Vogal_Aberta",
      vF1 = "Estilo_Vogal_Fechada",
      vF2 = "Estilo_Vogal_Fechada[1]"
      vC = "Estilo_Vogal_Aberta"

  Vogais_Escolhidas :
    vA1 = "A",
    vF1 = "I",
    vF2 = "O",
    vC = "A"

# FRASES
Frases :

 Macete_Ouro : A 1° Provoca da Dica e a 2° Revela Resolve a dica dada.

Definicoes_Frases
  FechaInicioSessao : Apontamento_DeixaNoAR
  FechaSessao : Apontamento_Resolve


Evitar : Evite gerundios [ que acabam com indo, amos, cava , exemplos ( ficava, estava, )]

Utilizar : [  Sacada, Frsses do dia a dia facil de decorar  ]

Por_Sessao:
  Padrao_Frase_1: Clima: Clima_Calmo,Conduta: Acao, QualConduta: OQuefoiFeito , Enderecada_A: $PersonagemSessao + (EAI) $FechaInicioSessao
   Padrao_Frase_2: Clima: InversoDoAnterior,Conduta: ReAcao, QualConduta: OQuefoiRecebido , Enderecada_A: $PersonagemInverso + (OQUE) $FechaSessao

Quantidade_Palavras :
  Causar_Descompasso: 4 Silabas Poeticas
  Causar_Viajem: 2 Silabas Poeticas
  Causar_Expectativa: 2 Silabas Poeticas - obs: poucas palavras pausadas ,
  Causar_Ponte: 4 Silabas Poeticas Cromaticas Rapidas
  Martelar_Resolucao_IN: 4 Silabas Poeticas Descompassada
  Martelar_Resolucao_OUT: 4 Silabas Poeticas Cromatica Rapida

Ritmos :
  Ritmo_Descompassado : 1_2
  Ritmo_Viajem : 1_2
  Ritmo_Descompassado_Termino : [ 1_2_34, ]
  Ritmo_Expectativa : [ 1234, ]
  Ritmo_Expectativa_Termino : [ 1_2_34, ]
  Ritmo_RefraoIN_Martelo : 1_2 - 1_2
  Ritmo_RefraoOUT_Martelo_Balada : 1_2 - 1_2 - lento
  Ritmo_Fecha_Cromatico : 1234 Rapido

  Sessoes_All : [ A, B, C, D, R_In, R_Out, OUTRO_OPCIONAL ]

  Personagem : [ Eu, Outro, Todos_Nos ]


---

# MUSICAL

Melodico : a nota demorada ou é no começo ou no fim da frase, quando acontece em uma não acontece na outra.

---

  # DEFINICOES_SESSOES

  Gatilhos_Condutas_Globais :

### Estrofe : Sessao_A

    Sessao : A,
    Conduta_Master: [ Acao_feita pelo_PersonagemdaSessao ]
    Gatilhos_Relacionados_A_Emocao_Central : [    ]
Faça Isso,
Narração de quem está no lugar/situação,
      OQuefoiFeito, O Problema é que,
O que a Personagem é, falar com alguém pelo fone ,
O que fez ?, O que faz ?,Hum ...,Já que , Já ,É que ..., Porque ...,Contar fatos pra alguém, Sempre, Quando,Dar ordem que mude algo, O que estava fazendo,
]
Nome_Formal : Verso 1

Conceito : "é a Abordagem,o que da sinais do Tema, o que leva ao tema, sem ficar falando do tema só sinais, mostre ao ouvinte o mesma sensação do Sentimento_Emocao_Central definido. Converso vc e o ouvinte sobre o Sentimento_Emocao_Central,faca-o se sentir como se esse sentimento é pra ele sobre ele,"

Objetivo_Descrever:  a $Emocao_Central

Tempo_Verbal : Presente

Personagem_da_Sessao:

Objetivo_Descrever : [ $O_Grande_Fato_Que_Esta_Acontecendo em relacao a $Emocao_Central,]

Musical :
  Andamento : Compassado Lento
  Melodico :
  Quantidade_Palavras : $Causar_Descompasso
  Ritmo : $Ritmo_Descompassado

FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vA1 , Detalhes_Vogais: alterna inicio Estilo_Vogal_Aberta - termina em Estilo_Vogal_F1
Frase: FF1, Detalhes_Frase: , VogalFinal: vF1 , Detalhes_Vogais:

---

### Estrofe : Sessao_B
    Sessao : B,
    Conduta_Master: [ Reacao do PersonagemInverso, ]
    Conduta_Opcoes : [
      O que retribuiu fez o outro, Desejo do PersonagemAfetado, O que acho disso,
      ]
    Nome_Formal : Verso F1 Hook
Conceito :
Tempo_Verbal : Presente
Personagem_da_Sessao:
Objetivo_Descrever : $Relatar_O_Que_Eu_Quero_Sobre_O_Grande_Fato_que_Esta_Acontecendo

Musical :
  Andamento : Baladinha suave de complemneto da parte A,
  Melodico : pode mudar a ultima nota para passagem pra outra sessao.
  Quantidade_Palavras : $Causar_Viajem
  Ritmo : Ritmo_Viajem
  obs: ,

FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vA1 , Detalhes_Vogais: Auto Rimas curtas
Frase: FF1, Detalhes_Frase: , VogalFinal: vA1 , Detalhes_Vogais: Auto Rimas curtas

---

### Estrofe : Sessao_C

    Sessao : C,
    Conduta_Master:  Dado aos fatos Passados ,
    Conduta_Opcoes : [
      Como ou Porque do EstadoAtual, Decisao, O que levou ai Estado Atual , Chamando Atencao,É claro que, Quantas vezes, Já que ,
  ]
Nome_Formal : Pre-Refrao
Conceito : gera no ouvinte: [ suspense, expectativa que será resolvida com gancho na solução do refrao  ]

Tempo_Verbal : Passado & Presente
Personagem_da_Sessao:
Objetivo_Descrever : $O_Que_Houve_Pra_Tudo_Estar_Assim

Musical :
  Andamento :
  Melodico :
  Quantidade_Palavras : $Causar_Expectativa
  Ritmo : $Ritmo_Expectativa
  obs: [ Dramatico, Suspense, notas tensas (menor, com setima, diminuto) oposto das partes anteriores ]

FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vF1 || vC , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

Frase: FF1, Detalhes_Frase: , VogalFinal: vF1  || vC , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

---

### Estrofe : Sessao_D

    Sessao : D,
    Conduta_Master: [ Decisao, passo a ser dado pra resultar na GrandeSolucao, ]
    Conduta_Opcoes : [
      Frase Marcante, O que fazer nesse Presente, pensando pról do Futuro,
     ]
Nome_Formal : Ponte, Ponte-Refrao
Conceito : Frase de impacto que deixa no ar Uma DEIXA pra grande solucao,

Tempo_Verbal : Pro_Futuro / Presente Pensando no Futuro
Personagem_da_Sessao:
Objetivo_Descrever : $O_Que_Precisa_Pra_Melhorar_O_Cenario

Musical :
  Andamento :
  Melodico :
  Quantidade_Palavras : $Causar_Ponte
  Ritmo : $Ritmo_Expectativa_Termino


FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vF2 , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

Frase: FF1, Detalhes_Frase: , VogalFinal: vF2 , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

---

### Estrofe : Sessao_REFRAO_IN

    Sessao : R_IN,
    Conduta_Master: [ $Coro_ResumoDaHistoria_TudoIssoPraDizer_GrandeSolucao_Pro_Futuro, ou $Tema, ]
    Conduta_Opcoes : [
      Solução do tema,  Atitude, Repetições de Silabas Poéticas, Tem (mos) que ?, Precisa (mos) ?...,  É que ...,  Um Pedido...,
    ]
Nome_Formal :
Conceito : Ouvinte tem que querer contar com 1 dedo levantado pro alto, tem que funcionar sozinho independe outras partes, sem gerar dúvidas. Tem que funcionar como um Loop do Começo combinando com o Fim e voltar em Loop mesmao sentido de historia

Tempo_Verbal :
Personagem_da_Sessao:
Objetivo_Descrever : $A_Grande_Solucao_Pro_Futuro

Musical :
  Andamento :
  Melodico :
  Quantidade_Palavras : $Martelar_Resolucao_IN
  Ritmo : $Ritmo_RefraoIN_Martelo
  obs: [ martelo,  em loop pra viciar, k final tem que chamar o começode novo pra entrar em loop, ]


FrasesFrase: F1, Detalhes_Frase: [grito de solucao_atitude], VogalFinal:  vA1, Detalhes_Vogais:

Frase: FF1, Detalhes_Frase: , VogalFinal:  vA1, Detalhes_Vogais:

---

### Estrofe : Sessao_REFRAO_OUT
    Sessao : R_OUT,
    Conduta_Master: $PraQue_da_Grande_Solucao_Pro_Futuro ou $Tema,
    Conduta_Opcoes : [   Uma Verdade relacionado a $Emocao_Central, Um desejo relacionado a $Emocao_Central, Declaracao, Sacada, Frsses do dia a dia facil de decorar, ]

Nome_Formal :
Conceito :
Tempo_Verbal :
Personagem_da_Sessao:
Objetivo_Descrever : + $PraQue_da_Grande_Solucao_Pro_Futuro

Musical :
  Andamento :
  Melodico :
  Quantidade_Palavras : $Martelar_Resolucao_OUT
  Ritmo : $Ritmo_RefraoOUT_Martelo_Balada + Ritmo_Fecha_Cromatico


FrasesFrase: F1, Detalhes_Frase: [ baladinha, abaixa a levada, cadência da espaço mais cadenciado pro cantor descansar ], VogalFinal: vF2 , Detalhes_Vogais:

Frase: FF1, Detalhes_Frase: , VogalFinal:  vF2, Detalhes_Vogais:

---

### Estrofe : Sessao_OUTRO_OPCIONAL
Sessao : OUTRO_OPCIONAL,
    Conduta_Master: ,
    Conduta_Opcoes : [ , ]

Nome_Formal :
Conceito :
Tempo_Verbal :
Personagem_da_Sessao:
Objetivo_Descrever :

Musical :
  Andamento :
  Melodico :  Quantidade_Palavras :
  Ritmo :
  VogalFinal:

Frases
Frase: F1
Detalhes_Frase:  ,  VogalFinal:  , Detalhes_Vogais:

Frase: FF1, Detalhes_Frase: , VogalFinal:  , Detalhes_Vogais:

---

# FIM