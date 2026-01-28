# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.3.3-r1","apk:binutils":"2.45.1-r0","apk:brotli-libs":"1.2.0-r0","apk:busybox":"1.37.0-r30","apk:busybox-binsh":"1.37.0-r30","apk:c-ares":"1.34.6-r0","apk:ca-certificates-bundle":"20251003-r0","apk:curl":"8.17.0-r1","apk:gcc":"15.2.0-r2","apk:git":"2.52.0-r0","apk:gmp":"6.3.0-r4","apk:go":"1.25.6-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14.1-r0","apk:libatomic":"15.2.0-r2","apk:libcrypto3":"3.5.5-r0","apk:libcurl":"8.17.0-r1","apk:libexpat":"2.7.3-r0","apk:libgcc":"15.2.0-r2","apk:libgomp":"15.2.0-r2","apk:libidn2":"2.3.8-r0","apk:libncursesw":"6.5_p20251123-r0","apk:libpsl":"0.21.5-r3","apk:libssl3":"3.5.5-r0","apk:libstdc++":"15.2.0-r2","apk:libunistring":"1.4.1-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.2-r0","apk:musl":"1.2.5-r21","apk:musl-dev":"1.2.5-r21","apk:ncurses-terminfo-base":"6.5_p20251123-r0","apk:nghttp2-libs":"1.68.0-r0","apk:nghttp3":"1.13.1-r0","apk:pcre2":"10.47-r0","apk:readline":"8.3.1-r0","apk:zlib":"1.3.1-r2","apk:zstd-libs":"1.5.7-r2","golang":"go1.25.6","image:alpine":"747f826a5f411a0e5c1e1d0af01828d8376fb90b3312a360a3a5d7a50f993982"}

FROM reg.c5h.io/alpine@sha256:747f826a5f411a0e5c1e1d0af01828d8376fb90b3312a360a3a5d7a50f993982


RUN set -eux; \
    apk add --no-cache --virtual .build-deps \
        brotli-libs=1.2.0-r0 \
        c-ares=1.34.6-r0 \
        ca-certificates-bundle=20251003-r0 \
        curl=8.17.0-r1 \
        libcrypto3=3.5.5-r0 \
        libcurl=8.17.0-r1 \
        libidn2=2.3.8-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.5.5-r0 \
        libunistring=1.4.1-r0 \
        musl=1.2.5-r21 \
        nghttp2-libs=1.68.0-r0 \
        nghttp3=1.13.1-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r2 \
        bash=5.3.3-r1 \
        busybox=1.37.0-r30 \
        busybox-binsh=1.37.0-r30 \
        libncursesw=6.5_p20251123-r0 \
        musl=1.2.5-r21 \
        ncurses-terminfo-base=6.5_p20251123-r0 \
        readline=8.3.1-r0 \
        binutils=2.45.1-r0 \
        gcc=15.2.0-r2 \
        gmp=6.3.0-r4 \
        go=1.25.6-r0 \
        isl25=0.25-r2 \
        jansson=2.14.1-r0 \
        libatomic=15.2.0-r2 \
        libgcc=15.2.0-r2 \
        libgomp=15.2.0-r2 \
        libstdc++=15.2.0-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.2-r0 \
        musl=1.2.5-r21 \
        musl-dev=1.2.5-r21 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r2 \
    ;


RUN set -eux; \
    apk add --no-cache \
        binutils=2.45.1-r0 \
        gcc=15.2.0-r2 \
        gmp=6.3.0-r4 \
        isl25=0.25-r2 \
        jansson=2.14.1-r0 \
        libatomic=15.2.0-r2 \
        libgcc=15.2.0-r2 \
        libgomp=15.2.0-r2 \
        libstdc++=15.2.0-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.2-r0 \
        musl=1.2.5-r21 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r2 \
        musl=1.2.5-r21 \
        musl-dev=1.2.5-r21 \
        brotli-libs=1.2.0-r0 \
        c-ares=1.34.6-r0 \
        ca-certificates-bundle=20251003-r0 \
        git=2.52.0-r0 \
        libcrypto3=3.5.5-r0 \
        libcurl=8.17.0-r1 \
        libexpat=2.7.3-r0 \
        libidn2=2.3.8-r0 \
        libpsl=0.21.5-r3 \
        libssl3=3.5.5-r0 \
        libunistring=1.4.1-r0 \
        musl=1.2.5-r21 \
        nghttp2-libs=1.68.0-r0 \
        nghttp3=1.13.1-r0 \
        pcre2=10.47-r0 \
        zlib=1.3.1-r2 \
        zstd-libs=1.5.7-r2 \
    ;

ARG ARCHIVE_URL="https://golang.org/dl/go1.25.6.src.tar.gz"
ARG ARCHIVE_SUM="58cbf771e44d76de6f56d19e33b77d745a1e489340922875e46585b975c2b059"

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
