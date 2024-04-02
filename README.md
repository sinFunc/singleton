Singleton 

purpose:
provide interface to get a global instance for any type.
There is only a file (singleton.go)

How to use

````package main

import (
    "fmt"
    "github.com/sinFunc/singleton"
)

type Addr struct {
    Ip   string
    port int
}

func main() {
    s := singleton.GetInstance[Addr]().(*Addr)
    s.Ip = "127.0.0.1"
    s.port = 7777
    fmt.Printf("instance1---->ip:port=%v:%v\n", s.Ip, s.port)

    s2 := singleton.GetInstance[Addr]().(*Addr)
    fmt.Printf("instance2---->ip:port=%v:%v\n", s2.Ip, s2.port)

    s3 := singleton.GetInstanceTemplate[Addr]()
    s3.Ip = "0.0.0.0"
    s3.port = 6666
    fmt.Printf("instance1---->ip:port=%v:%v;", s.Ip, s.port)
    fmt.Printf("instance2---->ip:port=%v:%v\n", s2.Ip, s2.port)

}````


