ALTER TABLE employee
DROP CONSTRAINT IF EXISTS employee_login_unique;

ALTER TABLE employee
DROP COLUMN IF EXISTS login,
DROP COLUMN IF EXISTS hashed_password; 