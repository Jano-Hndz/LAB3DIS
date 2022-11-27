package main

import (
	"context"
	"fmt"
	"time"
	"os"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)


//Funcion para comunicarse con el namenode y dependiendo del input de la funcion si es que realiza un fetch o un cierre de los programas.
func Comunicacion(DataTipo string) {
	//Conexion con namenode
	connS, err := grpc.Dial("dist041:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente := pb.NewMessageServiceClient(connS)

	if DataTipo == "CIERRE" {
		//Se realiza interacambio con la palabra CIERRE interpretada por Namenode para enviar a los datanode y cerrarlos
		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: DataTipo,
			})
		
		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}
		fmt.Println(res.Body) //respuesta del namenode
		//Luego se cierra el namenode directamente (esto se hace por separado para poder haber recibido el return del namenode antes de cerrarlo)
		_, _ = serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: "END",
			})
	} else {
		//Envia mensaje a Namenode con el tipo de solicitud
		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				//1 para lectura
				Body: "1:" + DataTipo,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}
		fmt.Println(res.Body) //respuesta del laboratorio
	}

	time.Sleep(2 * time.Second) //espera de 5 segundos
}

func main() {
	flag := true
	//Loop de la interfaz
	for flag {
		MenuInicio := "Rebeldes\n	[ 1 ] Consultar Datos\n	[ 2 ] Cerrar programa\nIngrese su opción:"
		MenuConsulta := "Consulta\n	[ 1 ] Consultar Datos de Logística\n	[ 2 ] Consultar Datos Financieros\n	[ 3 ] Consultar Datos Militares\n	[ 4 ] Volver\nIngrese su opción:"
		fmt.Print(MenuInicio)

		var eleccion string
		fmt.Scanln(&eleccion)
		//Un caso para hacer fetch y otro para cerrar los programas
		switch eleccion {
		case "1":
			//Se usa funcion Comunicacion() con el tipo correspondiente al numero seleccionado
			fmt.Print(MenuConsulta)
			var eleccionC string
			fmt.Scanln(&eleccionC)
			switch eleccionC {
			case "1":
				fmt.Println("Entregando Datos de Logistica:")
				Comunicacion("LOGISTICA")

			case "2":
				fmt.Println("Entregando Datos Financieron")
				Comunicacion("FINANCIERA")

			case "3":
				fmt.Println("Entregando Datos Militares")
				Comunicacion("MILITAR")
			case "4":
				fmt.Println("Volviendo")
			default:
				fmt.Println("Opcion no valida, volviendo al menu de inicio")
			}

		case "2":
			//Se usa la funcion Comunicacion() y luego se hace os.Exit(1) para salir del programa rebelde.go
			Comunicacion("CIERRE")
			os.Exit(1)
		default:
			fmt.Println("Opcion no valida")
		}
	}
	fmt.Println("Se acabo este programa")

}
