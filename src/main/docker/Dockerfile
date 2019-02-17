FROM registry.cn-shanghai.aliyuncs.com/ngtesting/ngtesting-web-base:2.0.0
VOLUME /tmp

ADD wait-for-it.sh /
RUN chmod +x /wait-for-it.sh

ADD postgres/setup.sh /setup.sh
RUN chmod +x /setup.sh

ADD launch.sh /
RUN chmod +x /launch.sh

ADD postgres/schema.sql /

RUN export PATH="/:$PATH"
ADD ngtesting-web-*.jar ngtesting-web.jar
RUN sh -c 'touch /ngtesting-web.jar'
EXPOSE 8080
ENV JAVA_OPTS=""

CMD [ "sh", "-c", "echo 'Wait for postgres lauching ...'" ]
