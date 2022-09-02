package module

import "github.com/consideritdone/polywrap-go-demo/wrap/types"

//go:generate polywrap build -v -m ../polywrap.yaml -o ../build

func Add(args *types.ArgsAdd) int32 {
	return args.A + args.B
}

func Sub(args *types.ArgsSub) int32 {
	return args.A - args.B
}
