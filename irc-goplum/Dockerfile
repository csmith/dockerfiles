# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66","image:golang":"5e07f9b2210c6981dca7906be8aa31aed0b75ae0bfc0d776b74e52b5b86b7736"}

FROM reg.c5h.io/golang@sha256:5e07f9b2210c6981dca7906be8aa31aed0b75ae0bfc0d776b74e52b5b86b7736 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:22cd0188413016b6119bf80c3e3be20bf03b54e9ce6b9ff1576ab07f64f6da66

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
