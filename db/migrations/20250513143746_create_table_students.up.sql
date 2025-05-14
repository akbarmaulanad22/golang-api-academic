CREATE TABLE students (
    npm VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    date_birth DATE,
    address TEXT,
    gender VARCHAR(10),
    class VARCHAR(20), -- pagi/malam/karyawan
    registration_wave INT,
    registration_date DATE,
    user_id INT,
    study_program_id INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (study_program_id) REFERENCES study_programs(id)
);