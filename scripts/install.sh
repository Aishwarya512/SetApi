#!/bin/bash
sudo yum -y update
sudo yum install -y golang
cd ~
mkdir go-projects
echo "export GOROOT=/usr/lib/golang" >> ~/.bash_profile
echo "export GOPATH=$HOME/go-projects" >> ~/.bash_profile
echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> ~/.bash_profile