package runner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func ExecuteGoFile(filename, input string) (string, error) {
	// ファイル存在チェック
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("file %s does not exist", filename)
	}

	// 一時バイナリ名
	tmpBinary := filepath.Join(os.TempDir(), fmt.Sprintf("gojudge_%d", time.Now().UnixNano()))
	defer os.Remove(tmpBinary)

	// コンパイル
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	buildCmd := exec.CommandContext(ctx, "go", "build", "-o", tmpBinary, filename)
	var buildErr bytes.Buffer
	buildCmd.Stderr = &buildErr
	if err := buildCmd.Run(); err != nil {
		return "", fmt.Errorf("compilation failed: %s", buildErr.String())
	}

	// 実行
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	runCmd := exec.CommandContext(ctx, tmpBinary)
	runCmd.Stdin = strings.NewReader(input)
	var out, runErr bytes.Buffer
	runCmd.Stdout = &out
	runCmd.Stderr = &runErr
	if err := runCmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out")
		}
		return "", fmt.Errorf("runtime error: %s", runErr.String())
	}

	return strings.TrimSpace(out.String()), nil
}
