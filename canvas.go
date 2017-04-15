package canvas

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rvolosatovs/canvas-go/internal/endpoints"

	"golang.org/x/oauth2"
)

var ErrUnauthorized = errors.New("Unauthorized")

type Client struct {
	client *http.Client
	addr   string
}

func OAuthEndpoint(addr string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  addr + "/" + endpoints.AuthURL(),
		TokenURL: addr + "/" + endpoints.TokenURL(),
	}
}

func NewClient(ctx context.Context, addr string, src oauth2.TokenSource) *Client {
	return &Client{
		client: oauth2.NewClient(ctx, src),
		addr:   addr,
	}
}

type Error struct {
	Status string `json:"status"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func processResponse(resp *http.Response, respVal interface{}) error {
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return dec.Decode(respVal)
	case http.StatusUnauthorized:
		respErr := &Error{}
		if err := dec.Decode(respErr); err != nil {
			return err
		}
		return fmt.Errorf("Unauthorized: %s", respErr.Errors)
	default:
		return fmt.Errorf("Error %s", resp.StatusCode)
	}
}

func (cl *Client) get(endpoint string, respVal interface{}) error {
	resp, err := cl.client.Get(cl.addr + "/" + endpoint)
	if err != nil {
		return err
	}
	return processResponse(resp, respVal)
}
func (cl *Client) postJSON(endpoint string, v interface{}, respVal interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	resp, err := cl.client.Post(cl.addr+"/"+endpoint, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	return processResponse(resp, respVal)
}
func (cl *Client) postForm(endpoint string, data url.Values, respVal interface{}) error {
	resp, err := cl.client.PostForm(cl.addr+"/"+endpoint, data)
	if err != nil {
		return err
	}
	return processResponse(resp, respVal)

}

type Course struct {
	ID                          int         `json:"id"`
	SisCourseID                 interface{} `json:"sis_course_id"`
	IntegrationID               interface{} `json:"integration_id"`
	Name                        string      `json:"name"`
	CourseCode                  string      `json:"course_code"`
	WorkflowState               string      `json:"workflow_state"`
	AccountID                   int         `json:"account_id"`
	RootAccountID               int         `json:"root_account_id"`
	EnrollmentTermID            int         `json:"enrollment_term_id"`
	GradingStandardID           int         `json:"grading_standard_id"`
	StartAt                     string      `json:"start_at"`
	EndAt                       string      `json:"end_at"`
	Locale                      string      `json:"locale"`
	Enrollments                 interface{} `json:"enrollments"`
	TotalStudents               int         `json:"total_students"`
	Calendar                    interface{} `json:"calendar"`
	DefaultView                 string      `json:"default_view"`
	SyllabusBody                string      `json:"syllabus_body"`
	NeedsGradingCount           int         `json:"needs_grading_count"`
	Term                        interface{} `json:"term"`
	CourseProgress              interface{} `json:"course_progress"`
	ApplyAssignmentGroupWeights bool        `json:"apply_assignment_group_weights"`
	Permissions                 struct {
		CreateDiscussionTopic bool `json:"create_discussion_topic"`
		CreateAnnouncement    bool `json:"create_announcement"`
	} `json:"permissions"`
	IsPublic                         bool   `json:"is_public"`
	IsPublicToAuthUsers              bool   `json:"is_public_to_auth_users"`
	PublicSyllabus                   bool   `json:"public_syllabus"`
	PublicSyllabusToAuth             bool   `json:"public_syllabus_to_auth"`
	PublicDescription                string `json:"public_description"`
	StorageQuotaMb                   int    `json:"storage_quota_mb"`
	StorageQuotaUsedMb               int    `json:"storage_quota_used_mb"`
	HideFinalGrades                  bool   `json:"hide_final_grades"`
	License                          string `json:"license"`
	AllowStudentAssignmentEdits      bool   `json:"allow_student_assignment_edits"`
	AllowWikiComments                bool   `json:"allow_wiki_comments"`
	AllowStudentForumAttachments     bool   `json:"allow_student_forum_attachments"`
	OpenEnrollment                   bool   `json:"open_enrollment"`
	SelfEnrollment                   bool   `json:"self_enrollment"`
	RestrictEnrollmentsToCourseDates bool   `json:"restrict_enrollments_to_course_dates"`
	CourseFormat                     string `json:"course_format"`
	AccessRestrictedByDate           bool   `json:"access_restricted_by_date"`
	TimeZone                         string `json:"time_zone"`
}

func (cl *Client) Courses() ([]Course, error) {
	courses := []Course{}
	return courses, cl.get(endpoints.Courses(), &courses)
}

type User struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	SortableName  string      `json:"sortable_name"`
	ShortName     string      `json:"short_name"`
	SisUserID     string      `json:"sis_user_id"`
	SisImportID   int         `json:"sis_import_id"`
	SisLoginID    interface{} `json:"sis_login_id"`
	IntegrationID string      `json:"integration_id"`
	LoginID       string      `json:"login_id"`
	Enrollments   interface{} `json:"enrollments"`
	Email         string      `json:"email"`
	Locale        string      `json:"locale"`
	LastLogin     time.Time   `json:"last_login"`
	TimeZone      string      `json:"time_zone"`
	Bio           string      `json:"bio"`
}

func (cl *Client) User(id int) (*User, error) {
	user := &User{}
	return user, cl.get(endpoints.User(id), user)
}

type Assignment struct {
	ID                             int         `json:"id"`
	Name                           string      `json:"name"`
	Description                    string      `json:"description"`
	CreatedAt                      string      `json:"created_at"`
	UpdatedAt                      string      `json:"updated_at"`
	DueAt                          string      `json:"due_at"`
	LockAt                         string      `json:"lock_at"`
	UnlockAt                       string      `json:"unlock_at"`
	HasOverrides                   bool        `json:"has_overrides"`
	AllDates                       interface{} `json:"all_dates"`
	CourseID                       int         `json:"course_id"`
	HTMLURL                        string      `json:"html_url"`
	SubmissionsDownloadURL         string      `json:"submissions_download_url"`
	AssignmentGroupID              int         `json:"assignment_group_id"`
	DueDateRequired                bool        `json:"due_date_required"`
	AllowedExtensions              []string    `json:"allowed_extensions"`
	MaxNameLength                  int         `json:"max_name_length"`
	TurnitinEnabled                bool        `json:"turnitin_enabled"`
	VericiteEnabled                bool        `json:"vericite_enabled"`
	TurnitinSettings               interface{} `json:"turnitin_settings"`
	GradeGroupStudentsIndividually bool        `json:"grade_group_students_individually"`
	ExternalToolTagAttributes      interface{} `json:"external_tool_tag_attributes"`
	PeerReviews                    bool        `json:"peer_reviews"`
	AutomaticPeerReviews           bool        `json:"automatic_peer_reviews"`
	PeerReviewCount                int         `json:"peer_review_count"`
	PeerReviewsAssignAt            string      `json:"peer_reviews_assign_at"`
	IntraGroupPeerReviews          bool        `json:"intra_group_peer_reviews"`
	GroupCategoryID                int         `json:"group_category_id"`
	NeedsGradingCount              int         `json:"needs_grading_count"`
	NeedsGradingCountBySection     []struct {
		SectionID         string `json:"section_id"`
		NeedsGradingCount int    `json:"needs_grading_count"`
	} `json:"needs_grading_count_by_section"`
	Position               int         `json:"position"`
	PostToSis              bool        `json:"post_to_sis"`
	IntegrationID          string      `json:"integration_id"`
	IntegrationData        string      `json:"integration_data"`
	Muted                  interface{} `json:"muted"`
	PointsPossible         int         `json:"points_possible"`
	SubmissionTypes        []string    `json:"submission_types"`
	GradingType            string      `json:"grading_type"`
	GradingStandardID      interface{} `json:"grading_standard_id"`
	Published              bool        `json:"published"`
	Unpublishable          bool        `json:"unpublishable"`
	OnlyVisibleToOverrides bool        `json:"only_visible_to_overrides"`
	LockedForUser          bool        `json:"locked_for_user"`
	LockInfo               interface{} `json:"lock_info"`
	LockExplanation        string      `json:"lock_explanation"`
	QuizID                 int         `json:"quiz_id"`
	AnonymousSubmissions   bool        `json:"anonymous_submissions"`
	DiscussionTopic        interface{} `json:"discussion_topic"`
	FreezeOnCopy           bool        `json:"freeze_on_copy"`
	Frozen                 bool        `json:"frozen"`
	FrozenAttributes       []string    `json:"frozen_attributes"`
	Submission             interface{} `json:"submission"`
	UseRubricForGrading    bool        `json:"use_rubric_for_grading"`
	RubricSettings         struct {
		PointsPossible string `json:"points_possible"`
	} `json:"rubric_settings"`
	Rubric               interface{} `json:"rubric"`
	AssignmentVisibility []int       `json:"assignment_visibility"`
	Overrides            interface{} `json:"overrides"`
	OmitFromFinalGrade   bool        `json:"omit_from_final_grade"`
}

func (cl *Client) Assignments(courseID int) ([]Assignment, error) {
	assignments := []Assignment{}
	return assignments, cl.get(endpoints.CourseAssignments(courseID), &assignments)
}
