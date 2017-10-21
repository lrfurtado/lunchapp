package main

import (
	"reflect"
	"testing"
)

func TestCalcGroupDistribution(t *testing.T) {
	tt := []struct {
		Name                 string
		EmployeeCount        int
		ExpectedDistribution map[GroupSize]int
		ExpectedError        error
	}{
		{"No Employees", 0, nil, ErrInvalidEmployeeCount},
		{"Invalid Non Zero Employee Count", 2, nil, ErrInvalidEmployeeCount},
		{"Group size Multiple of IdealGroupSize", 40, map[GroupSize]int{SmallestGroupSize: 0, IdealGroupSize: 10, BiggestGroupSize: 0}, nil},
		{"Group size Multiple of SmallestGroupSize", 30, map[GroupSize]int{SmallestGroupSize: 2, IdealGroupSize: 6, BiggestGroupSize: 0}, nil},
		{"Group size Multiple of BiggestGroupSize", 55, map[GroupSize]int{SmallestGroupSize: 1, IdealGroupSize: 13, BiggestGroupSize: 0}, nil},
		{"One group with the BiggestGroupSize", 41, map[GroupSize]int{SmallestGroupSize: 0, IdealGroupSize: 9, BiggestGroupSize: 1}, nil},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			dist, err := calcGroupDistribution(tc.EmployeeCount)

			if err != tc.ExpectedError {
				t.Error("Invalid Error expected ", tc.ExpectedError, " got ", err)
			}
			if err == nil && !reflect.DeepEqual(dist, tc.ExpectedDistribution) {
				t.Error("Invalid Distribution expected ", tc.ExpectedDistribution, " got ", dist)
			}

		})
	}
}
