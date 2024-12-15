
package gohelpers

import (
    "testing"
    "reflect"
    "math"
	"sort"
)

func TestMin(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 5, 3, 3},
        {"negative numbers", -5, -3, -5},
        {"mixed numbers", -5, 3, -5},
        {"equal numbers", 5, 5, 5},
        {"zero and positive", 0, 5, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Min(tt.a, tt.b); got != tt.expected {
                t.Errorf("Min(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
            }
        })
    }
}

func TestMax(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 5, 3, 5},
        {"negative numbers", -5, -3, -3},
        {"mixed numbers", -5, 3, 3},
        {"equal numbers", 5, 5, 5},
        {"zero and positive", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Max(tt.a, tt.b); got != tt.expected {
                t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
            }
        })
    }
}

func TestMinInSlice(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        expected int
        wantErr  bool
    }{
        {"normal slice", []int{5, 3, 7, 1, 9}, 1, false},
        {"single element", []int{5}, 5, false},
        {"negative numbers", []int{-5, -3, -7}, -7, false},
        {"empty slice", []int{}, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := MinInSlice(tt.slice)
            if (err != nil) != tt.wantErr {
                t.Errorf("MinInSlice() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && got != tt.expected {
                t.Errorf("MinInSlice() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestUnique(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        expected []int
    }{
        {"integers with duplicates", []int{1, 2, 2, 3, 3, 4}, []int{1, 2, 3, 4}},
        {"no duplicates", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
        {"empty slice", []int{}, []int{}},
        {"all duplicates", []int{1, 1, 1}, []int{1}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Unique(tt.slice)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Unique() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestReverse(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        expected []int
    }{
        {"normal slice", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
        {"single element", []int{1}, []int{1}},
        {"empty slice", []int{}, []int{}},
        {"two elements", []int{1, 2}, []int{2, 1}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Reverse(tt.slice)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Reverse() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestChunk(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        size     int
        expected [][]int
    }{
        {"normal chunk", []int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
        {"size equals length", []int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
        {"empty slice", []int{}, 2, [][]int{}},
        {"size zero", []int{1, 2, 3}, 0, [][]int{}},
        {"size one", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Chunk(tt.slice, tt.size)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Chunk() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestRange(t *testing.T) {
    tests := []struct {
        name          string
        start, end    int
        expected      []int
    }{
        {"normal range", 0, 5, []int{0, 1, 2, 3, 4}},
        {"negative to positive", -2, 2, []int{-2, -1, 0, 1}},
        {"single number range", 1, 2, []int{1}},
        {"invalid range", 5, 2, []int{}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Range(tt.start, tt.end)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Range() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestMap(t *testing.T) {
    double := func(x int) int { return x * 2 }
    
    tests := []struct {
        name     string
        slice    []int
        f        func(int) int
        expected []int
    }{
        {"double numbers", []int{1, 2, 3}, double, []int{2, 4, 6}},
        {"empty slice", []int{}, double, []int{}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Map(tt.slice, tt.f)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Map() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestRoundToDecimals(t *testing.T) {
    tests := []struct {
        name     string
        x        float64
        decimals int
        expected float64
    }{
        {"round to 2 decimals", 3.14159, 2, 3.14},
        {"round to 1 decimal", 3.14159, 1, 3.1},
        {"round up", 3.16, 1, 3.2},
        {"zero decimals", 3.14159, 0, 3.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := RoundToDecimals(tt.x, tt.decimals)
            if math.Abs(got-tt.expected) > 1e-10 {
                t.Errorf("RoundToDecimals() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestIntersection(t *testing.T) {
    tests := []struct {
        name     string
        a, b     []int
        expected []int
    }{
        {"normal intersection", []int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
        {"no intersection", []int{1, 2}, []int{3, 4}, []int{}},
        {"empty slices", []int{}, []int{}, []int{}},
        {"one empty slice", []int{1, 2}, []int{}, []int{}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Intersection(tt.a, tt.b)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Intersection() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestContains(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        element  int
        expected bool
    }{
        {"element exists", []int{1, 2, 3}, 2, true},
        {"element doesn't exist", []int{1, 2, 3}, 4, false},
        {"empty slice", []int{}, 1, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Contains(tt.slice, tt.element)
            if got != tt.expected {
                t.Errorf("Contains() = %v, want %v", got, tt.expected)
            }
        })
    }
}


func TestSum(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        expected int
    }{
        {"positive numbers", []int{1, 2, 3, 4, 5}, 15},
        {"negative numbers", []int{-1, -2, -3}, -6},
        {"mixed numbers", []int{-1, 0, 1}, 0},
        {"empty slice", []int{}, 0},
        {"single number", []int{5}, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Sum(tt.slice)
            if got != tt.expected {
                t.Errorf("Sum() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestAverage(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        expected float64
        wantErr  bool
    }{
        {"positive numbers", []int{1, 2, 3, 4, 5}, 3.0, false},
        {"negative numbers", []int{-1, -2, -3}, -2.0, false},
        {"mixed numbers", []int{-1, 0, 1}, 0.0, false},
        {"empty slice", []int{}, 0.0, true},
        {"single number", []int{5}, 5.0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Average(tt.slice)
            if (err != nil) != tt.wantErr {
                t.Errorf("Average() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !tt.wantErr && math.Abs(got-tt.expected) > 1e-10 {
                t.Errorf("Average() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestShuffle(t *testing.T) {
    original := []int{1, 2, 3, 4, 5}
    result := Shuffle(original)

    // Check length is same
    if len(result) != len(original) {
        t.Errorf("Shuffle() returned slice of length %v, want %v", len(result), len(original))
    }

    originalMap := make(map[int]bool)
    resultMap := make(map[int]bool)
    for _, v := range original {
        originalMap[v] = true
    }
    for _, v := range result {
        resultMap[v] = true
    }
    if !reflect.DeepEqual(originalMap, resultMap) {
        t.Errorf("Shuffle() returned slice with different elements")
    }
}

func TestFilter(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        f        func(int) bool
        expected []int
    }{
        {
            "even numbers",
            []int{1, 2, 3, 4, 5, 6},
            func(x int) bool { return x%2 == 0 },
            []int{2, 4, 6},
        },
        {
            "positive numbers",
            []int{-1, 0, 1, 2, -3},
            func(x int) bool { return x > 0 },
            []int{1, 2},
        },
        {
            "empty slice",
            []int{},
            func(x int) bool { return x > 0 },
            []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Filter(tt.slice, tt.f)
            if !reflect.DeepEqual(got, tt.expected) {
                t.Errorf("Filter() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestReduce(t *testing.T) {
    tests := []struct {
        name     string
        slice    []int
        initial  int
        f        func(int, int) int
        expected int
    }{
        {
            "sum reduction",
            []int{1, 2, 3, 4},
            0,
            func(acc, curr int) int { return acc + curr },
            10,
        },
        {
            "multiplication reduction",
            []int{1, 2, 3, 4},
            1,
            func(acc, curr int) int { return acc * curr },
            24,
        },
        {
            "empty slice",
            []int{},
            0,
            func(acc, curr int) int { return acc + curr },
            0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Reduce(tt.slice, tt.initial, tt.f)
            if got != tt.expected {
                t.Errorf("Reduce() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestJoinAndSplit(t *testing.T) {
    tests := []struct {
        name           string
        elements       []string
        separator      string
        keepEmpty     bool
        expectedJoin  string
        expectedSplit []string
    }{
        {
            "basic join and split",
            []string{"a", "b", "c"},
            ",",
            true,
            "a,b,c",
            []string{"a", "b", "c"},
        },
        {
            "empty strings",
            []string{"", "b", "", "c", ""},
            ",",
            true,
            ",b,,c,",
            []string{"", "b", "", "c", ""},
        },
        {
            "empty strings filtered",
            []string{"", "b", "", "c", ""},
            ",",
            false,
            ",b,,c,",
            []string{"b", "c"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            joined := Join(tt.elements, tt.separator)
            if joined != tt.expectedJoin {
                t.Errorf("Join() = %v, want %v", joined, tt.expectedJoin)
            }

            split := Split(joined, tt.separator, tt.keepEmpty)
            if !reflect.DeepEqual(split, tt.expectedSplit) {
                t.Errorf("Split() = %v, want %v", split, tt.expectedSplit)
            }
        })
    }
}

func TestIsNumeric(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {"only digits", "12345", true},
        {"mixed content", "123abc", false},
        {"empty string", "", true},
        {"special chars", "123.45", false},
        {"negative number", "-123", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := IsNumeric(tt.input)
            if got != tt.expected {
                t.Errorf("IsNumeric() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestKeysAndValues(t *testing.T) {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    t.Run("test Keys", func(t *testing.T) {
        keys := Keys(m)
        expectedLen := 3
        if len(keys) != expectedLen {
            t.Errorf("Keys() returned %d keys, want %d", len(keys), expectedLen)
        }
        // Check all keys exist
        for _, k := range keys {
            if _, exists := m[k]; !exists {
                t.Errorf("Keys() returned key %v that doesn't exist in map", k)
            }
        }
    })

    t.Run("test Values", func(t *testing.T) {
        values := Values(m)
        expectedLen := 3
        if len(values) != expectedLen {
            t.Errorf("Values() returned %d values, want %d", len(values), expectedLen)
        }
        // Create a map of values for easy lookup
        valueMap := make(map[int]bool)
        for _, v := range values {
            valueMap[v] = true
        }
        // Check all expected values exist
        for _, v := range []int{1, 2, 3} {
            if !valueMap[v] {
                t.Errorf("Values() missing value %v", v)
            }
        }
    })
}

func TestGroupBy(t *testing.T) {
    type Person struct {
        Name string
        Age  int
    }

    people := []Person{
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 25},
        {"David", 30},
    }

    tests := []struct {
        name     string
        slice    []Person
        keyFunc  func(Person) int
        expected map[int][]Person
    }{
        {
            "group by age",
            people,
            func(p Person) int { return p.Age },
            map[int][]Person{
                25: {{"Alice", 25}, {"Charlie", 25}},
                30: {{"Bob", 30}, {"David", 30}},
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := GroupBy(tt.slice, tt.keyFunc)
            if len(got) != len(tt.expected) {
                t.Errorf("GroupBy() returned map of size %d, want %d", len(got), len(tt.expected))
            }
            for k, v := range got {
                expectedGroup, exists := tt.expected[k]
                if !exists {
                    t.Errorf("GroupBy() returned unexpected key %v", k)
                    continue
                }
                if !reflect.DeepEqual(v, expectedGroup) {
                    t.Errorf("GroupBy()[%v] = %v, want %v", k, v, expectedGroup)
                }
            }
        })
    }
}

func TestUnion(t *testing.T) {
    tests := []struct {
        name     string
        a, b     []int
        expected []int
    }{
        {"different elements", []int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
        {"overlapping elements", []int{1, 2, 3}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
        {"empty slices", []int{}, []int{}, []int{}},
        {"one empty slice", []int{1, 2}, []int{}, []int{1, 2}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Union(tt.a, tt.b)
            // Sort both slices for comparison since Union doesn't guarantee order
            sort.Ints(got)
            expected := make([]int, len(tt.expected))
            copy(expected, tt.expected)
            sort.Ints(expected)
            if !reflect.DeepEqual(got, expected) {
                t.Errorf("Union() = %v, want %v", got, expected)
            }
        })
    }
}