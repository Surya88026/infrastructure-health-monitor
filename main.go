
package main
import (
	"fmt"
	
	"os"
	"strconv"

)
func main(){

	//fmt.Printf("Parsed Data: %+v\n",LoadServers())
	
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

		if  (status != "RUNNING" ) && (status != "DOWN") {
			fmt.Println("Status is incorrect. please choose correct status!")
			return
		}

		
		port,err := strconv.Atoi(os.Args[4])

		if err != nil {
			fmt.Println("Alert !")
			return
		}
		if port >65535 || port <1 {

			fmt.Println("Invalid port range value")
			return
		}
		servers.Addserver(ip,status,port)
		servers.list()
	case "health":
		fmt.Println("checking server health...")
		for i := range servers {
		
		if IsHealthy(servers[i]) {
			servers[i].CheckHealth()
			//activeServerCount++
		}
	}

	case "stats":
		fmt.Println("Showing server statistics...")
		fmt.Printf("Total Server Count %d\n",len(servers))
		fmt.Printf("Active Server Count %d\n",servers.ActiveServerCount())
		fmt.Printf("Inactive Server Count %d\n",len(servers)-servers.ActiveServerCount())

	default:
		fmt.Println("Unknown command:", command)
	}

	 

	
	//fmt.Printf("Number of inactive servers : %d", len(servers)-activeServerCount)

	
}
