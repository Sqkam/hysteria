FROM rockylinux:9

COPY build/hysteria-linux-amd64 /hysteria

ENTRYPOINT ["/hysteria"]
