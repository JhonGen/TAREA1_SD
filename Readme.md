# Distribucion  y Ejecución de máquinas
Al entrar a cada una de las maquinas, se debe entrar a la carpeta /TAREA1_SD
en la que se encontrarán las carpetas : 
- camiones
- cliente
- logistica
- financiero

en cada máquina, al momento de entrar, puede ejecutar "make run" y correrá un archivo con parametros predeterminados, de lo contrario utilizar lo siguiente

## Maquina 1: Cliente 



se puede escoger entre cliente "pymes" o  "retail" con un tiempo de entrega de solicitud delay a establecer.

- se entra a la carpeta TAREA1_SD/cliente : 

    - go run cliente.go -tipo_cliente=pymes -delay=1
    - go run cliente.go -tipo_cliente=reatil -delay=3


## Maquina 2: Camiones



Se puede escoger entre camion "pymes" o camion "retail"  con un tiempo de entrega a establecer, con numero de camiones 1,2 o 3
- se entra a la carpeta TAREA1_SD/camion :

    - run camion.go -tipo_camion=pymes -tiempo_entrega=2 -delay=2 -nro_camion=1

    - run camion.go -tipo_camion=retail -tiempo_entrega=2 -delay=2 -nro_camion=1

Deja guardado en un txt todas las ordenas que se entregaron y se trataron de entregar

## Maquina 3:Logística
- se entra a la carpeta TAREA1_SD/logistica:
    - go run logistica.go


## Maquina 4:Finanza
- se entra a la carpeta TAREA1_SD/financiero:
    - go run financiero.go 

guarda en un txt el balance final de las entregas


