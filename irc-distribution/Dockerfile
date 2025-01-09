# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"a326d0e1775b6525bf9968a89afbd9b957c15d92f33afc33bad4c11e38adb46b","image:golang":"8dec622c0522882adf3ed0d9e6bd7f0fe66ce0153729aec362f9a5d091e973e2"}

FROM reg.c5h.io/golang@sha256:8dec622c0522882adf3ed0d9e6bd7f0fe66ce0153729aec362f9a5d091e973e2 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.0 --single-branch https://github.com/csmith/irc-distribution /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:a326d0e1775b6525bf9968a89afbd9b957c15d92f33afc33bad4c11e38adb46b

COPY --from=build /work/out /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
