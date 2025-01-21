package dify

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// GetAppParameters 获取应用参数
func (c *Client) GetAppParameters() (*AppParameters, error) {
    req, err := http.NewRequest("GET", c.BaseURL+"/parameters", nil)
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
        return nil, fmt.Errorf("get app parameters failed with status code: %d", resp.StatusCode)
    }

    var result AppParameters
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}
