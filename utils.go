
package gohelpers

import (
    "math"
    "sort"
    "strings"
)

// Min returns the smallest of two integers
func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Max returns the largest of two integers
func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// MinInSlice returns the smallest number in a slice
func MinInSlice(nums []int) (int, error) {
    if len(nums) == 0 {
        return 0, fmt.Errorf("empty slice")
    }
    min := nums[0]
    for _, num := range nums[1:] {
        if num < min {
            min = num
        }
    }
    return min, nil
}

// MaxInSlice returns the largest number in a slice
func MaxInSlice(nums []int) (int, error) {
    if len(nums) == 0 {
        return 0, fmt.Errorf("empty slice")
    }
    max := nums[0]
    for _, num := range nums[1:] {
        if num > max {
            max = num
        }
    }
    return max, nil
}

// Sum returns the sum of a slice of integers
func Sum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

// Average returns the average of a slice of integers
func Average(nums []int) (float64, error) {
    if len(nums) == 0 {
        return 0, fmt.Errorf("empty slice")
    }
    sum := Sum(nums)
    return float64(sum) / float64(len(nums)), nil
}

// Contains checks if an element exists in a slice
func Contains[T comparable](slice []T, element T) bool {
    for _, v := range slice {
        if v == element {
            return true
        }
    }
    return false
}

// Unique returns a new slice with duplicate elements removed
func Unique[T comparable](slice []T) []T {
    seen := make(map[T]bool)
    result := make([]T, 0)
    
    for _, item := range slice {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }
    return result
}

// Reverse returns a new slice with elements in reverse order
func Reverse[T any](slice []T) []T {
    result := make([]T, len(slice))
    for i, v := range slice {
        result[len(slice)-1-i] = v
    }
    return result
}

// Shuffle randomly reorders elements in a slice
func Shuffle[T any](slice []T) []T {
    result := make([]T, len(slice))
    copy(result, slice)
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(result), func(i, j int) {
        result[i], result[j] = result[j], result[i]
    })
    return result
}

// Chunk splits a slice into chunks of specified size
func Chunk[T any](slice []T, size int) [][]T {
    if size <= 0 {
        return [][]T{}
    }
    
    chunks := make([][]T, 0, (len(slice)+size-1)/size)
    
    for size < len(slice) {
        slice, chunks = slice[size:], append(chunks, slice[0:size:size])
    }
    chunks = append(chunks, slice)
    return chunks
}

// Range creates a slice of numbers from start to end (exclusive)
func Range(start, end int) []int {
    if start >= end {
        return []int{}
    }
    
    result := make([]int, end-start)
    for i := range result {
        result[i] = start + i
    }
    return result
}

// Map applies a function to each element in a slice
func Map[T, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}

// Filter returns a new slice with elements that pass the test
func Filter[T any](slice []T, f func(T) bool) []T {
    result := make([]T, 0)
    for _, v := range slice {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

// Reduce reduces a slice to a single value using a function
func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
    result := initial
    for _, v := range slice {
        result = f(result, v)
    }
    return result
}

// Join concatenates strings with a separator (like Python's join)
func Join(elements []string, separator string) string {
    return strings.Join(elements, separator)
}

// Split splits a string by separator (with option to keep empty strings)
func Split(s, separator string, keepEmpty bool) []string {
    if keepEmpty {
        return strings.Split(s, separator)
    }
    return strings.FieldsFunc(s, func(r rune) bool {
        return string(r) == separator
    })
}

// IsNumeric checks if a string contains only numeric characters
func IsNumeric(s string) bool {
    for _, r := range s {
        if r < '0' || r > '9' {
            return false
        }
    }
    return true
}

// RoundToDecimals rounds a float64 to n decimal places
func RoundToDecimals(x float64, decimals int) float64 {
    pow := math.Pow10(decimals)
    return math.Round(x*pow) / pow
}

// Keys returns all keys from a map
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

// Values returns all values from a map
func Values[K comparable, V any](m map[K]V) []V {
    values := make([]V, 0, len(m))
    for _, v := range m {
        values = append(values, v)
    }
    return values
}

// GroupBy groups slice elements by a key function
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
    result := make(map[K][]T)
    for _, item := range slice {
        key := keyFunc(item)
        result[key] = append(result[key], item)
    }
    return result
}

// Intersection returns common elements between two slices
func Intersection[T comparable](a, b []T) []T {
    set := make(map[T]bool)
    result := make([]T, 0)
    
    for _, item := range a {
        set[item] = true
    }
    
    for _, item := range b {
        if set[item] {
            result = append(result, item)
        }
    }
    return result
}

// Union returns unique elements from both slices
func Union[T comparable](a, b []T) []T {
    set := make(map[T]bool)
    
    for _, item := range a {
        set[item] = true
    }
    for _, item := range b {
        set[item] = true
    }
    
    result := make([]T, 0, len(set))
    for item := range set {
        result = append(result, item)
    }
    return result
}