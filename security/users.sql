drop table interaction;
drop table userlog;
drop table userrole;
drop table user;
drop table role;

create table user (
  id integer primary key autoincrement,
  username text not null unique,
  password test null,
  name text not null,
  isadmin boolean not null,
  isactive boolean not null
);

create table userlog (
  id integer primary key autoincrement,
  userid integer not null,
  logdate datetime not null,
  action text not null,
  data text null,
  foreign key(userid) references user(id)
);

create table role (
  id integer primary key autoincrement,
  description text not null unique,
  data text null
);

create table userrole (
  userid integer not null,
  roleid integer not null,
  foreign key(userid) references user(id),
  foreign key(roleid) references role(id),
  primary key(userid,roleid)
);

create table interaction (
	 id integer primary key autoincrement,
   key text not null,
   action text not null,
   userid integer not null,
   actiondate datetime not null,
   isactive boolean not null,
   expiredate datetime not null,
   foreign key(userid) references user(id)
);

insert into user(username, password, name, isadmin, isactive) values ('bill', '', 'Bill', 0, 1);
insert into user(username, password, name, isadmin, isactive) values ('bob', '', 'Bob', 0, 1);
insert into user(username, password, name, isadmin, isactive) values ('tim', '', 'Tim', 0, 1);
insert into user(username, password, name, isadmin, isactive) values ('ted', '', 'Ted', 0, 1);

insert into role(description, data) values ('Normal', '');
insert into role(description, data) values ('Administrator', '');

insert into userrole(userid, roleid) values (1, 1);
insert into userrole(userid, roleid) values (1, 2);
insert into userrole(userid, roleid) values (2, 1);
insert into userrole(userid, roleid) values (3, 1);
insert into userrole(userid, roleid) values (4, 1);


