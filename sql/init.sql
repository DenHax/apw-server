DROP TABLE IF EXISTS UPLOAD;
DROP TABLE IF EXISTS SUBSYSTEM_CONDITION;
DROP TABLE IF EXISTS POWERUNIT_ASSESSMENT;
DROP TABLE IF EXISTS ENERGY;
DROP TABLE IF EXISTS CONTAINED;
DROP TABLE IF EXISTS EMPLOYEE;
DROP TABLE IF EXISTS SUBSYSTEM;
DROP TABLE IF EXISTS FUEL_ROAD;
DROP TABLE IF EXISTS FUEL_TYPE;
DROP TABLE IF EXISTS POWERUNIT;
DROP TABLE IF EXISTS EVENT;
DROP TABLE IF EXISTS ELEMENT;
create TABLE CONTAINED 
(
   IN_FUEL_ROAD_PERSENT float                          not null,
   ELEMENT_SYMBOL       varchar(10)                    not null,
   FUEL_ROAD_NUMBER     integer                        not null,
   constraint PK_CONTAINED primary key (IN_FUEL_ROAD_PERSENT)
);

create unique index CONTAINED_PK on CONTAINED (
IN_FUEL_ROAD_PERSENT ASC
);

create index CONTAINED_POWER_ROAD_FK on CONTAINED (
FUEL_ROAD_NUMBER ASC
);

create index CONTAINED_ELEM_FK on CONTAINED (
ELEMENT_SYMBOL ASC
);

create TABLE ELEMENT 
(
   ELEMENT_SYMBOL       varchar(10)                    not null,
   ELEMENT_NAME         varchar(20)                    not null,
   ELEMENT_STABILITY    varchar(18)                    not null,
   ELEMENT_RADIOACTIVITY varchar(18)                    not null,
   ELEMENT_TOXICITY     varchar(18)                    not null,
   constraint PK_ELEMENT primary key (ELEMENT_SYMBOL)
);

create unique index ELEMENT_PK on ELEMENT (
ELEMENT_SYMBOL ASC
);

create TABLE EMPLOYEE 
(
   EMPLOYEE_ID          SERIAL                        not null,
   SUBSYSTEM_ID         integer                        not null,
   EMPLOYEE_FIRSTNAME   varchar(20)                    not null,
   EMPLOYEE_SURNAME     varchar(20)                    not null,
   EMPLOYEE_LASTNAME    varchar(20)                    null,
   EMPLOYEE_TITLE       varchar(30)                    not null,
   EMPLOYEE_PHONE       char(12)                       not null,
   constraint PK_EMPLOYEE primary key (EMPLOYEE_ID)
);

create unique index EMPLOYEE_PK on EMPLOYEE (
EMPLOYEE_ID ASC
);

create index WORK_ON_FK on EMPLOYEE (
SUBSYSTEM_ID ASC
);

create index EMPLOYEE_SURNAME on EMPLOYEE (
EMPLOYEE_SURNAME ASC
);

create TABLE ENERGY 
(
   ENERGY_DATE          timestamp                      not null,
   EMPLOYEE_ID          integer                        not null,
   SUBSYSTEM_ID         integer                        not null,
   ENERGY_POOL          float                        not null,
   constraint PK_ENERGY primary key (ENERGY_DATE)
);

create unique index ENERGY_PK on ENERGY (
ENERGY_DATE ASC
);

create index FIXES_FK on ENERGY (
EMPLOYEE_ID ASC
);

create index GENERATE_FK on ENERGY (
SUBSYSTEM_ID ASC
);

create TABLE EVENT 
(
   EVENT_ID             SERIAL                        not null,
   EVENT_DESCRIPTION    varchar(200)                   not null,
   EVENT_DANGER_LEVEL   integer                        not null,
   EVENT_TYPE           varchar(50)                    not null,
   constraint PK_EVENT primary key (EVENT_ID)
);

create unique index EVENT_PK on EVENT (
EVENT_ID ASC
);

create index EVENT_TYPE on EVENT (
EVENT_TYPE ASC
);

create TABLE FUEL_ROAD 
(
   FUEL_ROAD_NUMBER     SERIAL                        not null,
   FUEL_TYPE_ID         integer                        not null,
   FUEL_ROAD_MASS       integer                        not null,
   FUEL_ROAD_CONDITION  varchar(64)                    not null,
   constraint PK_FUEL_ROAD primary key (FUEL_ROAD_NUMBER)
);

create unique index FUEL_ROAD_PK on FUEL_ROAD (
FUEL_ROAD_NUMBER ASC
);

create index REPRESENT_FK on FUEL_ROAD (
FUEL_TYPE_ID ASC
);

