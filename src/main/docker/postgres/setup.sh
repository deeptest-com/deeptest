#!/bin/bash

echo '1. 检测postgres状态'
export PGPASSWORD=P2ssw0rd;
count=`psql -h ngtesting-postgres -U postgres ngtesting-web  -c "select tablename from pg_tables where schemaname='public' | grep TstProject"`
echo "count=$count"

set -e

if [ "$count" != "" ]; then
    echo "数据已初始化，退出"
else
    echo '2. 创建用户....'
    psql -h ngtesting-postgres -U postgres -c "CREATE USER ngtesting WITH PASSWORD 'P2ssw0rd'"
    echo '2. 成功创建用户'

    echo '3. 创建数据库....'
    psql -h ngtesting-postgres -U postgres  -c 'CREATE DATABASE "ngtesting-web" OWNER ngtesting'
    echo '3. 成功创建数据库'

    echo '4. 导入数据....'
    psql -h ngtesting-postgres -U ngtesting "ngtesting-web" < /schema.sql
    echo '4. 成功导入数据'

    echo '5. 初始化全文检索....'
    psql -h ngtesting-postgres -U ngtesting ngtesting-web  -c 'CREATE EXTENSION zhparser;'
    psql -h ngtesting-postgres -U ngtesting ngtesting-web  -c 'CREATE TEXT SEARCH CONFIGURATION chinese_zh (PARSER = zhparser);'
    psql -h ngtesting-postgres -U ngtesting ngtesting-web  -c 'ALTER TEXT SEARCH CONFIGURATION chinese_zh ADD MAPPING FOR n,v,a,i,e,l WITH simple;'
    echo '4. 成功初始化全文检索'

    touch ~/init-success
fi