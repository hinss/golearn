package rpc_watcher

import "net/rpc"

const SERVICE_NAME = "path/to/pkg.KVStoreService"

type KVStoreServiceInterface interface {
	Get(string, *string) error
	Set([2]string, *struct{}) error
	Watch(int, *string) error
}

func RegisterKVStoreService(svc KVStoreServiceInterface) error {
	return rpc.RegisterName(SERVICE_NAME, svc)
}
