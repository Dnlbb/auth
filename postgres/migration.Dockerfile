FROM  alpine:3.20

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD ./postgres/migrations/*.sql migrations/
ADD ./postgres/migration.sh .
ADD ./postgres/.env .

RUN chmod +x migration.sh

ENTRYPOINT ["bash", "migration.sh"]