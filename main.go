package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	"github.com/urfave/cli"
	"log"
	"os"
)

type CmdLineOpts struct {
	etcdEndpoints             string
	etcdPrefix                string
	etcdKeyfile               string
	etcdCertfile              string
	etcdCAFile                string
	etcdUsername              string
	etcdPassword              string
	version                   bool
	kubeSubnetMgr             bool
	kubeApiUrl                string
	kubeAnnotationPrefix      string
	kubeConfigFile            string
	iface                     []string
	ifaceRegex                []string
	ipMasq                    bool
	ifaceCanReach             string
	subnetFile                string
	publicIP                  string
	publicIPv6                string
	subnetLeaseRenewMargin    int
	healthzIP                 string
	healthzPort               int
	iptablesResyncSeconds     int
	iptablesForwardRules      bool
	netConfPath               string
	setNodeNetworkUnavailable bool
	useMultiClusterCidr       bool
}

var (
	opts CmdLineOpts
)

func main() {
	app := &cli.App{
		Name: "rCni",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "etcd-endpoints",
				Value:       "http://127.0.0.1:4001,http://127.0.0.1:2379",
				Destination: &opts.etcdEndpoints,
			},
			&cli.StringFlag{
				Name:        "etcd-prefix",
				Value:       "/coreos.com/network",
				Destination: &opts.etcdPrefix,
			},
			&cli.StringFlag{
				Name:        "etcd-keyfile",
				Destination: &opts.etcdKeyfile,
			},
			&cli.StringFlag{
				Name:        "etcd-certfile",
				Destination: &opts.etcdCertfile,
			},
			&cli.StringFlag{
				Name:        "etcd-cafile",
				Destination: &opts.etcdCAFile,
			},
			&cli.StringFlag{
				Name:        "etcd-username",
				Destination: &opts.etcdUsername,
			},
			&cli.StringFlag{
				Name:        "etcd-password",
				Destination: &opts.etcdPassword,
			},
			&cli.StringFlag{
				Name:        "subnet-file",
				Value:       "/run/rcni/subnet.env",
				Destination: &opts.subnetFile,
			},
			&cli.IntFlag{
				Name:        "subnet-lease-renew-margin",
				Value:       60,
				Destination: &opts.subnetLeaseRenewMargin,
			},
			&cli.BoolFlag{
				Name:        "ip-masq",
				Destination: &opts.ipMasq,
			},
			&cli.BoolFlag{
				Name:        "version",
				Destination: &opts.version,
			},
		},
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, "")
}

func cmdAdd(args *skel.CmdArgs) error {
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	return nil
}
