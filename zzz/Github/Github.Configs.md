
Instalar_extensao: Git Extension Pack

Instalar_o_Git : `sudo apt-get install git`
versao: `git version`

Vincular_VsCode_e_Github :
    primeiro :iniciar git :  `git init`
    add_chave_ssh : `ssh-keygen -t rsa -C "reizao1977@gmail.com"`
    confirmacoes: confirmar onde vai salvar a chave `Pode dar ENTER - VAI CRIAR O ARQUIVO ID_RSA de etxto com a chave ssh em /.ssh/`,
     inventar um password pra chave ex: 77 ,depois:  confirma-lo repetindo a mesma senha criada
    comando_ativar_agente: `eval "$(ssh-agent -s)"`
    abrir_chave: para colar no github : `cat ~/.ssh/id_rsa.pub`  # use o cat para abrir o arquivo de texto no terminal e COPIE todo etxto do arquivo para colar no shh do github online.
    colar_chave_no_github : github.com / settings / ssh keys / add nova chave SSH colando a copiada.

    IMPORTANTE
    confirmar_conexao: `ssh -T git@github.com` // obs:dê `yes` e depois  `add a senha que criou neste processo` assim que pedir o enter passphrase

Configurar_usuario :
git config --global user.email "reizao1977@gmail.com"
git config --global user.name "reizzao"

Clonar_Repos :
Pode clonar via shh ao inves de https : se escolher shh no terminal vai pedir a senha de root que vc configurou nos passos de vinculo ssh.

opcao_vincular_via_https_com_configuracao_token :
no primeiro clone inserir o email de user e em vez da senha colar este configuracao_token:
`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`
como gerar este token : tuto : https://www.alura.com.br/artigos/nova-exigencia-do-git-de-autenticacao-por-token-o-que-e-o-que-devo-fazer

===
