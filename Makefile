path:=/Users/yc/work
gen:
	protoc -I proto -I $(path)/src/github.com/gogo/protobuf/protobuf \
         -I $(path)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				 -I $(path)/src \
         -I $(path)/src/github.com/gogo/protobuf \
				  proto/*.proto --gogo_out=plugins=grpc:pb \
					--grpc-gateway_out=:pb \
					--swagger_out=logtostderr=true:swagger


clean:
	rm pb/*.go; \
	rm swagger/*.json;