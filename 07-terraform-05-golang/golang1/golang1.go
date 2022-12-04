package main

import "fmt"

func main() {
        var input, output float64

        fmt.Print("Enter value (metters): ")
        _, err := fmt.Scanf("%f", &input)
        if err != nil {
                fmt.Print("Please input correct float value")
        }
        output = input/0.3048
        fmt.Println(output)
}
