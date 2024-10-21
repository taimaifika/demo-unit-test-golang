CREATE DATABASE gotest;

CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_date DATE,
    due_date DATE,
    status TINYINT NOT NULL,
    priority TINYINT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)  ENGINE=INNODB;

-- insert some data
INSERT INTO tasks (title, start_date, due_date, status, priority, description) VALUES
('Task 1', '2020-01-01', '2020-01-02', 1, 1, 'Description 1'),
('Task 2', '2020-02-01', '2020-02-02', 1, 1, 'Description 2'),
('Task 3', '2020-03-01', '2020-03-02', 1, 1, 'Description 3');
