package main

import (
	"os"
	"fmt"
	"encoding/json"
)

type Server struct {

	IP string
	Status string
	Port int
}

type Servers []Server


func (s Server) CheckHealth() {	
	fmt.Printf("checking %s ... %s... at Port number %d\n", s.IP,s.Status,s.Port)	
}
func (s *Servers) ActiveServerCount() int {
	activeServerCount := 0
	for _,slc := range *s {
		returnbool,_,err := IsHealthy(slc) 

		if err != nil {
			fmt.Printf("error : %v\n",err)
			
		}
		
		if returnbool {
			
			activeServerCount++
		}
		
	}
	return activeServerCount
}

func (s *Servers) list(){
	
	for _,slc := range *s {
		fmt.Printf("server IP: %s at Port : %d \n",slc.IP,slc.Port)
	}
}

// func IsHealthyv1(s Server) bool {
// 	return s.Status == "RUNNING" 
// }

func (s *Servers) Addserver(IP string , Status string, Port int){
	server := Server { IP:IP,Status:Status,Port:Port}

	*s =  append(*s, server)

	copyTobeSaved, err := json.MarshalIndent(*s, "", " ")
	if err != nil {
		fmt.Println("failed to convert servers to JSON")
		return

	}

	err = os.WriteFile("servers.json", copyTobeSaved, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	return 
	
}


