package main

import (
	"bufio"
	"dsuOnGo/dsu"
	"fmt"
	"os"
	"strings"
)

func printIntro() {
	println("Created default DSU of size 50")
	println("Supported commands:")
	println("CREATE N")
	println("GET_SETSG_NUM")
	println("GET_NODES_NUM")
	println("UNITE i j")
	println("IS_SAME N")
	println()
}

func runCreate(data **dsu.DsuDataStruct, cmd string) {
	var n int
	_, err := fmt.Sscanf(cmd, "CREATE %d", &n)
	if err != nil {
		fmt.Println("Error parsing arguments")
		return
	}
	*data = dsu.NewDsu(n)
	println("Created new DSU with requested size")
}

func runUnite(data *dsu.DsuDataStruct, cmd string) {
	var i1, i2 int
	_, err := fmt.Sscanf(cmd, "UNITE %d %d", &i1, &i2)
	if err != nil {
		fmt.Println("Error parsing arguments")
		return
	}
	if data.Unite(i1, i2) {
		fmt.Println("United")
	} else {
		fmt.Println("Were already united")
	}
}

func runGetNodesNum(data *dsu.DsuDataStruct) {
	fmt.Printf("There are %d nodes\n", data.GetNodesNum())
}

func runGetSetsNum(data *dsu.DsuDataStruct) {
	fmt.Printf("There are %d sets\n", data.GetSetsNum())
}

func runIsSame(data *dsu.DsuDataStruct, cmd string) {
	var i1, i2 int
	_, err := fmt.Sscanf(cmd, "IS_SAME %d %d", &i1, &i2)
	if err != nil {
		fmt.Println("Error parsing arguments")
		return
	}
	if data.Unite(i1, i2) {
		fmt.Println("In same set")
	} else {
		fmt.Println("Were already united")
	}
}

func main() {
	data := dsu.NewDsu(50)
	fmt.Println(data.Unite(5, 7))
	printIntro()

	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, isPref, err := reader.ReadLine()
		if isPref {
			fmt.Println("Too big line")
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		cmdLine := string(bytes)
		tokens := strings.Split(cmdLine, " ")
		if len(tokens) < 1 {
			fmt.Println("Invalid input: less than 1 token was found")
			continue
		}
		cmd := tokens[0]
		switch cmd {
		case "CREATE":
			runCreate(&data, cmdLine)
		case "UNITE":
			runUnite(data, cmdLine)
		case "GET_SETS_NUM":
			runGetSetsNum(data)
		case "GET_NODES_NUM":
			runGetNodesNum(data)
		case "IS_SAME":
			runIsSame(data, cmdLine)
		default:
			fmt.Println("Unknown command: ", cmd)
		}
	}
}
