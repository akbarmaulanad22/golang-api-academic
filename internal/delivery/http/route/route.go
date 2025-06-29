package route

import (
	controller "tugasakhir/internal/delivery/http"

	"github.com/gorilla/mux"
)

type RouteConfig struct {
	// router
	Router *mux.Router

	// middleware
	AuthMiddleware mux.MiddlewareFunc

	// all field controller
	UserController         *controller.UserController
	AttendanceController   *controller.AttendanceController
	ScheduleController     *controller.ScheduleController
	EnrollmentController   *controller.EnrollmentController
	GradeController        *controller.GradeController
	CourseController       *controller.CourseController
	StudentController      *controller.StudentController
	StudyProgramController *controller.StudyProgramController
	FacultyController      *controller.FacultyController
}

func (route *RouteConfig) Setup() {
	route.SetupGuestRoute()
	route.SetupAuthRoute()
}

func (route *RouteConfig) SetupGuestRoute() {
	// routes that do not require authentication
	route.Router.HandleFunc("/register", route.UserController.Register).Methods("POST")
	route.Router.HandleFunc("/login", route.UserController.Login).Methods("POST")
}

func (route *RouteConfig) SetupAuthRoute() {

	// Buat subrouter khusus untuk route yang butuh auth
	authRouter := route.Router.PathPrefix("/").Subrouter()
	authRouter.Use(route.AuthMiddleware)

	authRouter.HandleFunc("/logout", route.UserController.Logout).Methods("POST")

	authRouter = route.Router.PathPrefix("/api/v1/").Subrouter()
	authRouter.Use(route.AuthMiddleware)

	student := authRouter.PathPrefix("/student").Subrouter()

	student.HandleFunc("/attendance", route.AttendanceController.AttendStudent).Methods("POST")
	student.HandleFunc("/schedules", route.ScheduleController.ListByStudentUserID).Methods("GET")
	student.HandleFunc("/enrollments", route.EnrollmentController.ListByStudentUserID).Methods("GET")
	student.HandleFunc("/grades", route.GradeController.ListByStudentUserID).Methods("GET")

	lecturer := authRouter.PathPrefix("/lecturer").Subrouter()
	lecturer.HandleFunc("/attendance", route.AttendanceController.AttendLecturer).Methods("POST")
	lecturer.HandleFunc("/courses", route.CourseController.ListByLecturerUserID).Methods("GET")
	lecturer.HandleFunc("/courses/{courseCode}/students", route.StudentController.ListByCourseCode).Methods("GET")
	lecturer.HandleFunc("/courses/{courseCode}/students/{npm}/attendances", route.AttendanceController.ListByCourseCodeAndNpm).Methods("GET")
	lecturer.HandleFunc("/courses/{courseCode}/students/{npm}/attendances/{id}", route.AttendanceController.Update).Methods("PUT")
	lecturer.HandleFunc("/courses/{courseCode}/students/{npm}/grades", route.GradeController.ListByNpmAndCourseCode).Methods("GET")
	lecturer.HandleFunc("/schedules", route.ScheduleController.ListByLecturerUserID).Methods("GET")

	admin := authRouter.PathPrefix("/admin").Subrouter()
	admin.HandleFunc("/study-programs", route.StudyProgramController.Create).Methods("POST")
	admin.HandleFunc("/study-programs", route.StudyProgramController.List).Methods("GET")
	admin.HandleFunc("/study-programs/{id}", route.StudyProgramController.Update).Methods("PUT")
	admin.HandleFunc("/study-programs/{id}", route.StudyProgramController.Delete).Methods("DELETE")

	admin.HandleFunc("/faculties", route.FacultyController.Create).Methods("POST")
	admin.HandleFunc("/faculties", route.FacultyController.List).Methods("GET")
	admin.HandleFunc("/faculties/{id}", route.FacultyController.Update).Methods("PUT")
	admin.HandleFunc("/faculties/{id}", route.FacultyController.Delete).Methods("DELETE")

	admin.HandleFunc("/courses", route.CourseController.Create).Methods("POST")
	admin.HandleFunc("/courses", route.CourseController.List).Methods("GET")
	admin.HandleFunc("/courses/{code}", route.CourseController.Update).Methods("PUT")
	admin.HandleFunc("/courses/{code}", route.CourseController.Delete).Methods("DELETE")

	admin.HandleFunc("/enrollments", route.EnrollmentController.Create).Methods("POST")
	admin.HandleFunc("/enrollments", route.EnrollmentController.List).Methods("GET")
	admin.HandleFunc("/enrollments/{id}", route.EnrollmentController.Update).Methods("PUT")
	admin.HandleFunc("/enrollments/{id}", route.EnrollmentController.Delete).Methods("DELETE")

}
