package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/fxsjy/gonn/gonn"
)

func CreateNN() {
	// Создаём НС с 3 входными нейронами (столько же входных параметров),
	// 16 скрытыми нейронами и
	// 4 выходными нейронами (столько же вариантов ответа)
	nn := gonn.DefaultNetwork(3, 16, 3, false)

	// Создаём массив входящих параметров:
	// 1 параметр - количество здоровья (0.1 - 1.0)
	// 2 параметр - наличие оружия (0 - нет, 1 - есть)
	// 3 параметр - количество врагов

	//
	// Сумма на руках (0 - 5.0)
	// Стоимость (0-5.0)
	// Дней до ЗП (0-7.0)

	//

	input := [][]float64{
		[]float64{0, 1.0, 3.0}, []float64{0.9, 1, 2}, []float64{0.8, 0, 1},
		[]float64{0, 5.0, 0}, []float64{0.6, 1, 2}, []float64{0.4, 0, 1},
		[]float64{2.0, 1.0, 2.0}, []float64{0.6, 1, 4}, []float64{0.1, 0, 1},
		[]float64{0.6, 1, 0}, []float64{1, 0, 0}}

	// input := [][]float64{
	// 	[]float64{0.5, 1, 1}, []float64{0.9, 1, 2}, []float64{0.8, 0, 1},
	// 	[]float64{0.3, 1, 1}, []float64{0.6, 1, 2}, []float64{0.4, 0, 1},
	// 	[]float64{0.9, 1, 7}, []float64{0.6, 1, 4}, []float64{0.1, 0, 1},
	// 	[]float64{0.6, 1, 0}, []float64{1, 0, 0}}

	// Теперь создаём "цели" - те результаты, которые нужно получить
	target := [][]float64{
		[]float64{1, 0, 0, 0}, []float64{1, 0, 0, 0}, []float64{1, 0, 0, 0},
		[]float64{0, 1, 0, 0}, []float64{0, 1, 0, 0}, []float64{0, 1, 0, 0},
		[]float64{0, 0, 1, 0}, []float64{0, 0, 1, 0}, []float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1}, []float64{0, 0, 0, 1}}

	// Начинаем обучать нашу НС.
	// Количество итераций - 100000
	nn.Train(input, target, 100000)

	// Сохраняем готовую НС в файл.
	gonn.DumpNN("gonn", nn)
}

func GetResult(output []float64) string {
	var max float64 = -99999
	pos := -1
	// Ищем позицию нейрона с самым большим весом.
	for i, value := range output {
		if value > max {
			max = value
			pos = i
		}
	}

	// Теперь, в зависимости от позиции, возвращаем решение.
	switch pos {
	case 0:
		return "Покупай"
	case 1:
		return "Не покупай"
	case 2:
		return "Подумай"
	}
	return ""
}
func randFloats(min, max float64) float64 {
	var res float64
	res = math.Trunc((min+rand.Float64()*(max-min))*10) / 10
	return res

}
func main() {
	var par1, par2, par3 float64
	// var tar1, tar2, tar3 float64

	var input [][]float64
	var target [][]float64

	for i := 0; i < 100; i++ {
		par1 = randFloats(0, 5.0)
		par2 = randFloats(0, 5.0)
		par3 = randFloats(0, 7.0)

		if par1-par2 <= 0 {
			continue
		} else {
			if ((par1 - par2) / 0.5) > par3 {
				target = append(target, []float64{1, 0, 0})
			} else if ((par1 - par2) / 0.5) == par3 {
				target = append(target, []float64{0, 0, 1})
			} else {
				target = append(target, []float64{0, 1, 0})
			}
		}

		input = append(input, []float64{par1, par2, par3})
	}
	fmt.Print("[INPUT]")
	fmt.Println(len(input))
	fmt.Print("[TARGET]")
	fmt.Println(len(target))

	// nn := gonn.DefaultNetwork(3, 16, 3, false)
	// Начинаем обучать нашу НС.
	// Количество итераций - 100000
	// nn.Train(input, target, 100000)

	// Сохраняем готовую НС в файл.
	// gonn.DumpNN("gonn2", nn)

	// // CreateNN()
	// // Загружем НС из файла.
	nn := gonn.LoadNN("gonn2")

	// // // Записываем значения в переменные:
	// // hp - здоровье (0.1 - 1.0)
	// // weapon - наличие оружия (0 - нет, 1 - есть)
	// // enemyCount - количество врагов
	// 2.3 1.4 2
	var Wallet float64 = 2.0
	var Sale float64 = 2.0
	var Days float64 = 1.0

	// // Получаем ответ от НС (массив весов)
	out := nn.Forward([]float64{Wallet, Sale, Days})

	// // fmt.Println(out)
	// // Печатаем ответ на экран.
	fmt.Println(GetResult(out))
	qwrqwr
}
