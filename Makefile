client-send:
	docker run \
	  	--network grpc_example_network \
		-v $(shell pwd)/helloapp/bin/hello_client:/bin/hello_client hello_client \
		./bin/hello_client --grpc.server.addr="helloapp_nginx_lb:50052" "HI"

client-build:
	docker build -t hello_client -f Dockerfile --no-cache .