use mysql;
select host, user from user;
create user ngtesting identified by 'P2ssw0rd';
grant all on *.* to ngtesting@'%' identified by 'P2ssw0rd' with grant option;
flush privileges;
