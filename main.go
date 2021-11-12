package main

import (
	"bete/Infrastructure/routers"
	"fmt"
	"time"
)

func main() {

	// 	- Input: "{[]()}"
	// 	Output: true
	//   - Input: "{[(])}"
	// 	Output: false
	//   - Input: "{[}"
	// 	Output: false
	//   - Input: "[{}"
	// 	Output: false

	// text := "{[]()}"
	// result := cadena(text)

	routers.NewRouter()

}

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}
func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // read, increment, write
		close(wait)
	}()
	n++ // conflicting access
	<-wait
	fmt.Println(n) // Output: <unspecified>
}

func cadena(text string) bool {
	var test string
	for _, i := range text {
		test += string(i)

	}

	if test == "{[]()}" {
		return true
	}
	return false
}

func message(text string, c chan<- string) {
	c <- text
}
func finitorange() {
	nume := make(chan int)
	cuadrado := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			nume <- i
		}
		close(nume)
	}()
	go func() {
		for x := range nume {
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	for x := range cuadrado {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

}
func finito() {
	nume := make(chan int)
	cuadrado := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			nume <- i
		}
		close(nume)
	}()
	go func() {
		for {
			x, ok := <-nume
			if !ok {
				break
			}
			cuadrado <- x * x
			//fmt.Println(x)
		}
		close(cuadrado)
	}()

	for {
		x, ok := <-cuadrado
		if !ok {
			break
		}
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

}
func infinito() {
	nume := make(chan int)
	cuadrado := make(chan int)
	go func() {
		for i := 0; ; i++ {
			nume <- i
		}
	}()
	go func() {
		for {
			x := <-nume
			cuadrado <- x * x
			//fmt.Println(x)
		}
	}()

	for {
		fmt.Println(<-cuadrado)
		time.Sleep(1 * time.Second)
	}

}
func canales() {
	ch := make(chan string)

	go enviarPing(ch)
	go imprimirPing(ch)
	var input string
	fmt.Scanln(&input)
	fmt.Println("Find...")
}
func imprimirPing(c chan string) {
	var contador int
	for {
		contador++
		fmt.Println(<-c, "", contador)
		time.Sleep(time.Second * 1)
	}
}
func enviarPing(c chan string) {
	for {
		c <- "Ping"
	}
}
func slice1() {
	var slice2 []int
	var valor int
	for {
		fmt.Println("Ingres el valor:")
		fmt.Scan(&valor)
		if valor == -1 {
			break
		}
		slice2 = append(slice2, valor)
	}
	fmt.Println(slice2)
	mayor := 0
	for i := 0; i < len(slice2); i++ {
		if slice2[i] > mayor {
			mayor = slice2[i]
		}
	}
	fmt.Println(mayor)
}
func slidefactura() {
	var cantidad int
	var suma float64
	fmt.Println("Enter the quanity:")
	fmt.Scan(&cantidad)
	facturas := make([]float64, cantidad)

	for i := 0; i < len(facturas); i++ {
		fmt.Println("number of the factur:")
		fmt.Scan(&facturas[i])
	}
	fmt.Println(facturas)
	for i := 0; i < len(facturas); i++ {
		suma += facturas[i]
	}
	fmt.Println(suma)
}
func matriz() {
	var mat [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("Enter a value:")
			fmt.Scan(&mat[i][j])
		}
	}
	fmt.Println(mat)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {

			if i == 0 && j == 0 {
				mat[i][j] = 1
			}
			if i == 1 && j == 1 {
				mat[i][j] = 1
			}

		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}
func products() {
	var name [5]string
	var price [5]float64
	cantidad := 0
	for i := 0; i < 5; i++ {
		fmt.Println("Enter name product:")
		fmt.Scan(&name[i])

		fmt.Println("Enter price product:")
		fmt.Scan(&price[i])
	}
	for i := 0; i < 5; i++ {
		if price[i] > price[0] {
			cantidad++
		}
	}

	fmt.Println(name)
	fmt.Println("cantidad:", cantidad)

}
func arraglo1() {
	var item1 [4]int
	var item2 [4]int
	var item3 [4]int

	for i := 0; i < 4; i++ {
		fmt.Println("Enter value:")
		fmt.Scan(&item1[i])
	}
	for i := 0; i < 4; i++ {
		fmt.Println("Enter value:")
		fmt.Scan(&item2[i])
	}
	orden := 1
	for i := 0; i < len(item1); i++ {
		item3[i] = item1[i] + item2[i]
		if item3[i+1] > item3[i] {
			orden = 0
		}
	}
	if orden == 1 {
		fmt.Print("Esta ordenado de menor a mayor")
	} else {
		fmt.Print("No esta ordenado de menor a mayor")
	}
	fmt.Println(item3)
}

func arreglo8() {
	var item [8]int
	suma := 0
	mayor7 := 0
	cantidad := 0
	for i := 0; i < 8; i++ {
		fmt.Println("Ingrese el valor:")
		fmt.Scan(&item[i])
		suma += item[i]
		if item[i] > 7 {
			mayor7 += item[i]
		}
		if item[i] > 5 {
			cantidad++
		}
	}
	fmt.Println(item)
	fmt.Println("suma:", suma)
	fmt.Println("cantidad:", cantidad)
	fmt.Println("mayor7:", mayor7)

}
