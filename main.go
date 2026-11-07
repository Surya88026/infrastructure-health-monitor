package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type ServerConnectionDetail struct {
	serverInfo Server
	connTime   time.Duration
	err        error
}

type connectionInfo []ServerConnectionDetail

func main() {

	servers := LoadServers()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [list|add|health|stats]")
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		fmt.Println("Listing servers...")
		servers.list()
	case "add":
		fmt.Println("Adding servers...")
		if len(os.Args) < 5 {
			fmt.Println("not enough arguments: missing port")
			return
		}
		var ip string = os.Args[2]
		var status string = os.Args[3]

		if (status != "RUNNING") && (status != "DOWN") {
			fmt.Println("Status is incorrect. please choose correct status!")
			return
		}

		port, err := strconv.Atoi(os.Args[4])

		if err != nil {
			fmt.Println("Alert !")
			return
		}
		if port > 65535 || port < 1 {

			fmt.Println("Invalid port range value")
			return
		}
		servers.Addserver(ip, status, port)
		servers.list()

	case "health":

		fmt.Println("checking server health...")

		var failedserver connectionInfo
		var connectedserver connectionInfo
		for i := range servers {
			returnbool, returnValue, err := IsHealthy(servers[i])

			if returnbool {
				connectedserver = append(connectedserver, ServerConnectionDetail{servers[i], returnValue, err})
			} else {
				failedserver = append(failedserver, ServerConnectionDetail{servers[i], returnValue, err})
			}
		}

		if len(connectedserver) != 0 {
			fmt.Printf("Connected Server : \n")
			for _, s := range connectedserver {
				fmt.Printf("  - server : %s at Port : %d | Duration : %v\n", s.serverInfo.IP, s.serverInfo.Port, s.connTime)
			}
		}

		if len(failedserver) != 0 {
			fmt.Printf("Failed Connection : \n")
			for _, s := range failedserver {
				fmt.Printf("   - server : %s at Port : %d | Duration : %v\n Reason: %v\n", s.serverInfo.IP, s.serverInfo.Port, s.connTime, s.err)
			}
		}

	case "stats":
		active := servers.ActiveServerCount()
		successRate := 0.0

		if len(servers) > 0 {
			successRate = float64(active) / float64(len(servers)) * 100
		}

		fmt.Println("\nInfrastructure Monitoring Statistics")
		fmt.Println("-------------------------------------")
		fmt.Printf("Total Servers       : %d\n", len(servers))
		fmt.Printf("Active Servers      : %d\n", active)
		fmt.Printf("Inactive Servers    : %d\n", len(servers)-active)
		fmt.Printf("Health Success Rate : %.2f%%\n", successRate)

	default:
		fmt.Println("Unknown command:", command)
	}

	//fmt.Printf("Number of inactive servers : %d", len(servers)-activeServerCount)

}
