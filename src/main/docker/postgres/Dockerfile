FROM postgres:11.2

ADD wait-for-it.sh /etc/postgres/wait-for-it.sh
RUN chmod +x /etc/postgres/wait-for-it.sh

ADD setup.sh /etc/postgres/setup.sh
RUN chmod +x /etc/postgres/setup.sh

ADD schema.sql /etc/postgres/schema.sql

EXPOSE 5432
