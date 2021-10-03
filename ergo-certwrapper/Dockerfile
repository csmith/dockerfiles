# Generated from https://github.com/csmith/dockerfiles/blob/master/ergo-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v3.1.2","github:ergochat/ergo":"v2.7.0","image:base":"sha256:74295276a5989bf84e9bc17d2a44b826a4be6aa2fcfa8e53a6b33404c6f37a02","image:golang":"sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70"}

FROM reg.c5h.io/golang@sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70 AS build

ARG CERTTAG="v3.1.2"
ARG ERGOTAG="v2.7.0"

RUN set -eux; \
    # Build certwrapper
    git clone --depth=1 -b $CERTTAG --single-branch https://github.com/csmith/certwrapper /go/src/github.com/csmith/certwrapper; \
    cd /go/src/github.com/csmith/certwrapper && (\
        CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags httpreq -o /out/certwrapper .; \
        go run github.com/google/go-licenses@latest save ./... --save_path=/out/notices/certwrapper; \
    ); \
    # Build ergo
    git clone --depth=1 -b $ERGOTAG --single-branch https://github.com/ergochat/ergo /go/src/github.com/ergochat/ergo; \
    cd /go/src/github.com/ergochat/ergo && (\
      CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o /out/ergo .; \
      go run github.com/google/go-licenses@latest save ./... --save_path=/out/notices/ergo; \
    ); \
    # Create folders for volumes + copy over languages \
    install -o 65532 -g 65532 -d /out/ircd /out/certs /out/ircd-bin; \
    cp -r /go/src/github.com/ergochat/ergo/languages /out/ircd-bin/languages; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /out;

FROM reg.c5h.io/base@sha256:74295276a5989bf84e9bc17d2a44b826a4be6aa2fcfa8e53a6b33404c6f37a02

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/ircd", "/certs"]

ENTRYPOINT ["/certwrapper", "/ergo", "run"]
CMD ["--conf", "/ircd/ircd.yaml"]
