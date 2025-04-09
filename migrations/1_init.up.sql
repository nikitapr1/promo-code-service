create type promo_mode as enum ('COMMON', 'UNIQUE');

create table if not exists promo
(
    id        varchar unique primary key,
    company_id varchar,
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

create table if not exists promo_code (
    id varchar unique primary key,
    promo_id varchar,
    code varchar,
    is_activated bool,
    max_count int
)