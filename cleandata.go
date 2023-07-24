package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// CleanData 從指定日期的 CSV 檔案讀取資料，進行清理後將結果儲存為新的 CSV 檔案
func CleanData(stockNo, date string) {
	// 構建原始檔案路徑
	filename := fmt.Sprintf("./data/%s/STOCK_DAY_%s.csv", stockNo, date)

	// 檢查檔案是否為 CSV
	if !strings.HasSuffix(filename, ".csv") {
		panic("檔案格式不正確，必須是 CSV 檔案")
	}

	// 讀取檔案內容
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// 分割行
	lines := strings.Split(string(content), "\n")

	// 過濾行並重建內容
	var filteredLines []string
	for _, line := range lines {
		// 刪除頭尾空白
		line = strings.TrimSpace(line)

		// 檢查行是否有效的 UTF-8 字串
		if utf8.ValidString(line) {
			// 進行多個替代操作
			cleanedline := strings.Replace(line, ",", "", -1)
			cleanedline = strings.Replace(cleanedline, "\"\"", ",", -1)
			cleanedline = strings.Replace(cleanedline, "\"", "", -1)
			cleanedline = strings.Replace(cleanedline, " ", "", -1)
			cleanedline = strings.Replace(cleanedline, "X", "", -1)
			cleanedline = strings.Replace(cleanedline, "+", "", -1)

			cols := strings.Split(cleanedline, ",")
			// 轉換日期格式：112/06/01 變成 20230601
			if len(cols) > 0 {
				dateParts := strings.Split(cols[0], "/")
				if len(dateParts) == 3 {
					year, _ := strconv.Atoi(dateParts[0])
					month, _ := strconv.Atoi(dateParts[1])
					day, _ := strconv.Atoi(dateParts[2])

					// 將日期轉換成格式 "20060102"
					newDate := fmt.Sprintf("%04d%02d%02d", year+1911, month, day)
					cols[0] = newDate
				}
			}

			// 排除第 7 列的漲跌價差
			if len(cols) > 7 {
				cols = append(cols[:7], cols[8:]...)
			}

			// 重新组合新的行
			newLine := strings.Join(cols, ",")

			filteredLines = append(filteredLines, newLine)
		}
	}

	// 重建内容
	newContent := strings.Join(filteredLines, "\n")

	// 檔案名稱格式
	fileName := fmt.Sprintf("./data/%s/clean/CLEAN_DATA_%s.csv", stockNo, date)

	// 確認儲存檔案的目錄是否存在，若不存在則創建它
	dirPath := fmt.Sprintf("./data/%s/clean", stockNo)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 將資料寫入 CSV 檔案
	err = ioutil.WriteFile(fileName, []byte(newContent), 0644)
	if err != nil {
		panic(err)
	}
}
