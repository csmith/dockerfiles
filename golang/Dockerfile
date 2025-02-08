# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.37-r0","apk:binutils":"2.43.1-r1","apk:brotli-libs":"1.1.0-r2","apk:busybox":"1.37.0-r12","apk:busybox-binsh":"1.37.0-r12","apk:c-ares":"1.34.3-r0","apk:ca-certificates-bundle":"20241121-r1","apk:curl":"8.12.0-r0","apk:gcc":"14.2.0-r4","apk:git":"2.47.2-r0","apk:gmp":"6.3.0-r2","apk:go":"1.23.5-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"14.2.0-r4","apk:libcrypto3":"3.3.2-r5","apk:libcurl":"8.12.0-r0","apk:libexpat":"2.6.4-r0","apk:libgcc":"14.2.0-r4","apk:libgomp":"14.2.0-r4","apk:libidn2":"2.3.7-r0","apk:libncursesw":"6.5_p20241006-r3","apk:libpsl":"0.21.5-r3","apk:libssl3":"3.3.2-r5","apk:libstdc++":"14.2.0-r4","apk:libunistring":"1.2-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r8","apk:musl-dev":"1.2.5-r8","apk:ncurses-terminfo-base":"6.5_p20241006-r3","apk:nghttp2-libs":"1.64.0-r0","apk:pcre2":"10.43-r0","apk:readline":"8.2.13-r0","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.6-r2","github:google/go-licenses":"v1.6.0","golang":"go1.23.6","image:alpine":"0be8812dc5bc37867ac650dde5331475f4fabd765e7f2f37f3166338e36cf6de"}

FROM reg.c5h.io/alpine@sha256:0be8812dc5bc37867ac650dde5331475f4fabd765e7f2f37f3166338e36cf6de


RUN set -eux; \
    apk add --no-cache --virtual .build-deps \
        brotli-libs=1.1.0-r2 \
        c-ares=1.34.3-r0 \
        ca-certificates-bundle=20241121-r1 \
        curl=8.12.0-r0 \
        libcrypto3=3.3.2-r5 \
        libcurl=8.12.0-r0 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.3.2-r5 \
        libunistring=1.2-r0 \
        musl=1.2.5-r8 \
        nghttp2-libs=1.64.0-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
        bash=5.2.37-r0 \
        busybox=1.37.0-r12 \
        busybox-binsh=1.37.0-r12 \
        libncursesw=6.5_p20241006-r3 \
        musl=1.2.5-r8 \
        ncurses-terminfo-base=6.5_p20241006-r3 \
        readline=8.2.13-r0 \
        binutils=2.43.1-r1 \
        gcc=14.2.0-r4 \
        gmp=6.3.0-r2 \
        go=1.23.5-r0 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=14.2.0-r4 \
        libgcc=14.2.0-r4 \
        libgomp=14.2.0-r4 \
        libstdc++=14.2.0-r4 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r8 \
        musl-dev=1.2.5-r8 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
    ;


RUN set -eux; \
    apk add --no-cache \
        binutils=2.43.1-r1 \
        gcc=14.2.0-r4 \
        gmp=6.3.0-r2 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=14.2.0-r4 \
        libgcc=14.2.0-r4 \
        libgomp=14.2.0-r4 \
        libstdc++=14.2.0-r4 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r8 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
        musl=1.2.5-r8 \
        musl-dev=1.2.5-r8 \
        brotli-libs=1.1.0-r2 \
        c-ares=1.34.3-r0 \
        ca-certificates-bundle=20241121-r1 \
        git=2.47.2-r0 \
        libcrypto3=3.3.2-r5 \
        libcurl=8.12.0-r0 \
        libexpat=2.6.4-r0 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.3.2-r5 \
        libunistring=1.2-r0 \
        musl=1.2.5-r8 \
        nghttp2-libs=1.64.0-r0 \
        pcre2=10.43-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.6-r2 \
    ;

ARG GOLICENSETAG="v1.6.0"
ARG ARCHIVE_URL="https://golang.org/dl/go1.23.6.src.tar.gz"
ARG ARCHIVE_SUM="039c5b04e65279daceee8a6f71e70bd05cf5b801782b6f77c6e19e2ed0511222"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Source for bootstrapping go
    curl -sfLo go.tar.gz $ARCHIVE_URL; \
    echo "$ARCHIVE_SUM *go.tar.gz" | sha256sum -wc -; \
    tar -C /usr/local -xzf go.tar.gz; \
    # Bootstrap and install the standard library
    cd /usr/local/go/src && GOROOT_BOOTSTRAP="$(go env GOROOT)" GOHOSTOS="linux" GOHOSTARCH="amd64" ./make.bash; \
    go install std; \
    # Clean up a little
	apk del --no-network .build-deps; \
    rm -rf \
            /go.tar.gz \
            /root/.cache \
    		/usr/local/go/pkg/*/cmd \
    		/usr/local/go/pkg/bootstrap \
    		/usr/local/go/pkg/obj \
    		/usr/local/go/pkg/tool/*/api \
    		/usr/local/go/pkg/tool/*/go_bootstrap \
    		/usr/local/go/src/cmd/dist/dist; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH; \
    go install github.com/google/go-licenses@$GOLICENSETAG
