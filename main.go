package main

import "github.com/juanlubel/katas/marsrover/src"

/*
1. Iniciamos la instancia con una posición inicial
2. Recivirá una array de ordenes
3. Debe poder encararse con las ordendes L y R
4. Debe poder avanzar o retroceder con las ordenes F y B
5. Deberá conocer los limites del terreno
6. Debe poder detectar obstaculos
*/

func main() {
	rover := src.NewRover(src.InitialPosition)
	rover.StartsRover()
}
