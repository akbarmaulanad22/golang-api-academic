CREATE TABLE enrollments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    status VARCHAR(20) DEFAULT 'Aktif',
    academic_year VARCHAR(20),
    registration_date DATE,
    student_npm VARCHAR(20),
    schedule_id INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (student_npm) REFERENCES students(npm),
    FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);