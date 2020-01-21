# Kubestream
This project aims to make deploying [Streamlit](https://github.com/streamlit/streamlit) apps on Kubernetes simple. The main component is a CLI that takes care of automagically (you don't have to provide a Dockerfile) building docker images from your python scripts and pushing them to your cluster, while taking care of all the associated Kubernetes resources (deployments, services, ingresses, etc.) You can think of it as a mini PaaS for Streamlit apps.

Kubestream is heavily inspired by [Rancher Rio](https://github.com/rancher/rio)

## Installing
Binaries for most popular platforms are built on every [release](https://github.com/relingan/kubestream/releases). You can always build from sources if needed:

1. Clone this repo.
2. Install the build dependencies :
    - Golang 1.13+
    - [Goreleaser](https://github.com/goreleaser/goreleaser)
    - Make

3. Run `make build`. The binaries will be created in the `dist` folder.

## Usage


## Development
While you certainly can install the dev dependencies on your machine manually, we bundle `.devcontainer` definitions that can be used to automatically configure the project on Visual Studio Code. To use them just install the [Remote](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl) VScode extension and run the `Remote-Containers: Open Folder in Container` command.

If you don't use VScode, you can still take advantage of the bundled `Dockerfile` and mount a volume with the source code, or maybe attach your editor using other tools.

Copyright (c) 2020 [Relingan](https://relingan.com)
