-- +goose Up
-- +goose StatementBegin
insert into contents("text", "text_localized", "created_at", "updated_at") values
('hello world', '{"pt-PT": "olá mundo"}', now(), now()),
('Goodbye World!', '{"pt-PT": "Adeus Mundo!"}', now(), now());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table contents restart identity cascade;
-- +goose StatementEnd
