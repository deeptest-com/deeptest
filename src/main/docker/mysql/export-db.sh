rm init-* schema.sql

mysqldump -ungtesting -pP2ssw0rd -dR ngtesting-web > schema.sql

mysql -ungtesting -pP2ssw0rd -N information_schema -e "select table_name from tables where table_schema = 'ngtesting-web' and table_name like '%Define'" > init-tables.txt
# mysqldump -ungtesting -pP2ssw0rd ngtesting-web `cat init-tables.txt` --no-create-info --insert-ignore --complete-insert --skip-extended-insert > init-data.sql
mysqldump -ungtesting -pP2ssw0rd ngtesting-web `cat init-tables.txt` --no-create-info > init-data.sql

cat init-data.sql >> schema.sql

rm init-*
