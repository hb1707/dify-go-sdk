package dify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAppInfo 获取应用基本信息
func (c *Client) GetAppInfo() (*AppInfo, error) {
	req, err := http.NewRequest("GET", c.BaseURL+EndpointInfo, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get app info failed with status code: %d", resp.StatusCode)
	}

	var result AppInfo
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
