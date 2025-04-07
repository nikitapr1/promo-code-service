create type promo_mode as enum ('COMMON', 'UNIQUE');

create table if not exists promo
(
    id        varchar unique primary key,
    owner_id varchar unique,
    description text,
    image_url varchar,
    active_from timestamp,
    active_until timestamp,
    created_at timestamp,
    mode promo_mode,
    target_age_from int,
    target_age_until int,
    target_country varchar,
    target_categories varchar[]
);