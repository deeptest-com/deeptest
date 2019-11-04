#!/bin/bash

echo '1. 检测postgres状态'
export PGPASSWORD=postgres;
count=`psql -h ngtesting-postgres -U postgres ngtesting-web  -c "select tablename from pg_tables where schemaname='public'" | grep TstProject`
echo "count=$count"

set -e

if [ "$count" != "" ]; then
    echo "数据已初始化，退出"
else
    echo '2. 创建数据库....'
    psql -h ngtesting-postgres -U postgres  -c 'CREATE DATABASE "ngtesting-web" OWNER postgres'
    echo '2. 成功创建数据库'

    echo '3. 导入数据....'
    psql -h ngtesting-postgres -U postgres "ngtesting-web" < /schema.sql
    echo '3. 成功导入数据'

    touch ~/init-success
fi