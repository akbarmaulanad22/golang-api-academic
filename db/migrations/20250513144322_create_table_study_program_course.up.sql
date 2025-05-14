CREATE TABLE study_program_course (
    study_program_id INT,
    course_code VARCHAR(20),
    semester_wajib INT,
    tahun_akademik VARCHAR(20),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (study_program_id, course_code),
    FOREIGN KEY (study_program_id) REFERENCES study_programs(id),
    FOREIGN KEY (course_code) REFERENCES courses(code)
);