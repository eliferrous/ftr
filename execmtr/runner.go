package execmtr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/netip"
	"os/exec"
	"strings"

	"github.com/eliferrous/ftr"
)

type mtrRunner struct{}

func New() ftr.Runner { return &mtrRunner{} }

func (m *mtrRunner) Run(ctx context.Context, target string, count int) (*ftr.Report, error) {
	cmd := exec.CommandContext(
		ctx, "mtr",
		"--json",
		"--show-ips",
		"--aslookup",
		"--report-cycles", fmt.Sprint(count),
		target,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run mtr: %w\n%s", err, out)
	}
	return m.parse(out, target)
}

type rawHop struct {
	Count int     `json:"count"`
	Host  string  `json:"host"`
	ASN   string  `json:"ASN"`
	Loss  float64 `json:"Loss%"`
	Sent  int     `json:"Snt"`
	Last  float64 `json:"Last"`
	Avg   float64 `json:"Avg"`
	Best  float64 `json:"Best"`
	Worst float64 `json:"Wrst"`
	Stdev float64 `json:"StDev"`
}

type rawReport struct {
	Report struct {
		Hubs []rawHop `json:"hubs"`
	} `json:"report"`
}

func (m *mtrRunner) parse(b []byte, target string) (*ftr.Report, error) {
	var rr rawReport
	if err := json.Unmarshal(b, &rr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	rep := &ftr.Report{Target: target}
	for _, h := range rr.Report.Hubs {
		rep.Hops = append(rep.Hops, toDomain(h))
	}
	return rep, nil
}

func toDomain(h rawHop) ftr.Hop {

	hop := ftr.Hop{
		Count: h.Count,
		Host:  h.Host,
		ASN:   h.ASN,
		Ploss: h.Loss,
		Sent:  h.Sent,
		Last:  h.Last,
		Avg:   h.Avg,
		Best:  h.Best,
		Worst: h.Worst,
		Stdev: h.Stdev,
	}

	if h.Host == "???" {
		hop.Unanswered = true
		return hop
	}

	if ipStr, ok := extractIP(h.Host); ok {
		if ip, err := netip.ParseAddr(ipStr); err == nil {
			hop.Addr = &ip
		}
	}
	return hop
}

func extractIP(host string) (string, bool) {
	if start := strings.LastIndexByte(host, '('); start != -1 {
		if end := strings.LastIndexByte(host, ')'); end > start {
			ip := strings.TrimSpace(host[start+1 : end])
			return ip, true
		}
	}

	if _, err := netip.ParseAddr(host); err == nil {
		return host, true
	}

	return "", false
}
