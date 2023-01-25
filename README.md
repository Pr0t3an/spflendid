# spflendid

Overview

This code is a command-line tool that checks if a list of domains have a valid Sender Policy Framework (SPF) record configured. The tool reads a list of domains from a specified input file, checks each domain's SPF record, and writes the results to a specified output file in CSV format. The tool can be run with a specified number of concurrent threads to speed up the process.

Requirements

Go 1.15 or higher

Building

To build the code, navigate to the directory containing the code and run the following command:


go build


To run the code, use the following command:


./spflendid -i [input_file] -o [output_file] -c [concurrency]

input_file is the file containing the list of domains, one per line. Default is "domains.txt"
output_file is the file where the results will be written in CSV format. Default is "results.csv"
concurrency is the number of concurrent threads to use. Default is 10
For example, if you want to use a file called "mydomains.txt" as input, "results.csv" as output, and use 15 concurrent threads, the command would be:

./spflendid -I mydomains.txt -o results.csv -c 15


The output file will be in CSV format and will contain the following columns:

Domain: The domain that was checked
SPF Configured: "Yes" if the domain has a valid SPF record, "No" otherwise.


Note

The code is using net.LookupTXT function to check the domain's SPF record. This function may be blocked in some network environment.
