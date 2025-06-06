create table if not exists users (
  id bigint(11) not null auto_increment,
  user_name varchar(191) not null,
  email_address varchar(191) not null,
  password varchar(191) not null,
  is_admin boolean,
  can_login boolean,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into users values (1, "user", "user@email.com", "password", 1, 1, current_timestamp, current_timestamp);
insert into users values (2, "user2", "user2@email.com", "password", 1, 1, current_timestamp, current_timestamp);
insert into users values (3, "user3", "user3@email.com", "password", 1, 1, current_timestamp, current_timestamp);

create table if not exists meetings (
  id bigint(11) not null auto_increment,
  title varchar(191) not null,
  description text,
  is_onsite boolean not null,
  is_online boolean not null,
  place varchar(191),
  url varchar(191),
  all_participants_responded boolean not null,
  is_confirmed boolean not null,
  start_time timestamp null default null,
  end_time timestamp null default null,
  actual_start_time timestamp null default null,
  actual_end_time timestamp null default null,
  hour decimal not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into meetings values (1, "第1回ミーティング", "アイデア出し", true, false, "会議室1", "なし", 1, 1, "2022-09-19 19:00:00", "2022-09-19 22:00:00", null, null, 1, current_timestamp, current_timestamp);
insert into meetings values (2, "第2回ミーティング", "要件定義", false, true, "なし", "http://meeting2-zoom.com", 1, 0, null, null, null, null, 1, current_timestamp, current_timestamp);
insert into meetings values (3, "第3回ミーティング", "技術選定", true, true, "会議室", "http://meeting3-zoom.com", 0, 0, null, null, null, null, 1, current_timestamp, current_timestamp);
insert into meetings values (4, "第4回ミーティング", "DB設計", true, true, "会議室", "http://meeting4-zoom.com", 1, 1, "2022-09-19 19:00:00", "2022-09-19 20:00:00", null, null, 1, current_timestamp, current_timestamp);
insert into meetings values (5, "第5回ミーティング", "実装", true, true, "会議室", "http://meeting5-zoom.com", 0, 0, null, null, null, null, 1, current_timestamp, current_timestamp);
insert into meetings values (6, "第6回ミーティング", "デプロイ", true, true, "会議室", "http://meeting6-zoom.com", 0, 0, null, null, null, null, 1, current_timestamp, current_timestamp);

create table if not exists participants (
  id bigint(11) not null auto_increment,
  user_id bigint(11) not null,
  meeting_id bigint(11) not null,
  is_host boolean not null,
  has_responded boolean not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into participants values (1, 1, 1, 1, 1, current_timestamp, current_timestamp);
insert into participants values (2, 2, 1, 0, 1, current_timestamp, current_timestamp);
insert into participants values (3, 1, 2, 1, 1, current_timestamp, current_timestamp);
insert into participants values (4, 2, 2, 0, 1, current_timestamp, current_timestamp);
insert into participants values (5, 3, 2, 0, 1, current_timestamp, current_timestamp);
insert into participants values (6, 1, 3, 1, 1, current_timestamp, current_timestamp);
insert into participants values (7, 2, 3, 0, 0, current_timestamp, current_timestamp);
insert into participants values (8, 1, 4, 0, 1, current_timestamp, current_timestamp);
insert into participants values (9, 2, 4, 1, 1, current_timestamp, current_timestamp);
insert into participants values (10, 1, 5, 0, 1, current_timestamp, current_timestamp);
insert into participants values (11, 2, 5, 1, 1, current_timestamp, current_timestamp);
insert into participants values (12, 1, 6, 0, 0, current_timestamp, current_timestamp);
insert into participants values (13, 2, 6, 1, 1, current_timestamp, current_timestamp);


create table if not exists candidate_times (
  id bigint(11) not null auto_increment,
  user_id bigint(11) not null,
  meeting_id bigint(11) not null,
  start_time timestamp not null,
  end_time timestamp not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into candidate_times values (1, 1, 1, "2022-09-02 10:00:00", "2022-09-02 12:00:00", current_timestamp, current_timestamp);
insert into candidate_times values (2, 2, 1, "2022-09-02 11:00:00", "2022-09-02 13:00:00", current_timestamp, current_timestamp);
insert into candidate_times values (3, 1, 2, "2022-09-19 19:00:00", "2022-09-19 22:00:00", current_timestamp, current_timestamp);
insert into candidate_times values (4, 2, 2, "2022-09-19 18:00:00", "2022-09-19 21:00:00", current_timestamp, current_timestamp);
insert into candidate_times values (5, 3, 2, "2022-09-19 17:00:00", "2022-09-19 20:00:00", current_timestamp, current_timestamp);
insert into candidate_times values (6, 1, 3, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
insert into candidate_times values (7, 2, 3, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
insert into candidate_times values (8, 1, 4, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
insert into candidate_times values (9, 2, 4, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
insert into candidate_times values (10, 1, 5, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
insert into candidate_times values (11, 2, 5, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
insert into candidate_times values (12, 1, 2, "2022-09-19 21:00:00", "2022-09-19 23:00:00", current_timestamp, current_timestamp);

create table if not exists friends (
  id bigint(11) not null auto_increment,
  user_id bigint(11) not null,
  friend_user_id bigint(11) not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into friends values (1, 1, 2, current_timestamp, current_timestamp);
