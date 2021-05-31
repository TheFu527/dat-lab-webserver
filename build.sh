go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
GOPATH=`go env GOPATH`
export PATH=${PATH}:${GOPATH}/bin
rm -rf dat-lab-webserver/apigen
mkdir -p dat-lab-webserver/apigen
oapi-codegen -package apigen api/main.yaml > dat-lab-webserver/apigen/main.gen.go