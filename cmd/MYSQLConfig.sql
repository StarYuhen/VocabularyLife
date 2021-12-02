# 主用户表数据
create table adminaccounts
(
    ID                  smallint,
    uid                 smallint           not null,
    user                TEXT               not null,
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

# 修改字段类型
alter table adminaccounts
    modify uid varchar(512) not null;



# 关联主表的用户登录信息记录表
create table userinfos
(
    uid varchar(512),
    ip  text,
    jwt varchar(512),
    constraint userinfo_adminaccounts_uid_fk
        foreign key (uid) references adminaccounts (uid)
)
    comment '用户登录信息';

create unique index userinfo_uid_uindex
    on userinfos (uid);

alter table userinfos
    add constraint userinfo_pk
        primary key (uid);


# 增加级联更新及级联删除  ps：级联更新和删除就是另外一张表不允许有不同的主键，但并不会自动增加主键
alter table userinfos
    drop foreign key userinfo_adminaccounts_uid_fk;

alter table userinfos
    add constraint userinfo_adminaccounts_uid_fk
        foreign key (uid) references adminaccounts (uid)
            on update cascade on delete cascade;


