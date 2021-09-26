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
	//"io"
	"fmt"
	"prac1/com"
	"time"
)

// PRE: verdad
// POST: IsPrime devuelve verdad si n es primo y falso en caso contrario
func IsPrime(n int) (foundDivisor bool) {
	foundDivisor = false
	if n%2 == 0 {
		return true
	}
	for i := 3; (i < n/2+1) && !foundDivisor; i += 2 {
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

func main() {
	var times = 50
	var totaltime time.Duration = 0
	var start time.Time
	var interval = com.TPInterval{1000, 70000}
	for i := 0; i < times; i++ {
		start = time.Now()
		FindPrimes(interval)
		totaltime += time.Now().Sub(start)
		fmt.Print("*")
	}
	fmt.Println()
	fmt.Println("Executed ", times, "instances in ", totaltime, ", average time ", totaltime.Seconds()/float64(times))
}
