# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"sha256:daa3cfedaabdc411afb9e58693526d5e8d283fbc767abc6c2bcef929ea782566","image:golang":"sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70"}

FROM reg.c5h.io/golang@sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/irc-distribution/main;

FROM reg.c5h.io/base@sha256:daa3cfedaabdc411afb9e58693526d5e8d283fbc767abc6c2bcef929ea782566

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
