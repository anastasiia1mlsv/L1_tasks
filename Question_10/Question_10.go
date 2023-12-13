package main

import (
	"encoding/json"
	"fmt"
)

/*
	Дана последовательность температурных колебаний:
	-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	// Исходные температуры
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Map для группировки температур
	groupedTemps := make(map[int][]float64)

	// Группировка температур
	for _, temp := range temperatures {
		// Дата кастинг до инта, получаем целую часть
		group := int(temp/10) * 10                              // Вычисляем группу для текущей температуры
		groupedTemps[group] = append(groupedTemps[group], temp) // пушим ключ и значение
	}

	jsonTemps, err := json.MarshalIndent(groupedTemps, "", "\t")
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}
	fmt.Println(string(jsonTemps))
}
