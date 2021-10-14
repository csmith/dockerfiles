FROM reg.c5h.io/golang@sha256:d10adab145c50608e7fb95c0923830744dd674c981e383ac9d2300813f70e670 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/dotege /go/src/github.com/csmith/dotege; \
    cd /go/src/github.com/csmith/dotege; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' -X main.GitSHA=$(git rev-parse --short HEAD)" -trimpath -o main .; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/dotege/{main,templates} /notices;

FROM reg.c5h.io/base@sha256:74295276a5989bf84e9bc17d2a44b826a4be6aa2fcfa8e53a6b33404c6f37a02

COPY --from=build /go/src/github.com/csmith/dotege/main /dotege
COPY --from=build /notices /notices
COPY --from=build /go/src/github.com/csmith/dotege/templates /templates

VOLUME /data/config
VOLUME /data/certs
VOLUME /data/output

ENTRYPOINT ["/dotege"]
