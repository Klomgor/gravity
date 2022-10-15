package discovery

import (
	"time"

	"beryju.io/gravity/pkg/roles/discovery/types"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func (r *Role) startDiscovery(raw *mvccpb.KeyValue) {
	sub, err := r.subnetFromKV(raw)
	if err != nil {
		r.log.Warn("failed to parse subnet", zap.Error(err))
		return
	}
	go sub.RunDiscovery()
}

func (r *Role) startWatchSubnets() {
	prefix := r.i.KV().Key(types.KeyRole, types.KeySubnets).Prefix(true).String()
	subnets, err := r.i.KV().Get(r.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		r.log.Warn("failed to list initial subnets", zap.Error(err))
		time.Sleep(5 * time.Second)
		r.startWatchSubnets()
		return
	}
	for _, subnet := range subnets.Kvs {
		r.startDiscovery(subnet)
	}

	watchChan := r.i.KV().Watch(
		r.ctx,
		prefix,
		clientv3.WithPrefix(),
	)
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			if event.Type == clientv3.EventTypeDelete {
				continue
			}
			r.startDiscovery(event.Kv)
		}
	}
}
