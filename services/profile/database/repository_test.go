package database

import (
	"context"
	jlMocks "in-backend/services/joblisting/tests/mocks"
	"in-backend/services/profile/configs"
	"in-backend/services/profile/interfaces"
	"in-backend/services/profile/models"
	"in-backend/services/profile/tests/mocks"
	"strings"
	"testing"
	"time"

	pg "github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx      context.Context                  = context.Background()
	now      time.Time                        = time.Now()
	a        *mocks.Auth0Provider             = &mocks.Auth0Provider{}
	k        *mocks.KlentyProvider            = &mocks.KlentyProvider{}
	jlClient *jlMocks.JoblistingServiceClient = &jlMocks.JoblistingServiceClient{}
)

func TestNewRepository(t *testing.T) {
	want := &repository{
		DB:       &pg.DB{},
		auth0:    a,
		klenty:   k,
		jlClient: jlClient,
	}

	got := NewRepository(&pg.DB{}, a, k, jlClient)

	require.EqualValues(t, want, got)
}

func TestAllCRUD(t *testing.T) {
	testConfig, err := configs.LoadConfig(configs.TestFileName)
	require.NoError(t, err)

	opt := GetPgConnectionOptions(testConfig)

	c, err := setupPGContainer(opt)
	require.NoError(t, err)

	db, err := setupDB(c, opt, "../scripts/migrations/")
	defer cleanDb(db)
	defer cleanContainer(c)
	require.NoError(t, err)

	r := NewRepository(db, a, k, jlClient)

	testCreateCandidate(t, r, db)
	testGetAllCandidates(t, r, db)
	testGetCandidateByID(t, r, db)
	testUpdateCandidate(t, r, db)
	testDeleteCandidate(t, r, db)

	testCreateSkill(t, r, db)
	testGetAllSkills(t, r, db)
	testGetSkill(t, r, db)

	testCreateUserSkill(t, r, db)
	testDeleteUserSkill(t, r, db)

	testCreateInstitution(t, r, db)
	testGetAllInstitutions(t, r, db)
	testGetInstitution(t, r, db)

	testCreateCourse(t, r, db)
	testGetAllCourses(t, r, db)
	testGetCourse(t, r, db)

	testCreateAcademicHistory(t, r, db)
	testGetAcademicHistory(t, r, db)
	testUpdateAcademicHistory(t, r, db)
	testDeleteAcademicHistory(t, r, db)

	testCreateCompany(t, r, db)
	testGetAllCompanies(t, r, db)
	testGetCompany(t, r, db)

	testCreateDepartment(t, r, db)
	testGetAllDepartments(t, r, db)
	testGetDepartment(t, r, db)

	testCreateJobHistory(t, r, db)
	testGetJobHistory(t, r, db)
	testUpdateJobHistory(t, r, db)
	testDeleteJobHistory(t, r, db)
}

/* --------------- User --------------- */

