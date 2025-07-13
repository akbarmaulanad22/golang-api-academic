CREATE TABLE grades (
    id INT AUTO_INCREMENT PRIMARY KEY,
    score DECIMAL(5,2),
    enrollment_id INT,
    grade_component_id INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (enrollment_id) REFERENCES enrollments(id) ON DELETE CASCADE,
    FOREIGN KEY (grade_component_id) REFERENCES grade_components(id) ON DELETE CASCADE
);