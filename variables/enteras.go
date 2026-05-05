package variables

import "fmt"

func MuestrorEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Entero común:", intcomun)
	fmt.Println("Entero de 32 bits:", intde32)
	fmt.Println("Entero de 64 bits:", intde64)
}
