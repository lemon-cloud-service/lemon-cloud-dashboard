# generate usr-dto-golang
rm usr_dto/**.pb.go
protoc -I . --go_out=plugins=grpc:usr_dto grpc/usr_dto/*.proto
cp $(find usr_dto/ -type f -name "*.pb.go") usr_dto/
rm -rf usr_dto/github.com

# generate usr-service-golang
rm usr_service/**.pb.go
protoc -I . --go_out=plugins=grpc:usr_service grpc/usr_service/*.proto
cp $(find usr_service/ -type f -name "*.pb.go") usr_service/
rm -rf usr_service/github.com

# generate adm-dto-golang
rm adm_dto/**.pb.go
protoc -I . --go_out=plugins=grpc:adm_dto grpc/adm_dto/*.proto
cp $(find adm_dto/ -type f -name "*.pb.go") adm_dto/
rm -rf adm_dto/github.com

# generate adm-service-golang
rm adm_service/**.pb.go
protoc -I . --go_out=plugins=grpc:adm_service grpc/adm_service/*.proto
cp $(find adm_service/ -type f -name "*.pb.go") adm_service/
rm -rf adm_service/github.com

PROTOC_GEN_TS_PATH="../../lemon-cloud-dashboard-ui/node_modules/.bin/protoc-gen-ts"
OUTPUT_PATH_PREFIX_TS="protobuf_out/typescript/"

# generate typescript
rm -rf ${OUTPUT_PATH_PREFIX_TS}/*
protoc \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:${OUTPUT_PATH_PREFIX_TS}" \
    --ts_out="service=grpc-web:${OUTPUT_PATH_PREFIX_TS}" \
    grpc/**/*.proto
