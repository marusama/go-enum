/* NOTE: AUTO-GENERATED by midc, DON'T edit!! */

CREATE TABLE IF NOT EXISTS `user` (
	`id` BIGINT(20)   ,
	`name` VARCHAR(64) not null DEFAULT '' ,
	`age` BIGINT(20)   ,
	`addr_id` BIGINT(20)   ,
	`product_id` BIGINT(20)   ,
	`is_robot` TINYINT(1)  DEFAULT 0 ,
	PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `address` (
	`id` BIGINT(20) AUTO_INCREMENT  ,
	`addr` VARCHAR(1024) not null  ,
	PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `product` (
	`id` BIGINT(20) AUTO_INCREMENT  ,
	`price` INT(10) not null DEFAULT 1 ,
	`name` VARCHAR(256) not null  ,
	`image` VARCHAR(512) not null  ,
	`desc` VARCHAR(1024) not null DEFAULT 'hello' ,
	PRIMARY KEY (`id`)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

