<img src="https://github.com/ClemNTTS/net-cat-project/assets/161344481/3da84aaa-5ac3-4aaf-aade-cad8471e3f99" alt="image description" width="300"/>


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
