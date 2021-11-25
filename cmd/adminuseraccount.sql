create table adminaccounts
(
    ID                  smallint,
    uid                 smallint           not null,
    user                TEXT               not null ,
    userName            TEXT               not null,
    passWord            TEXT               not null,
    userLV              int      default 1 not null,
    imgUrl              LONGTEXT           not null,
    accountTime         DATE               not null,
    ContinuousCheckIn   SMALLINT default 0 not null,
    AccumulativeCheckIn SMALLINT default 0 not null,
    MaximumContinuous   SMALLINT default 0 not null
)
    comment '词汇人生账号主数据库';

create unique index adminaccounts_ID_uindex
    on adminaccounts (ID);

alter table adminaccounts
    add constraint adminaccounts_pk
        primary key (ID);

alter table adminaccounts
    modify ID smallint auto_increment;

alter table adminaccounts
    add WordLimit int default 600 not null;

alter table adminaccounts
    add PlannedWords int default 0 not null;

alter table adminaccounts
    add UsedWords int default 0 not null;

create unique index adminaccounts_uid_uindex
    on adminaccounts (uid);

