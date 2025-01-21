package dify

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

// StopResponse 停止响应
func (c *Client) StopResponse(taskID string, user string) error {
    data, err := json.Marshal(map[string]string{
        "user": user,
    })
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", fmt.Sprintf("%s%s/%s/stop", c.BaseURL, EndpointCompletion, taskID), bytes.NewBuffer(data))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+c.APIKey)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("stop response failed with status code: %d", resp.StatusCode)
    }

    return nil
}
