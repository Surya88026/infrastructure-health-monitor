package main

import (
	"net"
	"strconv"
	"time"
	//"fmt"

)

func IsHealthy(s Server) (bool,time.Duration,error){
	address := net.JoinHostPort(s.IP,strconv.Itoa(s.Port))

	start := time.Now()
	conn, err := net.DialTimeout("tcp",address,5*time.Second)

	duration := time.Since(start)

	if err != nil {
		// fmt.Printf("Error : faced connection issue ... %s",err)
		return false,duration,err
	}
	defer conn.Close()

	
	return true,duration,nil
}