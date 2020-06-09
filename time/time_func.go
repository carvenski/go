package main

import (
	"log"
	"math/big"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	log.Printf("Func = [%s], Time = [%s]", name, elapsed)
}

func BigIntFactorial(x *big.Int) *big.Int {
	// 使用defer给函数加上计时
	defer Duration(time.Now(), "IntFactorial")

	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}

	return x.Set(y)
}

func main() {
	log.Println(BigIntFactorial(big.NewInt(10)))
}
