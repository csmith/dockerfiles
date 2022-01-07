# Generated from https://github.com/csmith/dockerfiles/blob/master/soju-certwrapper/Dockerfile.gotpl
# BOM: {"github:csmith/certwrapper":"v4.0.0","github:emersion/soju":"v0.3.0","image:base":"e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903","image:golang":"52968a8b807cc37629d7bb688498c02998d5d2ef70caba40d5a048dbde4b7edf"}

FROM reg.c5h.io/golang@sha256:52968a8b807cc37629d7bb688498c02998d5d2ef70caba40d5a048dbde4b7edf AS build

ARG SOJUTAG="v0.3.0"
ARG CERTTAG="v4.0.0"

RUN set -eux; \
    # Build certwrapper
    git clone --depth=1 -b $CERTTAG --single-branch https://github.com/csmith/certwrapper /go/src/github.com/csmith/certwrapper; \
    cd /go/src/github.com/csmith/certwrapper && (\
        CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -tags lego_httpreq -o /out/certwrapper .; \
        go run github.com/google/go-licenses@latest save ./... --save_path=/out/notices/certwrapper; \
    ); \
    # Build soju
    git clone --depth=1 -b $SOJUTAG --single-branch https://github.com/emersion/soju /go/src/github.com/emersion/soju; \
    cd /go/src/github.com/emersion/soju && (\
      GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o /out/soju ./cmd/soju; \
      GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o /out/sojuctl ./cmd/sojuctl; \
      # TODO: find a license tool that doesn't hardcode Google's opinion on "forbidden" licenses :/
      # go run github.com/google/go-licenses@latest save ./... --save_path=/out/notices/soju; \
    ); \
    # Create folders for volumes
    install -o 65532 -g 65532 -d /out/data /out/certs; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /out;

FROM reg.c5h.io/base@sha256:e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903

ENV CERTWRAPPER_CERTIFICATE_PATH=/certs/certificate.pem \
    CERTWRAPPER_DNS_PROVIDER=httpreq \
    CERTWRAPPER_ISSUER_PATH=/certs/issuer.perm \
    CERTWRAPPER_PRIVATE_KEY_PATH=/certs/privatekey.pem \
    CERTWRAPPER_USER_PATH=/certs/user.json

COPY --from=build /out/ /
VOLUME ["/data", "/certs"]

ENTRYPOINT ["/certwrapper"]
CMD ["/soju"]
