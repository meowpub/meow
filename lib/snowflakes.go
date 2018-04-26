package lib

import (
	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

// GenSnowflake generates a snowflake ID. The nodeID is only used on the first call, to create a
// cached Node - the algorithm relies on there being a monotonous, global step counter.
func GenSnowflake(nodeID int64) (snowflake.ID, error) {
	if snowflakeNode == nil {
		node, err := snowflake.NewNode(nodeID)
		if err != nil {
			return 0, err
		}
		snowflakeNode = node
	}
	return snowflakeNode.Generate(), nil
}
