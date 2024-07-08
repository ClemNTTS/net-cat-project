![Logo of the project]([https://www.google.com/url?sa=i&url=https%3A%2F%2Fru.wikipedia.org%2Fwiki%2F%25D0%25A4%25D0%25B0%25D0%25B9%25D0%25BB%3ANetcat_logo.png&psig=AOvVaw0OlqCm4__hsuv9oS0AWrgf&ust=1720519343985000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCKijlZaYl4cDFQAAAAAdAAAAABAE](https://www.kali.org/tools/netcat/images/netcat-logo.svg))

# NET-CAT project
> by Clément Nuttens

This project could let you start a server and allowed people to connect on it, they could speak together when they use the netcat command : nc <$ip> <$port>

## Installing / Getting started

A quick introduction of the minimal setup you need to runn.

Launching the server
```shell
git clone https://github.com/ClemNTTS/net-cat-project.git
go build -o TCPChat main.go #build the executable file
./TCPChat $ip $port #start the server
```
Without arguments, the server is launch on a localhost at the 8080 port.

```shell
nc lohalhost 8080 #for base launching server
#or
nc <$ip> <$port>
```
Start a new user session

#### Argument 1
Type: `String`  
Default: `localhost`

with one argument, put the port, with two arguments, put th ip of the server

Example:
```bash
./TCPChat 8080
```

## Licensing

The code in this project is licensed under MIT license.
