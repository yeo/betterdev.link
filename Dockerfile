FROM ubuntu:16.04

# TODO: Switch to scrtach and multistage build
# TODO: Switch to non root user

RUN \
  apt-get update && apt-get install -y \
    curl unzip ca-certificates \
 && rm -rf /var/lib/apt/lists/*


WORKDIR /app
RUN mkdir -p /app

ADD content /app/content
ADD public /app/public
ADD static /app/static
ADD themes /app/themes
ADD linux /app/bd

# Once docker fix bugs about USER we can remove chown
RUN chown -R nobody /app

USER nobody
CMD ["/app/bd", "serve"]
