package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendFeedback 消息反馈（点赞）
func (c *Client) SendFeedback(messageID string, feedback *FeedbackRequest) error {
	data, err := json.Marshal(feedback)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s/%s%s", c.BaseURL, EndpointMessages, messageID, EndpointFeedbacks), bytes.NewBuffer(data))
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
		return fmt.Errorf("send feedback failed with status code: %d", resp.StatusCode)
	}

	return nil
}
