
--------------------------------------------------------
--  DDL for Table USERS
--------------------------------------------------------

CREATE TABLE "USERS" (
    "ID" NUMBER, 
    "USERNAME" VARCHAR2(200), 
    "USER_ROLE" VARCHAR2(20), 
    "USER_STATUS" VARCHAR2(20), 
    "TOKEN" VARCHAR2(500), 
    "SESSION_EXPIRY_DT" DATE, 
    "LAST_LOGIN_DT" DATE, 
    "LAST_LOGIN_IP" VARCHAR2(20)
    );
    
INSERT INTO USERS (ID,USERNAME,USER_ROLE,USER_STATUS,TOKEN,SESSION_EXPIRY_DT,LAST_LOGIN_DT,LAST_LOGIN_IP) 
VALUES (1,'einstein','admin','E',null,null,null,null);


--------------------------------------------------------
--  DDL for Table CUSTOMERS
--------------------------------------------------------

CREATE TABLE "CUSTOMERS" ("CUSTOMER_ID" VARCHAR2(20), "NAME" VARCHAR2(444), "EGN" VARCHAR2(20), "ADDRESS" VARCHAR2(444));

INSERT INTO CUSTOMERS (CUSTOMER_ID,NAME,EGN,ADDRESS) VALUES ('111111111','Ivan Petrov','8701126554','Sofia, 14 Tintiava Str.');
INSERT INTO CUSTOMERS (CUSTOMER_ID,NAME,EGN,ADDRESS) VALUES ('222222222','Maria Ivanova','9701126554','Sofia, 14 Smokinia Str.');


--------------------------------------------------------
--  DDL for Table ACCOUNTS
--------------------------------------------------------

CREATE TABLE "ACCOUNTS" ("IBAN" VARCHAR2(55), "BALANCE" NUMBER, "CUSTOMER_ID" VARCHAR2(20));

INSERT INTO ACCOUNTS (IBAN,BALANCE,CUSTOMER_ID) VALUES ('BG56FINV91501215766563',4425,'111111111');
INSERT INTO ACCOUNTS (IBAN,BALANCE,CUSTOMER_ID) VALUES ('BG56FINV91503453453455',244,'111111111');
INSERT INTO ACCOUNTS (IBAN,BALANCE,CUSTOMER_ID) VALUES ('BG56FINV91503643634635',55000,'111111111');
INSERT INTO ACCOUNTS (IBAN,BALANCE,CUSTOMER_ID) VALUES ('BG56FINV91503664564566',666,'111111111');
INSERT INTO ACCOUNTS (IBAN,BALANCE,CUSTOMER_ID) VALUES ('BG56FINV91507556577567',0,'111111111');