CREATE TABLE courses (
    code VARCHAR(20) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    sks INT,
    semester INT,
    total_meetings INT,
    lecturer_nidn INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (lecturer_nidn) REFERENCES lecturers(nidn)
);