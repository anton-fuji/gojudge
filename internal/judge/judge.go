package judge

import (
	"fmt"
	"os"
	"strings"

	"github.com/anton-fuji/gojudge/internal/problems"
	"github.com/anton-fuji/gojudge/internal/runner"
)

type Result = problems.Result

func CheckSolution(filename string, verbose bool, problemID string) (*Result, error) {
	problem, err := problems.GetProblemByID(problemID)
	if err != nil || problem == nil {
		return nil, fmt.Errorf("problem %s not found", problemID)
	}

	code, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if !isSafeCode(string(code)) {
		return nil, fmt.Errorf("unsafe code detected")
	}

	result := &Result{Problem: *problem, Passed: true, Details: make([]problems.TestResult, 0)}
	for i, tc := range problem.TestCases {
		output, err := runner.ExecuteGoFile(filename, tc.Input)
		testResult := problems.TestResult{TestCase: tc, Output: output}
		if err != nil {
			testResult.Error = err.Error()
			testResult.Passed = false
			result.Passed = false
			result.FailedCase = fmt.Sprintf("Test %d: %s", i+1, err.Error())
		} else if strings.TrimSpace(output) != strings.TrimSpace(tc.Expected) {
			testResult.Passed = false
			result.Passed = false
			result.FailedCase = fmt.Sprintf("Test %d: expected %q, got %q", i+1, tc.Expected, output)
		} else {
			testResult.Passed = true
		}
		result.Details = append(result.Details, testResult)
		if !testResult.Passed && !verbose {
			break
		}
	}
	return result, nil
}

func isSafeCode(code string) bool {
	return !strings.Contains(code, "import \"os\"") &&
		!strings.Contains(code, "import \"net\"") &&
		!strings.Contains(code, "import \"syscall\"")
}
