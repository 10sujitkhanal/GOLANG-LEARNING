package main

import "fmt"

type Address struct {
    city  string
    state string
}

type Person struct {
    name    string
    age     int
    address Address // embedded struct
}

func main() {
    p := Person{
        name: "Alice",
        age:  30,
        address: Address{
            city:  "New York",
            state: "NY",
        },
    }

    fmt.Println(p)
    fmt.Println(p.address.city)
}
