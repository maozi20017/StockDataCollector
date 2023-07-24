package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// GetData 從指定日期的 TWSE API 獲取股票資料並儲存為 CSV 檔案
func GetData(stockNo string, date string) {
	// 構建 TWSE API 的 URL
	url := fmt.Sprintf("https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=csv&date=%s&stockNo=%s", date, stockNo)

	// 發送 HTTP GET 請求
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 讀取 API 回傳的內容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 檔案名稱格式
	fileName := fmt.Sprintf("./data/%s/STOCK_DAY_%s.csv", stockNo, date)

	// 確認儲存檔案的目錄是否存在，若不存在則創建它
	dirPath := fmt.Sprintf("./data/%s", stockNo)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 將資料寫入 CSV 檔案
	err = ioutil.WriteFile(fileName, body, 0644)
	if err != nil {
		panic(err)
	}

	// 加入延遲，讓伺服器有時間反應
	time.Sleep(2 * time.Second)
}
