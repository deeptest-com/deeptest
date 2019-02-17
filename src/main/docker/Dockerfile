FROM java:8
VOLUME /tmp

ADD wait-for-it.sh /
RUN chmod +x /wait-for-it.sh

ADD postgres/setup.sh /setup.sh
RUN chmod +x /setup.sh

ADD postgres/schema.sql /

COPY sources.list /etc/apt/sources.list

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN export PATH="/:$PATH"
ADD ngtesting-web-*.jar ngtesting-web.jar
RUN sh -c 'touch /ngtesting-web.jar'
EXPOSE 8080
ENV JAVA_OPTS=""

CMD [ "sh", "-c", "echo 'Wait for postgres lauching ...'" ]
