package utils

import "github.com/bwmarrin/snowflake"

var (
	node *snowflake.Node
	err  error
)

func init() {
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}

func Gen() uint64 {
	return uint64(node.Generate().Int64())
}
