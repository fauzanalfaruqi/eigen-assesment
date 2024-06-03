-- Insert members
INSERT INTO member
    (code, username, password, role, total_books_borrowed, penalized_start_date, penalized_end_date, created_at, updated_at)
VALUES
    ('M001', 'Angga', '$2a$10$d3zy1Y3kbAOJMva8Du4/Ve/OTDFhYVbR14TtL1O4Iz5JeYNbyOtUS', 'MEMBER', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), -- pass: 12345
    ('M002', 'Ferry', '$2a$10$MYXE49b3KZjc4jbCrFTUzOcb3SowEeuqy6ol06B.lniSLiMfN7HaG', 'MEMBER', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), -- pass: 54321
    ('M003', 'Putri', '$2a$10$/a4YNcei6iqXZRxyA6hgN.ngclNEyQzOXuS5B1KX1s0HdX3xNV7bO', 'MEMBER', 0, '0001-01-01 00:00:00', '0001-01-01 00:00:00', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP); -- pass: 010101

-- Insert books
INSERT INTO book
    (code, title, author, stock, created_at, updated_at)
VALUES
    ('JK-45', 'Harry Potter', 'J.K Rowling', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('SHR-1', 'A Study in Scarlet', 'Arthur Conan Doyle', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('TW-11', 'Twilight', 'Stephenie Meyer', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('HOB-83', 'The Hobbit, or There and Back Again', 'J.R.R. Tolkien', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('NRN-7', 'The Lion, the Witch and the Wardrobe', 'C.S. Lewis', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);