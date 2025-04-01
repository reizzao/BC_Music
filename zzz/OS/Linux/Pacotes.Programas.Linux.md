# LOJAS DE APPS : RECURSOS - PACOTES - PROGRAMAS

Loja_de_Apps :

  Preferência: instale sempre pacotes Flatpack que são contsiners e  quando ele dar um bug não dá bug no sistema tido só no app
  acesso: https://flathub.org/

  instalador_de_pacotes : .deb: [
  tutoriais : (https://diolinux.com.br/tutoriais/instalar-arquivos-deb-no-ubuntu.html)
  opcao_na linha de comando : ` sudo dpkg -i PACOTE.DEB `
  opcao_com eddy do pop os , obs: mas ta faltando no cosmic beta
  opcao_com gdeb: `sudo apt install gdebi` o gdebi permite clicar com a direita encima do pacote .deb e instalar via interface grafica.

  pacotes_melhor_instalar_deb : [ vscode, ], conceito: instale esses via .deb para nao ter problemas com variaveis de ambiente.

  Outros : [
    Portateis : {
      app: AppImage ,
        Conceito: "AppImage são pacotes portateis, ",
        Instalacao: "Clica com esquerda no arquivo .appimage em Permissoes-> habilite para poder executar este arquivo como programa.
        Usar_Aux_ParaSempre_Abrir: auxiliar App Image Desktop Maker
        Usando_o_Auxiliar:
        ",  :link: https://www.youtube.com/watch?v=lNt6p9EQ7Ik
        pacotes_oficiais: https://appimage.github.io/apps/
        loja_auxiliar_pacotes: https://www.appimagehub.com
    }

    Snap : instalacao : sudo apt install Snap
  ]


instalar suporte flatpack para dustros que ainda não tem : https://youtu.be/sLkvroAgmsU?si=cPzZunfIO3rck-_a

Lista_de_Apps_tutoriais : [
[    ]
Diolinux 2023 - https://youtu.be/Lgg_etqPrOk?si=Y4Pr6WygFueGldM3

úteis 2924 - https://youtu.be/3RYe_WlxOMk?si=cQGxo99f6Ov5MCqq

]

lista de apps úteis:
  app: Steam : ??? acho que é algo de filmes [login: reizao1977 -- pass: Rz..30..# ]
  app: Flatseal ???



Bluetooth : para funcionar fones ou dispositivos bluetooth só pesquisar por bluetooth na loja de app (no mint ja vem instalado o bluetooth-manager )

docker : para usar containners

distrobox : para o sistema ficar sem dependencias desnecessarias.

Gerenciador de discos - particoes : `sudo apt install gnome-disk-utility`  -- apos instalado procure por disks em acessorios

Espelha celular no PC: tuto:
com áudio: https://youtu.be/V0Nb8-JQyzc?si=A5a5lNfvl12uSSAr
sem áudio: https://youtu.be/EQAHZbXzGSM?si=iKFC2_CwnQsk2h6X

 instalar :
 `
 # deixar o celular em modo desenvolvedor > com 7 cliques sobre o desenvolvedor
 # deixar habilitado a opção depuração usb

# instalacao :
 sudo snap install scrcpy
  sudo apt-get install adb # gerenciador de drives pro kenel ao espelhar
  # Observacao: na loja do mint - só procurar por scrcpy e instalar ele ja vai add o adb junto se vc autorizar.

  para rodar e aparecer a tela espelhada do celular digite no terminal :
  scrcpy


 `
