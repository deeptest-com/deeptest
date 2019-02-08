rm init-* schema.sql

pg_dump -h localhost -U postgres -c ngtesting-web -s > schema.sql

psql -h localhost -U postgres ngtesting-web  -c "select '-t \"' || tablename || '\"' from pg_tables where schemaname='public' and tablename like '%Define' order by tablename;" | grep Define  > init-tables.txt
pg_dump -h localhost -U postgres ngtesting-web -a `cat init-tables.txt` > init-data.sql

cat init-data.sql >> schema.sql

rm init-*

#psql ngtesting-test -U postgres < schema.sql
