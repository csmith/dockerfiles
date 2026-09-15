# Generated from https://github.com/csmith/dockerfiles/blob/master/soju/Dockerfile.gotpl
# BOM: {"github:https://github.com/emersion/soju":"v0.10.1","image:base":"9ab54133b898317b7cec48c724593554022dd5e993bc139ecbfb860ccd622383","image:golang":"b3e8ee60f06a5fd458baffa36cd6da269eede5ec30670575a2fa17e1981d0e66"}

FROM reg.c5h.io/golang@sha256:b3e8ee60f06a5fd458baffa36cd6da269eede5ec30670575a2fa17e1981d0e66 AS build

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

FROM reg.c5h.io/base@sha256:9ab54133b898317b7cec48c724593554022dd5e993bc139ecbfb860ccd622383

COPY --from=build /out/ /
VOLUME ["/data"]

ENTRYPOINT ["/soju"]
