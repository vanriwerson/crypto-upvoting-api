package main

import (
	"api/src/config"
	"fmt"
)

func main() {
	config.LoadEnvVars()

	fmt.Printf("Escutando na porta %d", config.Port)
}
