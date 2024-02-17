package main

import "flag"

func main() {
	host := flag.String("host", "localhost", "isi host")
	username := flag.String("username", "root", "isi username")
	passwrod := flag.String("password", "root", "isi password")
	port := flag.Int("port", 0, "isi port")

	flag.Parse()

	println("Username", *username)
	println("Password", *passwrod)
	println("Host", *host)
	println("Port", *port)
}
