# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.2.15-r5","apk:binutils":"2.40-r7","apk:brotli-libs":"1.0.9-r14","apk:busybox":"1.36.1-r2","apk:busybox-binsh":"1.36.1-r2","apk:ca-certificates":"20230506-r0","apk:curl":"8.2.1-r0","apk:gcc":"12.2.1_git20220924-r10","apk:git":"2.40.1-r0","apk:gmp":"6.2.1-r3","apk:go":"1.20.7-r0","apk:isl25":"0.25-r2","apk:libatomic":"12.2.1_git20220924-r10","apk:libcrypto3":"3.1.1-r3","apk:libcurl":"8.2.1-r0","apk:libexpat":"2.5.0-r1","apk:libgcc":"12.2.1_git20220924-r10","apk:libgomp":"12.2.1_git20220924-r10","apk:libidn2":"2.3.4-r1","apk:libncursesw":"6.4_p20230506-r0","apk:libssl3":"3.1.1-r3","apk:libstdc++":"12.2.1_git20220924-r10","apk:libunistring":"1.1-r1","apk:mpc1":"1.3.1-r1","apk:mpfr4":"4.2.0_p12-r0","apk:musl":"1.2.4-r1","apk:musl-dev":"1.2.4-r1","apk:ncurses-terminfo-base":"6.4_p20230506-r0","apk:nghttp2-libs":"1.55.1-r0","apk:pcre2":"10.42-r1","apk:readline":"8.2.1-r1","apk:zlib":"1.2.13-r1","apk:zstd-libs":"1.5.5-r4","github:google/go-licenses":"v1.6.0","golang":"go1.20.7","image:alpine":"0ce982c1a177d43800f731a14ddb968b2cfd1e94017d35961064e62634a7a5b8"}

FROM reg.c5h.io/alpine@sha256:0ce982c1a177d43800f731a14ddb968b2cfd1e94017d35961064e62634a7a5b8

ARG GOLICENSETAG="v1.6.0"
ARG ARCHIVE_URL="https:/golang.org/dl/go1.20.7.src.tar.gz"
ARG ARCHIVE_SUM="2c5ee9c9ec1e733b0dbbc2bdfed3f62306e51d8172bf38f4f4e542b27520f597"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps \
        bash=5.2.15-r5 \
        binutils=2.40-r7 \
        brotli-libs=1.0.9-r14 \
        busybox=1.36.1-r2 \
        busybox-binsh=1.36.1-r2 \
        ca-certificates=20230506-r0 \
        curl=8.2.1-r0 \
        gcc=12.2.1_git20220924-r10 \
        gmp=6.2.1-r3 \
        go=1.20.7-r0 \
        isl25=0.25-r2 \
        libatomic=12.2.1_git20220924-r10 \
        libcrypto3=3.1.1-r3 \
        libcurl=8.2.1-r0 \
        libgcc=12.2.1_git20220924-r10 \
        libgomp=12.2.1_git20220924-r10 \
        libidn2=2.3.4-r1 \
        libncursesw=6.4_p20230506-r0 \
        libssl3=3.1.1-r3 \
        libstdc++=12.2.1_git20220924-r10 \
        libunistring=1.1-r1 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.0_p12-r0 \
        musl=1.2.4-r1 \
        musl-dev=1.2.4-r1 \
        ncurses-terminfo-base=6.4_p20230506-r0 \
        nghttp2-libs=1.55.1-r0 \
        readline=8.2.1-r1 \
        zlib=1.2.13-r1 \
        zstd-libs=1.5.5-r4 \
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
        binutils=2.40-r7 \
        brotli-libs=1.0.9-r14 \
        busybox=1.36.1-r2 \
        busybox-binsh=1.36.1-r2 \
        ca-certificates=20230506-r0 \
        gcc=12.2.1_git20220924-r10 \
        git=2.40.1-r0 \
        gmp=6.2.1-r3 \
        isl25=0.25-r2 \
        libatomic=12.2.1_git20220924-r10 \
        libcrypto3=3.1.1-r3 \
        libcurl=8.2.1-r0 \
        libexpat=2.5.0-r1 \
        libgcc=12.2.1_git20220924-r10 \
        libgomp=12.2.1_git20220924-r10 \
        libidn2=2.3.4-r1 \
        libssl3=3.1.1-r3 \
        libstdc++=12.2.1_git20220924-r10 \
        libunistring=1.1-r1 \
        mpc1=1.3.1-r1 \
        mpfr4=4.2.0_p12-r0 \
        musl=1.2.4-r1 \
        musl-dev=1.2.4-r1 \
        nghttp2-libs=1.55.1-r0 \
        pcre2=10.42-r1 \
        zlib=1.2.13-r1 \
        zstd-libs=1.5.5-r4 \
        ; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH; \
    go install github.com/google/go-licenses@$GOLICENSETAG
