package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

// GetClusterNodes returns information about all the nodes participating in the cluster.
func (cl *Client) GetClusterNodes(ctx context.Context) (out []*GetClusterNodesResult, err error) {
	err = cl.rpcClient.CallForInto(ctx, &out, "getClusterNodes", nil)
	return
}

type GetClusterNodesResult struct {
	// Node public key.
	Pubkey solana.PublicKey `json:"pubkey"`

	// Gossip network address for the node.
	Gossip *string `json:"gossip"`

	// TPU network address for the node.
	TPU *string `json:"tpu"`

	// TODO: "" or nil ?

	// JSON RPC network address for the node, or empty if the JSON RPC service is not enabled.
	RPC *string `json:"rpc"`

	// The software version of the node, or empty if the version information is not available.
	Version *string `json:"version"`

	// TODO: what type is this?
	// The unique identifier of the node's feature set.
	FeatureSet int64 `json:"featureSet"`

	// The shred version the node has been configured to use.
	ShredVersion int64 `json:"shredVersion"`
}
