# Go Helpers

A collection of utility functions for Go, inspired by convenient helpers found in Python and JavaScript.

## Installation

```bash
go get github.com/yourusername/go_helpers
```

## Usage

```go
import "github.com/yourusername/go_helpers"
```

## Available Functions

### Math Operations
- `Min(a, b int)`: Returns the smaller of two integers
- `Max(a, b int)`: Returns the larger of two integers
- `MinInSlice(nums []int)`: Returns the smallest number in a slice
- `MaxInSlice(nums []int)`: Returns the largest number in a slice
- `Sum(nums []int)`: Returns the sum of all numbers in a slice
- `Average(nums []int)`: Calculates the average of numbers in a slice
- `RoundToDecimals(x float64, decimals int)`: Rounds a float to specified decimal places

### Slice Operations
- `Contains[T comparable](slice []T, element T)`: Checks if an element exists in a slice
- `Unique[T comparable](slice []T)`: Removes duplicate elements from a slice
- `Reverse[T any](slice []T)`: Reverses the order of elements in a slice
- `Shuffle[T any](slice []T)`: Randomly reorders elements in a slice
- `Chunk[T any](slice []T, size int)`: Splits a slice into smaller chunks of specified size
- `Range(start, end int)`: Creates a slice of numbers from start to end (exclusive)
- `Intersection[T comparable](a, b []T)`: Returns elements that exist in both slices
- `Union[T comparable](a, b []T)`: Returns unique elements from both slices

### Functional Programming
- `Map[T, U any](slice []T, f func(T) U)`: Applies a function to each element in a slice
- `Filter[T any](slice []T, f func(T) bool)`: Returns elements that pass a test function
- `Reduce[T, U any](slice []T, initial U, f func(U, T) U)`: Reduces a slice to a single value
- `GroupBy[T any, K comparable](slice []T, keyFunc func(T) K)`: Groups slice elements by a key function

### String Operations
- `Join(elements []string, separator string)`: Joins strings with a separator
- `Split(s, separator string, keepEmpty bool)`: Splits a string by separator
- `IsNumeric(s string)`: Checks if a string contains only numeric characters

### Map Operations
- `Keys[K comparable, V any](m map[K]V)`: Returns all keys from a map
- `Values[K comparable, V any](m map[K]V)`: Returns all values from a map

## Examples

```go
// Using Min/Max
min := gohelpers.Min(5, 3)  // Returns 3
max := gohelpers.Max(5, 3)  // Returns 5

// Using slice operations
nums := []int{1, 2, 2, 3, 3, 4}
unique := gohelpers.Unique(nums)  // Returns [1, 2, 3, 4]

// Using Map
doubled := gohelpers.Map([]int{1, 2, 3}, func(x int) int {
    return x * 2
})  // Returns [2, 4, 6]

// Using Filter
evens := gohelpers.Filter([]int{1, 2, 3, 4}, func(x int) bool {
    return x%2 == 0
})  // Returns [2, 4]
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.