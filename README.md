# customer-importer

## Run program

```
FILENAME=data/customers.csv make run          
```

## Benchmarks
Remember to use 
```
cd customerimporter   
go test -bench=. 
```
to execute benchmarks.

It runs 1000, 3000 and 10000 batches.

Seconds at the end are total time to execute all of them.
## Unit tests
```
cd customerimporter   
go test
```
