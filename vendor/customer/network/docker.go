package network

import (
	"net"
	"net/http"
	"time"

	etcd2client "github.com/coreos/etcd/client"
	"github.com/coreos/etcd/pkg/transport"

	utilnet "k8s.io/apimachinery/pkg/util/net"
	//"k8s.io/apiserver/pkg/storage"
	//"k8s.io/apiserver/pkg/storage/etcd"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"github.com/golang/glog"
)

// etcdHelper is the reference implementation of storage.Interface.
type Etcd2Helper struct {
	EtcdKeysAPI etcd2client.KeysAPI
}

func NewETCD2Client(c storagebackend.Config) (*Etcd2Helper, error) {
	tr, err := newTransportForETCD2(c.CertFile, c.KeyFile, c.CAFile)
	glog.V(1).Infof("20170927:in NewETCD2Client")
	if err != nil {
		return nil, err
	}
	client, err := newETCD2Client(tr, c.ServerList)
	if err != nil {
		return nil, err
	}
	var etcdhelp Etcd2Helper
	etcdhelp.EtcdKeysAPI = etcd2client.NewKeysAPI(client)
	//s := etcd.NewEtcdStorage(client, c.Codec, c.Prefix, c.Quorum, c.DeserializationCacheSize, c.Copier, etcd.IdentityTransformer)
	return &etcdhelp, nil
}

func newETCD2Client(tr *http.Transport, serverList []string) (etcd2client.Client, error) {
	cli, err := etcd2client.New(etcd2client.Config{
		Endpoints: serverList,
		Transport: tr,
	})
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func newTransportForETCD2(certFile, keyFile, caFile string) (*http.Transport, error) {
	info := transport.TLSInfo{
		CertFile: certFile,
		KeyFile:  keyFile,
		CAFile:   caFile,
	}
	cfg, err := info.ClientConfig()
	if err != nil {
		return nil, err
	}
	// Copied from etcd.DefaultTransport declaration.
	// TODO: Determine if transport needs optimization
	tr := utilnet.SetTransportDefaults(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConnsPerHost: 500,
		TLSClientConfig:     cfg,
	})
	return tr, nil
}
