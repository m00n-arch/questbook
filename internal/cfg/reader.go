package cfg

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/m00n-arch/questbook/internal/domain"
)

func OpenJson(path string) (*domain.DialogData, error) {

	// Открываем JSON-файл для чтения
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return nil, err
	}
	defer file.Close() // Важно закрыть файл после чтения

	// Создаем JSON декодер, используя io.Reader
	decoder := json.NewDecoder(file)

	// Создаем структуру для разбора данных из файла
	var data domain.DialogData

	// Разбираем JSON из файла
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return nil, err
	}

	// Теперь у вас есть данные из JSON-файла в структуре data
	return &data, nil
}
