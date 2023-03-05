package models

import "time"

type Customer struct {
	Cpf           int
	Private       int
	Incompleto    int
	UltimaCompra  time.Time
	TicketMedio   float64
	TicketUltComp float64
	LojaFrequente int
	LojaUltComp   int
}
