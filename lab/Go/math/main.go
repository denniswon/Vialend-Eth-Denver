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
	bigSqrtPX96, _ := new(big.Int).SetString(s, 10)

	fmt.Println("float to bigint:", bigSqrtPX96, "==", sqrtPriceX96)

	///------------ big int value test ， 深度copy bigint variable
	fmt.Println("\nbig int value test....")
	// 初始化两个变量: a = 1, b = 2
	a := big.NewInt(1)
	b := big.NewInt(2)

	// 打印交换前的数值
	fmt.Printf("a = %v   b = %v\n", a, b)
	// 使用中間變量法進行交換
	tmp := big.NewInt(0)
	tmp.Set(a)
	a.Set(b)
	b.Set(tmp)

	// 交換完成, 對中間變量加100
	tmp.Add(tmp, big.NewInt(100))

	// 打印交換後的結果
	fmt.Printf("a = %v    b = %v   tmp = %v\n", a, b, tmp)

	fmt.Println(x1E18(5000000000000005))
}

///  return big.NewInt(10 * 1e18) //10 * 10**18
func x1E18(x int64) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)
	bigx := big.NewInt(x)

	return bigx.Mul(bigx, e18)
}
