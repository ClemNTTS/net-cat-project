package serv

import (
	"app"
	"fmt"
	"log"
	"net"
	"os"
)

var liClients []app.Client
var messagesSlice []app.Message

func StartServ(ip, port string) {
	//set the log registration
	f, err := os.OpenFile("./server/logHistory.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("new session")

	ip, port = manageArgs(ip, port)
	//create a TCP listener
	listener, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	log.Println("Listening on " + ip + ":" + port + " port...")
	fmt.Println("Listening on " + ip + ":" + port + " port...")
	defer listener.Close()

	for {
		//Accepter les connections
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Erreur d'acces au serveur : ", err)
		} else {
			app.PrintPinguin(conn)
			go func() {
				client := app.Client{
					Name:       app.AskName(conn, &liClients),
					Socket:     conn,
					Connected:  true,
					TypingBool: false,
					Messagesch: make(chan string),
				}
				liClients = append(liClients, client)
				go app.HandleConnection(&client, &liClients, &messagesSlice) //go routine to manage each user

			}()
		}
	}
}

func manageArgs(ip, port string) (a string, b string) {
	if ip == "localhost" {
		a = ""
	} else {
		a = ip
	}
	b = port
	return a, b
}
