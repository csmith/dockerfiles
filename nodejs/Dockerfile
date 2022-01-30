# Generated from https://github.com/csmith/dockerfiles/blob/master/nodejs/Dockerfile.gotpl
# BOM: {"apk:bash":"5.1.8-r0","apk:brotli-libs":"1.0.9-r5","apk:busybox":"1.34.1-r3","apk:ca-certificates":"20211220-r0","apk:curl":"7.80.0-r0","apk:expat":"2.4.3-r0","apk:gcompat":"1.0.0-r4","apk:gdbm":"1.22-r0","apk:git":"2.34.1-r0","apk:gmp":"6.2.1-r1","apk:gnupg":"2.2.31-r1","apk:gnupg-dirmngr":"2.2.31-r1","apk:gnupg-gpgconf":"2.2.31-r1","apk:gnupg-utils":"2.2.31-r1","apk:gnupg-wks-client":"2.2.31-r1","apk:gnutls":"3.7.1-r0","apk:gpg":"2.2.31-r1","apk:gpg-agent":"2.2.31-r1","apk:gpg-wks-server":"2.2.31-r1","apk:gpgsm":"2.2.31-r1","apk:gpgv":"2.2.31-r1","apk:libacl":"2.2.53-r0","apk:libassuan":"2.5.5-r0","apk:libbz2":"1.0.8-r1","apk:libcap":"2.61-r0","apk:libcrypto1.1":"1.1.1l-r8","apk:libcurl":"7.80.0-r0","apk:libffi":"3.4.2-r1","apk:libgcc":"10.3.1_git20211027-r0","apk:libgcrypt":"1.9.4-r0","apk:libgpg-error":"1.42-r1","apk:libksba":"1.6.0-r0","apk:libldap":"2.6.0-r0","apk:libsasl":"2.1.27-r14","apk:libssl1.1":"1.1.1l-r8","apk:libstdc++":"10.3.1_git20211027-r0","apk:libtasn1":"4.18.0-r0","apk:libucontext":"1.1-r0","apk:libunistring":"0.9.10-r1","apk:lz4-libs":"1.9.3-r1","apk:musl":"1.2.2-r7","apk:musl-obstack":"1.2.2-r0","apk:ncurses-libs":"6.3_p20211120-r0","apk:ncurses-terminfo-base":"6.3_p20211120-r0","apk:nettle":"3.7.3-r0","apk:nghttp2-libs":"1.46.0-r0","apk:npth":"1.6-r1","apk:p11-kit":"0.24.0-r1","apk:pcre2":"10.39-r0","apk:pinentry":"1.2.0-r0","apk:popt":"1.18-r0","apk:readline":"8.1.1-r0","apk:rsync":"3.2.3-r5","apk:sqlite-libs":"3.36.0-r0","apk:tar":"1.34-r0","apk:xz":"5.2.5-r0","apk:xz-libs":"5.2.5-r0","apk:zlib":"1.2.11-r3","apk:zstd-libs":"1.5.0-r0","github:nodejs/node":"v17.4.0","image:alpine":"ef95559f475cdf56e55db71c168f1616d497e2215bbd9f6bc63398036ae546a1"}

FROM reg.c5h.io/alpine@sha256:ef95559f475cdf56e55db71c168f1616d497e2215bbd9f6bc63398036ae546a1 as build

ARG TAG="v17.4.0"

