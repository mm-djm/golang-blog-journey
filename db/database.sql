CREATE TABLE t_admin(
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    user_name varchar(255) DEFAULT NULL,
    user_password varchar(255) DEFAULT NULL,
    add_time int(10) unsigned DEFAULT '0',
    update_time int(10) unsigned DEFAULT '0',
    role varchar(255) DEFAULT NULL,
    is_deleted int(10) unsigned DEFAULT '0',
    email varchar(255) NOT NULL,
    PRIMARY KEY (`id`, `email`)
) ENGINE = InnoDB AUTO_INCREMENT = 325 DEFAULT CHARSET = utf8;

CREATE TABLE t_blog(
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) DEFAULT NULL,
    tag varchar(255) DEFAULT NULL,
    add_time int(10) unsigned DEFAULT '0',
    update_time int(10) unsigned DEFAULT '0',
    is_deleted int(10) unsigned DEFAULT '0',
    article_id varchar(255) DEFAULT NULL,
    read_count int(10) unsigned DEFAULT '0',
    content text,
    abstract text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `article_id` (`article_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 325 DEFAULT CHARSET = utf8;