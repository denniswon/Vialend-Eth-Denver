package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
)

func main() {

	x, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000", 10)
	fmt.Println(X1E18X(x))

	//	convert()
	//bigInt()

	//bigFloat()

	//	test()

}

func test() {
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

	fmt.Println(X1E18(5000000000000005))

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

func convert() {

	tick, _ := new(big.Int).SetString("-293422343333334", 10)

	fmt.Println(tick.Int64())
	fmt.Println(tick.IsInt64())

	fmt.Println(float64(tick.Int64()))

	tick, _ = new(big.Int).SetString("-29342223442332333333333333333333333334", 10)

	fmt.Println(tick.Int64())

	tick, _ = new(big.Int).SetString("-29342223423423423423423434333332333333333333333333333333333334", 10)

	fmt.Println(tick.Int64())
	fmt.Println(tick.IsInt64())
	fmt.Println(float64(tick.Int64())) //wrong because int overflow

	tick = big.NewInt(-23748)
	fmt.Println(tick.Quo(tick, big.NewInt(60)))
	tick = big.NewInt(-23748)
	fmt.Println(tick.Div(tick, big.NewInt(60)))

	tick = big.NewInt(23748)
	fmt.Println(tick.Quo(tick, big.NewInt(60)))
	tick = big.NewInt(23748)
	tickSpacing := big.NewInt(60)
	fmt.Println(tick.Div(tick, tickSpacing).Mul(tick, tickSpacing))

	tick, _ = new(big.Int).SetString("-23", 10)
	fmt.Println(new(big.Int).Quo(tick, big.NewInt(6)))
	tick, _ = new(big.Int).SetString("-23", 10)
	fmt.Println(tick.Div(tick, big.NewInt(6)))

}
func bigFloat() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........big float test .........  ")
	fmt.Println("----------------------------------------------")

	bigf, _ := new(big.Float).SetString("234923490230942903402349")

	fmt.Println(bigf)

	f1, _ := new(big.Float).SetString("2.32342343333333333433333333")

	fmt.Println(f1)

	f2, _ := new(big.Float).SetString("2323423433333333334333333332342342342342423433333333333333333333333333333333333333333333333333")
	fmt.Println(f2)

	a1, _ := new(big.Float).SetString("3")
	a2, _ := new(big.Float).SetString("2")
	fmt.Println(a1.Quo(a1, a2))

	fmt.Println(a1.Quo(bigf, f2))

	bigd, _ := new(big.Int).SetString("234923490230942903402349", 10)
	d1, _ := new(big.Int).SetString("23234234333333333343333333323423423423424234333333333333333333333333333333333333333333333333335", 10)

	fmt.Println(d1)

	fmt.Println(bigd.Div(bigd, d1))
}

func bigInt() {

	fmt.Println("----------------------------------------------")
	fmt.Println(".........big int test .........  ")
	fmt.Println("----------------------------------------------")

	fmt.Println("12345642342342341234564234234234123456423423423412345642342342341234564234234234")
	fmt.Println(FloatToBigInt(float64(12345642342342341234564234234234123456423423423412345642342342341234564234234234)))
	a, _ := new(big.Int).SetString("12345642342342341234564234234234123456423423423412345642342342341234564234234234", 10)
	b, _ := new(big.Int).SetString("4234123456423423423412345642342342341234564234234234", 10)
	c, _ := new(big.Int).SetString("28", 10)
	d, _ := new(big.Int).SetString("2", 10)
	e, _ := new(big.Int).SetString("3", 10)

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("a*b", new(big.Int).Mul(a, b))
	fmt.Println("a/b", new(big.Int).Div(a, b))
	fmt.Println("c*d/e", c.Mul(c, d).Div(c, e))

	fmt.Println("c", c)

	c, _ = new(big.Int).SetString("28", 10)
	d, _ = new(big.Int).SetString("2", 10)
	e, _ = new(big.Int).SetString("3", 10)

	fmt.Println("c+d+e", c.Add(c, d).Add(c, e))
	fmt.Println("(c+d)/e", c.Add(c, d).Div(c, e))
	fmt.Println("(c+d)/e", c.Add(c, d).Div(c, e))

	c, _ = new(big.Int).SetString("28", 10)
	d, _ = new(big.Int).SetString("2", 10)
	e, _ = new(big.Int).SetString("3", 10)

	fmt.Println("c+d/e", d.Div(d, e).Add(d, c))
	c, _ = new(big.Int).SetString("28", 10)
	d, _ = new(big.Int).SetString("2", 10)
	e, _ = new(big.Int).SetString("3", 10)
	fmt.Println("c+d/e", c.Add(c, d.Div(d, e)))

}

func FloatToBigInt(val float64) *big.Int {

	//price = int(sqrt(price) * (1 << 96))
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(val)
	bigf := newNum.FloatString(0)

	//fmt.Println("val, bigf:", val,  bigf)

	bigI, ok := new(big.Int).SetString(bigf, 10)
	if !ok {
		log.Fatal("float64 to bigInt err ", val, bigI)
	}

	return bigI

}

func X1E18(x int64) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)
	bigx := big.NewInt(x)

	return bigx.Mul(bigx, e18)
}

func X1E18X(bigx *big.Int) *big.Int {

	e18, _ := new(big.Int).SetString("1000000000000000000", 10)

	return bigx.Mul(bigx, e18)
}
