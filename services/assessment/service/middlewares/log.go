package middlewares

import (
	"context"
	"fmt"
	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"time"

	"github.com/go-kit/kit/log"
)

type logMiddleware struct {
	logger log.Logger
	next   interfaces.Service
}

// NewLogMiddleware creates and returns a new Log Middleware that implements the assessment Service interface
func NewLogMiddleware(logger log.Logger, svc interfaces.Service) interfaces.Service {
	return &logMiddleware{
		logger: logger,
		next:   svc,
	}
}

func (mw logMiddleware) log(method string, begin time.Time, input, output interface{}, err *error) {
	mw.logger.Log(
		"method", method,
		"input", fmt.Sprintf("%v", input),
		"output", fmt.Sprintf("%v", output),
		"err", err,
		"took", time.Since(begin),
	)
}

/* --------------- Candidate --------------- */

// CreateCandidate creates a new Candidate
func (mw logMiddleware) CreateCandidate(ctx context.Context, input *models.Candidate) (output *models.Candidate, err error) {
	defer mw.log("CreateCandidate", time.Now(), input, &output, &err)
	output, err = mw.next.CreateCandidate(ctx, input)
	return
}

// GetAllCandidates returns all Candidates
func (mw logMiddleware) GetAllCandidates(ctx context.Context, input models.CandidateFilters) (output []*models.Candidate, err error) {
	defer mw.log("GetAllCandidates", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllCandidates(ctx, input)
	return
}

// GetCandidateByID returns a Candidate by ID
func (mw logMiddleware) GetCandidateByID(ctx context.Context, input uint64) (output *models.Candidate, err error) {
	defer mw.log("GetCandidateByID", time.Now(), input, &output, &err)
	output, err = mw.next.GetCandidateByID(ctx, input)
	return
}

// UpdateCandidate updates a Candidate
func (mw logMiddleware) UpdateCandidate(ctx context.Context, input *models.Candidate) (output *models.Candidate, err error) {
	defer mw.log("UpdateCandidate", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateCandidate(ctx, input)
	return
}

// DeleteCandidate deletes a Candidate by ID
func (mw logMiddleware) DeleteCandidate(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteCandidate", time.Now(), input, nil, &err)
	err = mw.next.DeleteCandidate(ctx, input)
	return
}

/* --------------- Skill --------------- */

// CreateSkill creates a new Skill
func (mw logMiddleware) CreateSkill(ctx context.Context, input *models.Skill) (output *models.Skill, err error) {
	defer mw.log("CreateSkill", time.Now(), input, &output, &err)
	output, err = mw.next.CreateSkill(ctx, input)
	return
}

// GetSkill returns a Skill by ID
func (mw logMiddleware) GetSkill(ctx context.Context, input uint64) (output *models.Skill, err error) {
	defer mw.log("GetSkill", time.Now(), input, &output, &err)
	output, err = mw.next.GetSkill(ctx, input)
	return
}

// GetAllSkills returns all Skills
func (mw logMiddleware) GetAllSkills(ctx context.Context, input models.SkillFilters) (output []*models.Skill, err error) {
	defer mw.log("GetAllSkills", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllSkills(ctx, input)
	return
}

/* --------------- User Skill --------------- */

// CreateUserSkill creates a new UserSkill
func (mw logMiddleware) CreateUserSkill(ctx context.Context, input *models.UserSkill) (output *models.UserSkill, err error) {
	defer mw.log("CreateUserSkill", time.Now(), input, &output, &err)
	output, err = mw.next.CreateUserSkill(ctx, input)
	return
}

// DeleteUserSkill deletes a UserSkill by ID
func (mw logMiddleware) DeleteUserSkill(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteUserSkill", time.Now(), input, nil, &err)
	err = mw.next.DeleteUserSkill(ctx, input)
	return
}

/* --------------- Institution --------------- */

// CreateInstitution creates a new Institution
func (mw logMiddleware) CreateInstitution(ctx context.Context, input *models.Institution) (output *models.Institution, err error) {
	defer mw.log("CreateInstitution", time.Now(), input, &output, &err)
	output, err = mw.next.CreateInstitution(ctx, input)
	return
}

// GetInstitution returns a Institution by ID
func (mw logMiddleware) GetInstitution(ctx context.Context, input uint64) (output *models.Institution, err error) {
	defer mw.log("GetInstitution", time.Now(), input, &output, &err)
	output, err = mw.next.GetInstitution(ctx, input)
	return
}

// GetAllInstitutions returns all Institutions
func (mw logMiddleware) GetAllInstitutions(ctx context.Context, input models.InstitutionFilters) (output []*models.Institution, err error) {
	defer mw.log("GetAllInstitutions", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllInstitutions(ctx, input)
	return
}

/* --------------- Course --------------- */

// CreateCourse creates a new Course
func (mw logMiddleware) CreateCourse(ctx context.Context, input *models.Course) (output *models.Course, err error) {
	defer mw.log("CreateCourse", time.Now(), input, &output, &err)
	output, err = mw.next.CreateCourse(ctx, input)
	return
}

// GetCourse returns a Course by ID
func (mw logMiddleware) GetCourse(ctx context.Context, input uint64) (output *models.Course, err error) {
	defer mw.log("GetCourse", time.Now(), input, &output, &err)
	output, err = mw.next.GetCourse(ctx, input)
	return
}

// GetAllCourses returns all Courses
func (mw logMiddleware) GetAllCourses(ctx context.Context, input models.CourseFilters) (output []*models.Course, err error) {
	defer mw.log("GetAllCourses", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllCourses(ctx, input)
	return
}

/* --------------- Academic History --------------- */

// CreateAcademicHistory creates a new AcademicHistory
func (mw logMiddleware) CreateAcademicHistory(ctx context.Context, input *models.AcademicHistory) (output *models.AcademicHistory, err error) {
	defer mw.log("CreateAcademicHistory", time.Now(), input, &output, &err)
	output, err = mw.next.CreateAcademicHistory(ctx, input)
	return
}

// GetAcademicHistory returns a AcademicHistory by ID
func (mw logMiddleware) GetAcademicHistory(ctx context.Context, input uint64) (output *models.AcademicHistory, err error) {
	defer mw.log("GetAcademicHistory", time.Now(), input, &output, &err)
	output, err = mw.next.GetAcademicHistory(ctx, input)
	return
}

// UpdateAcademicHistory updates a AcademicHistory
func (mw logMiddleware) UpdateAcademicHistory(ctx context.Context, input *models.AcademicHistory) (output *models.AcademicHistory, err error) {
	defer mw.log("UpdateAcademicHistory", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateAcademicHistory(ctx, input)
	return
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (mw logMiddleware) DeleteAcademicHistory(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteAcademicHistory", time.Now(), input, nil, &err)
	err = mw.next.DeleteAcademicHistory(ctx, input)
	return
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (mw logMiddleware) CreateCompany(ctx context.Context, input *models.Company) (output *models.Company, err error) {
	defer mw.log("CreateCompany", time.Now(), input, &output, &err)
	output, err = mw.next.CreateCompany(ctx, input)
	return
}

// GetCompany returns a Company by ID
func (mw logMiddleware) GetCompany(ctx context.Context, input uint64) (output *models.Company, err error) {
	defer mw.log("GetCompany", time.Now(), input, &output, &err)
	output, err = mw.next.GetCompany(ctx, input)
	return
}

// GetAllCompanies returns all Companies
func (mw logMiddleware) GetAllCompanies(ctx context.Context, input models.CompanyFilters) (output []*models.Company, err error) {
	defer mw.log("GetAllCompanies", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllCompanies(ctx, input)
	return
}

/* --------------- Department --------------- */

// CreateDepartment creates a new Department
func (mw logMiddleware) CreateDepartment(ctx context.Context, input *models.Department) (output *models.Department, err error) {
	defer mw.log("CreateDepartment", time.Now(), input, &output, &err)
	output, err = mw.next.CreateDepartment(ctx, input)
	return
}

// GetDepartment returns a Department by ID
func (mw logMiddleware) GetDepartment(ctx context.Context, input uint64) (output *models.Department, err error) {
	defer mw.log("GetDepartment", time.Now(), input, &output, &err)
	output, err = mw.next.GetDepartment(ctx, input)
	return
}

// GetAllDepartments returns all Departments
func (mw logMiddleware) GetAllDepartments(ctx context.Context, input models.DepartmentFilters) (output []*models.Department, err error) {
	defer mw.log("GetAllDepartments", time.Now(), input, &output, &err)
	output, err = mw.next.GetAllDepartments(ctx, input)
	return
}

/* --------------- Job History --------------- */

// CreateJobHistory creates a new JobHistory
func (mw logMiddleware) CreateJobHistory(ctx context.Context, input *models.JobHistory) (output *models.JobHistory, err error) {
	defer mw.log("CreateJobHistory", time.Now(), input, &output, &err)
	output, err = mw.next.CreateJobHistory(ctx, input)
	return
}

// GetJobHistory returns a JobHistory by ID
func (mw logMiddleware) GetJobHistory(ctx context.Context, input uint64) (output *models.JobHistory, err error) {
	defer mw.log("GetJobHistory", time.Now(), input, &output, &err)
	output, err = mw.next.GetJobHistory(ctx, input)
	return
}

// UpdateJobHistory updates a JobHistory
func (mw logMiddleware) UpdateJobHistory(ctx context.Context, input *models.JobHistory) (output *models.JobHistory, err error) {
	defer mw.log("UpdateJobHistory", time.Now(), input, &output, &err)
	output, err = mw.next.UpdateJobHistory(ctx, input)
	return
}

// DeleteJobHistory deletes a JobHistory by ID
func (mw logMiddleware) DeleteJobHistory(ctx context.Context, input uint64) (err error) {
	defer mw.log("DeleteJobHistory", time.Now(), input, nil, &err)
	err = mw.next.DeleteJobHistory(ctx, input)
	return
}