package main

import (
	"context"
	"net"
	"os"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

//Estructura para usar con facilidad el server
type server struct {
	pb.UnimplementedMessageServiceServer
}

//Se maneja el intercambio de mensajes
func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {

	msn := ""

	Split_Msj := strings.Split(msg.Body, ":")
	//Se maneja peticiones de Rebeldes
	if Split_Msj[0] == "1" {
		msn = RetornarData(Split_Msj[1])
		println("Solicitud de NameNode recibida, mensaje enviado:" + msn)

	}
	//Se maneja peticiones de Combine
	if Split_Msj[0] == "0" {

		data := Split_Msj[1] + ":" + Split_Msj[2] + ":" + Split_Msj[3] + "\n"
		file.WriteString(data)
		msn = "Guardado"

		println("Dato guardado: " + data)

	}
	//Se maneja el cierre del programa
	if msg.Body == "CIERRE" {
		os.Exit(1)
	}

	return &pb.Message{Body: msn}, nil

}

//Main "DateNode Grunt"
func main() {

	listener, err := net.Listen("tcp", ":50051") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()

	defer file.Close()

	for {
		pb.RegisterMessageServiceServer(serv, &server{})
		if err = serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
	}

}
