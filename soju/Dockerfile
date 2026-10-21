# Generated from https://github.com/csmith/dockerfiles/blob/master/soju/Dockerfile.gotpl
# BOM: {"github:https://github.com/emersion/soju":"v0.10.1","image:base":"164f46a3cc86b9bdd26e51f8d4deb4688e97da19495b83cb1ff3bfc0ad3e31d1","image:golang":"0c7c8b0250d2f903a8c957d72e9a0c9172ac62f04a7c0594649db6b99549ece3"}

FROM reg.c5h.io/golang@sha256:0c7c8b0250d2f903a8c957d72e9a0c9172ac62f04a7c0594649db6b99549ece3 AS build

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

FROM reg.c5h.io/base@sha256:164f46a3cc86b9bdd26e51f8d4deb4688e97da19495b83cb1ff3bfc0ad3e31d1

COPY --from=build /out/ /
VOLUME ["/data"]

ENTRYPOINT ["/soju"]
