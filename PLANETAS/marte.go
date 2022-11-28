package main

import (
	"context"
	"net"
	"strconv"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

var reloj = []int{0, 0, 0}

func BuscarCantidad(Sector string, Base string) string {
	print("Buscando en " + Sector + "---" + Base + "\n")
	return "1111"
}

//Estructura para usar con facilidad el server
type server struct {
	pb.UnimplementedMessageServiceServer
}

//Se maneja el intercambio de mensajes
func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {

	Split_Msj := strings.Split(msg.Body, " ")
	var msn string

	switch Split_Msj[0] {
	case "0":
		//Vanguardia
		print("Peticion Vanguardia" + msg.Body + "\n")
		msn = "Modificacion Hecha"
	case "1":
		//Guardianes
		print("Peticion Guardianes" + msg.Body + "\n")
		Sector := Split_Msj[2]
		Base := Split_Msj[3]
		num := BuscarCantidad(Sector, Base)
		reloj := strconv.Itoa(reloj[0]) + "-" + strconv.Itoa(reloj[1]) + "-" + strconv.Itoa(reloj[2])
		msn = "ServidorMarte " + reloj + " " + num

	case "2":
		//Replicacion

	}

	return &pb.Message{Body: msn}, nil

}

func main() {

	listener, err := net.Listen("tcp", ":50052") //conexion sincrona
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
