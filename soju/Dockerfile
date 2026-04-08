# Generated from https://github.com/csmith/dockerfiles/blob/master/soju/Dockerfile.gotpl
# BOM: {"github:emersion/soju":"v0.10.1","image:base":"84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f","image:golang":"c18c857611620e0eccfbe4d1bb392260edbc75c24203d7f3c6a292091c142d72"}

FROM reg.c5h.io/golang@sha256:c18c857611620e0eccfbe4d1bb392260edbc75c24203d7f3c6a292091c142d72 AS build

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

FROM reg.c5h.io/base@sha256:84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f

COPY --from=build /out/ /
VOLUME ["/data"]

ENTRYPOINT ["/soju"]
