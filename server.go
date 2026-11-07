package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	IP     string
	Status string
	Port   int
}

type Servers []Server

func (s Server) CheckHealth() {
	fmt.Printf("checking %s ... %s... at Port number %d\n", s.IP, s.Status, s.Port)
}

func (s *Servers) ActiveServerCount() int {
	results := make([]bool, 0, len(*s))

	for _, server := range *s {
		isHealthy, _, _ := IsHealthy(server)
		results = append(results, isHealthy)
	}

	return CountHealthyServers(results)
}

func CountHealthyServers(results []bool) int {
	count := 0

	for _, healthy := range results {
		if healthy {
			count++
		}
	}
	return count
}

func (s *Servers) list() {

	for _, slc := range *s {
		fmt.Printf("server IP: %s at Port : %d \n", slc.IP, slc.Port)
	}
}

// func IsHealthyv1(s Server) bool {
// 	return s.Status == "RUNNING"
// }

func (s *Servers) Addserver(IP string, Status string, Port int) {
	server := Server{IP: IP, Status: Status, Port: Port}

	*s = append(*s, server)

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
