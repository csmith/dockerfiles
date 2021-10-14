FROM reg.c5h.io/alpine@sha256:f32b56d2f268e9a4dce68b95adac9f81b8ed8f9bc7fd33ba0d804e7426f603fd AS build

RUN set -eux; \
    # Sort out our dependencies (and sources)
    apk add --no-cache \
        busybox=1.33.1-r3 \
        ca-certificates=20191127-r5 \
        libacl=2.2.53-r0 \
        libcrypto1.1=1.1.1l-r0 \
        musl=1.2.2-r3 \
        popt=1.18-r0 \
        rsync=3.2.3-r4 \
        tzdata=2021c-r0 \
        zlib=1.2.11-r3 \
        zstd-libs=1.4.9-r1 \
        ; \
    # Set up directories and password file
    mkdir -p /rootfs/etc/ssl/certs /rootfs/tmp /rootfs/home/nonroot /rootfs/root; \
    echo "root:x:0:0:root:/root:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nobody:x:65534:65534:nobody:/nonexistent:/sbin/nologin" >> /rootfs/etc/passwd; \
    echo "nonroot:x:65532:65532:nonroot:/home/nonroot:/sbin/nologin" >> /rootfs/etc/passwd; \
    chown 65532:65532 /rootfs/home/nonroot; \
    chmod 01777 /tmp; \
    # Copy the tzdata files and musl
    apk info -qL tzdata | rsync -aq --files-from=- / /rootfs/; \
    apk info -qL musl | rsync -aq --files-from=- / /rootfs/; \
    # Generate thw CA certificates file and copy the result
    /usr/sbin/update-ca-certificates; \
    cp /etc/ssl/certs/ca-certificates.crt /rootfs/etc/ssl/certs/ca-certificates.crt; \
    # Clobber all the timestamps to make the build more reproducible
    find /rootfs -exec touch --date=@0 {} \;;

FROM scratch

COPY --from=build /rootfs/ /

ENV SSL_CERT_FILE /etc/ssl/certs/ca-certificates.crt
