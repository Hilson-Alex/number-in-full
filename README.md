# number-in-full

A rewrite of the [number_in_full](https://github.com/Hilson-Alex/number_in_full) repository in Go.

It takes a number between 0 and 18.446.744.073.709.551.615 and translates to its full form in brazilian portuguese.

## Download

To download you can either clone the repository and build it with
```
> go build github.com/Hilson-Alex/number-in-full/src
```
Or download the [executable](out/number-in-full.exe) directly

## Execute

It must be executed through a cmd like this:
```
> ./number-in-full.exe <number>
```

## Running the tests

To run the automated tests, just run:
```
> go test number-in-full
```

the `numeric_consts_test.go` file test the normalize function used to correct small spelling issues on the 
generated name, while the `parser_test.go` file tests most all the functions needed to translate the number

## Author

- [Hilson-Alex](https://github.com/Hilson-Alex)
