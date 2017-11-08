package network

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	etcd2client "github.com/coreos/etcd/client"
)
type mux interface {
	Handle(pattern string, handler http.Handler)
}


func InstallHandler(mux mux, etcdhelp *Etcd2Helper) {

	glog.V(1).Infof("20170927: dockernetwork InstallDockerNetworkHandle")
	mux.Handle("/docker/networkls", handleDockernetwork(etcdhelp))
}

func handleDockernetwork(etcdhelp *Etcd2Helper) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		glog.V(1).Infof("20170927: dockernetwork enter:handleDockernetwork")

		glog.V(1).Infof("etcdhelp.EtcdKeysAPI:%v", etcdhelp.EtcdKeysAPI)
		resp := getdockernetwork(etcdhelp)
		fmt.Fprint(w, resp)
		//fmt.Fprint(w, "dockernetwork handleDockernetwork handleDockernetwork\n")
	})
}

func getdockernetwork(etcdhelp *Etcd2Helper) string {
	ctx := context.Background()

	getOpts := &etcd2client.GetOptions{
		Quorum:    true,
		Recursive: true,
		Sort:      true,
	}
	resp, err := etcdhelp.EtcdKeysAPI.Get(ctx, "docker/network/v1.0/network/", getOpts)
	if err != nil {
		return "get error"
	}

	glog.V(1).Infof("resp:%v,err:%v", resp, err)

	var nws []network
	for _, n := range resp.Node.Nodes {
		var nw network
		glog.V(1).Infof("Etcd:Key:%v,Value:%v,LastIndex:%v", n.Key, n.Value, n.ModifiedIndex)

		json.Unmarshal([]byte(n.Value), &nw)
		nws = append(nws, nw)
	}
	var dns []dockernetwork
	for i, v := range nws {
		var dn dockernetwork
		dn.Name = v.Name
		dn.Driver = v.NetworkType
		dn.Resourceid = v.Id
		dn.Scope = "global"
		dns = append(dns, dn)
		glog.V(1).Infof("nws %d,%v", i, v)
	}
	b, err := json.Marshal(dns)
	glog.V(1).Infof("str:%s,err:%v", b, err)
	return string(b)
}

type network struct {
	Name        string
	NetworkType string
	Id          string
}
type dockernetwork struct {
	Resourceid string `json:"resourceid"`
	Name       string `json:"name"`
	Driver     string `json:"driver"`
	Scope      string `json:"scope"`
}
