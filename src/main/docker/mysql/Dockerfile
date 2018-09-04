FROM mysql:5.7

ADD my.cnf /etc/mysql/conf.d/mysql.cnf

ADD setup.sh /etc/mysql/setup.sh
ADD database.sql /etc/mysql/database.sql
ADD schema.sql /etc/mysql/schema.sql
ADD privileges.sql /etc/mysql/privileges.sql

EXPOSE 3306
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes

CMD ["sh", "/etc/mysql/setup.sh"]
