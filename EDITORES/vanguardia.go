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

func Revisar(camp string, ip string, solicitud string) string {

	solicitud_lista := strings.Split(solicitud, " ")
	revi := strings.Split(camp, " ")

	if solicitud_lista[1] == revi[0] && solicitud_lista[2] == revi[1] {
		solicitud = "0 " + ip + " " + solicitud
		return solicitud
	}
	solicitud = "0 0 " + solicitud
	return solicitud
}

func main() {

	Camp := "x x"
	ip := "1"

	for {

		connS, err := grpc.Dial("dist041:50051", grpc.WithInsecure())

		if err != nil {
			panic("No se pudo conectar con el servidor" + err.Error())
		}
		defer connS.Close()

		var input string

		input = SolicitarInput()
		println(input)

		input = Revisar(Camp, ip, input)

		lis_inp := strings.Split(input, " ")
		Camp = lis_inp[3] + " " + lis_inp[4]

		serviceCliente := pb.NewMessageServiceClient(connS)

		res, err := serviceCliente.Intercambio(context.Background(),
			&pb.Message{
				Body: input,
			})

		if err != nil {
			panic("No se puede crear el mensaje " + err.Error())
		}

		print(res.Body)
	}

}
