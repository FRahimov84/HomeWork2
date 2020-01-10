package main

func main() {

}

func sumOfBonuses(sales []int) (sum float64) {
	const minSale =	10_000
	const overSumPercent =  0.05
	for _,value := range sales {
		if value >= minSale {
			sum += float64(value) * overSumPercent
		}
	}
	return
}