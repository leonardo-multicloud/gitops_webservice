## Install golang

wget https://golang.org/dl/go1.20.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

## Docker build image
docker build -t leonardomulticloud/gitops_webservice:dev .
docker run --rm -p 8080:8080 leonardomulticloud/gitops_webservice:dev