package models

import (
	"in-backend/services/profile/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestUserToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.User{
		Id:            1,
		AuthId:        "authId",
		FirstName:     "first",
		LastName:      "last",
		Email:         "email",
		ContactNumber: "contact",
		Picture:       "picture",
		Gender:        "male",
		Roles:         []string{"Candidate", "Company"},
		CandidateId:   1,
		JobCompanyId:  1,
		Candidate: &pb.Candidate{
			Nationality:            "singapore",
			ResidenceCity:          "singapore",
			ExpectedSalaryCurrency: "SGD",
			ExpectedSalary:         1000,
			LinkedInUrl:            "linkedin",
			ScmUrl:                 "github",
			WebsiteUrl:             "website",
			EducationLevel:         "bachelor",
			Summary:                "summary",
			Birthday:               testPbTime,
			NoticePeriod:           1,
			PreferredRoles:         []string{"backend"},
			Skills: []*pb.Skill{
				{
					Id:   1,
					Name: "java",
				},
				{
					Id:   2,
					Name: "javascript",
				},
			},
			Academics: []*pb.AcademicHistory{
				{
					Id:            1,
					CandidateId:   1,
					InstitutionId: 1,
					CourseId:      1,
					YearObtained:  2020,
					Grade:         "first",
					CreatedAt:     testPbTime,
					UpdatedAt:     testPbTime,
					DeletedAt:     testPbTime,
				},
				{
					Id:            2,
					CandidateId:   1,
					InstitutionId: 1,
					CourseId:      2,
					YearObtained:  2020,
					Grade:         "first",
					CreatedAt:     testPbTime,
					UpdatedAt:     testPbTime,
					DeletedAt:     testPbTime,
				},
			},
			Jobs: []*pb.JobHistory{
				{
					Id:             1,
					CandidateId:    1,
					CompanyId:      1,
					DepartmentId:   1,
					Country:        "singapore",
					City:           "singapore",
					Title:          "software engineer",
					StartDate:      testPbTime,
					EndDate:        testPbTime,
					SalaryCurrency: "SGD",
					Salary:         1000,
					Description:    "worked hard",
					CreatedAt:      testPbTime,
					UpdatedAt:      testPbTime,
					DeletedAt:      testPbTime,
				},
				{
					Id:             2,
					CandidateId:    1,
					CompanyId:      1,
					DepartmentId:   1,
					Country:        "singapore",
					City:           "singapore",
					Title:          "senior software engineer",
					StartDate:      testPbTime,
					EndDate:        testPbTime,
					SalaryCurrency: "SGD",
					Salary:         2000,
					Description:    "worked hard",
					CreatedAt:      testPbTime,
					UpdatedAt:      testPbTime,
					DeletedAt:      testPbTime,
				},
			},
			CreatedAt: testPbTime,
			UpdatedAt: testPbTime,
			DeletedAt: testPbTime,
		},
		JobCompany: &pb.JoblistingCompany{
			Id:      1,
			Name:    "hubbedin",
			LogoUrl: "https://logo.jpg",
			Size:    50,
		},
		CreatedAt: testPbTime,
		UpdatedAt: testPbTime,
		DeletedAt: testPbTime,
	}

	expect := &User{
		ID:            1,
		AuthID:        "authId",
		FirstName:     "first",
		LastName:      "last",
		Email:         "email",
		ContactNumber: "contact",
		Picture:       "picture",
		Gender:        "male",
		Roles:         []string{"Candidate", "Company"},
		CandidateID:   1,
		JobCompanyID:  1,
		Candidate: &Candidate{
			Nationality:            "singapore",
			ResidenceCity:          "singapore",
			ExpectedSalaryCurrency: "SGD",
			ExpectedSalary:         1000,
			LinkedInURL:            "linkedin",
			SCMURL:                 "github",
			WebsiteURL:             "website",
			EducationLevel:         "bachelor",
			Summary:                "summary",
			Birthday:               &testTime,
			NoticePeriod:           1,
			PreferredRoles:         []string{"backend"},
			Skills: []*Skill{
				{
					ID:   1,
					Name: "java",
				},
				{
					ID:   2,
					Name: "javascript",
				},
			},
			Academics: []*AcademicHistory{
				{
					ID:            1,
					CandidateID:   1,
					InstitutionID: 1,
					CourseID:      1,
					YearObtained:  2020,
					CreatedAt:     &testTime,
					UpdatedAt:     &testTime,
					DeletedAt:     &testTime,
				},
				{
					ID:            2,
					CandidateID:   1,
					InstitutionID: 1,
					CourseID:      2,
					YearObtained:  2020,
					CreatedAt:     &testTime,
					UpdatedAt:     &testTime,
					DeletedAt:     &testTime,
				},
			},
			Jobs: []*JobHistory{
				{
					ID:             1,
					CandidateID:    1,
					CompanyID:      1,
					DepartmentID:   1,
					Country:        "singapore",
					City:           "singapore",
					Title:          "software engineer",
					StartDate:      &testTime,
					EndDate:        &testTime,
					SalaryCurrency: "SGD",
					Salary:         1000,
					Description:    "worked hard",
					CreatedAt:      &testTime,
					UpdatedAt:      &testTime,
					DeletedAt:      &testTime,
				},
				{
					ID:             2,
					CandidateID:    1,
					CompanyID:      1,
					DepartmentID:   1,
					Country:        "singapore",
					City:           "singapore",
					Title:          "senior software engineer",
					StartDate:      &testTime,
					EndDate:        &testTime,
					SalaryCurrency: "SGD",
					Salary:         2000,
					Description:    "worked hard",
					CreatedAt:      &testTime,
					UpdatedAt:      &testTime,
					DeletedAt:      &testTime,
				},
			},
			CreatedAt: &testTime,
			UpdatedAt: &testTime,
			DeletedAt: &testTime,
		},
		JobCompany: &JobCompany{
			ID:      1,
			Name:    "hubbedin",
			LogoURL: "https://logo.jpg",
			Size:    50,
		},
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
		DeletedAt: &testTime,
	}

	got := UserToORM(input)
	require.EqualValues(t, expect, got)
}

func TestCandidateToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.Candidate{
		Id:                     1,
		Nationality:            "singapore",
		ResidenceCity:          "singapore",
		ExpectedSalaryCurrency: "SGD",
		ExpectedSalary:         1000,
		LinkedInUrl:            "linkedin",
		ScmUrl:                 "github",
		WebsiteUrl:             "website",
		EducationLevel:         "bachelor",
		Summary:                "summary",
		Birthday:               testPbTime,
		NoticePeriod:           1,
		PreferredRoles:         []string{"backend"},
		Skills: []*pb.Skill{
			{
				Id:   1,
				Name: "java",
			},
			{
				Id:   2,
				Name: "javascript",
			},
		},
		Academics: []*pb.AcademicHistory{
			{
				Id:            1,
				CandidateId:   1,
				InstitutionId: 1,
				CourseId:      1,
				YearObtained:  2020,
				Grade:         "first",
				CreatedAt:     testPbTime,
				UpdatedAt:     testPbTime,
				DeletedAt:     testPbTime,
			},
			{
				Id:            2,
				CandidateId:   1,
				InstitutionId: 1,
				CourseId:      2,
				YearObtained:  2020,
				Grade:         "first",
				CreatedAt:     testPbTime,
				UpdatedAt:     testPbTime,
				DeletedAt:     testPbTime,
			},
		},
		Jobs: []*pb.JobHistory{
			{
				Id:             1,
				CandidateId:    1,
				CompanyId:      1,
				DepartmentId:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "software engineer",
				StartDate:      testPbTime,
				EndDate:        testPbTime,
				SalaryCurrency: "SGD",
				Salary:         1000,
				Description:    "worked hard",
				CreatedAt:      testPbTime,
				UpdatedAt:      testPbTime,
				DeletedAt:      testPbTime,
			},
			{
				Id:             2,
				CandidateId:    1,
				CompanyId:      1,
				DepartmentId:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "senior software engineer",
				StartDate:      testPbTime,
				EndDate:        testPbTime,
				SalaryCurrency: "SGD",
				Salary:         2000,
				Description:    "worked hard",
				CreatedAt:      testPbTime,
				UpdatedAt:      testPbTime,
				DeletedAt:      testPbTime,
			},
		},
		CreatedAt: testPbTime,
		UpdatedAt: testPbTime,
		DeletedAt: testPbTime,
	}

	expect := &Candidate{
		ID:                     1,
		Nationality:            "singapore",
		ResidenceCity:          "singapore",
		ExpectedSalaryCurrency: "SGD",
		ExpectedSalary:         1000,
		LinkedInURL:            "linkedin",
		SCMURL:                 "github",
		WebsiteURL:             "website",
		EducationLevel:         "bachelor",
		Summary:                "summary",
		Birthday:               &testTime,
		NoticePeriod:           1,
		PreferredRoles:         []string{"backend"},
		Skills: []*Skill{
			{
				ID:   1,
				Name: "java",
			},
			{
				ID:   2,
				Name: "javascript",
			},
		},
		Academics: []*AcademicHistory{
			{
				ID:            1,
				CandidateID:   1,
				InstitutionID: 1,
				CourseID:      1,
				YearObtained:  2020,
				CreatedAt:     &testTime,
				UpdatedAt:     &testTime,
				DeletedAt:     &testTime,
			},
			{
				ID:            2,
				CandidateID:   1,
				InstitutionID: 1,
				CourseID:      2,
				YearObtained:  2020,
				CreatedAt:     &testTime,
				UpdatedAt:     &testTime,
				DeletedAt:     &testTime,
			},
		},
		Jobs: []*JobHistory{
			{
				ID:             1,
				CandidateID:    1,
				CompanyID:      1,
				DepartmentID:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "software engineer",
				StartDate:      &testTime,
				EndDate:        &testTime,
				SalaryCurrency: "SGD",
				Salary:         1000,
				Description:    "worked hard",
				CreatedAt:      &testTime,
				UpdatedAt:      &testTime,
				DeletedAt:      &testTime,
			},
			{
				ID:             2,
				CandidateID:    1,
				CompanyID:      1,
				DepartmentID:   1,
				Country:        "singapore",
				City:           "singapore",
				Title:          "senior software engineer",
				StartDate:      &testTime,
				EndDate:        &testTime,
				SalaryCurrency: "SGD",
				Salary:         2000,
				Description:    "worked hard",
				CreatedAt:      &testTime,
				UpdatedAt:      &testTime,
				DeletedAt:      &testTime,
			},
		},
		CreatedAt: &testTime,
		UpdatedAt: &testTime,
		DeletedAt: &testTime,
	}

	got := CandidateToORM(input)
	require.EqualValues(t, expect, got)
}

