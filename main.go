package main

//1
//Begin12
//func main() {
//	var a float64
//	var b float64
//	fmt.Print("a=")
//	fmt.Scan(&a)
//	fmt.Print("b=")
//	fmt.Scan(&b)
//	var c float64 = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
//	fmt.Println("Гипотенуза = ", c)
//	p := a + b + c
//	fmt.Println("Периметр = ", p)
//}

// 2
// Begin15
//func main() {
//	const PI = 3.14
//	var s uint16
//	fmt.Print("Введите площадь круга = ")
//	fmt.Scan(&s)
//	r := math.Sqrt(float64(s) / PI)
//	l := 2 * PI * r
//	fmt.Println("Радиус = ", r)
//	fmt.Println("Длина окружности= ", l)
//}

// 3
// Begin20
//func main() {
//	var x1, x2, y1, y2 float64
//	fmt.Print("x1 = ")
//	fmt.Scan(&x1)
//	fmt.Print("x2 = ")
//	fmt.Scan(&x2)
//	fmt.Print("y1 = ")
//	fmt.Scan(&y1)
//	fmt.Print("y2 = ")
//	fmt.Scan(&y2)
//	c := math.Sqrt(math.Pow((x2-x1), 2) - math.Pow((y2-y1), 2))
//	fmt.Println("Расстояние между двумя точками = ", c)
//}

// 4
// Begin 37
//func main() {
//	var V1, V2, T float64
//	fmt.Print("Введите скорость первого автомобиля (км/ч): ")
//	fmt.Scan(&V1)
//	fmt.Print("Введите скорость второго автомобиля (км/ч): ")
//	fmt.Scan(&V2)
//	fmt.Print("Введите время (ч): ")
//	fmt.Scan(&T)
//
//	S := -2 * T * math.Min(V1, V2)
//	fmt.Printf("Расстояние между автомобилями: %.2f км", S)
//}

// 5
// Int11
//func main() {
//	var xyz int
//	fmt.Print("Введите трёхзначное число: ")
//	fmt.Scan(&xyz)
//	x := xyz / 100
//	y := (xyz / 10) % 10
//	z := xyz % 10
//	sum := x + y + z
//	pro := x * y * z
//	fmt.Println("Сумма цифр: ", sum)
//	fmt.Println("Произведение цифр", pro)
//}

// 6
// int15
//func main() {
//	var xyz int
//	fmt.Print("Введите трёхзначное число: ")
//	fmt.Scan(&xyz)
//	x := xyz / 100
//	y := (xyz / 10) % 10
//	z := xyz % 10
//	a := y * 100
//	b := x * 10
//	abc := a + b + z
//	fmt.Println("Ответ: ", abc)
//}

// 7
// int17
//func main() {
//	var abc int
//	fmt.Println("Введите число больше трехзначного: ")
//	fmt.Scan(&abc)
//	if abc > 999 {
//		x := (abc / 100) % 10
//		fmt.Println("Answer: ", x)
//	} else {
//		fmt.Print("Errow!")
//	}
//}

// 8
// int23
//func main() {
//	var N int
//	fmt.Print("Введите количество секунд: ")
//	fmt.Scan(&N)
//
//	minutes := (N / 60) % 60
//	fmt.Printf("Полных минут с начала последнего часа: %d", minutes)
//}

// 9
// bool12
//func main() {
//	var a, b, c int
//	fmt.Print("a = ")
//	fmt.Scan(&a)
//	fmt.Print("b = ")
//	fmt.Scan(&b)
//	fmt.Print("c = ")
//	fmt.Scan(&c)
//	if a > 0 && b > 0 && c > 0 {
//		fmt.Print("True")
//	} else {
//		fmt.Print("False")
//	}
//
//}

//10
//bool20

//func main() {
//	var abc int
//	fmt.Print("Введите трехзначное число: ")
//	fmt.Scan(&abc)
//	x := abc / 100
//	y := (abc / 10) % 10
//	z := abc % 10
//	if x != y && y != z && x != z {
//		fmt.Print("Все цифры данного числа различны!")
//	} else {
//		fmt.Print("Есть однозначные цифры! ")
//	}
//}

//11
//bool24

//func main() {
//	var a, b, c float64
//	fmt.Print("Введите а = ")
//	fmt.Scan(&a)
//	fmt.Print("b = ")
//	fmt.Scan(&b)
//	fmt.Print("c = ")
//	fmt.Scan(&c)
//	if a != 0 {
//		d := math.Pow(b, 2) - 4*a*c
//		fmt.Println("D = ", d)
//	} else {
//		fmt.Print("Errow!")
//	}
//}

//12
