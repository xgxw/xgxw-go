-- +goose Up
-- +goose StatementBegin
CREATE TABLE `xgxw`.`users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(40) DEFAULT NULL COMMENT 'file名称',
  `phone` char(11) DEFAULT NULL COMMENT 'size',
  `password` varchar(250) DEFAULT NULL COMMENT 'url地址',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
  `updated_at` TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间' ,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `xgxw`.`users`;
-- +goose StatementEnd
