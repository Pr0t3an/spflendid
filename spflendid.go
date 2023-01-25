package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	// Command-line flags
	inputFile := flag.String("i", "domains.txt", "Input file containing the domains, one per line")
	outputFile := flag.String("o", "results.csv", "Output file in CSV format")
	concurrency := flag.Int("c", 10, "Number of concurrent threads")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Open or create the output file
	csvFile, err := os.OpenFile(*outputFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	// Write the header row
	csvWriter.Write([]string{"Domain", "SPF Configured"})

	// Create a channel to receive domains from the input file
	domainChan := make(chan string, *concurrency)
	// Create a wait group to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// Start the specified number of concurrent goroutines
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			for domain := range domainChan {
				_, addrs, err := net.LookupTXT(domain)
				if err != nil {
					fmt.Println(err)
					continue
				}
				var spfStatus string
				for _, addr := range addrs {
					if addr == "v=spf1" {
						spfStatus = "Yes"
						break
					}
				}
				if spfStatus == "" {
					spfStatus = "No"
				}
				csvWriter.Write([]string{domain, spfStatus})
			}
			wg.Done()
		}()
	}

	// Read the input file and send the domains to the channel
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		domain := scanner.Text()
		domainChan <- domain
		fmt.Printf("\rProcessing domain %d", i)
	}
	close(domainChan)
	wg.Wait()

	fmt.Println("\nDone!")
}
