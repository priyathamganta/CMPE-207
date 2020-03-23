package main

import "net"
import "fmt"
import "bufio"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "94.142.241.111:23")
  connbuf := bufio.NewReader(conn)
  for{
    str, _ := connbuf.ReadString('\n')
    if len(str)>0 {
        fmt.Println(str)
    }
  }
}
