package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

//Se retorna un DatoNode con su ip al azar para guardar la data
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
	default:
		Nombre := "ServidorTitan"
		IP := "dist044:50051"
		return Nombre, IP
	}
}

//Se guarda el dato en un DateNode al azar
func GuardarDATA(data string) {

	Split_Msj := strings.Split(data, ":")
	Tipo := Split_Msj[0]
	ID := Split_Msj[1]

	NAMEDATENODE, IPNODE := DateNodeRandom()

	_, err := file.WriteString(Tipo + ":" + ID + ":" + NAMEDATENODE + "\n")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	err = file.Sync()

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	connS, err := grpc.Dial(IPNODE, grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente := pb.NewMessageServiceClient(connS)

	//envia el mensaje al laboratorio
	res, err := serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "0:" + data,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	fmt.Println(res.Body) //respuesta del laboratorio
	time.Sleep(1 * time.Second)

}

//Se consulta a cada DataNode por los datos de un tipo en especifico y retorna un string con todos
func Fetch_Rebeldes(tipo string) string {

	Respuesta := ""

	//CONEXION DATANODE 1
	connS, err := grpc.Dial("dist042:50051", grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	defer connS.Close()

	serviceCliente := pb.NewMessageServiceClient(connS)

	res, err := serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "1:" + tipo,
		})
	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	Respuesta = Respuesta + res.Body

	//CONEXION DATANODE 2
	connS, err = grpc.Dial("dist043:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente = pb.NewMessageServiceClient(connS)

	res, err = serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "1:" + tipo,
		})
	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	Respuesta = Respuesta + res.Body

	//CONEXION DATANODE 3
	connS, err = grpc.Dial("dist044:50051", grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente = pb.NewMessageServiceClient(connS)

	res, err = serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "1:" + tipo,
		})
	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	Respuesta = Respuesta + res.Body

	RetornarString := Ordenar(Respuesta, tipo)

	return RetornarString
}

//Se maneja el cierre de los programas
func Cierre() {

	//CONEXION DATANODE 1
	connS, err := grpc.Dial("dist042:50051", grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}
	defer connS.Close()

	serviceCliente := pb.NewMessageServiceClient(connS)

	_, _ = serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "CIERRE",
		})

	//CONEXION DATANODE 2
	connS, err = grpc.Dial("dist043:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente = pb.NewMessageServiceClient(connS)

	_, _ = serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "CIERRE",
		})

	//CONEXION DATANODE 3
	connS, err = grpc.Dial("dist044:50051", grpc.WithInsecure())
	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente = pb.NewMessageServiceClient(connS)

	_, _ = serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "CIERRE",
		})

}

//Estructura para usar con facilidad el server
type server struct {
	pb.UnimplementedMessageServiceServer
}

//Se maneja el intercambio de mensajes
func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {

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
