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

## generate usr-dto-web
#rm js_usr_dto/**_pb.js
#rm js_usr_dto/**.ts
#protoc -I . grpc/usr_dto/*.proto \
#  --js_out=import_style=commonjs,binary:js_usr_dto \
#  --grpc-web_out=import_style=typescript,mode=grpcwebtext:js_usr_dto
#cp $(find js_usr_dto/ -type f -name "*_pb.js") js_usr_dto/
#cp $(find js_usr_dto/ -type f -name "*.ts") js_usr_dto/
#rm -rf js_usr_dto/grpc
#
## generate usr-service-web
#rm js_usr_service/**_pb.js
#rm js_usr_service/**.ts
#protoc -I . grpc/usr_service/*.proto \
#  --js_out=import_style=commonjs,binary:js_usr_service \
#  --grpc-web_out=import_style=typescript,mode=grpcwebtext:js_usr_service
#cp $(find js_usr_service/ -type f -name "*_pb.js") js_usr_service/
#cp $(find js_usr_service/ -type f -name "*.ts") js_usr_service/
#rm -rf js_usr_service/grpc

PROTOC_GEN_TS_PATH="../../lemon-cloud-dashboard-ui/node_modules/.bin/protoc-gen-ts"
OUTPUT_PATH_PREFIX_TS="protobuf_out/typescript/"

# generate usr-dto-web
rm -rf ${OUTPUT_PATH_PREFIX_TS}/*
protoc \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:${OUTPUT_PATH_PREFIX_TS}" \
    --ts_out="service=grpc-web:${OUTPUT_PATH_PREFIX_TS}" \
    grpc/**/*.proto
#cp $(find js_usr_dto/ -type f -name "*.js") ${OUTPUT_PATH_PREFIX_TS}
#cp $(find js_usr_dto/ -type f -name "*.ts") ${OUTPUT_PATH_PREFIX_TS}
#rm -rf js_usr_dto/grpc

## generate usr-service-web
#rm js_usr_service/**_pb.js
#rm js_usr_service/**.ts
#protoc \
#    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
#    --js_out="import_style=commonjs,binary:js_usr_service" \
#    --ts_out="service=grpc-web:js_usr_service" \
#    grpc/usr_service/*.proto
#cp $(find js_usr_service/ -type f -name "*.js") js_usr_service/
#cp $(find js_usr_service/ -type f -name "*.ts") js_usr_service/
#rm -rf js_usr_service/grpc
