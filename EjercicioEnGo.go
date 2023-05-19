//Ejercicio en Go

//En el que en la consola pida ingresar las horas totales trabajadas por el empleado 
//y pida ingresar el costo  por hora trabajadas , y multiplicado estas dos variables ,
//muestre el pago total para el empleado



package main

import (
	"fmt"
	"log"
)

func main() 
{
	fmt.Print("Ingrese las horas totales trabajadas: ")
	var horasTrabajadas float64
	_, err := fmt.Scan(&horasTrabajadas)
	if err != nil
        {
		log.Fatal("Error al leer las horas trabajadas:", err)
	}

	fmt.Print("Ingrese el costo por hora trabajada: ")
	var costoPorHora float64
	_, err = fmt.Scan(&costoPorHora)
	if err != nil
        {
		log.Fatal("Error al leer el costo por hora:", err)
	}

	pagoTotal := horasTrabajadas * costoPorHora

	fmt.Println("El pago total para el empleado es:", pagoTotal)
}
