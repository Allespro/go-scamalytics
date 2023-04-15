# go-scamalytics
Check IP list on scamalytics.com exlude official API

![Console output](https://raw.githubusercontent.com/Allespro/go-scamalytics/master/img/1.png)


![CVS output](https://raw.githubusercontent.com/Allespro/go-scamalytics/master/img/2.png)

go-scamalytics is a command-line application that allows users to check the fraud score of a list of IP addresses using scamalytics.com. The application works by scraping data from the scamalytics.com website without using the official API, which allows users to perform this check for free.

The app is designed for security professionals and developers who need to quickly check the fraud score of multiple IP addresses and generate a report in CSV format. The application is written in Go and can be compiled for multiple platforms, including Windows, Linux, and macOS.

## Build
To install go-scamalytics, you will first need to clone the repository from GitHub and build them.

```
git clone https://github.com/Allespro/go-scamalytics
cd go-scamalytics
make
```

After running the make command, you can find the compiled binary files in the `build` directory within the project directory. These binary files can be executed directly on their respective platforms.

## Usage
Once compiled, you can run the application from the command line. To see the available options, run `go-scamalytics --help`.

```
go-scamalytics -i inputFile.txt -o outputFile.csv -u useragents.txt
```

## License

This project is licensed under the MIT License - see the file [LICENSE](https://github.com/Allespro/go-scamalytics/blob/main/LICENSE) for details.
