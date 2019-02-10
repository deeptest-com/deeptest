rm init-* schema.sql

pg_dump -h localhost -U ngtesting -c ngtesting-test -s > schema.sql

psql -h localhost -U ngtesting ngtesting-test  -c "select '-t \"' || tablename || '\"' from pg_tables where schemaname='public' and tablename like '%Define' order by tablename;" | grep Define  > init-tables.txt
pg_dump -h localhost -U ngtesting ngtesting-test -a `cat init-tables.txt` > init-data.sql

cat init-data.sql >> schema.sql

rm init-*

# psql ngtesting-web -U ngtesting < schema.sql
