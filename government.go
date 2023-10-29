package main

import (
	"sync"
)

// GovernmentInfo represents information about the government.
type GovernmentInfo struct {
	Name        string
	President   string
	Capital     string
	Established int
}

// governmentSingleton is a struct that holds the single instance of GovernmentInfo.
type governmentSingleton struct {
	instance *GovernmentInfo
	once     sync.Once
}

var singletonInstance *governmentSingleton

// GetGovernmentInfo returns the single instance of GovernmentInfo, creating it if necessary.
func (gs *governmentSingleton) GetGovernmentInfo() *GovernmentInfo {
	gs.once.Do(func() {
		gs.instance = &GovernmentInfo{
			Name:        "United States of America",
			President:   "Joe Biden",
			Capital:     "Washington, D.C.",
			Established: 1776,
		}
	})
	return gs.instance
}

func newGovernmentSingleton() *governmentSingleton {
	return &governmentSingleton{}
}
