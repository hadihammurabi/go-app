CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
INSERT INTO users(id, email, password) VALUES
(uuid_generate_v4(), 'tes@mail.com', '$2a$10$IlXDNwWn56AuCnCKM7dvouYY8sc5xRex9ADlfL5lXbhSfzCrzOSr.');