# Generated from https://github.com/csmith/dockerfiles/blob/master/soju-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.2.0","github:emersion/soju":"v0.7.0","image:base":"7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f","image:golang":"5e230c03f6aaf29b89f73ae91102d1b2740706fc131fff0dedc3f72bf189adba"}

FROM reg.c5h.io/golang@sha256:5e230c03f6aaf29b89f73ae91102d1b2740706fc131fff0dedc3f72bf189adba AS build

ARG SOJUTAG="v0.7.0"
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

FROM reg.c5h.io/base@sha256:7442771fc9e4d70fbf8f9d592e28b9ec7b9b81af30ddb87cbaaba90f2cc0064f

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/data", "/certs"]

ENTRYPOINT ["/certwrapper"]
CMD ["/soju"]
