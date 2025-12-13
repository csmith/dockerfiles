# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.1.0","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef"}

FROM reg.c5h.io/golang@sha256:18a4339b2b4f079668919a005c982400f8cf8ce18f4752336395732498bddeef AS build

RUN set -eux; \
    git clone --depth=1 -b v1.1.0 --single-branch https://github.com/csmith/irc-distribution /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /work/out /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
