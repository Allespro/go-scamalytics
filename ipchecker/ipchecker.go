package ipchecker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
)

func CheckIP(ip string, userAgentsList []string) map[string]any {
	userAgent := userAgentsList[rand.Intn(len(userAgentsList))]
	url := fmt.Sprintf("https://scamalytics.com/ip/%s", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Can`t init request: %s\n", err.Error())
		os.Exit(1)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Host = "scamalytics.com"
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to get response from the server: %s\n", err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Can't check IP %s - Response code: %d\n", ip, response.StatusCode)
		return nil
	}
	fmt.Printf("Checking IP %s\n", ip)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err.Error())
		os.Exit(1)
	}
	pattern := regexp.MustCompile(`(?s)<pre[^>]*>(.*?)<\/pre>`)
	matches := pattern.FindStringSubmatch(string(body))
	if len(matches) != 2 {
		fmt.Printf("Failed to find pattern in the response body\n")
		os.Exit(1)
	}

	rawData := []byte(matches[1])
	rawData = bytes.Replace(rawData, []byte("..."), []byte(""), 1)
	rawData = bytes.Replace(rawData, []byte("false,"), []byte("false"), 1)
	rawData = bytes.Replace(rawData, []byte("true,"), []byte("true"), 1)

	rawData = append([]byte("{"), rawData...)
	rawData = append(rawData, '}')

	var data map[string]any
	if err := json.Unmarshal(rawData, &data); err != nil {
		fmt.Printf("Failed to parse JSON response: %s\n", err.Error())
		os.Exit(1)
	}
	return data
}
