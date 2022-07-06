# Generated from https://github.com/csmith/dockerfiles/blob/master/golang/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.16-r2","apk:binutils":"2.38-r2","apk:brotli-libs":"1.0.9-r6","apk:busybox":"1.35.0-r14","apk:ca-certificates":"20211220-r0","apk:curl":"7.83.1-r2","apk:expat":"2.4.8-r0","apk:gcc":"11.2.1_git20220219-r2","apk:git":"2.36.1-r0","apk:gmp":"6.2.1-r2","apk:go":"1.18.2-r0","apk:isl22":"0.22-r0","apk:libatomic":"11.2.1_git20220219-r2","apk:libcrypto1.1":"1.1.1q-r0","apk:libcurl":"7.83.1-r2","apk:libgcc":"11.2.1_git20220219-r2","apk:libgomp":"11.2.1_git20220219-r2","apk:libssl1.1":"1.1.1q-r0","apk:libstdc++":"11.2.1_git20220219-r2","apk:mpc1":"1.2.1-r0","apk:mpfr4":"4.1.0-r0","apk:musl":"1.2.3-r0","apk:musl-dev":"1.2.3-r0","apk:ncurses-libs":"6.3_p20220521-r0","apk:ncurses-terminfo-base":"6.3_p20220521-r0","apk:nghttp2-libs":"1.47.0-r0","apk:pcre2":"10.40-r0","apk:readline":"8.1.2-r0","apk:zlib":"1.2.12-r1","github:google/go-licenses":"v1.2.1","golang":"go1.18.3","image:alpine":"e1af61331d46ce7a5770c5083e1c81e297fbf3fcc97b63542b0ac62a8f3900a3"}

FROM reg.c5h.io/alpine@sha256:e1af61331d46ce7a5770c5083e1c81e297fbf3fcc97b63542b0ac62a8f3900a3

ARG GOLICENSETAG="v1.2.1"
ARG ARCHIVE_URL="https:/golang.org/dl/go1.18.3.src.tar.gz"
ARG ARCHIVE_SUM="0012386ddcbb5f3350e407c679923811dbd283fcdc421724931614a842ecbc2d"

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

RUN set -eux; \
    # Dependencies and source for bootstrapping go
    apk add --no-cache --virtual .build-deps \
        bash=5.1.16-r2 \
        binutils=2.38-r2 \
        brotli-libs=1.0.9-r6 \
        busybox=1.35.0-r14 \
        ca-certificates=20211220-r0 \
        curl=7.83.1-r2 \
        gcc=11.2.1_git20220219-r2 \
        gmp=6.2.1-r2 \
        go=1.18.2-r0 \
        isl22=0.22-r0 \
        libatomic=11.2.1_git20220219-r2 \
        libcrypto1.1=1.1.1q-r0 \
        libcurl=7.83.1-r2 \
        libgcc=11.2.1_git20220219-r2 \
        libgomp=11.2.1_git20220219-r2 \
        libssl1.1=1.1.1q-r0 \
        libstdc++=11.2.1_git20220219-r2 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r0 \
        musl-dev=1.2.3-r0 \
        ncurses-libs=6.3_p20220521-r0 \
        ncurses-terminfo-base=6.3_p20220521-r0 \
        nghttp2-libs=1.47.0-r0 \
        readline=8.1.2-r0 \
        zlib=1.2.12-r1 \
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
        binutils=2.38-r2 \
        brotli-libs=1.0.9-r6 \
        busybox=1.35.0-r14 \
        ca-certificates=20211220-r0 \
        expat=2.4.8-r0 \
        gcc=11.2.1_git20220219-r2 \
        git=2.36.1-r0 \
        gmp=6.2.1-r2 \
        isl22=0.22-r0 \
        libatomic=11.2.1_git20220219-r2 \
        libcrypto1.1=1.1.1q-r0 \
        libcurl=7.83.1-r2 \
        libgcc=11.2.1_git20220219-r2 \
        libgomp=11.2.1_git20220219-r2 \
        libssl1.1=1.1.1q-r0 \
        libstdc++=11.2.1_git20220219-r2 \
        mpc1=1.2.1-r0 \
        mpfr4=4.1.0-r0 \
        musl=1.2.3-r0 \
        musl-dev=1.2.3-r0 \
        nghttp2-libs=1.47.0-r0 \
        pcre2=10.40-r0 \
        zlib=1.2.12-r1 \
        ; \
    # Set up the go path
    mkdir -p $GOPATH/src $GOPATH/bin; \
    chmod -R 777 $GOPATH; \
    go install github.com/google/go-licenses@$GOLICENSETAG
