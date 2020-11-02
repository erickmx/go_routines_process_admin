package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	opc := ""
	done := make(chan bool)
	scan := bufio.NewScanner(os.Stdin)
	procecess := []*Process{}
	processAdmin := &ProcessAdmin{
		Procecess: procecess,
	}

	var processIDIncrement uint64
	var processIDString string
	var processID uint64
	var newProcess *Process

	for opc != "4" {
		fmt.Println("Menu principal")
		fmt.Println("1. Agregar proceso")
		fmt.Println("2. Mostrat procesos")
		fmt.Println("3. Eliminar proceso")
		fmt.Println("4. Salir")
		scan.Scan()
		opc = scan.Text()

		switch opc {
		case "1":
			newProcess = NewProcess(processIDIncrement)
			processAdmin.AddProcess(newProcess)
			processIDIncrement += 1
			go newProcess.Start()
			fmt.Println("Proceso creado")
			break
		case "2":
			if processAdmin.ProcessLength != 0 {
				go DisplayProcecesConcurrently(processAdmin, done)
				scan.Scan()
				done <- true
			}
			break
		case "3":
			fmt.Println("Ingrese el ID a eliminar")
			scan.Scan()
			processIDString = scan.Text()
			processID, _ = strconv.ParseUint(processIDString, 10, 64)
			if processAdmin.KillProcess(processID) {
				fmt.Println("Proceso eliminado exitosamente.")
			}
			processID = 0
			break
		case "4":
			processAdmin.KillAllProcecess()
			fmt.Println("Gracias")
			break
		default:
			fmt.Println("Opcion incorrecta")
		}
		fmt.Println()
		fmt.Println("Presione enter para continuar...")
		scan.Scan()
	}
}

func DisplayProcecesConcurrently(processAdmin *ProcessAdmin, done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			processAdmin.ShowProcecess()
			time.Sleep(time.Millisecond * 500)
		}
	}
}