func TestSkillToORM(t *testing.T) {
	input := &pb.Skill{
		Id:   1,
		Name: "skill",
	}

	expect := &Skill{
		ID:   1,
		Name: "skill",
	}

	got := SkillToORM(input)
	require.EqualValues(t, expect, got)
}

func TestUserkillToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.UserSkill{
		Id:          1,
		CandidateId: 1,
		SkillId:     1,
		CreatedAt:   testPbTime,
		UpdatedAt:   testPbTime,
	}

	expect := &UserSkill{
		ID:          1,
		CandidateID: 1,
		SkillID:     1,
		CreatedAt:   &testTime,
		UpdatedAt:   &testTime,
	}

	got := UserSkillToORM(input)
	require.EqualValues(t, expect, got)
}

func TestInstitutionToORM(t *testing.T) {
	input := &pb.Institution{
		Id:      1,
		Name:    "institution",
		Country: "singapore",
	}

	expect := &Institution{
		ID:      1,
		Name:    "institution",
		Country: "singapore",
	}

	got := InstitutionToORM(input)
	require.EqualValues(t, expect, got)
}

func TestCourseToORM(t *testing.T) {
	input := &pb.Course{
		Id:    1,
		Name:  "course",
		Level: "bachelor",
	}

	expect := &Course{
		ID:    1,
		Name:  "course",
		Level: "bachelor",
	}

	got := CourseToORM(input)
	require.EqualValues(t, expect, got)
}

func TestAcademicHistoryToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.AcademicHistory{
		Id:            1,
		CandidateId:   1,
		InstitutionId: 1,
		CourseId:      1,
		YearObtained:  2020,
		CreatedAt:     testPbTime,
		UpdatedAt:     testPbTime,
		DeletedAt:     testPbTime,
	}

	expect := &AcademicHistory{
		ID:            1,
		CandidateID:   1,
		InstitutionID: 1,
		CourseID:      1,
		YearObtained:  2020,
		CreatedAt:     &testTime,
		UpdatedAt:     &testTime,
		DeletedAt:     &testTime,
	}

	got := AcademicHistoryToORM(input)
	require.EqualValues(t, expect, got)
}

func TestCompanyToORM(t *testing.T) {
	input := &pb.Company{
		Id:   1,
		Name: "company",
	}

	expect := &Company{
		ID:   1,
		Name: "company",
	}

	got := CompanyToORM(input)
	require.EqualValues(t, expect, got)
}

func TestDepartmentToORM(t *testing.T) {
	input := &pb.Department{
		Id:   1,
		Name: "department",
	}

	expect := &Department{
		ID:   1,
		Name: "department",
	}

	got := DepartmentToORM(input)
	require.EqualValues(t, expect, got)
}

func TestJobHistoryToORM(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &pb.JobHistory{
		Id:             1,
		CandidateId:    1,
		CompanyId:      1,
		DepartmentId:   1,
		Country:        "singapore",
		City:           "singapore",
		Title:          "software engineer",
		StartDate:      testPbTime,
		EndDate:        testPbTime,
		SalaryCurrency: "SGD",
		Salary:         1000,
		Description:    "worked hard",
		CreatedAt:      testPbTime,
		UpdatedAt:      testPbTime,
		DeletedAt:      testPbTime,
	}

	expect := &JobHistory{
		ID:             1,
		CandidateID:    1,
		CompanyID:      1,
		DepartmentID:   1,
		Country:        "singapore",
		City:           "singapore",
		Title:          "software engineer",
		StartDate:      &testTime,
		EndDate:        &testTime,
		SalaryCurrency: "SGD",
		Salary:         1000,
		Description:    "worked hard",
		CreatedAt:      &testTime,
		UpdatedAt:      &testTime,
		DeletedAt:      &testTime,
	}

	got := JobHistoryToORM(input)
	require.EqualValues(t, expect, got)
}
