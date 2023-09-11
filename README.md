<!--
SPDX-FileCopyrightText: The RamenDR authors
SPDX-License-Identifier: Apache-2.0
-->

# oc clusterset plugin

The oc clusterset plugin manages logins for clusterset.

Typical use case is to log in to the hub and managed clusters when
working with ACM managed clusters.

## Requirements

- oc - OpenShift Command Line Interface (CLI)

## Install

Copy the plugin to directory in the PATH

```
cp oc-clusterset /usr/local/bin/
```

## Clusterset configuration

To use this tool, prepare a yaml file with the cluster details:

```
$ cat config.yaml
clusters:
- name: hub
  url: cluster1.example.com:8443
  username: kubeadmin
  password: PeSkM-R6YcH-LyPZa-oTOO1
- name: c1
  url: cluster2.example.com:8443
  username: kubeadmin
  password: ZjIZn-SFUyR-aE4gI-fJcfL
- name: c2
  url: cluster3.example.com:8443
  username: kubeadmin
  password: 7C700-oVS3Q-25rtx-YMew5
  current-context: hub
```

## Log in to cluterset

To login to all clusters in config.yaml:

```
oc clusterset login --config config.yaml
```

## Log out from cluterset

To log out from all clusters in config.yaml:

```
oc clusterset --config config.yaml
```

## License

oc-clusterset is released under the Apache 2.0 license. See [LICENSE](LICENSE)
