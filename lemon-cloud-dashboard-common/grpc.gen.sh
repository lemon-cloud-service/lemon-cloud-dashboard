# generate adm-golang
rm grpc_adm/**.pb.go
protoc -I . \
  --micro_out=grpc_adm \
  --go_out=grpc_adm \
  grpc/adm/*.proto
cp $(find grpc_adm/ -type f -name "*.go") grpc_adm/
rm -rf grpc_adm/github.com