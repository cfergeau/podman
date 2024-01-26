package types

import (
	commonTypes "github.com/containers/common/libnetwork/types"
)

// NetworkPruneReport containers the name of network and an error
// associated in its pruning (removal)
// swagger:model NetworkPruneReport
type NetworkPruneReport struct {
	Name  string
	Error error
}

// NetworkReloadReport describes the results of reloading a container network.
type NetworkReloadReport struct {
	//nolint:stylecheck,revive
	Id  string
	Err error
}

// NetworkConnectOptions describes options for connecting
// a container to a network
type NetworkConnectOptions struct {
	Container string `json:"container"`
	commonTypes.PerNetworkOptions
}

// NetworkRmReport describes the results of network removal
type NetworkRmReport struct {
	Name string
	Err  error
}

type NetworkCreateReport struct {
	Name string
}
