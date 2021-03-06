package models

// CandidateFilters define filters for Candidate model
type CandidateFilters struct {
	ID              []uint64
	FirstName       string
	LastName        string
	Email           string
	ContactNumber   string
	Gender          []string
	Nationality     []string
	ResidenceCity   []string
	MinSalary       uint32
	MaxSalary       uint32
	EducationLevel  []string
	MaxNoticePeriod uint32
	Roles           []string
}

// SkillFilters define filters for Skill model
type SkillFilters struct {
	ID   []uint64
	Name []string
}

// InstitutionFilters define filters for Institution model
type InstitutionFilters struct {
	Name    []string
	Country []string
}

// CourseFilters define filters for Course model
type CourseFilters struct {
	Name  []string
	Level []string
}

// AcademicHistoryFilters define filters for AcademicHistory model
type AcademicHistoryFilters struct {
	CandidateID   []uint64
	InstitutionID []uint64
	CourseID      []uint64
	YearObtained  []uint32
	Grade         []string
}

// CompanyFilters define filters for Company model
type CompanyFilters struct {
	Name []string
}

// DepartmentFilters define filters for Department model
type DepartmentFilters struct {
	Name []string
}

// JobHistoryFilters define filters for JobHistory model
type JobHistoryFilters struct {
	CandidateID  []uint64
	CompanyID    []uint64
	DepartmentID []uint64
	Country      []string
	City         []string
	Title        []string
}
