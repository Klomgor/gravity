package dns_test

import (
	"testing"

	"beryju.io/gravity/pkg/instance"
	"beryju.io/gravity/pkg/roles/dns"
	"beryju.io/gravity/pkg/roles/dns/types"
	"beryju.io/gravity/pkg/tests"
	d "github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
)

func TestRoleDNS_SOA(t *testing.T) {
	defer tests.Setup(t)()
	rootInst := instance.New()
	ctx := tests.Context()
	inst := rootInst.ForRole("dns", ctx)
	tests.PanicIfError(inst.KV().Put(
		ctx,
		inst.KV().Key(
			types.KeyRole,
			types.KeyZones,
			".",
		).String(),
		tests.MustJSON(dns.Zone{
			Authoritative: true,
			HandlerConfigs: []map[string]interface{}{
				{
					"type": "forward_ip",
					"to":   "127.0.0.1:1053",
				},
			},
		}),
	))
	role := dns.New(inst)
	assert.NotNil(t, role)
	assert.Nil(t, role.Start(ctx, RoleConfig()))
	defer role.Stop()

	fw := NewNullDNSWriter()
	role.Handler(fw, &d.Msg{
		Question: []d.Question{
			{
				Name:   ".",
				Qtype:  d.TypeSOA,
				Qclass: d.ClassINET,
			},
		},
	})
	ans := fw.Msg().Answer[0]
	assert.Equal(t, ".", ans.(*d.SOA).Ns)
}
