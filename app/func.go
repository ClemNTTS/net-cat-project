package app

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Message struct {
	Sender  Client
	Content string
	Time    string
}

type Client struct {
	Socket     net.Conn
	Name       string
	Connected  bool
	Typing     chan bool
	TypingBool bool
	Messagesch chan string
	Hist       int
}

// main part of the programm process, manage messages
func HandleConnection(client *Client, clientsSlice *[]Client, messageList *[]Message) {
	defer func() {
		client.Socket.Close()
	}()
	conn := client.Socket

	//manage the number of users
	if len(*clientsSlice) > 10 {
		conn.Write([]byte("To much users, retry later\n"))
		return
	}

	//Notice other users and serv
	log.Println("Nouvelle connection de ", client.Name)
	fmt.Println("Nouvelle connection de ", client.Name)
	broadCast(*clientsSlice, Message{Content: "\r" + time.Now().Format("15:04:05") + " " + client.Name + " s'est connecté\n", Sender: *client, Time: time.Now().Format("15:04:05")})

	//say hello and print History
	conn.Write([]byte("Bonjour " + client.Name + " !\n\n"))
	PrintMessages(messageList, client)

	//send user message as he's connected
	for client.Connected {
		TakeInput(client, clientsSlice, messageList)
	}
}

// Manage the input of the user, and send them to server and other users
func TakeInput(client *Client, clientsSlice *[]Client, messageList *[]Message) {
	reader := bufio.NewReader(client.Socket)

	client.Socket.SetReadDeadline(time.Now().Add(time.Minute * 2))
	mssg, err := reader.ReadString('\n')
	if err != nil {
		broadCast(*clientsSlice, Message{Content: client.Name + " s'est déconnecté", Sender: *client, Time: time.Now().Format("15:04:05")})
		log.Println(client.Name + " s'est déconnecté")
		fmt.Println(client.Name + " s'est déconnecté")
		client.Connected = false
		removeClient(*client, clientsSlice)
		return
	} else if strings.HasPrefix(mssg, "/rename") || strings.HasPrefix(mssg, "/commands") {
		commandsHandler(client, *clientsSlice, mssg)
		return
	}
	log.Print(time.Now().Format("15:04:05"), " ", client.Name, " : ", mssg)
	fmt.Print(time.Now().Format("15:04:05"), " ", client.Name, " : ", mssg)
	to_send := Message{Sender: *client, Content: mssg, Time: time.Now().Format("15:04:05")}
	*messageList = append(*messageList, to_send)
	client.Socket.Write([]byte("Envoyé!\n-------\n"))
	broadCast(*clientsSlice, to_send)
}

// function to send to other users
func broadCast(liClients []Client, message Message) {
	for _, cl := range liClients {
		if cl.Name != message.Sender.Name {
			cl.Socket.Write([]byte("\r" + message.Time + " " + message.Sender.Name + " : " + message.Content + "-------\n"))
		}
	}
}

// Remove a client
func removeClient(cl Client, li *[]Client) []Client {
	var l []Client
	for _, val := range *li {
		if val.Name != cl.Name {
			l = append(l, val)
		}
	}
	*li = l
	return *li
}

// Print all passed messages
func PrintMessages(li *[]Message, cl *Client) {
	if len(*li) <= 0 {
		return
	}
	cl.Socket.Write([]byte("=====HISTORY=====\n"))
	for _, msg := range *li {
		if msg.Sender.Name != cl.Name {
			cl.Socket.Write([]byte(msg.Time + " " + msg.Sender.Name + " : " + msg.Content))
		}
	}
	cl.Socket.Write([]byte("=================\n\n"))
	cl.Hist = len(*li)

}

func commandsHandler(client *Client, clientsSlice []Client, com string) {
	if strings.HasPrefix(com, "/commands") {
		client.Socket.Write([]byte("-----\nCOMMANDS:\n/rename [Your New Name]\n-----\n"))
	} else if strings.HasPrefix(com, "/rename") {
		if len(com) <= 8 {
			client.Socket.Write([]byte("-----\nWrong command using, check commands with /commands\n-----\n"))
		} else {
			new := strings.Replace(com, "/rename", "", 1)
			new = strings.Replace(new, "\n", "", 1)
			oldname := client.Name

			for i, cl := range clientsSlice {
				if cl.Name == oldname {
					clientsSlice[i].Name = new
				}
			}
			client.Name = new
			client.Socket.Write([]byte("-----\nName swap done! New name : " + client.Name + "\n-----\n"))
			broadCast(clientsSlice, Message{Content: "\r" + time.Now().Format("15:04:05") + " " + oldname + " changed his/her name by " + client.Name + "\n", Sender: *client, Time: time.Now().Format("15:04:05")})
		}
	}
}

//NOT CALL HERE

// Ask the username and manage if it's already used
func AskName(socket net.Conn, clientList *[]Client) string {
	for {
		socket.Write([]byte("Wsh rentre ton nom : "))
		scan := bufio.NewScanner(socket)
		if !scan.Scan() {
			log.Fatalln("Error to get name.")
		}
		name := scan.Text()
		channel := make(chan bool) //channel to manage the checking before
		go checkName(channel, clientList, name)
		result := <-channel
		switch result {
		case true:
			socket.Write([]byte("Name already use.\n"))
			continue
		case false:
			return name
		}

	}
}

func checkName(chk chan<- bool, li *[]Client, s string) {
	for _, cl := range *li {
		if cl.Name == s && s != "" {
			chk <- true
		}
	}

	chk <- false
}

func PrintPinguin(conn net.Conn) {
	content, err := os.ReadFile("./sources/ping.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		conn.Write(content)
	}
}
