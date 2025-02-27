package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ConversationsDel 删除会话
func (c *Client) ConversationsDel(conversationId string, user string) error {
	data, err := json.Marshal(map[string]string{
		"user": user,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s/%s", c.BaseURL, EndpointConversations, conversationId), bytes.NewBuffer(data))
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
