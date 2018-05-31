create table user (
  id integer primary key autoincrement,
  username text not null,
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
  description text not null
);

create table userrole (
  userid integer not null,
  roleid integer not null,
  foreign key(userid) references user(id)
  foreign key(roleid) references role(id)
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
