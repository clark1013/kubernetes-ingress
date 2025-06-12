package ingress

import (
	"github.com/haproxytech/kubernetes-ingress/v3/pkg/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/v3/pkg/haproxy/rules"
	"github.com/haproxytech/kubernetes-ingress/v3/pkg/store"
)

type ReqSetHost struct {
	rules *rules.List
	name  string
}

func NewReqSetHost(n string, r *rules.List) *ReqSetHost {
	return &ReqSetHost{name: n, rules: r}
}

func (a *ReqSetHost) GetName() string {
	return a.name
}

func (a *ReqSetHost) Process(k store.K8s, annotations ...map[string]string) (err error) {
	input := common.GetValue(a.GetName(), annotations...)
	if input == "" {
		return
	}
	a.rules.Add(&rules.SetHdr{
		HdrName:   "Host",
		HdrFormat: input,
	})
	return
}
