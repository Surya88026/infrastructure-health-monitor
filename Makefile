.PHONY: health list stats

health:
	docker run --rm --network monitoring -v "$$(pwd)/servers.json:/app/servers.json" healthchecker health

list:
	docker run --rm --network monitoring -v "$$(pwd)/servers.json:/app/servers.json" healthchecker list

stats:
	docker run --rm --network monitoring -v "$$(pwd)/servers.json:/app/servers.json" healthchecker stats

