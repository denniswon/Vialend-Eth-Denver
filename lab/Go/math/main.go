package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	var price float64
	var sqrtPriceX96 float64

	fmt.Println("----------------------------------------------")
	fmt.Println(".........移位 左移， 右移 .........  ")
	fmt.Println("----------------------------------------------")
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

	fmt.Println("=================================================================================")

	fmt.Println("----------------------------------------------")
	fmt.Println(".........深度copy 大数运算 big.Int 一次赋值引发的血案 .........  ")
	fmt.Println("----------------------------------------------")
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

	fmt.Println("=================================================================================")

	fmt.Println("----------------------------------------------")
	fmt.Println("......... big float Quo 除法 ， big int Div.........  ")
	fmt.Println("----------------------------------------------")
	//var bv = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(25))
	var bv, _ = new(big.Int).SetString("25729324269165216041", 10)
	fmt.Println("big int 整数", bv)
	fbalance := new(big.Float)
	fbalance.SetString(bv.String())
	fValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(fValue) // 25.729324269165216041

	fmt.Println("----------------------------------------------")
	fmt.Println(".........big.float 百分比计算 .........  ")
	fmt.Println("----------------------------------------------")
	// calculate percentable   a * 100 / B = x %
	bal := big.NewInt(125)
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	fmt.Println(bal, "*", 100, "/", fbal)
	fmt.Println(new(big.Float).Quo(fbal, big.NewFloat(1200)))
	x := new(big.Float).Quo(new(big.Float).Mul(fbal, big.NewFloat(100)), big.NewFloat(1200))
	fmt.Println(x, "%")

	fmt.Println("=================================================================================")

	fmt.Println("----------------------------------------------")
	fmt.Println(".........向下取整， 向上取整， 四舍五入 .........  ")
	fmt.Println("----------------------------------------------")

	var f = float64(4.5)
	fmt.Println(f, "向上取整= ", math.Ceil(f))           // 5
	fmt.Println(f, "向下取整= ", math.Floor(f))          // 4
	fmt.Println(f, "四舍五入= ", int(math.Floor(f+0.5))) //
}

///  return big.NewInt(10 * 1e18) //10 * 10**18
func x1E18(x int64) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)
	bigx := big.NewInt(x)

	return bigx.Mul(bigx, e18)
}
