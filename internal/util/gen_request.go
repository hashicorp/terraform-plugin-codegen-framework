package util

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

func Request(command func(timestamp, accessKey, signature string) *exec.Cmd, method, url, accessKey, secretKey, requestBody string) (map[string]interface{}, error) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	signature := makeSignature(method, url, timestamp, accessKey, secretKey)

	cmd := command(timestamp, accessKey, signature)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, err
	}

	// code 200 but error occurs
	if result["error"] != nil {
		return result, fmt.Errorf("error with code 200: %s", result["error"])
	}

	return result, nil
}
