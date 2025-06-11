# Tetris Optimizer

This project assembles a list of tetrominoes into the smallest square
possible. The program reads a text file describing the pieces and
outputs the arrangement using capital letters.

## Building

```
go build
```

## Running

Provide a single argument, the path to the file containing the
tetrominoes. A sample input is provided in `sample.txt`:

```
go run . sample.txt
```

If the input file contains invalid data or the program encounters an
error, it prints `ERROR`.

## File Format

Pieces are written as 4 lines of 4 characters (`#` or `.`) and separated
by a blank line. For example:

```
#...
#...
#...
#...

....
....
..##
..##
```

## Testing

Run the unit tests with:

```
go test ./...
```
