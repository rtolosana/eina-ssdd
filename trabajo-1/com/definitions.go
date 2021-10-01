/*
* AUTOR: Rafael Tolosana Calasanz
* ASIGNATURA: 30221 Sistemas Distribuidos del Grado en Ingeniería Informática
*			Escuela de Ingeniería y Arquitectura - Universidad de Zaragoza
* FECHA: septiembre de 2021
* FICHERO: definitions.go
* DESCRIPCIÓN: contiene las definiciones de estructuras de datos necesarias para
*			el trabajo 1
*/
package com

import "time"

type TPInterval struct {
    A int
    B int
}

type Request struct {
    Id int		// identificador único de la petición, espera respuesta con el mismo id
    Interval TPInterval
}


type Reply struct {
    Id int		// es el identificador de la respuesta, correspondiente a una petición
    Primes []int
}
