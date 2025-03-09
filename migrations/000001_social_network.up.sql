create extension if not exists "uuid-ossp";

create table "user" (
    id uuid primary key default uuid_generate_v4(),
    username varchar(50) not null unique,
    password text not null,
    created_at timestamptz default current_timestamp
);

create table post (
    id uuid primary key default uuid_generate_v4(),
    user_id uuid not null,
    title text not null,
    content text not null,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    foreign key (user_id) references "user"(id) on delete cascade
);

create table follower (
    follower_id uuid not null,
    following_id uuid not null,
    primary key (follower_id, following_id),
    foreign key (follower_id) references "user"(id) on delete cascade,
    foreign key (following_id) references "user"(id) on delete cascade
);

create table user_reaction (
    user_id uuid not null,
    post_id uuid not null,
    reaction text not null,
    created_at timestamptz default current_timestamp,
    primary key (user_id, post_id),
    foreign key (user_id) references "user"(id) on delete cascade,
    foreign key (post_id) references post(id) on delete cascade
);