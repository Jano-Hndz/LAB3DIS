package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func Revisar(camp string, ip string, solicitud string) (string, int) {

	solicitud_lista := strings.Split(solicitud, " ")
	revi := strings.Split(camp, " ")

	if solicitud_lista[1] == revi[0] && solicitud_lista[2] == revi[1] {
		solicitud = "0 " + ip + " " + solicitud
		return solicitud, 0
	}
	solicitud = "0 0 " + solicitud
	return solicitud, 1
}

func main() {

	Camp := "x x"
	ip := "1"

	for {
		var b int
		var input string

		input = SolicitarInput()
		println(input)

		input, b = Revisar(Camp, ip, input)

		lis_inp := strings.Split(input, " ")
		Camp = lis_inp[3] + " " + lis_inp[4]

		if b == 1 {
			_, ipp := ServidorRandom()
			ip = ipp
		}

		println(input)

		println(Camp)

		println(ip)
	}

}
