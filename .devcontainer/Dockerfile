FROM golang:1.13.6

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils=1.8.* dialog=1.3* 2>&1 \
    #
    # Verify git, process tools, lsb-release (common
    # in install instructions for CLIs) installed
    git=1:2.20.1-2+deb10u1 iproute2=4.20.* \
    procps=2:3.* lsb-release=10.2019051400 \
    #
    # [Optional] Add sudo support
    sudo=1.8.27-1+deb10u1 \
    #
    # Install Go tools
    && go get -u -v \
        golang.org/x/tools/gopls \
        github.com/cweill/gotests/... \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/lint/golint \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/stamblerre/gocode \
        github.com/rogpeppe/godef \
        github.com/sqs/goreturns \
        github.com/markbates/pkger/cmd/pkger \
        github.com/go-delve/delve/cmd/dlv 2>&1 \
    #
    # Create a non-root user to use if preferred
    && groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* /go/src

# Install Goreleaser
RUN wget -q https://github.com/goreleaser/goreleaser/releases/download/v0.124.1/goreleaser_Linux_x86_64.tar.gz \
  && tar -xzf goreleaser_Linux_x86_64.tar.gz goreleaser \
  && chmod +x goreleaser && mv goreleaser /usr/local/bin

# Install golangci-lint
RUN wget -q https://github.com/golangci/golangci-lint/releases/download/v1.23.0/golangci-lint-1.23.0-linux-amd64.tar.gz \
  && tar -xzf golangci-lint-1.23.0-linux-amd64.tar.gz golangci-lint-1.23.0-linux-amd64 \
  && chmod +x golangci-lint-1.23.0-linux-amd64 \
  && mv golangci-lint-1.23.0-linux-amd64/golangci-lint /usr/local/bin/golangci-lint

# Switch back to dialog for any ad-hoc use of apt-get
ENV DEBIAN_FRONTEND=dialog

