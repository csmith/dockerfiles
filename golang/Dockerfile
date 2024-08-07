# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.26-r0","apk:binutils":"2.42-r0","apk:brotli-libs":"1.1.0-r2","apk:busybox":"1.36.1-r29","apk:busybox-binsh":"1.36.1-r29","apk:c-ares":"1.28.1-r0","apk:ca-certificates":"20240705-r0","apk:curl":"8.9.0-r0","apk:gcc":"13.2.1_git20240309-r0","apk:git":"2.45.2-r0","apk:gmp":"6.3.0-r1","apk:go":"1.22.5-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20240309-r0","apk:libcrypto3":"3.3.1-r3","apk:libcurl":"8.9.0-r0","apk:libexpat":"2.6.2-r0","apk:libgcc":"13.2.1_git20240309-r0","apk:libgomp":"13.2.1_git20240309-r0","apk:libidn2":"2.3.7-r0","apk:libncursesw":"6.4_p20240420-r0","apk:libpsl":"0.21.5-r1","apk:libssl3":"3.3.1-r3","apk:libstdc++":"13.2.1_git20240309-r0","apk:libunistring":"1.2-r0","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.5-r0","apk:musl-dev":"1.2.5-r0","apk:ncurses-terminfo-base":"6.4_p20240420-r0","apk:nghttp2-libs":"1.62.1-r0","apk:pcre2":"10.43-r0","apk:readline":"8.2.10-r0","apk:zlib":"1.3.1-r1","apk:zstd-libs":"1.5.6-r0","github:google/go-licenses":"v1.6.0","golang":"go1.22.6","image:alpine":"e9dd4558ca431c8f52d46083be243dda993140051e5ce71bc883f74bf330558f"}

FROM reg.c5h.io/alpine@sha256:e9dd4558ca431c8f52d46083be243dda993140051e5ce71bc883f74bf330558f

ARG GOLICENSETAG="v1.6.0"
ARG ARCHIVE_URL="https:/golang.org/dl/go1.22.6.src.tar.gz"
ARG ARCHIVE_SUM="9e48d99d519882579917d8189c17e98c373ce25abaebb98772e2927088992a51"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps \
        bash=5.2.26-r0 \
        binutils=2.42-r0 \
        brotli-libs=1.1.0-r2 \
        busybox=1.36.1-r29 \
        busybox-binsh=1.36.1-r29 \
        c-ares=1.28.1-r0 \
        ca-certificates=20240705-r0 \
        curl=8.9.0-r0 \
        gcc=13.2.1_git20240309-r0 \
        gmp=6.3.0-r1 \
        go=1.22.5-r0 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20240309-r0 \
        libcrypto3=3.3.1-r3 \
        libcurl=8.9.0-r0 \
        libgcc=13.2.1_git20240309-r0 \
        libgomp=13.2.1_git20240309-r0 \
        libidn2=2.3.7-r0 \
        libncursesw=6.4_p20240420-r0 \
        libpsl=0.21.5-r1 \
        libssl3=3.3.1-r3 \
        libstdc++=13.2.1_git20240309-r0 \
        libunistring=1.2-r0 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r0 \
        musl-dev=1.2.5-r0 \
        ncurses-terminfo-base=6.4_p20240420-r0 \
        nghttp2-libs=1.62.1-r0 \
        readline=8.2.10-r0 \
        zlib=1.3.1-r1 \
        zstd-libs=1.5.6-r0 \
        ; \
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
    # Dependencies commonly needed for building go apps
    apk add --no-cache \
        binutils=2.42-r0 \
        brotli-libs=1.1.0-r2 \
        busybox=1.36.1-r29 \
        busybox-binsh=1.36.1-r29 \
        c-ares=1.28.1-r0 \
        ca-certificates=20240705-r0 \
        gcc=13.2.1_git20240309-r0 \
        git=2.45.2-r0 \
        gmp=6.3.0-r1 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20240309-r0 \
        libcrypto3=3.3.1-r3 \
        libcurl=8.9.0-r0 \
        libexpat=2.6.2-r0 \
        libgcc=13.2.1_git20240309-r0 \
        libgomp=13.2.1_git20240309-r0 \
        libidn2=2.3.7-r0 \
        libpsl=0.21.5-r1 \
        libssl3=3.3.1-r3 \
        libstdc++=13.2.1_git20240309-r0 \
        libunistring=1.2-r0 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.5-r0 \
        musl-dev=1.2.5-r0 \
        nghttp2-libs=1.62.1-r0 \
        pcre2=10.43-r0 \
        zlib=1.3.1-r1 \
        zstd-libs=1.5.6-r0 \
        ; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH; \
    go install github.com/google/go-licenses@$GOLICENSETAG
