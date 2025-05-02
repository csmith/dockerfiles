# Generated from https://github.com/csmith/dockerfiles/blob/master/soju/Dockerfile.gotpl
# BOM: {"github:emersion/soju":"v0.9.0","image:base":"4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363","image:golang":"8db722d311cd3d21ae8cf75d8889dae3f4c17c8911456b2c8bd375bf716a6588"}

FROM reg.c5h.io/golang@sha256:8db722d311cd3d21ae8cf75d8889dae3f4c17c8911456b2c8bd375bf716a6588 AS build

RUN set -eux; \
    git clone --depth=1 -b v0.9.0 --single-branch https://github.com/emersion/soju /go/src/github.com/emersion/soju

WORKDIR /go/src/github.com/emersion/soju


RUN set -eux; \
    GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags=moderncsqlite -o /out/soju ./cmd/soju; \
    GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags=moderncsqlite -o /out/sojuctl ./cmd/sojuctl; \
    # TODO: find a license tool that doesn't hardcode Google's opinion on "forbidden" licenses :/
    # go-licenses save ./... --save_path=/notices; \
    # Create folders for volumes
    install -o 65532 -g 65532 -d /out/data;

FROM reg.c5h.io/base@sha256:4fea2559a3c315f01c6241c689da902344f4694dbdbdab4f5c0a1d24a8a78363

COPY --from=build /out/ /
VOLUME ["/data"]

ENTRYPOINT ["/soju"]
