FROM nim65s/matrix-webhook:latest

ADD diff.patch /tmp
RUN patch -p1 < /tmp/diff.patch
