use mysql;
select host, user from user;
grant all on *.* to ngtesting@'%' identified by 'P2ssw0rd';
flush privileges;
