/wait-for-it.sh ngtesting-postgres:5432 -t 10;
/setup.sh;
/usr/bin/java -Djava.security.egd=file:/dev/./urandom -jar /ngtesting-web.jar --spring.profiles.active=docker;