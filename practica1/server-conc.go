/*
* AUTOR: Rafael Tolosana Calasanz
* ASIGNATURA: 30221 Sistemas Distribuidos del Grado en Ingeniería Informática
*			Escuela de Ingeniería y Arquitectura - Universidad de Zaragoza
* FECHA: septiembre de 2021
* FICHERO: server.go
* DESCRIPCIÓN: contiene la funcionalidad esencial para realizar los servidores
*				correspondientes a la práctica 1
 */
package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	//"io"
	"prac1/com"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// PRE: verdad
// POST: IsPrime devuelve verdad si n es primo y falso en caso contrario
func IsPrime(n int) (foundDivisor bool) {
	foundDivisor = false
	if n%2 == 0 {
		return true
	}
	for i := 3; (i < n) && !foundDivisor; i += 2 {
		foundDivisor = (n%i == 0)
	}
	return !foundDivisor
}

// PRE: interval.A < interval.B
// POST: FindPrimes devuelve todos los números primos comprendidos en el
// 		intervalo [interval.A, interval.B]
func FindPrimes(interval com.TPInterval) (primes []int) {
	for i := interval.A; i <= interval.B; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func AnswerRequest(encoder *gob.Encoder, req com.Request) {
	var reply com.Reply
	reply.Id = req.Id
	reply.Primes = FindPrimes(req.Interval)
	encoder.Encode(reply)
}

func main() {

	fmt.Println("hello!\n")

	listener, err := net.Listen("tcp", "127.0.0.1:30000")
	checkError(err)

	conn, err := listener.Accept()
	defer conn.Close()
	checkError(err)
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	var req com.Request
	fmt.Println("Server on\n")
	for {
		err := decoder.Decode(&req)
		checkError(err)
		go AnswerRequest(encoder, req)
	}
}
