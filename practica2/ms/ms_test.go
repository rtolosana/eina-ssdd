/*
* AUTOR: Rafael Tolosana Calasanz
* ASIGNATURA: 30221 Sistemas Distribuidos del Grado en Ingeniería Informática
*			Escuela de Ingeniería y Arquitectura - Universidad de Zaragoza
* FECHA: septiembre de 2021
* FICHERO: ms_test.go
* DESCRIPCIÓN: Implementación de un sistema de mensajería asíncrono, insipirado en el Modelo Actor
*/
package ms

import (
	"testing"
)

type Request struct {
	Id int
}

type Reply struct{
	Response string
}

func TestSendReceiveMessage(t *testing.T) {
	p1 := New(1, "./users.txt", []Message{Request{}, Reply{}})
	p2 := New(2, "./users.txt", []Message{Request{}, Reply{}})
	p1.Send(2, Request{1})
	request := p2.Receive().(Request)
	
	if(request.Id != 1) {
		t.Errorf("P1 envio Request{1}, pero P2 ha recibido::Request{%d}; se esperaba Request{1}", request.Id)
	} else {
		p2.Send(1, Reply{"received"}) 
		msg := p1.Receive().(Reply)
		if msg.Response != "received"{
			t.Errorf("P2 envio Reply{received}, pero P1 ha recibido::Reply{%s}; se esperaba Reply{received}", msg.Response)
		}
	}
}


