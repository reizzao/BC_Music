# COMPOSICAO DADOS

Composicao_Dados :
  DirecoesTematicas :
    Tema_Titulo
    Logica_Sentimental - Triste ? NaoTriste
    Emocao_Central
    Destino_da_Solicitacao

  Sinopse :
    O_Grande_Fato_Que_Esta_Acontecendo
    Relatar_O_Que_Eu_Quero_Sobre_O_Grande_Fato_que_Esta_Acontecendo
    O_Que_Houve_Pra_Tudo_Estar_Assim
    O_Que_Precisa_Pra_Melhorar_O_Cenario
    A_Grande_Solucao_Pro_Futuro
    PraQue_da_Grande_Solucao_Pro_Futuro

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


---

## ESTROFES

---

### Estrofe : Sessao_A

Nome_Formal : Verso 1

Conceito : "é a Abordagem,o que da sinais do Tema, o que leva ao tema, sem ficar falando do tema só sinais, mostre ao ouvinte o mesma sensação do Sentimento_Emocao_Central definido. Converso vc e o ouvinte sobre o Sentimento_Emocao_Central,faca-o se sentir como se esse sentimento é pra ele sobre ele,"

Objetivo_Descrever:  a $Emocao_Central

Tempo_Verbal : Presente

Objetivo_Descrever : [ $O_Grande_Fato_Que_Esta_Acontecendo em relacao a $Emocao_Central,]

Musical :
  Andamento : Compassado Lento
  Melodico :


Gatilhos : [
O Grande problema é que,
O que a Personagem é .. cheia de manias...,
falar com alguém pelo fone ,
O que fez ?,
O que faz ?,
Hum ...
Já que ..., Já ...,
É que ...
Porque ...
Contar fatos pra alguém,
Toda vez,
Sempre,
Quando,
Dar ordem que mude algo,
]

FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vA1 , Detalhes_Vogais: alterna inicio Estilo_Vogal_Aberta - termina em Estilo_Vogal_F1
Frase: FF1, Detalhes_Frase: , VogalFinal: vF1 , Detalhes_Vogais:


---

### Estrofe : Sessao_B

Nome_Formal : Verso F1 Hook

Conceito :

Tempo_Verbal : Presente

Objetivo_Descrever : $Relatar_O_Que_Eu_Quero_Sobre_O_Grande_Fato_que_Esta_Acontecendo

Musical :
  Andamento : Baladinha suave de complemneto da parte A,
  Melodico : O mesmo da sessao anterior, pode mudar a ultima nota para passagem pra outra sessao.


Gatilhos : [
O que acho disso?,
]

FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vA1 , Detalhes_Vogais: Auto Rimas curtas
Frase: FF1, Detalhes_Frase: , VogalFinal: vA1 , Detalhes_Vogais: Auto Rimas curtas

---

### Estrofe : Sessao_C

Nome_Formal : Pre-Refrao

Conceito : gera no ouvinte: [ suspense, expectativa que será resolvida com gancho na solução do refrao  ]

Tempo_Verbal : Passado & Presente

Objetivo_Descrever : $O_Que_Houve_Pra_Tudo_Estar_Assim

Musical :
  Andamento :
  Melodico : [ Dramatico, Suspense, notas tensas (menor, com setima, diminuto) oposto das partes anteriores ]


Gatilhos : [
Como ou Porque do EstadoAtual,
o que levou ai Estado Atual , Chamando Atencao,
falando algo do passado,
É claro que...,
Quantas vezes ...
Já que ,

]

FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vF1 || vC , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

Frase: FF1, Detalhes_Frase: , VogalFinal: vF1  || vC , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

---

### Estrofe : Sessao_D

Nome_Formal : Ponte, Ponte-Refrao

Conceito : Frase de impacto que deixa no ar Uma DEIXA pra grande solucao,

Tempo_Verbal : Pro_Futuro / Presente Pensando no Futuro

Objetivo_Descrever : $O_Que_Precisa_Pra_Melhorar_O_Cenario

Musical :
  Andamento :
  Melodico :

Gatilhos : [
Frase Marcante,
O que fazer nesse Presente pensando pról do Futuro

]

FrasesFrase: F1, Detalhes_Frase: , VogalFinal: vF2 , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

Frase: FF1, Detalhes_Frase: , VogalFinal: vF2 , Detalhes_Vogais: Nessa fase usar vogais do mesmo Estilo_Vogal no inicio e termino

---


### Estrofe : Sessao_REFRAO_IN

Nome_Formal :

Conceito : Ouvinte tem que querer contar com 1 dedo levantado pro alto, tem que funcionar sozinho independe outras partes, sem gerar dúvidas. Tem que funcionar como um Loop do Começo combinando com o Fim e voltar em Loop mesmao sentido de historia

Tempo_Verbal :

Objetivo_Descrever : $A_Grande_Solucao_Pro_Futuro

Musical :
  Andamento :
  Melodico : [ martelo,  em loop pra viciar, k final tem que chamar o começode novo pra entrar em loop, ]

Gatilhos : [
  Solução do tema,
  Atitude,
  repetições de Silabas Poéticas,
  Tem (mos) que ?...,
  Precisa (mos) ?...,
  É que ...,
  Um Pedido...,
]

FrasesFrase: F1, Detalhes_Frase: [grito de solucao_atitude], VogalFinal:  vA1, Detalhes_Vogais:

Frase: FF1, Detalhes_Frase: , VogalFinal:  vA1, Detalhes_Vogais:


### Estrofe : Sessao_REFRAO_OUT

Nome_Formal :

Conceito :

Tempo_Verbal :

Objetivo_Descrever : + $PraQue_da_Grande_Solucao_Pro_Futuro

Musical :
  Andamento :
  Melodico :

Gatilhos : [

]

FrasesFrase: F1, Detalhes_Frase: [ baladinha, abaixa a levada, cadência da espaço mais cadenciado pro cantor descansar ], VogalFinal: vF2 , Detalhes_Vogais:

Frase: FF1, Detalhes_Frase: , VogalFinal:  vF2, Detalhes_Vogais:


### Estrofe : Sessao_OUTRO_OPCIONAL

Nome_Formal :

Conceito :

Tempo_Verbal :

Objetivo_Descrever :

Musical :
  Andamento :
  Melodico :
  VogalFinal:

Gatilhos : [

]

Frases
Frase: F1
Detalhes_Frase:  ,  VogalFinal:  , Detalhes_Vogais:

Frase: FF1, Detalhes_Frase: , VogalFinal:  , Detalhes_Vogais:

---

