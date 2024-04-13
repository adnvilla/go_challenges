package algorithm

// Función para encontrar todas las combinaciones posibles de sumas de los números primos dados, permitiendo duplicación
func FindCombinationsWithGoroutines(primes []int, target int, initComb []int, findOnce bool) (int, [][]int) {
	count := 0
	combis := [][]int{}

	c := make(chan []int)
	//[]int{3, 5, 7} Se establece para que al menos este presente cada numero primo en la combinacion que se encuentre
	// Si desea que se permita combinaciones que satisfagan la suma con un solo numero primo, establecer []int{}
	go combinationsWithGoroutines(primes, target, initComb, findOnce, &count, c)

	for currentComb := range c {
		tmp := make([]int, len(currentComb))
		copy(tmp, currentComb)
		combis = append(combis, tmp)
	}

	return count, combis
}

// Función recursiva para generar todas las combinaciones posibles con duplicación
func combinationsWithGoroutines(primes []int, target int, currentComb []int, findOnce bool, count *int, c chan []int) {
	sum := 0
	for _, num := range currentComb {
		sum += num
	}
	if sum == target {
		//fmt.Println(currentComb) // printSortedComb para revisar las combinaciones
		tmp := make([]int, len(currentComb))
		copy(tmp, currentComb)
		c <- tmp
		*count++
		close(c)
		return
	}
	if sum > target {
		close(c)
		return
	}
	for i := 0; i < len(primes); i++ {
		tmp := make([]int, len(currentComb))
		copy(tmp, currentComb)

		cc := make(chan []int)
		go combinationsWithGoroutines(primes, target, append(tmp, primes[i]), findOnce, count, cc)
		for x := range cc {
			tmp := make([]int, len(x))
			copy(tmp, x)
			c <- x
		}
		if findOnce && *count > 0 {
			close(c)
			return // Si ya se encontró una combinación válida, terminar la verificación de combinaciones
		}
	}

	close(c)
}
