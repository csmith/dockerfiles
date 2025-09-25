# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.37-r0","apk:binutils":"2.44-r3","apk:brotli-libs":"1.1.0-r2","apk:busybox":"1.37.0-r19","apk:busybox-binsh":"1.37.0-r19","apk:c-ares":"1.34.5-r0","apk:ca-certificates-bundle":"20250619-r0","apk:curl":"8.14.1-r1","apk:gcc":"14.2.0-r6","apk:git":"2.49.1-r0","apk:gmp":"6.3.0-r3","apk:go":"1.24.7-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14.1-r0","apk:libatomic":"14.2.0-r6","apk:libcrypto3":"3.5.3-r0","apk:libcurl":"8.14.1-r1","apk:libexpat":"2.7.2-r0","apk:libgcc":"14.2.0-r6","apk:libgomp":"14.2.0-r6","apk:libidn2":"2.3.7-r0","apk:libncursesw":"6.5_p20250503-r0","apk:libpsl":"0.21.5-r3","apk:libssl3":"3.5.3-r0","apk:libstdc++":"14.2.0-r6","apk:libunistring":"1.3-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1_p1-r0","apk:musl":"1.2.5-r10","apk:musl-dev":"1.2.5-r10","apk:ncurses-terminfo-base":"6.5_p20250503-r0","apk:nghttp2-libs":"1.65.0-r0","apk:pcre2":"10.43-r1","apk:readline":"8.2.13-r1","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.7-r0","golang":"go1.25.1","image:alpine":"660c3cb1032bcacc15b95b8be733562d35b7c6b66be87fc02e55cfe5bcce3bd8"}

FROM reg.c5h.io/alpine@sha256:660c3cb1032bcacc15b95b8be733562d35b7c6b66be87fc02e55cfe5bcce3bd8


RUN set -eux; \
    apk add --no-cache --virtual .build-deps \
        brotli-libs=1.1.0-r2 \
        c-ares=1.34.5-r0 \
        ca-certificates-bundle=20250619-r0 \
        curl=8.14.1-r1 \
        libcrypto3=3.5.3-r0 \
        libcurl=8.14.1-r1 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.5.3-r0 \
        libunistring=1.3-r0 \
        musl=1.2.5-r10 \
        nghttp2-libs=1.65.0-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
        bash=5.2.37-r0 \
        busybox=1.37.0-r19 \
        busybox-binsh=1.37.0-r19 \
        libncursesw=6.5_p20250503-r0 \
        musl=1.2.5-r10 \
        ncurses-terminfo-base=6.5_p20250503-r0 \
        readline=8.2.13-r1 \
        binutils=2.44-r3 \
        gcc=14.2.0-r6 \
        gmp=6.3.0-r3 \
        go=1.24.7-r0 \
        isl25=0.25-r2 \
        jansson=2.14.1-r0 \
        libatomic=14.2.0-r6 \
        libgcc=14.2.0-r6 \
        libgomp=14.2.0-r6 \
        libstdc++=14.2.0-r6 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1_p1-r0 \
        musl=1.2.5-r10 \
        musl-dev=1.2.5-r10 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
    ;


RUN set -eux; \
    apk add --no-cache \
        binutils=2.44-r3 \
        gcc=14.2.0-r6 \
        gmp=6.3.0-r3 \
        isl25=0.25-r2 \
        jansson=2.14.1-r0 \
        libatomic=14.2.0-r6 \
        libgcc=14.2.0-r6 \
        libgomp=14.2.0-r6 \
        libstdc++=14.2.0-r6 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1_p1-r0 \
        musl=1.2.5-r10 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
        musl=1.2.5-r10 \
        musl-dev=1.2.5-r10 \
        brotli-libs=1.1.0-r2 \
        c-ares=1.34.5-r0 \
        ca-certificates-bundle=20250619-r0 \
        git=2.49.1-r0 \
        libcrypto3=3.5.3-r0 \
        libcurl=8.14.1-r1 \
        libexpat=2.7.2-r0 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.5.3-r0 \
        libunistring=1.3-r0 \
        musl=1.2.5-r10 \
        nghttp2-libs=1.65.0-r0 \
        pcre2=10.43-r1 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r0 \
    ;

ARG ARCHIVE_URL="https://golang.org/dl/go1.25.1.src.tar.gz"
ARG ARCHIVE_SUM="d010c109cee94d80efe681eab46bdea491ac906bf46583c32e9f0dbb0bd1a594"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Source for bootstrapping go \
    curl -sfLo go.tar.gz $ARCHIVE_URL; \
    echo "$ARCHIVE_SUM *go.tar.gz" | sha256sum -wc -; \
    tar -C /usr/local -xzf go.tar.gz; \
    # Bootstrap and install the standard library \
    cd /usr/local/go/src && GOROOT_BOOTSTRAP="$(go env GOROOT)" GOHOSTOS="linux" GOHOSTARCH="amd64" ./make.bash; \
    go install std; \
    # Clean up a little \
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
    # Set up the go path \
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH; \
    go install github.com/google/go-licenses@v1.6.0
