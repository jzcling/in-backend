package interfaces

import (
	"context"
	"in-backend/services/profile/models"

	"github.com/go-pg/pg/v10"
)

// Repository declares the repository for candidate profiles
type Repository interface {
	/* --------------- User --------------- */

	// CreateUser creates a new User
	CreateUser(ctx context.Context, m *models.User) (*models.User, error)

	// UpdateUser updates a User
	UpdateUser(ctx context.Context, c *models.User) (*models.User, error)

	// GetUserByID finds and returns a user by ID
	GetUserByID(ctx context.Context, id uint64) (*models.User, error)

	// GetUserByEmail finds and returns a user by email
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// DeleteUser deletes a User by ID
	DeleteUser(ctx context.Context, id uint64) error

	/* --------------- Candidate --------------- */

	// CreateCandidate creates a new candidate
	CreateCandidate(ctx context.Context, tx *pg.Tx, c *models.Candidate) (*models.Candidate, error)

	// GetAllCandidates returns all candidates
	GetAllCandidates(ctx context.Context, f models.CandidateFilters) ([]*models.User, error)

	// GetCandidateByID finds and returns a candidate by ID
	GetCandidateByID(ctx context.Context, id uint64) (*models.User, error)

	// UpdateCandidate updates a candidate
	UpdateCandidate(ctx context.Context, tx *pg.Tx, c *models.Candidate) (*models.Candidate, error)

	// DeleteCandidate deletes a candidate by ID
	DeleteCandidate(ctx context.Context, id uint64) error

	/* --------------- Skill --------------- */

	// CreateSkill creates a new Skill
	CreateSkill(ctx context.Context, s *models.Skill) (*models.Skill, error)

	// GetSkill returns a Skill by ID
	GetSkill(ctx context.Context, id uint64) (*models.Skill, error)

	// GetAllSkills returns all Skills
	GetAllSkills(ctx context.Context, f models.SkillFilters) ([]*models.Skill, error)

	/* --------------- User Skill --------------- */

	// CreateUserSkill creates a new UserSkill
	CreateUserSkill(ctx context.Context, us *models.UserSkill) (*models.UserSkill, error)

	// DeleteUserSkill deletes a UserSkill by ID
	DeleteUserSkill(ctx context.Context, cid, sid uint64) error

	/* --------------- Institution --------------- */

	// CreateInstitution creates a new Institution
	CreateInstitution(ctx context.Context, i *models.Institution) (*models.Institution, error)

	// GetInstitution returns a Institution by ID
	GetInstitution(ctx context.Context, id uint64) (*models.Institution, error)

	// GetAllInstitutions returns all Institutions
	GetAllInstitutions(ctx context.Context, f models.InstitutionFilters) ([]*models.Institution, error)

	/* --------------- Course --------------- */

	// CreateCourse creates a new Course
	CreateCourse(ctx context.Context, c *models.Course) (*models.Course, error)

	// GetCourse returns a Course by ID
	GetCourse(ctx context.Context, id uint64) (*models.Course, error)

	// GetAllCourses returns all Courses
	GetAllCourses(ctx context.Context, f models.CourseFilters) ([]*models.Course, error)

	/* --------------- Academic History --------------- */

	// CreateAcademicHistory creates a new AcademicHistory
	CreateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error)

	// GetAcademicHistory returns a AcademicHistory by ID
	GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error)

	// UpdateAcademicHistory updates a AcademicHistory
	UpdateAcademicHistory(ctx context.Context, a *models.AcademicHistory) (*models.AcademicHistory, error)

	// DeleteAcademicHistory deletes a AcademicHistory by ID
	DeleteAcademicHistory(ctx context.Context, cid, ahid uint64) error

	/* --------------- Company --------------- */

	// CreateCompany creates a new Company
	CreateCompany(ctx context.Context, c *models.Company) (*models.Company, error)

	// GetCompany returns a Company by ID
	GetCompany(ctx context.Context, id uint64) (*models.Company, error)

	// GetAllCompanies returns all Companies
	GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error)

	/* --------------- Department --------------- */

	// CreateDepartment creates a new Department
	CreateDepartment(ctx context.Context, c *models.Department) (*models.Department, error)

	// GetDepartment returns a Department by ID
	GetDepartment(ctx context.Context, id uint64) (*models.Department, error)

	// GetAllDepartments returns all Departments
	GetAllDepartments(ctx context.Context, f models.DepartmentFilters) ([]*models.Department, error)

	/* --------------- Job History --------------- */

	// CreateJobHistory creates a new JobHistory
	CreateJobHistory(ctx context.Context, a *models.JobHistory) (*models.JobHistory, error)

	// GetJobHistory returns a JobHistory by ID
	GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error)

	// UpdateJobHistory updates a JobHistory
	UpdateJobHistory(ctx context.Context, j *models.JobHistory) (*models.JobHistory, error)

	// DeleteJobHistory deletes a JobHistory by ID
	DeleteJobHistory(ctx context.Context, cid, jhid uint64) error
}
