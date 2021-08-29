package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	var price float64
	var sqrtPriceX96 float64

	/// bit shift
	fmt.Println("i<<j = i * (1<<j) ", 3<<2, "==", 3.0*(1<<2))
	fmt.Println("i>>j = i/(2^j) ", 164>>4, "==", math.Floor(164/math.Pow(2, 4)))

	/// example
	price = 1e18 / 2000e6

	fmt.Println("price:", price)

	/// price to sqrtPx96   sqrt(price) << 96
	sqrtPriceX96 = math.Floor(math.Sqrt(price) * (1 << 96))
	fmt.Println("sqrPriceX96:", sqrtPriceX96)

	/// sqrtPx96 to price  sqrtPriceX96 ** 2  >> 96*2
	sqrtPx962Price := math.Floor((sqrtPriceX96 * sqrtPriceX96) / math.Pow(2, 2*96))
	fmt.Println("sqrtP*96 -> Price ", sqrtPx962Price, " == ", price)

	/// convert float64 to big.int
	s := fmt.Sprintf("%.0f", sqrtPriceX96)
	biSqrtPX96, _ := new(big.Int).SetString(s, 10)

	fmt.Println("float to bigint:", biSqrtPX96, "==", sqrtPriceX96)

}
