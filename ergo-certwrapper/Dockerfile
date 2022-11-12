# Generated from https://github.com/csmith/dockerfiles/blob/master/ergo-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.1.0","github:ergochat/ergo":"v2.10.0","image:base":"2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665","image:golang":"dc4eac488c6c21ead1b8d08cf44dfa9e8d33ccec71eb063fe8b3e2ede109d082"}

FROM reg.c5h.io/golang@sha256:dc4eac488c6c21ead1b8d08cf44dfa9e8d33ccec71eb063fe8b3e2ede109d082 AS build

ARG CERTTAG="v4.1.0"
ARG ERGOTAG="v2.10.0"

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

FROM reg.c5h.io/base@sha256:2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/ircd", "/certs"]

ENTRYPOINT ["/certwrapper", "/ergo", "run"]
CMD ["--conf", "/ircd/ircd.yaml"]
