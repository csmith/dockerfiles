# Generated from https://github.com/csmith/dockerfiles/blob/master/soju/Dockerfile.gotpl
# BOM: {"github:emersion/soju":"v0.10.1","image:base":"66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd","image:golang":"d364769391d3a61e0c701264d331028c3249cb84cc8316dfffe5d882339f1feb"}

FROM reg.c5h.io/golang@sha256:d364769391d3a61e0c701264d331028c3249cb84cc8316dfffe5d882339f1feb AS build

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

FROM reg.c5h.io/base@sha256:66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd

COPY --from=build /out/ /
VOLUME ["/data"]

ENTRYPOINT ["/soju"]
