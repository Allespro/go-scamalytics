package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/allespro/go-scamalytics/ipchecker"
)

func addToCSV(filename string, ipInfo map[string]any, delimiter string) {
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

func Start() {
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

	userAgents, err := os.ReadFile(*useragentsFile)

	if err != nil {
		fmt.Printf("Failed to read useragents file: %s\n", err.Error())
		os.Exit(1)
	}

	userAgentsList := strings.Split(string(userAgents), "\n")

	flag.Parse()

	if flag.NArg() > 0 {
		flag.Usage()
		os.Exit(1)
	}

	addToCSV(*outputFile, map[string]any{"ip": "ip", "score": "score", "risk": "risk"}, *delimiter)
	ips := loadIPFile(*inputFile)
	for _, ip := range ips {
		ipInfo := ipchecker.CheckIP(ip, userAgentsList)
		if ipInfo != nil {
			addToCSV(*outputFile, ipInfo, *delimiter)
		} else {
			return
		}
	}
}
