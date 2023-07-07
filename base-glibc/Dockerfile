# Generated from https://github.com/csmith/dockerfiles/blob/master/base-glibc/Dockerfile.gotpl
# BOM: {"apk:busybox":"1.36.1-r1","apk:busybox-binsh":"1.36.1-r1","apk:ca-certificates":"20230506-r0","apk:libacl":"2.3.1-r3","apk:libc6-compat":"1.2.4-r0","apk:libcrypto3":"3.1.1-r1","apk:libgcc":"12.2.1_git20220924-r10","apk:libstdc++":"12.2.1_git20220924-r10","apk:libxxhash":"0.8.1-r4","apk:lz4-libs":"1.9.4-r4","apk:musl":"1.2.4-r0","apk:popt":"1.19-r2","apk:rsync":"3.2.7-r4","apk:tzdata":"2023c-r1","apk:zlib":"1.2.13-r1","apk:zstd-libs":"1.5.5-r4","image:alpine":"0ce982c1a177d43800f731a14ddb968b2cfd1e94017d35961064e62634a7a5b8"}

FROM reg.c5h.io/alpine@sha256:0ce982c1a177d43800f731a14ddb968b2cfd1e94017d35961064e62634a7a5b8 AS build

RUN set -eux; \
    # Sort out our dependencies
    apk add --no-cache \
        busybox=1.36.1-r1 \
        busybox-binsh=1.36.1-r1 \
        ca-certificates=20230506-r0 \
        libacl=2.3.1-r3 \
        libcrypto3=3.1.1-r1 \
        libxxhash=0.8.1-r4 \
        lz4-libs=1.9.4-r4 \
        musl=1.2.4-r0 \
        popt=1.19-r2 \
        rsync=3.2.7-r4 \
        zlib=1.2.13-r1 \
        zstd-libs=1.5.5-r4 \
        ; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy our data
    apk add --no-cache libc6-compat=1.2.4-r0; \
    apk info -qL libc6-compat | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libgcc=12.2.1_git20220924-r10; \
    apk info -qL libgcc | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache libstdc++=12.2.1_git20220924-r10; \
    apk info -qL libstdc++ | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache musl=1.2.4-r0; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    apk add --no-cache tzdata=2023c-r1; \
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    # Generate the CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Create mount points for various special things
    mkdir /rootfs/proc /rootfs/dev /rootfs/sys;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
