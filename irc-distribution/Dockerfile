# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874","image:golang":"4c46b35f960d889b4b388a6501c507357c5f16948eb3dc73411ac85dbd72520e"}

FROM reg.c5h.io/golang@sha256:4c46b35f960d889b4b388a6501c507357c5f16948eb3dc73411ac85dbd72520e AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:33b47501f2e31be55c1af6b21082de3c3efaae9ec0062beb1d3255a9fd739874

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
