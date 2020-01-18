package main

func main() {

}

func sumOfBonuses(sales []int64) (sum int64) {
	const minSale =	10_000
	const overSumPercent = 5
	for _,value := range sales {
		if value >= minSale {
			sum += (value - minSale) * overSumPercent / 100
		}
	}
	return
}