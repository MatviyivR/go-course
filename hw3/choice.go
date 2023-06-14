package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name     string
	Inventory []string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	player := Player{Name: "Vanya"}
	player.Inventory = []string{"black package", "bottle", "knife"}

	fmt.Printf("%s woke up and do not remember anything only his name\n", player.Name)
  fmt.Printf("%s faund some box near him, look inside and saw %s\n", player.Name, player.Inventory)
  fmt.Printf("After what he found, %s start reminding that he is russia soldier, and he understood that have only 3 choise \n", player.Name)
	fmt.Println("1. Sit on the", player.Inventory[1])
	fmt.Println("2. Kill himself with", player.Inventory[2])
	fmt.Println("3. Hide in", player.Inventory[0])
  fmt.Printf("Help %s to make a choice\n", player.Name)

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSuffix(choice, "\n")

	switch choice {
	case "1":
		fmt.Printf("Thanks, %s fill %s very deeply\n", player.Name, player.Inventory[1])

  case "2":
		fmt.Printf("Thanks, %s finaly dead\n", player.Name)

  case "3":
		fmt.Printf("Thanks, %s was sent to home\n", player.Name)

  default:
		fmt.Printf("Ok, %s then %s =)\n.", player.Inventory[1], player.Inventory[1])
  }
}
