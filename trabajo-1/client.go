/*
* AUTOR: Rafael Tolosana Calasanz
* ASIGNATURA: 30221 Sistemas Distribuidos del Grado en Ingeniería Informática
*			Escuela de Ingeniería y Arquitectura - Universidad de Zaragoza
* FECHA: septiembre de 2021
* FICHERO: client.go
* DESCRIPCIÓN: cliente por completar para el trabajo 1
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
    defer conn.Close()
    
    // TODO completar el código para que el cliente envíe el intervalo al servidor
    // y el servidor de devuelva los primos que hay dentro del intervalo

    // la variable conn es de tipo *net.TCPconn
    // net.TCPconn es un tipo de dato que permite realizar la comunicación TCP / IP full duplex
    // en Golang. Entre otros TCPconn tiene dos métodos que se utilizan tanto por
    // el cliente como por el servidor:
    //
    // Write escribe en el canal TCP de comunicación los datos codificados en b
    // func (c *TCPConn) Write(b []byte) (n int, err os.Error)
    //
    // Read lee del canal los datos y los guarda en b
    // func (c *TCPConn) Read(b []byte) (n int, err os.Error)
    
}
