#!/usr/bin/env bash

# Author: Mohammad Badruzzaman (https://github.com/sakib37)

echo "#################################################"
echo "   This script install Golang in a Linux amd64"
echo "#################################################"

# Change the version number according to your requirement
GOLANG_VERSION=1.11.5

cd $HOME

wget https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz

sudo tar -C /usr/local -xvzf go${GOLANG_VERSION}.linux-amd64.tar.gz

sudo rm go${GOLANG_VERSION}.linux-amd64.tar.gz

# This scripts considers that $HOME/go will be the workspace for golang
mkdir -p $HOME/go/{bin,src,pkg}

echo "Add the following environment variables in ~/.bashrc or ~/.bash_profile or ~/.profile file"
echo "##########################################################################################"

echo  -e "\texport GOROOT=/usr/local/go"
echo  -e "\texport GOPATH=\$HOME/go		# Change this if your workspace is not \$HOME/go"
echo  -e "\texport GOBIN=\$GOPATH/bin"

echo "Also add the following paths to your \$PATH variable"
echo  -e "\t\t\$GOROOT/bin:\$GOPATH/bin"
