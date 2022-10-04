# Prácticas de la Asignatura 30221 - Sistemas Distribuidos
## Práctica 1
En términos generales, puede decirse que el problema fundamental de los sistemas distribuidos consiste en asignar tareas de una aplicación distribuida a recursos computacionales. Por recurso computacional, se entiende en esencia, CPU, red de comunicación y almacenamiento. El objetivo es utilizar toda la capacidad computacional de los recursos para poder satisfacer los requisitos de la aplicación. De manera que es fundamental conocer los recursos computacionales para poder construir una arquitectura software distribuida de forma adecuada. En esta práctica vamos a analizar cómo construir arquitecturas cliente servidor y master-worker, para una aplicación muy sencilla que calcula los números primos en un intervalo dado. Para ello, utilizaremos los recursos computacionales, fundamentalmente las CPUs, del Laboratorio L1.02. Crearemos cuatro escenarios, cuatro escenarios, arquitecturas cliente servidor / máster worker. Para ello, en el directorio práctica 1 podéis encontrar los fuentes auxiliares para realizar la práctica. En particular:
- *client.go*: el cliente completo para los cuatro escenarios de la práctica
- *server.go*: el servidor que hay que completar y que será la base para el diseño y la implementación de los cuatro escenarios
- *plot.sh*: script para gnuplot que toma como entrada un fichero output.txt (fichero que recoge la salida del client.go) y muestra gráficamente si se ha violado el QoS.

## Práctica 2
Relojes, tiempo y estado en un sistema distribuido. En esta práctica se va a diseñar e implementar en Go una aplicación de lectores / escritores distribuidos. Deberéis utilizar:
- el modelo actor para la comunicación entre procesos (se proporciona un código auxiliar)
- el Algoritmo de Ricart-Agrawala generalizado
