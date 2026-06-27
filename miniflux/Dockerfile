# Generated from https://github.com/csmith/dockerfiles/blob/master/miniflux/Dockerfile.gotpl
# BOM: {"github:miniflux/v2":"2.3.2","image:base":"5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38","image:golang":"ff41b0741faec698a6e43ebcd978c29e86d0c186483503e59e655535bc14aa63"}

FROM reg.c5h.io/golang@sha256:ff41b0741faec698a6e43ebcd978c29e86d0c186483503e59e655535bc14aa63 AS build

RUN set -eux; \
    git clone --depth=1 -b 2.3.2 --single-branch https://github.com/miniflux/v2 /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out .; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38

COPY --from=build /work/out /miniflux
COPY --from=build /notices /notices
ENTRYPOINT ["/miniflux"]
