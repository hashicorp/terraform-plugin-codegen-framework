package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

func Request(command func(timestamp, accessKey, signature string) *exec.Cmd, requestBody string) (map[string]interface{}, error) {
	timestamp := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))
	accessKey := "hello world"
	secretKey := "bye world"
	signature := generateSignature(requestBody, secretKey)

	cmd := command(timestamp, accessKey, signature)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// generateSignature creates a HMAC-SHA256 signature for the given data
func generateSignature(data, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
