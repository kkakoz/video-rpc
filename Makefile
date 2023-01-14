include .env

SEED := $(shell perl -e "print int(rand(1000000))")



.PHONY: docker-run
docker-run:
	docker build . --tag videoweb:${VERSION}-${SEED}
	docker rm -f videoweb
	docker run -d --restart=always -p 9010:9010 --name videoweb -v /mnt/e/code/video_web/configs/conf.yaml:/app/configs/conf.yaml videoweb:${VERSION}-${SEED}

.PHONY: docker-push-user
docker-push-user:
	docker build . --tag video-rpc-user:${VERSION}-${SEED} -f ./Dockerfile-user
	docker tag video-rpc-user:${VERSION}-${SEED} ${ADDR_USER}:${VERSION}-${SEED}
	echo ${PASSWORD}  |  docker login --username=${USERNAME} registry.cn-hangzhou.aliyuncs.com --password-stdin
	docker push ${ADDR_USER}:${VERSION}-${SEED}


.PHONY: protoc
protoc:
		protoc -I. -I ./pb/third_party -I ./pb \
			--go_out . --go_opt paths=source_relative \
			--go-grpc_out . --go-grpc_opt paths=source_relative \
			./pb/common/*.proto

.PHONY: protoc-user
protoc-user:
		protoc -I. -I ./pb/third_party -I ./pb  \
			--go_out . --go_opt paths=source_relative \
			--go-grpc_out . --go-grpc_opt paths=source_relative \
			./pb/user/*.proto
