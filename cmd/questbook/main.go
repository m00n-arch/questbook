package main

import (
	"fmt"

	"github.com/m00n-arch/questbook/internal/cfg"
	"github.com/m00n-arch/questbook/internal/game"
)

func main() {
	dialogData, err := cfg.OpenJson("scriptstemplate/chapter1/act1.json")
	if err != nil {
		fmt.Println("Ошибка при чтении JSON:", err)
		return
	}

	game.DisplayDialog(dialogData)

	fmt.Println("Диалог завершен.")
}
