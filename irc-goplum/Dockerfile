# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd","image:golang":"79c4a958212bde36a835f9b7c9dd2f721090e1a214699a11cfddb72858bfac6d"}

FROM reg.c5h.io/golang@sha256:79c4a958212bde36a835f9b7c9dd2f721090e1a214699a11cfddb72858bfac6d AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