create TABLE FUEL_TYPE 
(
   FUEL_TYPE_ID         SERIAL                        not null,
   FUEL_TYPE_SHELL      varchar(64)                    not null,
   FUEL_TYPE_CONTACT    varchar(64)                    not null,
   FUEL_TYPE_FORM       varchar(64)                    not null,
   constraint PK_FUEL_TYPE primary key (FUEL_TYPE_ID)
);

create unique index FUEL_TYPE_PK on FUEL_TYPE (
FUEL_TYPE_ID ASC
);

create TABLE POWERUNIT 
(
   POWERUNIT_NUMBER     SERIAL                        not null,
   constraint PK_POWERUNIT primary key (POWERUNIT_NUMBER)
);

create unique index POWERUNIT_PK on POWERUNIT (
POWERUNIT_NUMBER ASC
);

CREATE TABLE POWERUNIT_ASSESSMENT 
(
   POWERUNIT_ASSESSMENT_DATE timestamp NOT NULL,
   POWERUNIT_NUMBER     integer NOT NULL,
   EMPLOYEE_ID          integer NOT NULL,
   POWERUNIT_ASSESSMENT_ENERGY_GENERATION_ASSESSMENT integer NOT NULL,
   POWERUNIT_ASSESSMENT_CONDITION integer NOT NULL,
   POWERUNIT_ASSESSMENT_ENERGY_GENERATION_AVERAGE FLOAT,
   CONSTRAINT PK_POWERUNIT_ASSESSMENT PRIMARY KEY (POWERUNIT_ASSESSMENT_DATE)
);

CREATE OR REPLACE FUNCTION update_energy_generation_average() 
RETURNS TRIGGER AS $$
BEGIN
   NEW.POWERUNIT_ASSESSMENT_ENERGY_GENERATION_AVERAGE := (
      SELECT AVG(E.ENERGY_POOL) 
      FROM ENERGY E
      WHERE E.ENERGY_DATE BETWEEN 
            (SELECT MAX(PA2.POWERUNIT_ASSESSMENT_DATE) 
             FROM POWERUNIT_ASSESSMENT PA2 
             WHERE PA2.POWERUNIT_NUMBER = NEW.POWERUNIT_NUMBER 
               AND PA2.POWERUNIT_ASSESSMENT_DATE < NEW.POWERUNIT_ASSESSMENT_DATE)
            AND NEW.POWERUNIT_ASSESSMENT_DATE
   );
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_update_energy_generation_average
BEFORE INSERT OR UPDATE ON POWERUNIT_ASSESSMENT
FOR EACH ROW EXECUTE FUNCTION update_energy_generation_average();

create unique index POWERUNIT_ASSESSMENT_PK on POWERUNIT_ASSESSMENT (
POWERUNIT_ASSESSMENT_DATE ASC
);

create index EVALUTES_FK on POWERUNIT_ASSESSMENT (
EMPLOYEE_ID ASC
);

create index BEING_EVALUATED_FK on POWERUNIT_ASSESSMENT (
POWERUNIT_NUMBER ASC
);

create TABLE SUBSYSTEM 
(
   SUBSYSTEM_ID         SERIAl                        not null,
   POWERUNIT_NUMBER     integer                        not null,
   SUBSYSTEM_NAME      varchar(30)                    not null,
   SUBSYSTEM_STATUS     varchar(20)                    not null,
   constraint PK_SUBSYSTEM primary key (SUBSYSTEM_ID)
);

create unique index SUBSYSTEM_PK on SUBSYSTEM (
SUBSYSTEM_ID ASC
);

create index BELONG_FK on SUBSYSTEM (
POWERUNIT_NUMBER ASC
);

create index SUBSYSTEM_NAME on SUBSYSTEM (
SUBSYSTEM_NAME ASC
);

create TABLE SUBSYSTEM_CONDITION 
(
   SUBSYSTEM_CONDITION_DATE timestamp                      not null,
   EVENT_ID             integer                        null,
   SUBSYSTEM_ID         integer                        not null,
   EMPLOYEE_ID          integer                        not null,
   SUBSYSTEM_CONDITION_TECHNICAL varchar(20)                    not null,
   SUBSYSTEM_CONDITION_TEMPERATURE float                          not null,
   constraint PK_SUBSYSTEM_CONDITION primary key (SUBSYSTEM_CONDITION_DATE)
);

create unique index SUBSYSTEM_CONDITION_PK on SUBSYSTEM_CONDITION (
SUBSYSTEM_CONDITION_DATE ASC
);

create index REGISTERS_FK on SUBSYSTEM_CONDITION (
EMPLOYEE_ID ASC
);

create index TRIGGERS_FK on SUBSYSTEM_CONDITION (
EVENT_ID ASC
);

create index IN_STATE_FK on SUBSYSTEM_CONDITION (
SUBSYSTEM_ID ASC
);

