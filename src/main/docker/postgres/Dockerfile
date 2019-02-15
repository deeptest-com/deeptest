FROM postgres:11.2

ADD wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

ADD setup.sh /setup.sh
RUN chmod +x /setup.sh

ADD schema.sql /schema.sql

EXPOSE 5432
