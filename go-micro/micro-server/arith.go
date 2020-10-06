package main

import (
	"context"
	"github.com/ha666/benchmark/go-micro/proto/arith"
)

type Arith struct{}

func (e *Arith) Mul(ctx context.Context, req *arith.Args, res *arith.Reply) error {
	res.C = req.A * req.B
	return nil
}
