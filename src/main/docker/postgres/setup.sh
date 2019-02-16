#!/bin/bash

echo '1. 安装postgres客户端'
apt-get install postgresql-client
echo '1. 成功安装postgres客户端'

echo '2. 检测postgres状态'
export PGPASSWORD=P2ssw0rd;
count=`psql -h ngtesting-postgres -U ngtesting ngtesting-web  -c "select tablename from pg_tables where schemaname='public'" | grep TstProject`
echo "count=$count"

#set -e

if [ "$count" != "" ]; then
    echo "数据已初始化，退出"
else
    echo '3. 创建用户....'
    psql -h ngtesting-postgres -U postgres -c "CREATE USER ngtesting WITH PASSWORD 'P2ssw0rd'
    echo '3. 成功创建用户'

    echo '4. 创建数据库....'
    psql -h ngtesting-postgres -U postgres  -c 'CREATE DATABASE "ngtesting-web" OWNER ngtesting'
    echo '4. 成功创建数据库'

    echo '5. 导入数据....'
    psql -h ngtesting-postgres -U ngtesting "ngtesting-web" < /schema.sql
    echo '5. 成功导入数据'

    touch ~/init-success
fi

tail -f /dev/null