package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TextToSpeech 文字转语音
func (c *Client) TextToSpeech(request *TTSRequest) ([]byte, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+EndpointAudio, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("text to speech failed with status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
