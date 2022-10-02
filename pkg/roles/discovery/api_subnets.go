package discovery

import (
	"context"

	"beryju.io/gravity/pkg/roles/discovery/types"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type APISubnet struct {
	Name string `json:"name" required:"true"`

	CIDR         string `json:"subnetCidr" required:"true"`
	DNSResolver  string `json:"dnsResolver" required:"true"`
	DiscoveryTTL int    `json:"discoveryTTL" required:"true"`
}
type APISubnetsGetOutput struct {
	Subnets []APISubnet `json:"subnets"`
}

func (r *Role) APISubnetsGet() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *APISubnetsGetOutput) error {
		prefix := r.i.KV().Key(types.KeyRole, types.KeySubnets).Prefix(true).String()
		subnets, err := r.i.KV().Get(ctx, prefix, clientv3.WithPrefix())
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		for _, rawSub := range subnets.Kvs {
			sub, err := r.subnetFromKV(rawSub)
			if err != nil {
				r.log.WithError(err).Warning("failed to parse subnet")
				continue
			}
			output.Subnets = append(output.Subnets, APISubnet{
				Name:         sub.Identifier,
				CIDR:         sub.CIDR,
				DNSResolver:  sub.DNSResolver,
				DiscoveryTTL: sub.DiscoveryTTL,
			})
		}
		return nil
	})
	u.SetName("discovery.get_subnets")
	u.SetTitle("Discovery subnets")
	u.SetTags("roles/discovery")
	u.SetExpectedErrors(status.Internal)
	return u
}

type APISubnetsPutInput struct {
	Name string `query:"identifier" required:"true" maxLength:"255"`

	SubnetCIDR   string `json:"subnetCidr" required:"true" maxLength:"40"`
	DNSResolver  string `json:"dnsResolver" required:"true" maxLength:"255"`
	DiscoveryTTL int    `json:"discoveryTTL" required:"true"`
}

func (r *Role) APISubnetsPut() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input APISubnetsPutInput, output *struct{}) error {
		s := r.NewSubnet(input.Name)
		s.CIDR = input.SubnetCIDR
		s.DiscoveryTTL = input.DiscoveryTTL
		s.DNSResolver = input.DNSResolver
		err := s.put()
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		return nil
	})
	u.SetName("discovery.put_subnets")
	u.SetTitle("Discovery Subnets")
	u.SetTags("roles/discovery")
	u.SetExpectedErrors(status.Internal, status.InvalidArgument)
	return u
}

type APISubnetsStartInput struct {
	Name string `query:"identifier" required:"true"`
}

func (r *Role) APISubnetsStart() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input APISubnetsStartInput, output *struct{}) error {
		rawSub, err := r.i.KV().Get(ctx, r.i.KV().Key(
			types.KeyRole,
			types.KeySubnets,
			input.Name,
		).String())
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		if len(rawSub.Kvs) < 1 {
			return status.Wrap(err, status.NotFound)
		}
		s, err := r.subnetFromKV(rawSub.Kvs[0])
		if err != nil {
			r.log.WithError(err).Warning("failed to parse subnet from KV")
			return status.Wrap(err, status.Internal)
		}
		go s.RunDiscovery()
		return nil
	})
	u.SetName("discovery.subnet_start")
	u.SetTitle("Discovery Subnets")
	u.SetTags("roles/discovery")
	u.SetExpectedErrors(status.Internal, status.NotFound)
	return u
}

type APISubnetsDeleteInput struct {
	Name string `query:"identifier"`
}

func (r *Role) APISubnetsDelete() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input APISubnetsDeleteInput, output *struct{}) error {
		_, err := r.i.KV().Delete(ctx, r.i.KV().Key(
			types.KeyRole,
			types.KeySubnets,
			input.Name,
		).String())
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		return nil
	})
	u.SetName("discovery.delete_subnets")
	u.SetTitle("Discovery Subnets")
	u.SetTags("roles/discovery")
	u.SetExpectedErrors(status.Internal, status.InvalidArgument)
	return u
}
