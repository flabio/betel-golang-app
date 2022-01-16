package main

import (
	"bete/Infrastructure/routers"
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
