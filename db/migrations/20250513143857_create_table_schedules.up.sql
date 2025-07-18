CREATE TABLE schedules (
    id INT AUTO_INCREMENT PRIMARY KEY,
    date DATE,
    start_at TIME,
    end_at TIME,
    course_code VARCHAR(20),
    lecturer_nidn INT,
    classroom_id INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (course_code) REFERENCES courses(code) ON DELETE CASCADE,
    FOREIGN KEY (lecturer_nidn) REFERENCES lecturers(nidn) ON DELETE CASCADE,
    FOREIGN KEY (classroom_id) REFERENCES classrooms(id) ON DELETE CASCADE
);