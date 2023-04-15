package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func checkIP(ip, useragentsFile string) map[string]interface{} {
	userAgents, err := ioutil.ReadFile(useragentsFile)
	if err != nil {
		fmt.Printf("Failed to read useragents file: %s\n", err.Error())
		os.Exit(1)
	}
	userAgentsList := strings.Split(string(userAgents), "\n")
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
	body, err := ioutil.ReadAll(response.Body)
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
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(matches[1]), &data); err != nil {
		fmt.Printf("Failed to parse JSON response: %s\n", err.Error())
		os.Exit(1)
	}
	return data
}

func addToCSV(filename string, ipInfo map[string]interface{}, delimiter string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s\n", filename, err.Error())
		os.Exit(1)
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "%s%s%s%s%s\n",
		ipInfo["ip"].(string), delimiter,
		ipInfo["score"].(string), delimiter,
		ipInfo["risk"].(string))
	if err != nil {
		fmt.Printf("Failed to write to file %s: %s\n", filename, err.Error())
		os.Exit(1)
	}
}

func loadIPFile(filename string) []string {
	var ipAddresses []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s\n", filename, err.Error())
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ipAddresses = append(ipAddresses, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file %s: %s\n", filename, err.Error())
		os.Exit(1)
	}
	return ipAddresses
}

func main() {
	banner := `::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
:::::::::::'###::::'##:::::'##:'########::'########:::::::::::
::::::::::'## ##::: ##:'##: ##:... ##..::'##. ##. ##::::::::::
:::::::::'##:. ##:: ##: ##: ##:::: ##:::: ##: ##:..:::::::::::
::::::::'##:::. ##: ##: ##: ##:::: ##::::. ########:::::::::::
:::::::: #########: ##: ##: ##:::: ##:::::... ##. ##::::::::::
:::::::: ##.... ##: ##: ##: ##:::: ##::::'##: ##: ##::::::::::
:::::::: ##:::: ##:. ###. ###::::: ##::::. ########:::::::::::
::::::::..:::::..:::...::...::::::..::::::........::::::::::::
:                                                            :
: Base Files: check_ip.txt / checked_ip.csv / useragents.txt :
: Version 0.1                                                :
: Check ip list on scamalytics.com exlude official API       :
:                                                            :
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
`
	fmt.Println(banner)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Check IP list on scamalytics.com exlude official API")
		fmt.Fprintln(os.Stderr, "Base Files: check_ip.txt / checked_ip.csv / useragents.txt")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}

	inputFile := flag.String("i", "check_ip.txt", "Input file")
	outputFile := flag.String("o", "checked_ip.csv", "Output file")
	useragentsFile := flag.String("u", "useragents.txt", "Useragents file (for html parsing)")
	delimiter := flag.String("d", ";", "Delimiter for CSV")

	flag.Parse()

	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}
	addToCSV(*outputFile, map[string]interface{}{"ip": "ip", "score": "score", "risk": "risk"}, *delimiter)
	ips := loadIPFile(*inputFile)
	for _, ip := range ips {
		ipInfo := checkIP(ip, *useragentsFile)
		if ipInfo != nil {
			addToCSV(*outputFile, ipInfo, *delimiter)
		} else {
			return
		}
	}
}
