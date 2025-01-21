package dify

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
)

// UploadFile 上传文件
func (c *Client) UploadFile(filePath string, user string) (*FileUploadResponse, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    // 添加文件
    part, err := writer.CreateFormFile("file", filepath.Base(filePath))
    if err != nil {
        return nil, err
    }
    _, err = io.Copy(part, file)
    if err != nil {
        return nil, err
    }

    // 添加user字段
    err = writer.WriteField("user", user)
    if err != nil {
        return nil, err
    }

    err = writer.Close()
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", c.BaseURL+EndpointFiles+"/upload", body)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Authorization", "Bearer "+c.APIKey)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("upload file failed with status code: %d", resp.StatusCode)
    }

    var result FileUploadResponse
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}
