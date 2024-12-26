
Meu SDD atual : 120 GB

Meu Pendrive Cinza com Ventoy : 4 Gb

Sistema_Particoes: [
  exFAT : serve para todos sistemas operacionais
  NTFS : serve para windows
]

Para dual Boot
Debian_11 : EspacoMinimo : 50 gb, pendriveMInimo: 4 gb, SistemaDisco: NTFS
Windows_10, EspacoMinimo : 50 gb (Win10) , pendriveMInimo: 8 gb, SistemaDisco: NTFS

hardware

Pendrive :
Acessar_pendrive :
tutorial: https://www.vivaolinux.com.br/topico/Iniciantes-no-Linux/pendrive-no-debian#:~:text=Voc%C3%AA%20pode%20usar%20o%20comando,um%20diret%C3%B3rio%20de%20sua%20escolha.
comando : lsblk

Pendrive Bootavel :
Ventoy: versao win e Linux : baixar: v1.0.91 usado no tuto: https://sourceforge.net/projects/ventoy/files/v1.0.91/ventoy-1.0.91-linux.tar.gz/download  - novas: ja é iso nao sei usar: https://www.ventoy.net/en/download.html

tutorial: ventoy , no souce nao baixe a versao live cd baixe a tar.gz para executaveis
video explicando 20224 : https://www.youtube.com/watch?v=MQ4ZO0dZQiM

Como_usar_ventoy : dentro da pasta com seus arquivos ventoy executa o .sh instalador com -i de instalacao e o endereço onde esta a partição correta mostrada no gerenciador de discos ex: /dev/sdb sem o numero ok. : `sudo sh Ventoy2Disk.sh -i /dev/sdb
`
- vai criar uma pasta do ventoy , para acessarmos o que pode ser colocado dentro em: /media/rzj/Ventoy/
- Depois vai ver a pasta do drive no explore copie e cole nela suas isos.
- Importante só remova o pendrive com injeção segura , para não corrompe-lo.

Drives
Terabox: tenho o terabox com a conta reizao1977 , tutorial : baixar arquivos no terabox: https://www.youtube.com/watch?v=Ba8Gih-y_Mc&t=194s