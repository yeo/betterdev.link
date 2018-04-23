FROM ubuntu:16.04
# TODO: Switch to scrtach and multistage build
# TODO: Switch to non root user

WORKDIR /app
RUN mkdir -p /var/log/bd/ && \
    mkdir -p /app

USER nobody

ADD content /app/content
ADD public /app/public
ADD static /app/static
ADD themes /app/themes
ADD bd /app

CMD ["/app/bd", "serve"]
