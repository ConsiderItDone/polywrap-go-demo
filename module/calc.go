package module

import "github.com/consideritdone/polywrap-go-demo/wrap/types"

//go:generate polywrap codegen

func Add(args *types.ArgsAdd) int32 {
	return args.A + args.B
}

func Sub(args *types.ArgsSub) int32 {
	return args.A - args.B
}
