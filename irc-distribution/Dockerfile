# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.1.0","image:base":"a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024","image:golang":"8c0d652f5dc30757c4c51c834dd3142d4c102d4999030bb5203f1fa1ddd760c4"}

FROM reg.c5h.io/golang@sha256:8c0d652f5dc30757c4c51c834dd3142d4c102d4999030bb5203f1fa1ddd760c4 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.1.0 --single-branch https://github.com/csmith/irc-distribution /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024

COPY --from=build /work/out /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
