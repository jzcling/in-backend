package database

import (
	"context"
	"fmt"
	"strings"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"

	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
	"in-backend/services/assessment/providers"
)

// Repository implements the assessment Repository interface
type repository struct {
	DB          *pg.DB
	auth0       providers.Auth0Provider
	hubbedlearn providers.HubbedLearnProvider
	klenty      providers.KlentyProvider
}

// NewRepository declares a new Repository that implements assessment Repository
func NewRepository(db *pg.DB, a providers.Auth0Provider, hl providers.HubbedLearnProvider, k providers.KlentyProvider) interfaces.Repository {
	return &repository{
		DB:          db,
		auth0:       a,
		hubbedlearn: hl,
		klenty:      k,
	}
}

/* --------------- Candidate --------------- */

// CreateCandidate creates a new Candidate
func (r *repository) CreateCandidate(ctx context.Context, m *models.Candidate) (*models.Candidate, error) {
	if m == nil {
		return nil, errors.New("Input parameter candidate is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation(relCandidateSkill).
		Relation(relCandidateAcademic).Relation(relCandidateAcademicInstitution).Relation(relCandidateAcademicCourse).
		Relation(relCandidateJob).Relation(relCandidateJobCompany).Relation(relCandidateJobDepartment).
		Returning("*").
		Where("c.email = ?", m.Email).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert candidate %v", m)
		return nil, err
	}

	err = r.updateAuth0User(m)
	if err != nil {
		err = errors.Wrapf(err, "Failed to update auth0 user %v", m.AuthID)
		return nil, err
	}

	pw, err := r.createHubbedLearnUser(m)
	if err != nil {
		err = errors.Wrapf(err, "Failed to create Hubbed Learn user %v", m.Email)
		return nil, err
	}

	err = r.triggerCRMRegistrationWorkflow(m, *pw)
	if err != nil {
		err = errors.Wrapf(err, "Failed to trigger CRM workflow %v", m.Email)
		return nil, err
	}

	return m, nil
}

func (r *repository) updateAuth0User(m *models.Candidate) error {
	token, err := r.auth0.GetToken()
	if err != nil {
		return err
	}
	t := token["access_token"].(string)
	err = r.auth0.UpdateUser(t, m)
	if err != nil {
		return err
	}
	err = r.auth0.SetUserRole(t, m.AuthID, "Candidate")
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) createHubbedLearnUser(m *models.Candidate) (*string, error) {
	pw, err := r.hubbedlearn.CreateUser(m)
	if err != nil {
		return nil, err
	}
	return pw, nil
}

func (r *repository) triggerCRMRegistrationWorkflow(m *models.Candidate, pw string) error {
	err := r.klenty.CreateProspect(m, pw)
	if err != nil {
		return err
	}
	err = r.klenty.StartCadence(m.Email)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCandidates returns all Candidates
func (r *repository) GetAllCandidates(ctx context.Context, f models.CandidateFilters) ([]*models.Candidate, error) {
	var m []*models.Candidate
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("c.id in (?)", pg.In(f.ID))
	}
	if f.FirstName != "" {
		q = q.Where("lower(c.first_name) like ?", "%"+strings.ToLower(f.FirstName)+"%")
	}
	if f.LastName != "" {
		q = q.Where("lower(c.last_name) like ?", "%"+strings.ToLower(f.LastName)+"%")
	}
	if f.Email != "" {
		q = q.Where("lower(c.email) like ?", "%"+strings.ToLower(f.Email)+"%")
	}
	if f.ContactNumber != "" {
		q = q.Where("lower(c.contact_number) like ?", "%"+strings.ToLower(f.ContactNumber)+"%")
	}
	if len(f.Gender) > 0 {
		q = q.Where("c.gender in (?)", pg.In(f.Gender))
	}
	if len(f.Nationality) > 0 {
		q = q.Where("c.nationality in (?)", pg.In(f.Nationality))
	}
	if len(f.ResidenceCity) > 0 {
		q = q.Where("c.residence_city in (?)", pg.In(f.ResidenceCity))
	}
	if f.MinSalary > 0 {
		q = q.Where("c.expected_salary >= ?", f.MinSalary)
	}
	if f.MaxSalary > 0 {
		q = q.Where("c.expected_salary <= ?", f.MaxSalary)
	}
	if len(f.EducationLevel) > 0 {
		q = q.Where("c.education_level in (?)", pg.In(f.EducationLevel))
	}
	if f.MaxNoticePeriod > 0 {
		q = q.Where("c.notice_period <= ?", f.MaxNoticePeriod)
	}
	err := q.Relation(relCandidateSkill).
		Relation(relCandidateAcademic).Relation(relCandidateAcademicInstitution).Relation(relCandidateAcademicCourse).
		Relation(relCandidateJob).Relation(relCandidateJobCompany).Relation(relCandidateJobDepartment).
		Returning("*").
		Select()
	return m, err
}

// GetCandidateByID returns a Candidate by ID
func (r *repository) GetCandidateByID(ctx context.Context, id uint64) (*models.Candidate, error) {
	m := models.Candidate{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where(filCandidateID, id).
		Relation(relCandidateSkill).
		Relation(relCandidateAcademic).Relation(relCandidateAcademicInstitution).Relation(relCandidateAcademicCourse).
		Relation(relCandidateJob).Relation(relCandidateJobCompany).Relation(relCandidateJobDepartment).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateCandidate updates a Candidate
func (r *repository) UpdateCandidate(ctx context.Context, m *models.Candidate) (*models.Candidate, error) {
	if m == nil {
		return nil, errors.New("Candidate is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relCandidateSkill).
		Relation(relCandidateAcademic).Relation(relCandidateAcademicInstitution).Relation(relCandidateAcademicCourse).
		Relation(relCandidateJob).Relation(relCandidateJobCompany).Relation(relCandidateJobDepartment).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update candidate with id %v", m.ID))
	}

	return m, nil
}

// DeleteCandidate deletes a Candidate by ID
func (r *repository) DeleteCandidate(ctx context.Context, id uint64) error {
	m := &models.Candidate{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete candidate with id %v", id))
	}
	return nil
}

/* --------------- Skill --------------- */

// CreateSkill creates a new Skill
func (r *repository) CreateSkill(ctx context.Context, m *models.Skill) (*models.Skill, error) {
	if m == nil {
		return nil, errors.New("Input parameter skill is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Returning("*").
		Where(filNameEquals, m.Name).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert skill %v", m)
	}

	return m, nil
}

// GetSkill returns a Skill by ID
func (r *repository) GetSkill(ctx context.Context, id uint64) (*models.Skill, error) {
	m := models.Skill{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).Where(filSkillID, id).First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// GetAllSkills returns all Skills
func (r *repository) GetAllSkills(ctx context.Context, f models.SkillFilters) ([]*models.Skill, error) {
	var m []*models.Skill
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.Name) > 0 {
		q = q.Where(filNameIn, pg.In(f.Name))
	}
	err := q.Select()
	return m, err
}

/* --------------- User Skill --------------- */

// CreateUserSkill creates a new UserSkill
func (r *repository) CreateUserSkill(ctx context.Context, us *models.UserSkill) (*models.UserSkill, error) {
	if us == nil {
		return nil, errors.New("Input parameter user skill is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(us).Returning("*").
		Where("candidate_id = ?", us.CandidateID).
		Where("skill_id = ?", us.SkillID).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert user skill %v", us)
	}

	return us, nil
}

// DeleteUserSkill deletes a UserSkill by ID
func (r *repository) DeleteUserSkill(ctx context.Context, id uint64) error {
	us := &models.UserSkill{ID: id}
	_, err := r.DB.WithContext(ctx).Model(us).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete user skill with id %v", id))
	}
	return nil
}

/* --------------- Institution --------------- */

// CreateInstitution creates a new Institution
func (r *repository) CreateInstitution(ctx context.Context, m *models.Institution) (*models.Institution, error) {
	if m == nil {
		return nil, errors.New("Input parameter institution is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Returning("*").
		Where(filNameEquals, m.Name).
		Where(filCountryEquals, m.Country).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert institution %v", m)
	}

	return m, nil
}

// GetInstitution returns a Institution by ID
func (r *repository) GetInstitution(ctx context.Context, id uint64) (*models.Institution, error) {
	m := models.Institution{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).Where(filInstitutionID, id).First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// GetAllInstitutions returns all Institutions
func (r *repository) GetAllInstitutions(ctx context.Context, f models.InstitutionFilters) ([]*models.Institution, error) {
	var m []*models.Institution
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.Name) > 0 {
		q = q.Where(filNameIn, pg.In(f.Name))
	}
	if len(f.Country) > 0 {
		q = q.Where(filCountryIn, pg.In(f.Country))
	}
	err := q.Select()
	return m, err
}

/* --------------- Course --------------- */

// CreateCourse creates a new Course
func (r *repository) CreateCourse(ctx context.Context, m *models.Course) (*models.Course, error) {
	if m == nil {
		return nil, errors.New("Input parameter course is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Returning("*").
		Where(filNameEquals, m.Name).
		Where(filLevelEquals, m.Level).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert course %v", m)
	}

	return m, nil
}

// GetCourse returns a Course by ID
func (r *repository) GetCourse(ctx context.Context, id uint64) (*models.Course, error) {
	m := models.Course{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).Where(filCourseID, id).First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// GetAllCourses returns all Courses
func (r *repository) GetAllCourses(ctx context.Context, f models.CourseFilters) ([]*models.Course, error) {
	var m []*models.Course
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.Name) > 0 {
		q = q.Where(filNameIn, pg.In(f.Name))
	}
	if len(f.Level) > 0 {
		q = q.Where(filLevelIn, pg.In(f.Level))
	}
	err := q.Select()
	return m, err
}

/* --------------- Academic History --------------- */

// CreateAcademicHistory creates a new AcademicHistory
func (r *repository) CreateAcademicHistory(ctx context.Context, m *models.AcademicHistory) (*models.AcademicHistory, error) {
	if m == nil {
		return nil, errors.New("Input parameter academic history is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation("Institution").
		Relation("Course").
		Returning("*").
		Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert academic history %v", m)
	}

	return m, nil
}

// GetAcademicHistory returns a AcademicHistory by ID
func (r *repository) GetAcademicHistory(ctx context.Context, id uint64) (*models.AcademicHistory, error) {
	m := models.AcademicHistory{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where(filAcademicID, id).
		Relation("Institution").
		Relation("Course").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateAcademicHistory updates a AcademicHistory
func (r *repository) UpdateAcademicHistory(ctx context.Context, m *models.AcademicHistory) (*models.AcademicHistory, error) {
	if m == nil {
		return nil, errors.New("AcademicHistory is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation("Institution").
		Relation("Course").
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update academic history with id %v", m.ID))
	}

	return m, nil
}

// DeleteAcademicHistory deletes a AcademicHistory by ID
func (r *repository) DeleteAcademicHistory(ctx context.Context, id uint64) error {
	m := &models.AcademicHistory{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete academic history with id %v", id))
	}
	return nil
}

/* --------------- Company --------------- */

// CreateCompany creates a new Company
func (r *repository) CreateCompany(ctx context.Context, m *models.Company) (*models.Company, error) {
	if m == nil {
		return nil, errors.New("Input parameter company is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Returning("*").
		Where(filNameEquals, m.Name).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert company %v", m)
	}

	return m, nil
}

// GetCompany returns a Company by ID
func (r *repository) GetCompany(ctx context.Context, id uint64) (*models.Company, error) {
	m := models.Company{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).Where(filCompanyID, id).First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// GetAllCompanies returns all Companies
func (r *repository) GetAllCompanies(ctx context.Context, f models.CompanyFilters) ([]*models.Company, error) {
	var m []*models.Company
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.Name) > 0 {
		q = q.Where(filNameIn, pg.In(f.Name))
	}
	err := q.Select()
	return m, err
}

/* --------------- Department --------------- */

// CreateDepartment creates a new Department
func (r *repository) CreateDepartment(ctx context.Context, m *models.Department) (*models.Department, error) {
	if m == nil {
		return nil, errors.New("Input parameter department is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).Returning("*").
		Where(filNameEquals, m.Name).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert department %v", m)
	}

	return m, nil
}

// GetDepartment returns a Department by ID
func (r *repository) GetDepartment(ctx context.Context, id uint64) (*models.Department, error) {
	m := models.Department{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).Where(filDepartmentID, id).First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// GetAllDepartments returns all Departments
func (r *repository) GetAllDepartments(ctx context.Context, f models.DepartmentFilters) ([]*models.Department, error) {
	var m []*models.Department
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.Name) > 0 {
		q = q.Where(filNameIn, pg.In(f.Name))
	}
	err := q.Select()
	return m, err
}

/* --------------- Job History --------------- */

// CreateJobHistory creates a new JobHistory
func (r *repository) CreateJobHistory(ctx context.Context, m *models.JobHistory) (*models.JobHistory, error) {
	if m == nil {
		return nil, errors.New("Input parameter job history is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation("Company").
		Relation("Department").
		Returning("*").
		Insert()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to insert job history %v", m)
	}

	return m, nil
}

// GetJobHistory returns a JobHistory by ID
func (r *repository) GetJobHistory(ctx context.Context, id uint64) (*models.JobHistory, error) {
	m := models.JobHistory{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).Where(filJobID, id).
		Relation("Company").
		Relation("Department").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateJobHistory updates a JobHistory
func (r *repository) UpdateJobHistory(ctx context.Context, m *models.JobHistory) (*models.JobHistory, error) {
	if m == nil {
		return nil, errors.New("JobHistory is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation("Company").
		Relation("Department").
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update job history with id %v", m.ID))
	}

	return m, nil
}

// DeleteJobHistory deletes a JobHistory by ID
func (r *repository) DeleteJobHistory(ctx context.Context, id uint64) error {
	m := &models.JobHistory{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete job history with id %v", id))
	}
	return nil
}