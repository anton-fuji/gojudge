package problems

type Problem struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Difficulty  string     `json:"difficulty"`
	TestCases   []TestCase `json:"test_cases"`
	Template    string     `json:"template,omitempty"`
}

// テストケース用のストラクト
type TestCase struct {
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Name     string `json:"name,omitempty"`
}

// 実行用のストラクト
type Result struct {
	Problem    Problem
	Passed     bool
	FailedCase string
	Details    []TestResult
}

// 単一のテストケース用のストラクト
type TestResult struct {
	TestCase TestCase
	Output   string
	Passed   bool
	Error    string
}
