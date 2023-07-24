package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func MergeCSVFiles(stockNo string) error {
	// 讀取資料夾內的所有檔案
	folderPath := fmt.Sprintf("./data/%s/clean", stockNo)
	fileList, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return err
	}

	var mergedData [][]string

	// 新增欄位標題行
	headerRow := []string{
		"日期", "成交股數", "成交金額", "開盤價", "最高價", "最低價", "收盤價", "成交筆數",
	}
	mergedData = append(mergedData, headerRow)

	// 篩選出 CSV 檔案並讀取內容
	for _, file := range fileList {
		if strings.HasSuffix(file.Name(), ".csv") {
			filePath := filepath.Join(folderPath, file.Name())
			data, err := readCSVFile(filePath)
			if err != nil {
				fmt.Printf("無法讀取檔案 %s: %v\n", file.Name(), err)
				continue
			}
			mergedData = append(mergedData, data...)
		}
	}

	// 儲存檔案目錄路徑
	mergeFolderPath := fmt.Sprintf("./data/%s/clean/merge", stockNo)

	// 確認儲存檔案的目錄是否存在，若不存在則創建它
	err = os.MkdirAll(mergeFolderPath, os.ModePerm)
	if err != nil {
		return err
	}

	// 儲存檔案的路徑
	outputFilePath := filepath.Join(mergeFolderPath, "merged.csv")

	// 寫入合併後的結果到新的 CSV 檔案
	err = writeCSVFile(outputFilePath, mergedData)
	if err != nil {
		return err
	}
	return nil
}

// 讀取 CSV 檔案並回傳資料
func readCSVFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 將資料寫入 CSV 檔案
func writeCSVFile(filePath string, data [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, row := range data {
		// 移除頭尾的雙引號
		for i, value := range row {
			row[i] = strings.Trim(value, "\"")
		}
		writer.Write(row)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	return nil
}
