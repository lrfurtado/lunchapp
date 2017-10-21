package main

import "errors"

type GroupSize int

const (
	SmallestGroupSize GroupSize = iota + 3 //3
	IdealGroupSize                         //4
	BiggestGroupSize                       //5
)

var ErrInvalidEmployeeCount = errors.New("Invalid Employee Count")

func calcGroupDistribution(employeeCount int) (map[GroupSize]int, error) {
	dist := make(map[GroupSize]int)
	if employeeCount < int(SmallestGroupSize) {
		return nil, ErrInvalidEmployeeCount
	}

	dist[IdealGroupSize] = employeeCount / int(IdealGroupSize)
	dist[SmallestGroupSize] = 0
	dist[BiggestGroupSize] = 0
	mod := employeeCount % int(IdealGroupSize)
	switch mod {
	case 1:
		dist[BiggestGroupSize] = 1
		dist[IdealGroupSize] -= 1
	case 2:
		dist[SmallestGroupSize] = 2
		dist[IdealGroupSize] -= 1
	case 3:
		dist[SmallestGroupSize] = 1
	}

	return dist, nil
}
