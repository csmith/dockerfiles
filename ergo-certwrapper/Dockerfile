# Generated from https://github.com/csmith/dockerfiles/blob/master/ergo-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.2.0","github:ergochat/ergo":"v2.13.1","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"99f92c3368d01fa05c26a46857d73679f8b1cce3ba3606a7ac072f6baad69eae"}

FROM reg.c5h.io/golang@sha256:99f92c3368d01fa05c26a46857d73679f8b1cce3ba3606a7ac072f6baad69eae AS build

ARG CERTTAG="v4.2.0"
ARG ERGOTAG="v2.13.1"

RUN set -eux; \
    # Build certwrapper
    git clone --depth=1 -b $CERTTAG --single-branch https://github.com/csmith/certwrapper /go/src/github.com/csmith/certwrapper; \
    cd /go/src/github.com/csmith/certwrapper && (\
        CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags lego_httpreq -o /out/certwrapper .; \
        go-licenses save ./... --save_path=/out/notices/certwrapper; \
    ); \
    # Build ergo
    git clone --depth=1 -b $ERGOTAG --single-branch https://github.com/ergochat/ergo /go/src/github.com/ergochat/ergo; \
    cd /go/src/github.com/ergochat/ergo && (\
      CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o /out/ergo .; \
      go-licenses save ./... --save_path=/out/notices/ergo; \
    ); \
    # Create folders for volumes + copy over languages \
    install -o 65532 -g 65532 -d /out/ircd /out/certs /out/ircd-bin; \
    cp -r /go/src/github.com/ergochat/ergo/languages /out/ircd-bin/languages;

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/ircd", "/certs"]

ENTRYPOINT ["/certwrapper", "/ergo", "run"]
CMD ["--conf", "/ircd/ircd.yaml"]
