package ftr

import (
	"context"
)

type Hop struct {
	Host string
	Avg  float64
	Loss float64
}

type Report struct {
	Target string
	Hops   []Hop
}

type Runner interface {
	Run(ctx context.Context, target string, count int) (*Report, error)
}

type mtrRunner struct{}

func NewMTRRunner() Runner {
	return &mtrRunner{}
}

func (m *mtrRunner) Run(ctx context.Context, target string, count int) (*Report, error) {
	// Implementation of the Run method
	return &Report{Target: target, Hops: []Hop{}}, nil
}