create TABLE UPLOAD 
(
   LOAD_DATE            timestamp                      not null,
   EMPLOYEE_ID          integer                        not null,
   SUBSYSTEM_ID         integer                        null,
   FUEL_ROAD_NUMBER     integer                        not null,
   constraint PK_UPLOAD primary key (LOAD_DATE)
);

create unique index UPLOAD_PK on UPLOAD (
LOAD_DATE ASC
);

create index UPLOAD_EMPLOYEE_FK on UPLOAD (
EMPLOYEE_ID ASC
);

create index UPLOAD_POWER_ROAD_FK on UPLOAD (
FUEL_ROAD_NUMBER ASC
);

/*==============================================================*/
/* Index: UPLOAD_SUBSYSTEM_FK                                   */
/*==============================================================*/
create index UPLOAD_SUBSYSTEM_FK on UPLOAD (
SUBSYSTEM_ID ASC
);

alter TABLE CONTAINED
   add constraint FK_CONTAINE_CONTAINED_ELEMENT foreign key (ELEMENT_SYMBOL)
      references ELEMENT (ELEMENT_SYMBOL)
      on update restrict
      on delete restrict;

alter TABLE CONTAINED
   add constraint FK_CONTAINE_CONTAINED_FUEL_ROA foreign key (FUEL_ROAD_NUMBER)
      references FUEL_ROAD (FUEL_ROAD_NUMBER)
      on update cascade
      on delete cascade;

alter TABLE EMPLOYEE
   add constraint FK_EMPLOYEE_WORK_ON_SUBSYSTE foreign key (SUBSYSTEM_ID)
      references SUBSYSTEM (SUBSYSTEM_ID)
      on update cascade
      on delete set null;

alter TABLE ENERGY
   add constraint FK_ENERGY_FIXES_EMPLOYEE foreign key (EMPLOYEE_ID)
      references EMPLOYEE (EMPLOYEE_ID)
      on update cascade
      on delete restrict;

alter TABLE ENERGY
   add constraint FK_ENERGY_GENERATE_SUBSYSTE foreign key (SUBSYSTEM_ID)
      references SUBSYSTEM (SUBSYSTEM_ID)
      on update restrict
      on delete restrict;

alter TABLE FUEL_ROAD
   add constraint FK_FUEL_ROA_REPRESENT_FUEL_TYP foreign key (FUEL_TYPE_ID)
      references FUEL_TYPE (FUEL_TYPE_ID)
      on update cascade
      on delete restrict;

alter TABLE POWERUNIT_ASSESSMENT
   add constraint FK_POWERUNI_BEING_EVA_POWERUNI foreign key (POWERUNIT_NUMBER)
      references POWERUNIT (POWERUNIT_NUMBER)
      on update restrict
      on delete restrict;

alter TABLE POWERUNIT_ASSESSMENT
   add constraint FK_POWERUNI_EVALUTES_EMPLOYEE foreign key (EMPLOYEE_ID)
      references EMPLOYEE (EMPLOYEE_ID)
      on update cascade
      on delete restrict;

alter TABLE SUBSYSTEM
   add constraint FK_SUBSYSTE_BELONG_POWERUNI foreign key (POWERUNIT_NUMBER)
      references POWERUNIT (POWERUNIT_NUMBER)
      on update restrict
      on delete restrict;

alter TABLE SUBSYSTEM_CONDITION
   add constraint FK_SUBSYSTE_IN_STATE_SUBSYSTE foreign key (SUBSYSTEM_ID)
      references SUBSYSTEM (SUBSYSTEM_ID)
      on update restrict
      on delete restrict;

alter TABLE SUBSYSTEM_CONDITION
   add constraint FK_SUBSYSTE_REGISTERS_EMPLOYEE foreign key (EMPLOYEE_ID)
      references EMPLOYEE (EMPLOYEE_ID)
      on update cascade
      on delete restrict;

alter TABLE SUBSYSTEM_CONDITION
   add constraint FK_SUBSYSTE_TRIGGERS_EVENT foreign key (EVENT_ID)
      references EVENT (EVENT_ID)
      on update restrict
      on delete restrict;

alter TABLE UPLOAD
   add constraint FK_UPLOAD_UPLOAD_EM_EMPLOYEE foreign key (EMPLOYEE_ID)
      references EMPLOYEE (EMPLOYEE_ID)
      on update cascade
      on delete restrict;

alter TABLE UPLOAD
   add constraint FK_UPLOAD_UPLOAD_PO_FUEL_ROA foreign key (FUEL_ROAD_NUMBER)
      references FUEL_ROAD (FUEL_ROAD_NUMBER)
      on update cascade
      on delete restrict;

alter TABLE UPLOAD
   add constraint FK_UPLOAD_UPLOAD_SU_SUBSYSTE foreign key (SUBSYSTEM_ID)
      references SUBSYSTEM (SUBSYSTEM_ID)
      on update restrict
      on delete restrict;

