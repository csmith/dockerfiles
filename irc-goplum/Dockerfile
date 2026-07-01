# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38","image:golang":"08fdcff7769508d1d6de4829951a3600b4aaf14a53b7d68f65a14977edfcd38a"}

FROM reg.c5h.io/golang@sha256:08fdcff7769508d1d6de4829951a3600b4aaf14a53b7d68f65a14977edfcd38a AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
