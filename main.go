package main

import (
	"fmt"
	"time"
)

func main() {

	var stockNo string
	var numMonths int

	// 獲取使用者輸入的股票代號和月份數
	fmt.Print("輸入股票代號：")
	fmt.Scan(&stockNo)

	fmt.Print("輸入要收集的月份數：")
	fmt.Scan(&numMonths)

	// 獲取當天的日期
	currentTime := time.Now()

	// 將當前日期改為該月的1號
	currentTime = time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, time.UTC)

	// 使用 for 迴圈往前推算指定個月數的日期
	for i := 0; i < numMonths; i++ {
		// 從當前日期減去 i 個月
		date := currentTime.AddDate(0, -i, 0).Format("20060102")
		GetData(stockNo, date)
		CleanData(stockNo, date)
	}
	fmt.Printf("檔案已保存到 data/%s/\n", stockNo)
	fmt.Printf("CLEAN 檔案已保存到 data/%s/clean/\n", stockNo)
	// 合併 CSV 檔案
	err := MergeCSVFiles(stockNo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("CSV 檔案已合併到 data/%s/clean/merge/\n", stockNo)
}
