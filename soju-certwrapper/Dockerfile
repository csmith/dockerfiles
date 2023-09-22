# Generated from https://github.com/csmith/dockerfiles/blob/master/soju-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.2.0","github:emersion/soju":"v0.6.2","image:base":"8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f","image:golang":"3073df7a16320b2bb862d08160cf5d53a3146655efc6cca2b1d84665274074ce"}

FROM reg.c5h.io/golang@sha256:3073df7a16320b2bb862d08160cf5d53a3146655efc6cca2b1d84665274074ce AS build

ARG SOJUTAG="v0.6.2"
ARG CERTTAG="v4.2.0"

RUN set -eux; \
    # Build certwrapper
    git clone --depth=1 -b $CERTTAG --single-branch https://github.com/csmith/certwrapper /go/src/github.com/csmith/certwrapper; \
    cd /go/src/github.com/csmith/certwrapper && (\
        CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags lego_httpreq -o /out/certwrapper .; \
        go-licenses save ./... --save_path=/notices; \
    ); \
    # Build soju
    git clone --depth=1 -b $SOJUTAG --single-branch https://github.com/emersion/soju /go/src/github.com/emersion/soju; \
    cd /go/src/github.com/emersion/soju && (\
      GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o /out/soju ./cmd/soju; \
      GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o /out/sojuctl ./cmd/sojuctl; \
      # TODO: find a license tool that doesn't hardcode Google's opinion on "forbidden" licenses :/
      # go-licenses save ./... --save_path=/notices; \
    ); \
    # Create folders for volumes
    install -o 65532 -g 65532 -d /out/data /out/certs;

FROM reg.c5h.io/base@sha256:8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/data", "/certs"]

ENTRYPOINT ["/certwrapper"]
CMD ["/soju"]
