package main

import (
	"context"
	"net"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

//Se crea variable file para que pueda ser accedida desde todo el codigo

//Estructura para usar con facilidad el server
type server struct {
	pb.UnimplementedMessageServiceServer
}

//Se maneja el intercambio de mensajes
func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {

	Split_Msj := strings.Split(msg.Body, " ")
	print(Split_Msj[0])
	var msn string

	switch Split_Msj[0] {
	case "0":
		//Vanguardia
		print(msg.Body)
		msn = "Modificacion Hecha"
	case "1":
		//Guardianes

	case "2":
		//Replicacion

	}

	return &pb.Message{Body: msn}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":50051") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()

	for {
		pb.RegisterMessageServiceServer(serv, &server{})
		if err = serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
	}

}
