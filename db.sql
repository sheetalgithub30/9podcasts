CREATE TABLE categories (
  id serial primary key,
  title varchar(20)
);

CREATE TABLE keywords (
  item varchar(20) primary key
);

CREATE TABLE podcasts(
  id serial primary key,
  title varchar(200),
  description varchar(4000),
  website_address varchar(255),
  category_id integer,
  -- related_categories_ids array(integer)
  language varchar(6),
  is_explicit boolean,
  cover_art_id integer,
  author_name varchar(40),
  author_email varchar(75),
  copyright varchar(40),
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE episodes(
  id serial primary key,
  podcast_id integer,
  title varchar(50),
  description text,
  season_no smallint,
  episode_no smallint,
  type_of_episode integer,
  is_explicit boolean,
  episode_art_id integer,
  episode_content_id integer,
  -- keyword_ids array(integer)
  published boolean,
  published_at timestamp,
  updated_at timestamp,
  created_at timestamp
);


CREATE TABLE users(
  id serial primary key,
  name varchar(40),
  email varchar(75),
  password varchar(255),
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE uploads(
  id serial primary key,
  name varchar(255),
  path varchar(400)
);

CREATE TABLE feed (
  podcast_id serial primary key,
  rss_feed text 
);