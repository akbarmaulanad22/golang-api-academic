CREATE TABLE schedules (
    id INT AUTO_INCREMENT PRIMARY KEY,
    date DATE,
    start_at TIME,
    end_at TIME,
    course_code VARCHAR(20),
    lecturer_nidn VARCHAR(50),
    classroom_id INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (course_code) REFERENCES courses(code),
    FOREIGN KEY (lecturer_nidn) REFERENCES lecturers(nidn),
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id)
);