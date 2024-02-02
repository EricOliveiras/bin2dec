# Bin2Dec

**Tier:** 1-Beginner

Binary is the number system all digital computers are based on.
Therefore it's important for developers to understand binary, or base 2,
mathematics. The purpose of Bin2Dec is to provide practice and
understanding of how binary calculations.

Bin2Dec allows the user to enter strings of up to 8 binary digits, 0's
and 1's, in any sequence and then displays its decimal equivalent.

This challenge requires that the developer implementing it follow these
constraints:

-   Arrays may not be used to contain the binary digits entered by the user
-   Determining the decimal equivalent of a particular binary digit in the
    sequence must be calculated using a single mathematical function, for
    example the natural logarithm. It's up to you to figure out which function
    to use.

## User Stories

-   [x] User can enter up to 8 binary digits in one input field
-   [x] User must be notified if anything other than a 0 or 1 was entered
-   [x] User views the results in a single output field containing the decimal (base 10) equivalent of the binary number that was entered

## Bonus features

-   [x] User can enter a variable number of binary digits


# How to run

## Prerequisites
You need to have the golang programming language installed - [Install here](https://go.dev/)

## Installation

1 - Clone de project repository:
```
git clone https://github.com/EricOliveiras/bin2dec.git
```
2 - Navigate to the project directory:
```
cd bin2dec
```

## Execution
Now you are ready to run Bin2Dec. Use the following command:
```
go run cmd/main.go
```

**Note**: This project is inspired by the ideas from this repository: [app-ideas by florinpop17](https://github.com/florinpop17/app-ideas)