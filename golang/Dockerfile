# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.21-r0","apk:binutils":"2.41-r0","apk:brotli-libs":"1.1.0-r1","apk:busybox":"1.36.1-r15","apk:busybox-binsh":"1.36.1-r15","apk:c-ares":"1.27.0-r0","apk:ca-certificates":"20240226-r0","apk:curl":"8.5.0-r0","apk:gcc":"13.2.1_git20231014-r0","apk:git":"2.43.0-r0","apk:gmp":"6.3.0-r0","apk:go":"1.21.9-r0","apk:isl25":"0.25-r2","apk:jansson":"2.14-r4","apk:libatomic":"13.2.1_git20231014-r0","apk:libcrypto3":"3.1.4-r6","apk:libcurl":"8.5.0-r0","apk:libexpat":"2.6.2-r0","apk:libgcc":"13.2.1_git20231014-r0","apk:libgomp":"13.2.1_git20231014-r0","apk:libidn2":"2.3.4-r4","apk:libncursesw":"6.4_p20231125-r0","apk:libssl3":"3.1.4-r6","apk:libstdc++":"13.2.1_git20231014-r0","apk:libunistring":"1.1-r2","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.1-r0","apk:musl":"1.2.4_git20230717-r4","apk:musl-dev":"1.2.4_git20230717-r4","apk:ncurses-terminfo-base":"6.4_p20231125-r0","apk:nghttp2-libs":"1.58.0-r0","apk:pcre2":"10.42-r2","apk:readline":"8.2.1-r2","apk:zlib":"1.3.1-r0","apk:zstd-libs":"1.5.5-r8","github:google/go-licenses":"v1.6.0","golang":"go1.22.2","image:alpine":"8a2be8ea03cd3dfada335f805eafa28768098b44255a76310f8da133be9901ef"}

FROM reg.c5h.io/alpine@sha256:8a2be8ea03cd3dfada335f805eafa28768098b44255a76310f8da133be9901ef

ARG GOLICENSETAG="v1.6.0"
ARG ARCHIVE_URL="https:/golang.org/dl/go1.22.2.src.tar.gz"
ARG ARCHIVE_SUM="374ea82b289ec738e968267cac59c7d5ff180f9492250254784b2044e90df5a9"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps \
        bash=5.2.21-r0 \
        binutils=2.41-r0 \
        brotli-libs=1.1.0-r1 \
        busybox=1.36.1-r15 \
        busybox-binsh=1.36.1-r15 \
        c-ares=1.27.0-r0 \
        ca-certificates=20240226-r0 \
        curl=8.5.0-r0 \
        gcc=13.2.1_git20231014-r0 \
        gmp=6.3.0-r0 \
        go=1.21.9-r0 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20231014-r0 \
        libcrypto3=3.1.4-r6 \
        libcurl=8.5.0-r0 \
        libgcc=13.2.1_git20231014-r0 \
        libgomp=13.2.1_git20231014-r0 \
        libidn2=2.3.4-r4 \
        libncursesw=6.4_p20231125-r0 \
        libssl3=3.1.4-r6 \
        libstdc++=13.2.1_git20231014-r0 \
        libunistring=1.1-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.4_git20230717-r4 \
        musl-dev=1.2.4_git20230717-r4 \
        ncurses-terminfo-base=6.4_p20231125-r0 \
        nghttp2-libs=1.58.0-r0 \
        readline=8.2.1-r2 \
        zlib=1.3.1-r0 \
        zstd-libs=1.5.5-r8 \
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
        binutils=2.41-r0 \
        brotli-libs=1.1.0-r1 \
        busybox=1.36.1-r15 \
        busybox-binsh=1.36.1-r15 \
        c-ares=1.27.0-r0 \
        ca-certificates=20240226-r0 \
        gcc=13.2.1_git20231014-r0 \
        git=2.43.0-r0 \
        gmp=6.3.0-r0 \
        isl25=0.25-r2 \
        jansson=2.14-r4 \
        libatomic=13.2.1_git20231014-r0 \
        libcrypto3=3.1.4-r6 \
        libcurl=8.5.0-r0 \
        libexpat=2.6.2-r0 \
        libgcc=13.2.1_git20231014-r0 \
        libgomp=13.2.1_git20231014-r0 \
        libidn2=2.3.4-r4 \
        libssl3=3.1.4-r6 \
        libstdc++=13.2.1_git20231014-r0 \
        libunistring=1.1-r2 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.1-r0 \
        musl=1.2.4_git20230717-r4 \
        musl-dev=1.2.4_git20230717-r4 \
        nghttp2-libs=1.58.0-r0 \
        pcre2=10.42-r2 \
        zlib=1.3.1-r0 \
        zstd-libs=1.5.5-r8 \
        ; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH; \
    go install github.com/google/go-licenses@$GOLICENSETAG