RUN set -eux; \
    apk add --no-cache \
        bash=5.1.8-r0 \
        brotli-libs=1.0.9-r5 \
        busybox=1.34.1-r3 \
        ca-certificates=20211220-r0 \
        curl=7.80.0-r0 \
        gcompat=1.0.0-r4 \
        gdbm=1.22-r0 \
        gmp=6.2.1-r1 \
        gnupg=2.2.31-r1 \
        gnupg-dirmngr=2.2.31-r1 \
        gnupg-gpgconf=2.2.31-r1 \
        gnupg-utils=2.2.31-r1 \
        gnupg-wks-client=2.2.31-r1 \
        gnutls=3.7.1-r0 \
        gpg=2.2.31-r1 \
        gpg-agent=2.2.31-r1 \
        gpg-wks-server=2.2.31-r1 \
        gpgsm=2.2.31-r1 \
        gpgv=2.2.31-r1 \
        libacl=2.2.53-r0 \
        libassuan=2.5.5-r0 \
        libbz2=1.0.8-r1 \
        libcap=2.61-r0 \
        libcrypto1.1=1.1.1l-r8 \
        libcurl=7.80.0-r0 \
        libffi=3.4.2-r1 \
        libgcc=10.3.1_git20211027-r0 \
        libgcrypt=1.9.4-r0 \
        libgpg-error=1.42-r1 \
        libksba=1.6.0-r0 \
        libldap=2.6.0-r0 \
        libsasl=2.1.27-r14 \
        libssl1.1=1.1.1l-r8 \
        libstdc++=10.3.1_git20211027-r0 \
        libtasn1=4.18.0-r0 \
        libucontext=1.1-r0 \
        libunistring=0.9.10-r1 \
        lz4-libs=1.9.3-r1 \
        musl=1.2.2-r7 \
        musl-obstack=1.2.2-r0 \
        ncurses-libs=6.3_p20211120-r0 \
        ncurses-terminfo-base=6.3_p20211120-r0 \
        nettle=3.7.3-r0 \
        nghttp2-libs=1.46.0-r0 \
        npth=1.6-r1 \
        p11-kit=0.24.0-r1 \
        pinentry=1.2.0-r0 \
        popt=1.18-r0 \
        readline=8.1.1-r0 \
        rsync=3.2.3-r5 \
        sqlite-libs=3.36.0-r0 \
        tar=1.34-r0 \
        xz=5.2.5-r0 \
        xz-libs=5.2.5-r0 \
        zlib=1.2.11-r3 \
        zstd-libs=1.5.0-r0 \
        ;\
    curl -sfLo /checksums.txt.asc https://nodejs.org/dist/$TAG/SHASUMS256.txt.asc; \
    gpg --keyserver hkps://keys.openpgp.org --keyserver-options auto-key-retrieve --verify /checksums.txt.asc; \
    curl -sfLo /node.tar.xz https://nodejs.org/dist/$TAG/node-$TAG-linux-x64.tar.xz; \
    grep ".*linux-x64.tar.xz" checksums.txt.asc | echo $(awk '{print $1}') \*node.tar.xz | sha256sum -wc -; \
    mkdir -p /src; \
    mkdir -p /usr/local; \
    tar -xf /node.tar.xz -C /src; \
    rsync -ap /src/node*/ /usr/local/; \
    rm -rf /usr/local/CHANGELOG.md; \
    rm -rf /usr/local/LICENSE; \
    rm -rf /usr/local/README.md;

FROM reg.c5h.io/alpine@sha256:ef95559f475cdf56e55db71c168f1616d497e2215bbd9f6bc63398036ae546a1

RUN set -eux; \
    apk add --no-cache \
        brotli-libs=1.0.9-r5 \
        busybox=1.34.1-r3 \
        ca-certificates=20211220-r0 \
        expat=2.4.3-r0 \
        gcompat=1.0.0-r4 \
        git=2.34.1-r0 \
        libcrypto1.1=1.1.1l-r8 \
        libcurl=7.80.0-r0 \
        libgcc=10.3.1_git20211027-r0 \
        libssl1.1=1.1.1l-r8 \
        libstdc++=10.3.1_git20211027-r0 \
        libucontext=1.1-r0 \
        musl=1.2.2-r7 \
        musl-obstack=1.2.2-r0 \
        nghttp2-libs=1.46.0-r0 \
        pcre2=10.39-r0 \
        zlib=1.2.11-r3 \
        ; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /etc/passwd; \
    mkdir /home/nonroot; \
    chown 65532:65532 /home/nonroot;

COPY --from=build --chown=0:0 /usr/local /usr/local

CMD ["/usr/local/bin/node"]