func testCreateUser(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoAuthID := &models.User{
		FirstName:     "first",
		LastName:      "last",
		Email:         "first@last.com",
		ContactNumber: "+6591234567",
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	test := *testNoAuthID
	test.AuthID = "authId"

	testDupEmail := test

	test2 := test
	test2.AuthID = "authId2"
	test2.Email = "test@test.com"
	test2.ContactNumber = "+6587654321"

	auth0GetOutput := make(map[string]interface{})
	auth0GetOutput["access_token"] = "test"
	auth0UpdateInput := "test"

	type args struct {
		ctx              context.Context
		input            *models.User
		auth0UpdateInput string
	}

	type expect struct {
		output            *models.User
		err               error
		auth0GetOutput    map[string]interface{}
		auth0UpdateOutput error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil, auth0UpdateInput}, expect{nil, errors.New("Input parameter user is nil"), auth0GetOutput, nil}},
		{"failed not null", args{ctx, testNoAuthID, auth0UpdateInput}, expect{nil, errors.New("Failed to insert user"), auth0GetOutput, nil}},
		{"valid", args{ctx, &test, auth0UpdateInput}, expect{&test, nil, auth0GetOutput, nil}},
		{"failed unique", args{ctx, &testDupEmail, auth0UpdateInput}, expect{nil, errors.New("Failed to insert user"), auth0GetOutput, nil}},
		{"invalid auth0", args{ctx, &test2, auth0UpdateInput}, expect{nil, errors.New("Failed to update auth0 user"), auth0GetOutput, errors.New("Failed to update auth0 user")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a.On("GetAuth0Token").Return(tt.exp.auth0GetOutput, nil)
			a.On("UpdateAuth0User", tt.args.auth0UpdateInput, tt.args.input).Return(tt.exp.auth0UpdateOutput)
			got, err := r.CreateUser(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Candidate --------------- */

func testCreateCandidate(t *testing.T, r interfaces.Repository, db *pg.DB) {
	test := &models.Candidate{
		Nationality:            "Singapore",
		ExpectedSalaryCurrency: "SGD",
		ExpectedSalary:         1000,
		PreferredRoles:         []string{"Frontend"},
		CreatedAt:              &now,
		UpdatedAt:              &now,
	}

	type args struct {
		ctx   context.Context
		tx    *pg.Tx
		input *models.Candidate
	}

	type expect struct {
		output *models.Candidate
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil, nil}, expect{nil, errors.New("Input parameter candidate is nil")}},
		{"valid", args{ctx, nil, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateCandidate(tt.args.ctx, tt.args.tx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCandidates(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Candidate)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.CandidateFilters
	}

	type expect struct {
		cnt int
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, &models.CandidateFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllCandidates(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCandidateByID(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Candidate{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Candidate
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Candidate{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetCandidateByID(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testUpdateCandidate(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Candidate{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.ResidenceCity = "Singapore"

	type args struct {
		ctx   context.Context
		input *models.Candidate
	}

	type expect struct {
		output *models.Candidate
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Candidate is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.Candidate{ID: 10000}}, expect{nil, errors.New("Cannot update candidate with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateCandidate(tt.args.ctx, nil, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteCandidate(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Candidate{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, existing.ID}, expect{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.DeleteCandidate(tt.args.ctx, tt.args.id)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Skill --------------- */

func testCreateSkill(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoName := &models.Skill{
		ID: 1,
	}

	test := &models.Skill{
		Name: "skill",
	}

	type args struct {
		ctx   context.Context
		input *models.Skill
	}

	type expect struct {
		output *models.Skill
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter skill is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert skill")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateSkill(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllSkills(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Skill)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.SkillFilters
	}

	type expect struct {
		cnt int
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, &models.SkillFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllSkills(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetSkill(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Skill{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Skill
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Skill{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetSkill(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- User Skill --------------- */

func testCreateUserSkill(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existingC := &models.Candidate{}
	err := db.WithContext(ctx).Model(existingC).First()
	require.NoError(t, err)

	existingS := &models.Skill{}
	err = db.WithContext(ctx).Model(existingS).First()
	require.NoError(t, err)

	testNoCID := &models.UserSkill{
		SkillID: existingS.ID,
	}

	test := &models.UserSkill{
		CandidateID: existingC.ID,
		SkillID:     existingS.ID,
	}

	type args struct {
		ctx   context.Context
		input *models.UserSkill
	}

	type expect struct {
		output *models.UserSkill
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter user skill is nil")}},
		{"failed no foreign key", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert user skill")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateUserSkill(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteUserSkill(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.UserSkill{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		cid uint64
		sid uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, existing.CandidateID, existing.SkillID}, expect{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.DeleteUserSkill(tt.args.ctx, tt.args.cid, tt.args.sid)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Institution --------------- */

func testCreateInstitution(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoName := &models.Institution{
		ID: 1,
	}

	test := &models.Institution{
		Name:    "institution",
		Country: "singapore",
	}

	type args struct {
		ctx   context.Context
		input *models.Institution
	}

	type expect struct {
		output *models.Institution
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter institution is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert institution")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateInstitution(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllInstitutions(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Institution)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.InstitutionFilters
	}

	type expect struct {
		cnt int
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, &models.InstitutionFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllInstitutions(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetInstitution(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Institution{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Institution
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Institution{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetInstitution(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Course --------------- */

func testCreateCourse(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoName := &models.Course{
		ID: 1,
	}

	test := &models.Course{
		Name:  "course",
		Level: "bachelor",
	}

	type args struct {
		ctx   context.Context
		input *models.Course
	}

	type expect struct {
		output *models.Course
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter course is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert course")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateCourse(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCourses(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Course)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.CourseFilters
	}

	type expect struct {
		cnt int
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, &models.CourseFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllCourses(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCourse(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Course{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Course
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Course{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetCourse(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- AcademicHistory --------------- */

func testCreateAcademicHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existingC := &models.Candidate{}
	err := db.WithContext(ctx).Model(existingC).First()
	require.NoError(t, err)

	existingI := &models.Institution{}
	err = db.WithContext(ctx).Model(existingI).First()
	require.NoError(t, err)

	existingCr := &models.Course{}
	err = db.WithContext(ctx).Model(existingCr).First()
	require.NoError(t, err)

	testNoCID := &models.AcademicHistory{
		InstitutionID: existingI.ID,
		CourseID:      existingCr.ID,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	testMissingFK := &models.AcademicHistory{
		CandidateID:   10000,
		InstitutionID: 10000,
		CourseID:      10000,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	test := &models.AcademicHistory{
		CandidateID:   existingC.ID,
		InstitutionID: existingI.ID,
		CourseID:      existingCr.ID,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	type args struct {
		ctx   context.Context
		input *models.AcademicHistory
	}

	type expect struct {
		output *models.AcademicHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter academic history is nil")}},
		{"failed not null", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert academic history")}},
		{"failed missing fk", args{ctx, testMissingFK}, expect{nil, errors.New("Failed to insert academic history")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateAcademicHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAcademicHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.AcademicHistory{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.AcademicHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.AcademicHistory{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAcademicHistory(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testUpdateAcademicHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.AcademicHistory{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.YearObtained = 2020

	type args struct {
		ctx   context.Context
		input *models.AcademicHistory
	}

	type expect struct {
		output *models.AcademicHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("AcademicHistory is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.AcademicHistory{ID: 10000}}, expect{nil, errors.New("Cannot update academic history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateAcademicHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteAcademicHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.AcademicHistory{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx  context.Context
		cid  uint64
		ahid uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, existing.CandidateID, existing.ID}, expect{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.DeleteAcademicHistory(tt.args.ctx, tt.args.cid, tt.args.ahid)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Company --------------- */

func testCreateCompany(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoName := &models.Company{
		ID: 1,
	}

	test := &models.Company{
		Name: "company",
	}

	type args struct {
		ctx   context.Context
		input *models.Company
	}

	type expect struct {
		output *models.Company
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter company is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert company")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateCompany(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllCompanies(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Company)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.CompanyFilters
	}

	type expect struct {
		cnt int
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, &models.CompanyFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllCompanies(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetCompany(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Company{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Company
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Company{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetCompany(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- Department --------------- */

func testCreateDepartment(t *testing.T, r interfaces.Repository, db *pg.DB) {
	testNoName := &models.Department{
		ID: 1,
	}

	test := &models.Department{
		Name: "department",
	}

	type args struct {
		ctx   context.Context
		input *models.Department
	}

	type expect struct {
		output *models.Department
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter department is nil")}},
		{"failed not null", args{ctx, testNoName}, expect{nil, errors.New("Failed to insert department")}},
		{"valid", args{ctx, test}, expect{test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateDepartment(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetAllDepartments(t *testing.T, r interfaces.Repository, db *pg.DB) {
	count, err := db.WithContext(ctx).Model((*models.Department)(nil)).Count()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		f   *models.DepartmentFilters
	}

	type expect struct {
		cnt int
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"no filter", args{ctx, &models.DepartmentFilters{}}, expect{count, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetAllDepartments(tt.args.ctx, *tt.args.f)
			assert.Equal(t, tt.exp.cnt, len(got))
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetDepartment(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.Department{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.Department
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.Department{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetDepartment(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

/* --------------- JobHistory --------------- */

func testCreateJobHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existingC := &models.Candidate{}
	err := db.WithContext(ctx).Model(existingC).First()
	require.NoError(t, err)

	existingCo := &models.Company{}
	err = db.WithContext(ctx).Model(existingCo).First()
	require.NoError(t, err)

	existingD := &models.Department{}
	err = db.WithContext(ctx).Model(existingD).First()
	require.NoError(t, err)

	start := time.Date(2020, 11, 10, 13, 0, 0, 0, time.Local)

	testNoCID := &models.JobHistory{
		CompanyID:    existingCo.ID,
		DepartmentID: existingD.ID,
		Country:      "singapore",
		Title:        "software engineer",
		StartDate:    &start,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}

	test := *testNoCID
	test.CandidateID = existingC.ID

	testMissingFK := test
	testMissingFK.CandidateID = 10000
	testMissingFK.CompanyID = 10000
	testMissingFK.DepartmentID = 10000

	type args struct {
		ctx   context.Context
		input *models.JobHistory
	}

	type expect struct {
		output *models.JobHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("Input parameter job history is nil")}},
		{"failed not null", args{ctx, testNoCID}, expect{nil, errors.New("Failed to insert job history")}},
		{"failed missing fk", args{ctx, &testMissingFK}, expect{nil, errors.New("Failed to insert job history")}},
		{"valid", args{ctx, &test}, expect{&test, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.CreateJobHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testGetJobHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.JobHistory{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx context.Context
		id  uint64
	}

	type expect struct {
		output *models.JobHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id exists", args{ctx, existing.ID}, expect{&models.JobHistory{ID: existing.ID}, nil}},
		{"id 10000", args{ctx, 10000}, expect{nil, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetJobHistory(tt.args.ctx, tt.args.id)
			if tt.exp.output != nil && got != nil {
				assert.Equal(t, tt.exp.output.ID, got.ID)
			} else {
				assert.Equal(t, tt.exp.output, got)
			}
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testUpdateJobHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.JobHistory{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	updated := *existing
	updated.Country = "indonesia"

	type args struct {
		ctx   context.Context
		input *models.JobHistory
	}

	type expect struct {
		output *models.JobHistory
		err    error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"nil", args{ctx, nil}, expect{nil, errors.New("JobHistory is nil")}},
		{"id existing", args{ctx, &updated}, expect{&updated, nil}},
		{"id 10000", args{ctx, &models.JobHistory{ID: 10000}}, expect{nil, errors.New("Cannot update job history with id")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UpdateJobHistory(tt.args.ctx, tt.args.input)
			assert.Condition(t, func() bool { return tt.exp.output.IsEqual(got) })
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}

func testDeleteJobHistory(t *testing.T, r interfaces.Repository, db *pg.DB) {
	existing := &models.JobHistory{}
	err := db.WithContext(ctx).Model(existing).First()
	require.NoError(t, err)

	type args struct {
		ctx  context.Context
		cid  uint64
		jhid uint64
	}

	type expect struct {
		err error
	}

	var tests = []struct {
		name string
		args args
		exp  expect
	}{
		{"id existing", args{ctx, existing.CandidateID, existing.ID}, expect{nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.DeleteJobHistory(tt.args.ctx, tt.args.cid, tt.args.jhid)
			if tt.exp.err != nil && err != nil {
				assert.Condition(t, func() bool { return strings.Contains(err.Error(), tt.exp.err.Error()) })
			} else {
				assert.Equal(t, tt.exp.err, err)
			}
		})
	}
}
