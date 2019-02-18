FROM java:8
VOLUME /tmp

COPY postgres/sources.list /etc/apt/sources.list

RUN apt-get update
RUN apt-get -y install postgresql-client

CMD [ "sh", "-c", "echo 'success'" ]
