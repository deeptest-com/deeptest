FROM postgres:11.2

ADD wait-for-it.sh /wait-for-it.sh
ADD setup.sh /setup.sh
ADD schema.sql /schema.sql

EXPOSE 5432

CMD ["sh", "/setup.sh"]
