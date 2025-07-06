package main

import (
	"fmt"
)

func buscaItem(itens map[string]string, item string) (string, error) {

	status, ok := itens[item]
	if ok {
		return status, nil
	} else {
		return "", fmt.Errorf("Item %s não encontrado", item)
	}

	
}

func main() {
	itens := map[string]string {
		"cadeira":   "Disponível",
		"mesa": "Reservado",
		"porta": "Disponível",
	}

	item, err := buscaItem(itens, "mesa")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O status do item é: %s\n", item)
	}
		
}
