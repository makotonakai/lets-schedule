create table if not exists users (
  `id` bigint(11) not null auto_increment,
  `username` varchar(191) not null,
  `email_address` varchar(191) not null,
  `password` varchar(191) not null,
  `is_admin` boolean,
  `can_login` boolean,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into users values (1, "makoto", "makoto@email.com", "password", 1, 1, current_timestamp, current_timestamp);
insert into users values (2, "minoru", "minoru@email.com", "password", 1, 1, current_timestamp, current_timestamp);


create table if not exists meetings (
  `id` bigint(11) not null auto_increment,
  `title` varchar(191) not null,
  `description` text,
  `start_time` timestamp not null,
  `end_time` timestamp not null,
  `type` varchar(191),
  `meeting_place` varchar(191),
  `meeting_url` varchar(191),
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into meetings values (1, "AQUA meeting", "Just a weekly meeting", current_timestamp, current_timestamp, "meeting", "delta", "http://keio-zoom.com", current_timestamp, current_timestamp);
insert into meetings values (2, "Meeting with Aram-san", "Progress report", current_timestamp, current_timestamp, "meeting", "", "http://gaiax-zoom.com", current_timestamp, current_timestamp);

create table if not exists participants (
  `id` bigint(11) not null auto_increment,
  `meeting_id` bigint(11) not null,
  `user_id` bigint(11) not null,
  `is_host` boolean not null,
  `has_responded` boolean not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into participants values (1, 1, 1, 1, 1, current_timestamp, current_timestamp);
insert into participants values (2, 1, 2, 0, 0, current_timestamp, current_timestamp);


create table if not exists candidate_times (
  `id` bigint(11) not null auto_increment,
  `meeting_id` bigint(11) not null,
  `user_id` bigint(11) not null,
  `is_host` boolean not null,
  `has_responded` boolean not null,
  `start_time` timestamp not null,
  `end_time` timestamp not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into candidate_times values (1, 1, 1, 1, 1, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
insert into candidate_times values (2, 1, 1, 0, 0, current_timestamp, current_timestamp, current_timestamp, current_timestamp);
