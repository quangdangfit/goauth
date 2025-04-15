create table users
(
    id           bigint primary key,
    created_at   datetime(3)  null,
    updated_at   datetime(3)  null,
    deleted_at   datetime(3)  null,
    username   longtext     null,
    amount       double       null,
    from_account longtext     null,
    to_account   longtext     null,
    description  longtext     null,
    note         longtext     null,
    status       longtext     null,
    delivery_at  datetime(3)  null,
    constraint idx_transactions_uuid
        unique (uuid)
);

create index idx_transactions_deleted_at on
    transactions(deleted_at);

create index idx_transactions_status on
    transactions(status);
