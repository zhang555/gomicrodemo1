FROM mysql:5.7.25

RUN cp /usr/share/zoneinfo/PRC /etc/localtime

COPY micro.sql /docker-entrypoint-initdb.d
