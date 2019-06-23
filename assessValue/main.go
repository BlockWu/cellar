package main

import "fmt"

func main() {
	AssessValueNormal()
}

// 贴现估算法
func AssessValueNormal() {
	var (
		Year             int     = 10   // 预测年数
		FirstYearProfit  float64 = 1074 // 当年利润
		DCFRate          float64 = 0.08 // 贴现率
		ProfitGrowthRate float64 = 0.1  // 预测增长率
		StockShare       float64 = 183  // 总股本

		Sum                float64
		YearProfitSlice    []float64
		DCFYearProfitSlice []float64
	)

	for i := 0; i < Year; i++ {
		FirstYearProfit = FirstYearProfit * (1 + ProfitGrowthRate)
		YearProfitSlice = append(YearProfitSlice, FirstYearProfit)
	}

	for k, v := range YearProfitSlice {
		count := k
		value := v

		for {
			if count == -1 {
				break
			}

			value = value / (1 + DCFRate)
			count--
		}

		DCFYearProfitSlice = append(DCFYearProfitSlice, value)
	}

	for _, v := range DCFYearProfitSlice {
		Sum = Sum + v
	}

	fmt.Printf("使用贴现率为%v,预测增长率为%v,得到的合理估值为%v,合理价格为%v\n", DCFRate, ProfitGrowthRate, Sum, Sum/StockShare)
}
