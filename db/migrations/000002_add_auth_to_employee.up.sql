ALTER TABLE employee
ADD COLUMN login VARCHAR(50),
ADD COLUMN hashed_password VARCHAR(255);

ALTER TABLE employee
ADD CONSTRAINT employee_login_unique UNIQUE (login); 