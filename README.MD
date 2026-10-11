# ESEP Grade Calculator

Go library that calculates the final letter grade of a given course.

## Runtime requirements:
- Go 1.24+

## Grade weights
- Assignments: 50%
- Exams: 35%
- Essays: 15%

## Letter scale

A >= 90
B >= 80
C >= 70 (C is passing)
D >= 60
F <  60

## Usage
```go
package main

import (
    "fmt"

    esepunittests "esep/grade-calculator"
)

func main() {
    gc := esepunittests.NewGradeCalculator()
    gc.AddGrade("Homework Assignment", 100, esepunittests.Assignment)
    gc.AddGrade("Exam", 99, esepunittests.Exam)
    gc.AddGrade("Essay", 98, esepunittests.Essay)
    fmt.Println(gc.GetFinalGrade())
}
```