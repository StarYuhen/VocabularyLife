create table adminuseraccount
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

create unique index adminuseraccount_ID_uindex
    on adminuseraccount (ID);

alter table adminuseraccount
    add constraint adminuseraccount_pk
        primary key (ID);

alter table adminuseraccount
    modify ID smallint auto_increment;

alter table adminuseraccount
    add WordLimit int default 600 not null;

alter table adminuseraccount
    add PlannedWords int default 0 not null;

alter table adminuseraccount
    add UsedWords int default 0 not null;



