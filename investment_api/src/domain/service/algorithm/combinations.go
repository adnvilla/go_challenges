package algorithm

// Función para encontrar todas las combinaciones posibles de sumas de los números primos dados, permitiendo duplicación
func FindCombinations(primes []int, target int, initComb []int, findOnce bool) (int, [][]int) {
	count := 0
	combis := [][]int{}
	//[]int{3, 5, 7} Se establece para que al menos este presente cada numero primo en la combinacion que se encuentre
	// Si desea que se permita combinaciones que satisfagan la suma con un solo numero primo, establecer []int{}
	combinations(primes, target, initComb, findOnce, &count, &combis)
	return count, combis
}

// Función recursiva para generar todas las combinaciones posibles con duplicación
func combinations(primes []int, target int, currentComb []int, findOnce bool, count *int, combis *[][]int) {
	sum := 0
	for _, num := range currentComb {
		sum += num
	}
	if sum == target {
		//fmt.Println(currentComb) // printSortedComb para revisar las combinaciones
		tmp := make([]int, len(currentComb))
		copy(tmp, currentComb)
		*combis = append(*combis, tmp)
		*count++
		return
	}
	if sum > target {
		return
	}
	for i := 0; i < len(primes); i++ {
		combinations(primes, target, append(currentComb, primes[i]), findOnce, count, combis)
		if findOnce && *count > 0 {
			return // Si ya se encontró una combinación válida, terminar la verificación de combinaciones
		}
	}
}
