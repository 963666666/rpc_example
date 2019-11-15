package main

import (

	"fmt"
	"io/ioutil"

	"log"


)

func main() {
	/*cert, err := tls.LoadX509KeyPair("../../conf/server/server.pem", "../../conf/server/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err %s", err.Error())
	}

	certPool := x509.NewCertPool()*/
	ca, err := ioutil.ReadFile("../../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err %s", err.Error())
	}
	fmt.Println(ca)

	//serve := grpc.NewServer()

}
