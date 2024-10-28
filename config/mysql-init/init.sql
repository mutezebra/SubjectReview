CREATE TABLE IF NOT EXISTS user (
    id  BIGINT UNSIGNED COMMENT 'PK',
    user_name varchar(255) NOT NULL DEFAULT 0 COMMENT '用户名称',
    email varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
    password_digest varchar(255) NOT NULL DEFAULT '' COMMENT '加密后的密码',
    avatar varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像的相对路径',
    created_at BIGINT UNSIGNED COMMENT '创建时间',
    deleted_at BIGINT UNSIGNED COMMENT '删除时间',
    PRIMARY KEY pk_user(id),
    INDEX idx_deleted_at(deleted_at),
    INDEX idl_email_at(email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户基础表' ;

CREATE TABLE IF NOT EXISTS subject (
    id BIGINT UNSIGNED COMMENT 'PK',
    name varchar(255) NOT NULL DEFAULT '' COMMENT '题目',
    answer text NOT NULL COMMENT '回答',
    subject_type smallint NOT NULL DEFAULT 0 COMMENT '题目类型',
    creator_id  BIGINT UNSIGNED COMMENT 'PK',
    created_at BIGINT UNSIGNED COMMENT '创建时间',
    PRIMARY KEY pk_subject(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='subject表' ;

CREATE TABLE IF NOT EXISTS private_subject (
    id BIGINT UNSIGNED COMMENT 'PK',
    name varchar(255) NOT NULL DEFAULT '' COMMENT '题目',
    answer varchar(255) NOT NULL DEFAULT '' COMMENT '回答',
    subject_type smallint NOT NULL DEFAULT 0 COMMENT '题目类型',
    creator_id  BIGINT UNSIGNED COMMENT 'PK',
    created_at BIGINT UNSIGNED COMMENT '创建时间',
    deleted_at BIGINT UNSIGNED COMMENT '删除时间',
    PRIMARY KEY pk_subject(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='private_subject' ;

CREATE TABLE IF NOT EXISTS user_with_subject (
    user_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    subject_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'subject id',
    subject_type smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT '题目类型',
    phase smallint NOT NULL DEFAULT 0 COMMENT '阶段',
    learn_times smallint UNSIGNED NOT NULL DEFAULT 0 COMMENT 'review次数',
    last_review_at BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '上一次review的时间',
    INDEX idx_user_id(user_id),
    INDEX idx_subject_id(subject_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='user_with_subject' ;

CREATE TABLE IF NOT EXISTS remind (
    user_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    subject_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'subject id',
    remind BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '提醒时间',
    index idx_user_id (user_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='remind' ;

