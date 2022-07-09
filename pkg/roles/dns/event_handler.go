package dns

import (
	"net/netip"
	"strings"

	"beryju.io/ddet/pkg/roles"
	"beryju.io/ddet/pkg/roles/dns/types"
	"beryju.io/ddet/pkg/roles/dns/utils"
	"github.com/miekg/dns"
)

func (r *DNSRole) eventHandlerDHCPLeaseGiven(ev *roles.Event) {
	if ev.Payload.Data["hostname"] == "" {
		return
	}
	// Forward record
	r.eventCreateForward(ev)
	// Reverse record
	r.eventCreateReverse(ev)
}

func (r *DNSRole) eventCreateForward(ev *roles.Event) {
	hostname := ev.Payload.Data["hostname"].(string)
	fqdn := ev.Payload.Data["fqdn"].(string)
	forwardZone := r.findZone(fqdn)
	if forwardZone == nil {
		r.log.WithField("event", ev).WithField("fqdn", fqdn).Debug("No zone found for hostname")
		return
	}

	rawAddr := ev.Payload.Data["address"].(string)
	ip, err := netip.ParseAddr(rawAddr)
	if err != nil {
		r.log.WithError(err).Warning("failed to parse address to add dns record")
		return
	}
	var rec *Record
	if ip.Is4() {
		rec = forwardZone.newRecord(hostname, types.DNSRecordTypeA)
	} else {
		rec = forwardZone.newRecord(hostname, types.DNSRecordTypeA)
	}
	rec.Data = ip.String()
	rec.TTL = forwardZone.DefaultTTL
	err = rec.put(0, ev.Payload.RelatedObjectOptions...)
	if err != nil {
		r.log.WithError(err).Warning("failed to save dns record")
	}
}

func (r *DNSRole) eventCreateReverse(ev *roles.Event) {
	fqdn := ev.Payload.Data["fqdn"].(string)
	rawAddr := ev.Payload.Data["address"].(string)
	ip, err := netip.ParseAddr(rawAddr)
	if err != nil {
		r.log.WithError(err).Warning("failed to parse address to add dns record")
		return
	}

	rev, err := dns.ReverseAddr(ip.String())
	if err != nil {
		r.log.WithError(err).Warning("failed to get reverse of ip")
		return
	}

	forwardZone := r.findZone(rev)
	if forwardZone == nil {
		r.log.WithField("event", ev).WithField("fqdn", fqdn).Debug("No zone found for hostname")
		return
	}

	relName := strings.TrimSuffix(rev, utils.EnsureLeadingPeriod(forwardZone.Name))
	rec := forwardZone.newRecord(relName, types.DNSRecordTypePTR)
	rec.Data = fqdn
	rec.TTL = forwardZone.DefaultTTL
	err = rec.put(0, ev.Payload.RelatedObjectOptions...)
	if err != nil {
		r.log.WithError(err).Warning("failed to save dns record")
	}
}
