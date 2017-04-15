package endpoints

import "fmt"

const (
	prefix = "api/v1"

	courses        = "courses"
	users          = "users"
	searchUsers    = "search_users"
	permissions    = "permissions"
	accounts       = "accounts"
	students       = "students"
	files          = "files"
	assignments    = "assignments"
	recentStudents = "recent_students"
	previewHTML    = "preview_html"
)

func AuthURL() string {
	return "login/oauth2/auth"
}

func TokenURL() string {
	return "login/oauth2/token"
}

// /api/v1/courses
func Courses() string {
	return fmt.Sprintf("%s/%s", prefix, courses)
}

// /api/v1/courses/:id
func Course(id int) string {
	return fmt.Sprintf("%s/%d", Courses(), id)
}

// /api/v1/users
func Users() string {
	return fmt.Sprintf("%s/%s", prefix, users)
}

// /api/v1/users/:id
func User(id int) string {
	return fmt.Sprintf("%s/%d", Users(), id)
}

// /api/v1/accounts
func Accounts() string {
	return fmt.Sprintf("%s/%s", prefix, accounts)
}

// /api/v1/accounts/:id
func Account(id int) string {
	return fmt.Sprintf("%s/%d", Accounts(), id)
}

// /api/v1/users/:user_id/courses
func CoursesForUser(userID int) string {
	return fmt.Sprintf("%s/%s", User(userID), courses)
}

// /api/v1/accounts/:account_id/courses
func CoursesForAccount(accountID int) string {
	return fmt.Sprintf("%s/%s", Account(accountID), courses)
}

// /api/v1/courses/:course_id/files
func FilesForCourse(courseID int) string {
	return fmt.Sprintf("%s/%s", Course(courseID), files)
}

// /api/v1/courses/:course_id/students
func StudentsForCourse(courseID int) string {
	return fmt.Sprintf("%s/%s", Course(courseID), students)
}

// /api/v1/courses/:course_id/users
func UsersForCourse(courseID int) string {
	return fmt.Sprintf("%s/%s", Course(courseID), users)
}

// api/v1/courses/:course_id/search_users
func SearchUsersForCourse(courseID int) string {
	return fmt.Sprintf("%s/%s", Course(courseID), searchUsers)
}

// /api/v1/courses/:course_id/recent_students
func RecentStudentsForCourse(courseID int) string {
	return fmt.Sprintf("%s/%s", Course(courseID), recentStudents)
}

// /api/v1/courses/:course_id/users/:id
func UserForCourse(courseID, userID int) string {
	return fmt.Sprintf("%s/%s", CourseUsers(courseID), userID)
}

// /api/v1/courses/:course_id/assignments
func AssignmentsForCourse(courseID int) string {
	return fmt.Sprintf("%s/%s", Course(courseID), assignments)
}

// /api/v1/users/:user_id/courses
func CoursesForUser(userID int) string {
	return fmt.Sprintf("%s/%s", User(userID), courses)
}

// /api/v1/users/:user_id/courses/:course_id
func CourseForUser(userID, courseID int) string {
	return fmt.Sprintf("%s/%d", CoursesForUser(userID), courseID)
}

// /api/v1/users/:user_id/courses/:course_id/assignments
func CourseAssignmentsForUser(userID, courseID int) string {
	return fmt.Sprintf("%s/%s", CourseForUser(userID, courseID), assignments)
}

// /api/v1/courses/:course_id/preview_html
func CoursePreviewHTML(courseID int) string {
	return fmt.Sprintf("%s/%s", Course(courseID), previewHTML)
}

//GET /api/v1/courses/:course_id/activity_stream
//GET /api/v1/courses/:course_id/activity_stream/summary
//GET /api/v1/courses/:course_id/todo
//DELETE /api/v1/courses/:id
//GET /api/v1/courses/:course_id/settings
//PUT /api/v1/courses/:course_id/settings
//GET /api/v1/accounts/:account_id/courses/:id
//PUT /api/v1/accounts/:account_id/courses
//POST /api/v1/courses/:course_id/reset_content
//GET /api/v1/courses/:course_id/effective_due_dates
//GET /api/v1/courses/:course_id/permissions
//GET /api/v1/courses/:course_id/course_copy/:id
//POST /api/v1/courses/:course_id/course_copy
