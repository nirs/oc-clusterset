<!--
SPDX-FileCopyrightText: The RamenDR authors
SPDX-License-Identifier: Apache-2.0
-->

# oc clusterset plugin.

## Requirements

- oc - OpenShift Command Line Interface (CLI)

## Install

Copy the plugin to directory in the PATH

```
cp oc-clusterset /usr/loca/bin/
```

## Clusterset configuration

To use this tool, prepare a yaml files with the cluster details:

    $ cat config.yaml
    clusters:
      - name: hub
        url: api.perf1.example.com
        username: kubeadmin
        password: password-for-perf1
      - name: c1
        url: api.perf2.example.com
        username: kubeadmin
        password: password-for-perf2
      - name: c2
        url: api.perf3.example.com
        username: kubeadmin
        password: password-for-perf3
    current-context: hub

## Log in to cluterset

To login to all clusters in config.yaml:

```
oc clusterset login --config config.yaml
```

## Log out from cluterset

To log out from all clusters config.yaml:

```
oc clusterset --config config.yaml
```
