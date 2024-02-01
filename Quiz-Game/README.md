## Questions I had

1. Why there is no comma in betwwen the slice while printing the list?
   `[[5+5 10] [1+1 2] [8+3 11] [1+2 3] [8+6 14] [3+1 4] [1+4 5] [5+1 6] [2+3 5] [3+3 6] [2+4 6] [5+2 7]]`

```
Ans:
When you print a 2D slice in Go using fmt.Println,
it prints the inner slices without commas between
elements.

In your code, you are printing the first element
of the lines slice, which is itself a slice.
The default behavior of fmt.Println for
slices is to print the elements without commas.
```
