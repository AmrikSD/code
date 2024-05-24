# Manual setup

This cluster was manually boostrapped by some manual ssh-ing and typing some commands in

## k3s

I followed the quick-start guide [here](https://docs.k3s.io/quick-start)

On the first node
```sh
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC='--flannel-backend=none --disable-network-policy' sh -
```
(The command line args are there so that we can install cilium later.)


On the other nodes

```sh
curl -sfL https://get.k3s.io | K3S_URL=https://myserver:6443 K3S_TOKEN=mynodetoken sh -
```

Replace $myserver with the IP address of node #1
Replace $mynodetoken with the value at `/var/lib/rancher/k3s/server/node-token` in node 1

Remember to get your kubectl config from `/etc/rancher/k3s/k3s.yaml`!

## Cilium


I followed the quick-start guide [here](https://docs.cilium.io/en/stable/gettingstarted/k8s-install-default/)

After I had kubectl setup and working it's a very easy process

`brew install cilium-cli && cilium install --version 1.15.5`

