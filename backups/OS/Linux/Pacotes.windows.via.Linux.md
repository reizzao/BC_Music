# PACOTES WINDOWS

> Capcut
obs: procurar por arquivo .exe pra instaalr senao o bottles nao consegue com o .dmg que vem do site oficial
opcao : mais viavel - pedir aos vendedores de capcut desbloqueado um arquivo instaaldor em .exe
opcao : da web : https://capcut.br.uptodown.com/windows/download




# INSTALADORES

## BOTTLES
Prioridade : 1,
tutoriais: [ (https://www.youtube.com/watch?v=96b1CdB6OHE&t=406s), ]
Instalacao: via Flatpack - procure na loja por Bottles - é o pacote de garrafinhas.
Detalhes: Executores: o Bottles utiliza o Soda por padrão. componentes são libs que o programa windows pode requerer é só procurar em componentes e baixa-la.
Local_arquivos : ~/.var/app/com.usebottles.bottles/



---

## WINE
Prioridade: opcao 2 - dê priopridade ao Bottles.
Instalar: o wine com `sudo apt install wine`
ver_versao_wine : `wine --version`

Instalar_pacotes_de_outro_SistemaOperacional_com_wine :

baixar_instalador:  vc tem que ir no site do pacote em windows e baixar ele pra window , vAÍ onde vc guardou e executa o instalador do programa.exe com o wine

- INSTALAR PROGRAMAS WINDOWS NO LINUX :
abrir_programa_com_wine: VÁ na pasta onde o .exe foi baixado e `wine < NOME DO ARQUIVO.exe >` // VAI criar na area de trabalho o atalho do programa desejado.

instalar : vAÍ abrir instalador igual do windows , então vAÍ no modo next next
onde_sera_instalado: na raiz, na pasta .wine que é o C: do windows no linux.

pos_instalado:  depois de instalado o arquivo.exe que eu tinha baixado no inicio pra instaalr JÁ posso deletar.

---
Local onde achar programas instalados com wine no linux : /home/rzj/.wine/drive_c/Program Files/ NOME_EMPRESA_/PROGRAMA ALVO
