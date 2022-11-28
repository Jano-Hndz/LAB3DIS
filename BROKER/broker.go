package main

import (
	"context"
	"math/rand"
	"net"
	"strings"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func EnviarPeticion(ip string, peticion string) string {

	connS, err := grpc.Dial(ip, grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	defer connS.Close()

	serviceCliente := pb.NewMessageServiceClient(connS)

	res, err := serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: peticion,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	return res.Body
}

//Se retorna  un servidor randmo
func ServidorRandom() (Nombre_DateNode string, IP string) {
	rand.Seed(time.Now().UnixNano())
	switch os := rand.Intn(3); os {
	case 0:
		Nombre := "ServidorTierra"
		IP := "dist042:50051"
		return Nombre, IP
	case 1:
		Nombre := "ServidorMarte"
		IP := "dist043:50051"
		return Nombre, IP
	case 3:
		Nombre := "ServidorTitan"
		IP := "dist044:50051"
		return Nombre, IP
	}
	return " ", " "
}

func NombreServer(ip string) string {
	switch ip {
	case "dist042:50051":
		Nombre := "ServidorTierra"
		return Nombre
	case "dist043:50051":
		Nombre := "ServidorMarte"
		return Nombre
	case "dist044:50051":
		Nombre := "ServidorTitan"
		return Nombre
	}
	return " "

}

//Estructura para usar con facilidad el server
type server struct {
	pb.UnimplementedMessageServiceServer
}

//Se maneja el intercambio de mensajes
func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	var msn string
	Split_Msj := strings.Split(msg.Body, " ")

	if Split_Msj[0] == "0" {

		//Peticion Vanguardia
		var nombre string
		var ip string

		if Split_Msj[1] == "0" {
			nombre, ip = ServidorRandom()
		} else {
			ip = Split_Msj[1]
			nombre = NombreServer(ip)
		}

		print(msg.Body + "\n" + ip + "->" + nombre)

		msn := EnviarPeticion(ip, msg.Body)

		print(msn)

	} else {
		//Peticion Guardianes
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
