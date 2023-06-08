package main

import "fmt"

type Animal struct {
    Name  string
}

type Cage struct {
    Number   int
    Capacity int
    Animal   Animal
}

type Zookeeper struct {
    Name    string
    Age     int
    Cages   []Cage
    Captured int
}

func main() {
    monkey := Animal{Name: "monkey"}
    tiger := Animal{Name: "tiger"}
    elephant := Animal{Name: "elephant"}
    giraffe := Animal{Name: "giraffe"}
    zebra := Animal{Name: "Zebra"}

    cage1 := Cage{Number: 1, Capacity: 1, Animal: monkey}
    cage2 := Cage{Number: 2, Capacity: 1, Animal: tiger}
    cage3 := Cage{Number: 3, Capacity: 1, Animal: elephant}
    cage4 := Cage{Number: 4, Capacity: 1, Animal: giraffe}
    cage5 := Cage{Number: 5, Capacity: 1, Animal: zebra}

    cages := []Cage{cage1, cage2, cage3, cage4, cage5}

    zookeeper := Zookeeper{Name: "Ruslan", Age: 30, Cages: cages, Captured: 0}

    fmt.Printf("Zookeeper: %s\n", zookeeper.Name)
    fmt.Printf("Age: %d\n", zookeeper.Age)
    fmt.Printf("Number of cages: %d\n", len(zookeeper.Cages))


    zookeeper.Captured = 4

    fmt.Printf("\nNumber of captured animals: %d\n", zookeeper.Captured)
}
