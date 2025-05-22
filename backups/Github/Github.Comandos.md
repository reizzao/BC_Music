# Github Comandos

acesso : (https://github.com/joshnh/Git-Commands)

conceitos:  [
  a cada comando/passo comita,
  a cada ação dar um git status para ver a situação atual,
]

Acoes :

Puxar_Alteracoes_Remotas
`
git pull # Atualizar repositório local para o commit mais recente no remoto.
git commit -m "contexto [local] acao" # comentar dar titulo as modificacoes
`

Fluxo_Guardar_Alteracoes
`
git merge ## unir alteraçoes remotas com as locais
git pull ##
git push origin main ## sincroizar >> enviar as mudancas locais para o servidor remoto, especificando a ramificação lembrada (a branch)
git commit -m "contexto [local] acao" ## comentar dar titulo as modificacoes
`
> chamar atualizacao remota e local de uma vez : `git pull && git push && git merge`