package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintAAFromTxt(filename string) {
	// 実行ファイルのディレクトリを取得
	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		return
	}

	// 実行ファイルのディレクトリから相対パスを構築
	execDir := filepath.Dir(execPath)
	fullPath := filepath.Join(execDir, "static", "aa", filename)

	text := readFile(fullPath)
	fmt.Print(text)
}

func readFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return ""
	}
	return string(data)
}
