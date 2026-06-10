# Generated from https://github.com/csmith/dockerfiles/blob/master/soju/Dockerfile.gotpl
# BOM: {"github:emersion/soju":"v0.10.1","image:base":"932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929","image:golang":"f6a5f3b1a6a705f2608ba10bb77e9d9e855775f50c8ccd6d159ecea6fdb8d1f6"}

FROM reg.c5h.io/golang@sha256:f6a5f3b1a6a705f2608ba10bb77e9d9e855775f50c8ccd6d159ecea6fdb8d1f6 AS build

RUN set -eux; \
    git clone --depth=1 -b v0.10.1 --single-branch https://github.com/emersion/soju /go/src/github.com/emersion/soju

WORKDIR /go/src/github.com/emersion/soju


RUN set -eux; \
    GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags=moderncsqlite -o /out/soju ./cmd/soju; \
    GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags=moderncsqlite -o /out/sojuctl ./cmd/sojuctl; \
    # TODO: find a license tool that doesn't hardcode Google's opinion on "forbidden" licenses :/
    # go-licenses save ./... --save_path=/notices; \
    # Create folders for volumes
    install -o 65532 -g 65532 -d /out/data;

FROM reg.c5h.io/base@sha256:932d79427476d41a4cc015dae576efd446b46fe22463ecf15a62f8a5b54b9929

COPY --from=build /out/ /
VOLUME ["/data"]

ENTRYPOINT ["/soju"]
