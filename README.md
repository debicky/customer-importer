# customer-importer

## Features
- Reads customer data from a specified CSV file.
- Counts the number of customers using each email domain.
- Sorts and displays email domains along with their respective customer counts.
- Includes unit tests and performance benchmarks.

## Run the program
To run the program, use the following command:


```
FILENAME=data/customers.csv make run          
```
you can always provide your own dataset.

### Output
The program outputs a sorted list of email domains and the count of customers associated with each domain.

## Benchmarks
To evaluate the performance of the program, especially when handling large datasets, use the provided benchmark tests.

Remember to use 
```
cd customerimporter   
go test -bench=. 
```
to execute benchmarks.

The benchmarks will run with batches of 1000, 3000, and 10000 records. The output will include the time taken for each batch and the total time for all benchmarks.
## Unit tests
The program includes unit tests for validating the core functionalities.


```
cd customerimporter   
go test
```
