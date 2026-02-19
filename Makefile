VERSION=v1.146.0

all: build-main push-main build-e2e push-e2e

build-main:
	DOCKER_BUILDKIT=1 docker build -t ghcr.io/code-tool/matrix-stack/synapse:${VERSION} --build-arg SYNAPSE_PKG_VER=${VERSION} --target main -f build/Dockerfile build

push-main: build-main
	docker push ghcr.io/code-tool/matrix-stack/synapse:${VERSION}

build-e2e:
	DOCKER_BUILDKIT=1 docker build -t ghcr.io/code-tool/matrix-stack/synapse:${VERSION}-e2e-optimized --build-arg SYNAPSE_PKG_VER=${VERSION} --target e2e -f build/Dockerfile build

push-e2e: build-e2e
	docker push ghcr.io/code-tool/matrix-stack/synapse:${VERSION}-e2e-optimized
