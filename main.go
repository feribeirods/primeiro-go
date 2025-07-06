package main

import (
	"fmt"
)

func buscaItem(itens map[string]string, item string) (string, error) {
	
	for itens, status := range itens {
		if itens == item {
			return status, nil
		}
	}

	return "", fmt.Errorf("Item %s não encontrado", item)
}

func main() {
	var itens map[string]string
	itens = make(map[string]string)

	itens["mesa"] = "Disponível"
	itens["cadeira"] = "Reservado"
	itens["porta"] = "Disponível"

	item, err := buscaItem(itens, "porta")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O status do item é: %s\n", item)
	}
		
}
