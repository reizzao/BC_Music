# ACOES - COMANDOS

Sistema : {
  conceito: em alguams distros só procurar por "system Info",

informações:
  versao_do_kernel: `uname -r`,
  classica: `sudo lshw -class CPU`
  cpu: `lscpu`,
  dispositivos: `lshw`,
  hardwares_organizada: `lspci`,
  memoria: free -h

 

  tutoriais: [
    sempreupdate - (https://sempreupdate.com.br/linux/comandos/5-maneiras-de-verificar-as-informacoes-da-cpu-no-linux/)

  ]

}

Propagar_configuracoes_no_bash : `source ~/.bashrc` ---> mudei para propag no Personal_Bash

DIRETORIOS_ARQUIVOS
- Remover pasta com tudo dentro : rm -r DIRETORIO_ALVO/

Arquivos :
  abrir_arquivos_para_leitura: cat ARQUIVO.extensao (vai abrir o conteudo do arquivo no terminal)