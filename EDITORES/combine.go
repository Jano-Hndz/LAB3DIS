package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func SolicitarInput() []string {
	for {
		fmt.Println("> POR FAVOR ESCRIBIR SOLICITUD : ")
		var solicitud string
		fmt.Scanln(&solicitud)
		var solicitud_lista = strings.Split(solicitud, ":")

		//Verificar formato de input sea correcto, primero largo
		if len(solicitud_lista) == 3 {
			//Luego si los tipos calzan
			if solicitud_lista[0] == "LOGISTICA" || solicitud_lista[0] == "FINANCIERA" || solicitud_lista[0] == "MILITAR" {
				//Si el ID es un numero
				if _, err := strconv.Atoi(solicitud_lista[1]); err == nil {
					return solicitud_lista
				}
			}
		}

	}
}

func main() {

	for {

		var input []string

		input = SolicitarInput()
		
		//Conexion con namenode
		connS, err := grpc.Dial("dist041:50051", grpc.WithInsecure())

		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}

		defer connS.Close()

		serviceCliente := pb.NewMessageServiceClient(connS)

		//envia el mensaje al Namenode
		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: "0:" + input[0] + ":" + input[1] + ":" + input[2],
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}

		fmt.Println(res.Body) //respuesta de Namenode
	}

}
