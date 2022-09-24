package dns

import (
	"context"
	"strings"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (r *Role) apiHandlerZonesGet() usecase.Interactor {
	type zone struct {
		Name           string              `json:"name" required:"true"`
		Authoritative  bool                `json:"authoritative" required:"true"`
		HandlerConfigs []map[string]string `json:"handlerConfigs" required:"true"`
		DefaultTTL     uint32              `json:"defaultTTL" required:"true"`
	}
	type zonesOutput struct {
		Zones []zone `json:"zones" required:"true"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *zonesOutput) error {
		for name, _zone := range r.zones {
			output.Zones = append(output.Zones, zone{
				Name:          name,
				Authoritative: _zone.Authoritative,
			})
		}
		return nil
	})
	u.SetName("dns.get_zones")
	u.SetTitle("DNS Zones")
	u.SetTags("roles/dns")
	return u
}

func (r *Role) apiHandlerZonesPut() usecase.Interactor {
	type zoneInput struct {
		Name           string              `query:"zone" required:"true"`
		Authoritative  bool                `json:"authoritative" required:"true"`
		HandlerConfigs []map[string]string `json:"handlerConfigs" required:"true"`
		DefaultTTL     uint32              `json:"defaultTTL" required:"true"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input zoneInput, output *struct{}) error {
		z := r.newZone(input.Name)
		z.Name = input.Name
		if !strings.HasSuffix(z.Name, ".") {
			z.Name += "."
		}
		z.Authoritative = input.Authoritative
		z.HandlerConfigs = input.HandlerConfigs
		z.DefaultTTL = input.DefaultTTL
		err := z.put(ctx)
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		return nil
	})
	u.SetName("dns.put_zones")
	u.SetTitle("DNS Zones")
	u.SetTags("roles/dns")
	u.SetExpectedErrors(status.Internal, status.InvalidArgument)
	return u
}

func (r *Role) apiHandlerZonesDelete() usecase.Interactor {
	type zoneInput struct {
		Zone string `query:"zone"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input zoneInput, output *struct{}) error {
		z, ok := r.zones[input.Zone]
		if !ok {
			return status.InvalidArgument
		}
		_, err := r.i.KV().Delete(ctx, z.etcdKey, clientv3.WithPrefix())
		if err != nil {
			return status.Wrap(err, status.Internal)
		}
		return nil
	})
	u.SetName("dns.delete_zones")
	u.SetTitle("DNS Zones")
	u.SetTags("roles/dns")
	u.SetExpectedErrors(status.Internal, status.InvalidArgument)
	return u
}
