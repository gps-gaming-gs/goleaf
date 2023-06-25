env:
	cd tools/proto-compile; \
	npm install

grpc:
	cd tools/proto-compile; \
	npm run build-proto