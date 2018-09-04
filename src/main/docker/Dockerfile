FROM java:8
VOLUME /tmp
ADD ngtesting-web-*.jar ngtesting-web.jar
RUN sh -c 'touch /ngtesting-web.jar'
EXPOSE 8080
ENV JAVA_OPTS=""
ENTRYPOINT [ "sh", "-c", "echo 'Wait for MySQL lauching ...'; sleep 3; java $JAVA_OPTS -Djava.security.egd=file:/dev/./urandom -jar /ngtesting-web.jar --spring.profiles.active=docker" ]
