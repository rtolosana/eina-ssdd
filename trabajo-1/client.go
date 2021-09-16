/*
* AUTOR: Rafael Tolosana Calasanz
* ASIGNATURA: 30221 Sistemas Distribuidos del Grado en Ingeniería Informática
*			Escuela de Ingeniería y Arquitectura - Universidad de Zaragoza
* FECHA: septiembre de 2021
* FICHERO: client.go
* DESCRIPCIÓN: cliente completo para los cuatro escenarios de la práctica 1
*/
package main

import (
    "fmt"
    "time"
    "os"
    "net"
    "./com"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}



func main(){
    endpoint := "155.210.154.200:30000"
    
    // TODO: crear el intervalo solicitando dos números por teclado
    interval := com.TPInterval{1000, 70000}

    tcpAddr, err := net.ResolveTCPAddr("tcp", endpoint)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    // la variable conn es de tipo *net.TCPconn
}
