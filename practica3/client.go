/*
* AUTOR: Rafael Tolosana Calasanz
* ASIGNATURA: 30221 Sistemas Distribuidos del Grado en Ingeniería Informática
*			Escuela de Ingeniería y Arquitectura - Universidad de Zaragoza
* FECHA: septiembre de 2021
* FICHERO: client.go
* DESCRIPCIÓN: cliente completo para los cuatro escenarios de la práctica 3
*/
package main

import (
    "fmt"
    "time"
    "practica3/com"
    "os"
    "net/rpc"
    "log"
    "math/rand"
    "sync"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// sendRequest realiza una petición RPC al servidor. Cada petición 
// envía únicamente el intervalo en el cual se desea que el servidor encuentre los
// números primos. La invocación RPC devuelve un slice de enteros
// sendRequest escribe por pantalla id_peticion tiempo_observado
func sendRequest(endpoint string, id int, interval com.TPInterval, wg *sync.WaitGroup){
    defer wg.Done()
	start := time.Now()
	client, err := rpc.DialHTTP("tcp", endpoint)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply []int
	err = client.Call("PrimesImpl.FindPrimes", interval, &reply)
	if err != nil {
		log.Fatal("primes error:", err)
	}
	fmt.Println(id, " ", time.Since(start))
}


func main(){
    var tts int
    var wg *sync.WaitGroup = new(sync.WaitGroup)

    if len(os.Args) == 2 {
        endpoint := os.Args[1]
        numIt := 100
        maxIntvl := 70000
        minIntvl := 1000
        maxSegundos := 5000
        minSegundos := 1000
        wg.Add(numIt)
        for i := 1; i <= numIt; i++ {
        	if i%10 == 1 {
			    tts = rand.Intn(maxSegundos-minSegundos) + minSegundos
		    }
		    n := rand.Intn(maxIntvl-minIntvl*2) + minIntvl*2
		    interval := com.TPInterval{minIntvl, n}
            go sendRequest(endpoint, i, interval, wg)
            time.Sleep(time.Duration(tts) * time.Millisecond)
        }
        wg.Wait()
    } else {
        fmt.Println("Usage: go run client.go <ip_server:port>")
    }
}
