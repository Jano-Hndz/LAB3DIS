package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	println(solicitud)
	solicitud_lista := strings.Split(solicitud, " ")
	revi := strings.Split(camp, " ")

	if solicitud_lista[1] == revi[0] && solicitud_lista[2] == revi[1] {
		solicitud = "0 " + ip + " " + solicitud
		return solicitud
	}
	solicitud = "0 1 " + solicitud
	return solicitud
}

func main() {

	Camp := "x x"
	ip := "1"

	for {

		var input string

		input = SolicitarInput()
		println(input)

		input = Revisar(Camp, ip, input)

		print(input)

	}

}
