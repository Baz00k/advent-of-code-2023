# Advent of Code 2023

My solutions to the [Advent of Code 2023](https://adventofcode.com/2023) puzzles.
Please, bear in mind that these are first lines of GO I've ever written, so the solutions are probably not the most elegant ones.

## Usage

To run the solution for a given day, just run the following command:

```bash
go run . -day <day> solve
```

If you want to run the tests for a given day, run the following command:

```bash
go test ./day<day>
```

If you want to scaffold a new day, run the following command:

```bash
go run . -day <day> scaffold
```

## Visual Studio Code

Launching: press F5 to launch the debugger.

Testing: press Ctrl+Shift+P and select "Go: Test Package" to run the tests for the current day.

Scalfolding: press Ctrl+Shift+P and select "Tasks: Run Task" and then "Scaffold Day" to scaffold a new day.
