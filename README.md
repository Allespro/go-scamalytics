# go-scamalytics
<p align="center">
	<a href="https://golang.org"><img src="https://img.shields.io/badge/made%20with-Go-brightgreen"></a>
	<a href="https://goreportcard.com/report/github.com/allespro/go-scamalytics"><img src="https://goreportcard.com/badge/github.com/allespro/go-scamalytics"></a>
</p>

Check IP list on scamalytics.com exlude official API

![CVS output](https://raw.githubusercontent.com/Allespro/go-scamalytics/master/img/1.png)

go-scamalytics is a command-line application and package that allows users to check the fraud score of a list of IP addresses using scamalytics.com. The application works by scraping data from the scamalytics.com website without using the official API, which allows users to perform this check for free.

The app is designed for security professionals and developers who need to quickly check the fraud score of multiple IP addresses and generate a report in CSV format. The application is written in Go and can be compiled for multiple platforms, including Windows, Linux, and macOS.

## Package Usage
To use go-scamalytics package, use this example.

```go
package main

import (
	"fmt"
	goscamalytics_ipchecker "github.com/allespro/go-scamalytics/ipchecker"
)

func main() {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.3",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.1.1 Safari/605.1.1",
	}
	result := goscamalytics_ipchecker.CheckIP("8.8.8.8", userAgents)

	fmt.Printf("IP: %s\n", result["ip"])
	fmt.Printf("Score: %s\n", result["score"])
	fmt.Printf("Risk: %s\n", result["risk"])
}
```

## Cli Usage
To use go-scamalytics cli, use this example.

```go
package main

import (
	goscamalytics_cli "github.com/allespro/go-scamalytics/cli"
)

func main() {
	goscamalytics_cli.Start()
}

```

Once compiled, you can run the application from the command line. To see the available options, run `main --help`.

```
main -i inputFile.txt -o outputFile.csv -u useragents.txt
```

## License

This project is licensed under the MIT License - see the file [LICENSE](https://github.com/Allespro/go-scamalytics/blob/main/LICENSE) for details.
