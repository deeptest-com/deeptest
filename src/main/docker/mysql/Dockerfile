FROM mysql:5.7

VOLUME my.cnf:/etc/mysql/my.cnf mysql

ADD setup.sh /etc/mysql/setup.sh
ADD schema.sql /etc/mysql/schema.sql
ADD privileges.sql /etc/mysql/privileges.sql

EXPOSE 3306
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes

CMD ["sh", "/etc/mysql/setup.sh"]
