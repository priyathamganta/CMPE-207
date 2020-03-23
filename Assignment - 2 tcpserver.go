package main

import (
    "fmt"
    "net"
    "os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "23"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn , err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  conn.Write([]byte("          _______  _        _        _______             _______  _______  _        ______" + "\n"))
  conn.Write([]byte("|\\     /|(  ____ \\/ \\      ( \\      (  ___  )  |\\     /|(  ___  )(  ____ )( \\      (  __  \\" + "\n"))
  conn.Write([]byte("| )   ( || (    \\/| (      | (      | (   ) |  | )   ( || (   ) || (    )|| (      | (  \\  )" + "\n"))
  conn.Write([]byte("| (___) || (__    | |      | |      | |   | |  | | _ | || |   | || (____)|| |      | |   ) |" + "\n"))
  conn.Write([]byte("|  ___  ||  __)   | |      | |      | |   | |  | |( )| || |   | ||     __)| |      | |   | |" + "\n"))
  conn.Write([]byte("| (   ) || (      | |      | |      | |   | |  | || || || |   | || (\\ (   | |      | |   ) |" + "\n"))
  conn.Write([]byte("| )   ( || (____/\\| (____/\\| (____/\\| (___) |  | () () || (___) || ) \\ \\__| (____/\\| (__/  )" + "\n"))
  conn.Write([]byte("|/     \\|(_______/(_______/(_______/(_______)  (_______)(_______)|/   \\__/(_______/(______/" + "\n"))
  // Close the connection when you're done with it.
  conn.Close()
}
