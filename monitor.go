package main

import (
	"net"
	"strconv"
	"time"
	"fmt"

)

func IsHealthy(s Server) bool{
	address := net.JoinHostPort(s.IP,strconv.Itoa(s.Port))

	conn, err := net.DialTimeout("tcp",address,3*time.Second)

	if err != nil {
		fmt.Printf("Error : faced connection ... %s",err)
		return false
	}
	defer conn.Close()
	return true
}