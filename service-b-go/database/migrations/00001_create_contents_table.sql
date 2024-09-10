-- +goose Up
-- +goose StatementBegin
create table if not exists
    contents (
                 id bigserial,
                 text text not null,
                 text_localized json not null,
                 created_at timestamp not null,
                 updated_at timestamp not null,
                 primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists contents;
-- +goose StatementEnd
