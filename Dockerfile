FROM busybox:1.35.0
LABEL maintainer="Nho Luong <luongutnho@hotmail.com>"
COPY _output/linux/amd64/jb /

ENTRYPOINT ["/jb"]
