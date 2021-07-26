FROM reg.c5h.io/golang@sha256:eae319ad8b654dc4b903a23ab012a71b840472cf418cfee7c4e176fa9f96faeb AS build

ARG TAG="v1.7.3"

RUN set -eux; \
    # Dependencies
    apk add --no-cache git libcap; \
    # Check out and build
    git clone --depth=1 -b $TAG --single-branch https://github.com/hashicorp/vault /src; \
    cd /src; \
    go build -tags release -ldflags "-X github.com/hashicorp/vault/sdk/version.GitCommit=$(git rev-parse HEAD)" -o /vault  github.com/hashicorp/vault; \
    # Make some directories with the correct ownership to copy into the final image, and a listener config snippet
    mkdir -p /out/logs /out/file /out/config; \
    echo -e 'listener "tcp" {\n  address = "0.0.0.0:8080"\n  tls_disable = 1\n}' > /out/config/listener.hcl; \
    chown -R 65532:65532 /out; \
    # Allow the vault process to prevent its memory regions being swapped out to disk
    /usr/sbin/setcap CAP_IPC_LOCK=+ep /vault;

FROM reg.c5h.io/base@sha256:eb133a7d9c9ee3de40aa43e67fd1c79f7e317743efe85c33cbf45eae548ea8da

COPY --from=build /vault /bin/
COPY --from=build /out/ /vault/
EXPOSE 8080
VOLUME /vault/logs /vault/file
ENTRYPOINT ["/bin/vault"]
ENV VAULT_ADDRESS http://127.0.0.1:8080
CMD ["server", "-config=/vault/config/"]
