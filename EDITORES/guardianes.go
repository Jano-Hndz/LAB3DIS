package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func SolicitarInput() string {
	fmt.Println("> POR FAVOR ESCRIBIR SOLICITUD : ")
	reader := bufio.NewReader(os.Stdin)
	solicitud, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)

	}

	var solicitud_lista = strings.Split(solicitud, " ")

	//Verificar formato de input sea correcto, primero largo
	if len(solicitud_lista) == 3 && solicitud_lista[0] == "AgregarBase" {
		solicitud = solicitud + " 0"
	}

	return solicitud
}

func main() {
	Guardado := ""

	for {

		connS, err := grpc.Dial("dist041:50051", grpc.WithInsecure())

		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
		defer connS.Close()

		var input string

		input = SolicitarInput()
		println(input)

		serviceCliente := pb.NewMessageServiceClient(connS)

		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: "1 " + input,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}

		print("Respuesta->" + res.Body)
		Guardado = Guardado + res.Body + "\n"
		print(Guardado)

	}
	fmt.Println("Se acabo este programa")

}
