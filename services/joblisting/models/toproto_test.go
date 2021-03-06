package models

import (
	"in-backend/services/joblisting/pb"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/require"
)

func TestJobPostToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &JobPost{
		ID:              1,
		CompanyID:       1,
		HRContactID:     1,
		HiringManagerID: 1,
		JobPlatformID:   1,
		Title:           "software engineer",
		Description:     "description",
		SeniorityLevel:  "entry",
		YearsExperience: 2,
		EmploymentType:  "full-time",
		FunctionID:      1,
		IndustryID:      1,
		Location:        "singapore",
		Remote:          true,
		SalaryCurrency:  "SGD",
		MinSalary:       1000,
		MaxSalary:       2000,
		SkillID:         []uint64{1, 2},
		CreatedAt:       &testTime,
		UpdatedAt:       &testTime,
		StartAt:         &testTime,
		ExpireAt:        &testTime,
		Company: &Company{
			ID:      1,
			Name:    "hubbedin",
			LogoURL: "https://logo.jpg",
			Size:    50,
		},
		HRContact: &KeyPerson{
			ID:            1,
			CompanyID:     1,
			Name:          "average joe",
			ContactNumber: "+6512345678",
			Email:         "email",
			JobTitle:      "hr manager",
			UpdatedAt:     &testTime,
		},
		HiringManager: &KeyPerson{
			ID:            2,
			CompanyID:     2,
			Name:          "plain jane",
			ContactNumber: "+6512345678",
			Email:         "email",
			JobTitle:      "ceo",
			UpdatedAt:     &testTime,
		},
	}

	expect := &pb.JobPost{
		Id:              1,
		CompanyId:       1,
		HrContactId:     1,
		HiringManagerId: 1,
		JobPlatformId:   1,
		Title:           "software engineer",
		Description:     "description",
		SeniorityLevel:  "entry",
		YearsExperience: 2,
		EmploymentType:  "full-time",
		FunctionId:      1,
		IndustryId:      1,
		Location:        "singapore",
		Remote:          true,
		SalaryCurrency:  "SGD",
		MinSalary:       1000,
		MaxSalary:       2000,
		SkillId:         []uint64{1, 2},
		CreatedAt:       testPbTime,
		UpdatedAt:       testPbTime,
		StartAt:         testPbTime,
		ExpireAt:        testPbTime,
		Company: &pb.JobCompany{
			Id:      1,
			Name:    "hubbedin",
			LogoUrl: "https://logo.jpg",
			Size:    50,
		},
		HrContact: &pb.KeyPerson{
			Id:            1,
			CompanyId:     1,
			Name:          "average joe",
			ContactNumber: "+6512345678",
			Email:         "email",
			JobTitle:      "hr manager",
			UpdatedAt:     testPbTime,
		},
		HiringManager: &pb.KeyPerson{
			Id:            2,
			CompanyId:     2,
			Name:          "plain jane",
			ContactNumber: "+6512345678",
			Email:         "email",
			JobTitle:      "ceo",
			UpdatedAt:     testPbTime,
		},
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestCompanyToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &Company{
		ID:      1,
		Name:    "hubbedin",
		LogoURL: "https://logo.jpg",
		Size:    50,
		Industries: []*Industry{
			{
				ID:   1,
				Name: "tech",
			},
			{
				ID:   2,
				Name: "bank",
			},
		},
		JobPosts: []*JobPost{
			{
				ID:              1,
				CompanyID:       1,
				HRContactID:     1,
				HiringManagerID: 1,
				JobPlatformID:   1,
				Title:           "software engineer",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 2,
				EmploymentType:  "full-time",
				FunctionID:      1,
				IndustryID:      1,
				Location:        "singapore",
				Remote:          true,
				SalaryCurrency:  "SGD",
				MinSalary:       1000,
				MaxSalary:       2000,
				SkillID:         []uint64{1, 2},
				CreatedAt:       &testTime,
				UpdatedAt:       &testTime,
				StartAt:         &testTime,
				ExpireAt:        &testTime,
			},
			{
				ID:              2,
				CompanyID:       2,
				HRContactID:     2,
				HiringManagerID: 2,
				JobPlatformID:   2,
				Title:           "product manager",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 5,
				EmploymentType:  "full-time",
				FunctionID:      1,
				IndustryID:      1,
				Location:        "singapore",
				Remote:          true,
				MinSalary:       3000,
				MaxSalary:       4000,
				SkillID:         []uint64{1, 2},
				CreatedAt:       &testTime,
				UpdatedAt:       &testTime,
				StartAt:         &testTime,
				ExpireAt:        &testTime,
			},
		},
		KeyPersons: []*KeyPerson{
			{
				ID:            1,
				CompanyID:     1,
				Name:          "average joe",
				ContactNumber: "+6512345678",
				Email:         "email",
				JobTitle:      "cto",
				UpdatedAt:     &testTime,
			},
			{
				ID:            2,
				CompanyID:     2,
				Name:          "average jane",
				ContactNumber: "+6512345678",
				Email:         "email",
				JobTitle:      "ceo",
				UpdatedAt:     &testTime,
			},
		},
	}

	expect := &pb.JobCompany{
		Id:      1,
		Name:    "hubbedin",
		LogoUrl: "https://logo.jpg",
		Size:    50,
		Industries: []*pb.Industry{
			{},
			{},
		},
		JobPosts: []*pb.JobPost{
			{
				Id:              1,
				CompanyId:       1,
				HrContactId:     1,
				HiringManagerId: 1,
				JobPlatformId:   1,
				Title:           "software engineer",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 2,
				EmploymentType:  "full-time",
				FunctionId:      1,
				IndustryId:      1,
				Location:        "singapore",
				Remote:          true,
				SalaryCurrency:  "SGD",
				MinSalary:       1000,
				MaxSalary:       2000,
				SkillId:         []uint64{1, 2},
				CreatedAt:       testPbTime,
				UpdatedAt:       testPbTime,
				StartAt:         testPbTime,
				ExpireAt:        testPbTime,
			},
			{
				Id:              2,
				CompanyId:       2,
				HrContactId:     2,
				HiringManagerId: 2,
				JobPlatformId:   2,
				Title:           "product manager",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 5,
				EmploymentType:  "full-time",
				FunctionId:      1,
				IndustryId:      1,
				Location:        "singapore",
				Remote:          true,
				MinSalary:       3000,
				MaxSalary:       4000,
				SkillId:         []uint64{1, 2},
				CreatedAt:       testPbTime,
				UpdatedAt:       testPbTime,
				StartAt:         testPbTime,
				ExpireAt:        testPbTime,
			},
		},
		KeyPersons: []*pb.KeyPerson{
			{
				Id:            1,
				CompanyId:     1,
				Name:          "average joe",
				ContactNumber: "+6512345678",
				Email:         "email",
				JobTitle:      "cto",
				UpdatedAt:     testPbTime,
			},
			{
				Id:            2,
				CompanyId:     2,
				Name:          "plain jane",
				ContactNumber: "+6512345678",
				Email:         "email",
				JobTitle:      "ceo",
				UpdatedAt:     testPbTime,
			},
		},
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestIndustryToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &Industry{
		ID:   1,
		Name: "tech",
		Companies: []*Company{
			{
				ID:      1,
				Name:    "hubbedin",
				LogoURL: "https://logo.jpg",
				Size:    50,
			},
			{
				ID:      2,
				Name:    "linkedin",
				LogoURL: "https://logolinkedin.jpg",
				Size:    50,
			},
		},
		JobPosts: []*JobPost{
			{
				ID:              1,
				CompanyID:       1,
				HRContactID:     1,
				HiringManagerID: 1,
				JobPlatformID:   1,
				Title:           "software engineer",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 2,
				EmploymentType:  "full-time",
				FunctionID:      1,
				IndustryID:      1,
				Location:        "singapore",
				Remote:          true,
				SalaryCurrency:  "SGD",
				MinSalary:       1000,
				MaxSalary:       2000,
				SkillID:         []uint64{1, 2},
				CreatedAt:       &testTime,
				UpdatedAt:       &testTime,
				StartAt:         &testTime,
				ExpireAt:        &testTime,
			},
			{
				ID:              2,
				CompanyID:       2,
				HRContactID:     2,
				HiringManagerID: 2,
				JobPlatformID:   2,
				Title:           "product manager",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 5,
				EmploymentType:  "full-time",
				FunctionID:      1,
				IndustryID:      1,
				Location:        "singapore",
				Remote:          true,
				MinSalary:       3000,
				MaxSalary:       4000,
				SkillID:         []uint64{1, 2},
				CreatedAt:       &testTime,
				UpdatedAt:       &testTime,
				StartAt:         &testTime,
				ExpireAt:        &testTime,
			},
		},
	}

	expect := &pb.Industry{
		Id:   1,
		Name: "tech",
		Companies: []*pb.JobCompany{
			{
				Id:      1,
				Name:    "hubbedin",
				LogoUrl: "https://logo.jpg",
				Size:    50,
			},
			{
				Id:      2,
				Name:    "linkedin",
				LogoUrl: "https://logolinkedin.jpg",
				Size:    50,
			},
		},
		JobPosts: []*pb.JobPost{
			{
				Id:              1,
				CompanyId:       1,
				HrContactId:     1,
				HiringManagerId: 1,
				JobPlatformId:   1,
				Title:           "software engineer",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 2,
				EmploymentType:  "full-time",
				FunctionId:      1,
				IndustryId:      1,
				Location:        "singapore",
				Remote:          true,
				SalaryCurrency:  "SGD",
				MinSalary:       1000,
				MaxSalary:       2000,
				SkillId:         []uint64{1, 2},
				CreatedAt:       testPbTime,
				UpdatedAt:       testPbTime,
				StartAt:         testPbTime,
				ExpireAt:        testPbTime,
			},
			{
				Id:              2,
				CompanyId:       2,
				HrContactId:     2,
				HiringManagerId: 2,
				JobPlatformId:   2,
				Title:           "product manager",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 5,
				EmploymentType:  "full-time",
				FunctionId:      1,
				IndustryId:      1,
				Location:        "singapore",
				Remote:          true,
				MinSalary:       3000,
				MaxSalary:       4000,
				SkillId:         []uint64{1, 2},
				CreatedAt:       testPbTime,
				UpdatedAt:       testPbTime,
				StartAt:         testPbTime,
				ExpireAt:        testPbTime,
			},
		},
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestKeyPersonToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &KeyPerson{
		ID:            1,
		CompanyID:     1,
		Name:          "average joe",
		ContactNumber: "+6512345678",
		Email:         "email",
		JobTitle:      "cto",
		UpdatedAt:     &testTime,
		Company: &Company{
			ID:      1,
			Name:    "hubbedin",
			LogoURL: "https://logo.jpg",
			Size:    50,
		},
	}

	expect := &pb.KeyPerson{
		Id:            1,
		CompanyId:     1,
		Name:          "average joe",
		ContactNumber: "+6512345678",
		Email:         "email",
		JobTitle:      "cto",
		UpdatedAt:     testPbTime,
		Company: &pb.JobCompany{
			Id:      1,
			Name:    "hubbedin",
			LogoUrl: "https://logo.jpg",
			Size:    50,
		},
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}

func TestJobPlatformToProto(t *testing.T) {
	testPbTime := ptypes.TimestampNow()
	testTime, err := ptypes.Timestamp(testPbTime)
	require.NoError(t, err)

	input := &JobPlatform{
		ID:      1,
		Name:    "indeed",
		BaseURL: "https://indeed.com",
		JobPosts: []*JobPost{
			{
				ID:              1,
				CompanyID:       1,
				HRContactID:     1,
				HiringManagerID: 1,
				JobPlatformID:   1,
				Title:           "software engineer",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 2,
				EmploymentType:  "full-time",
				FunctionID:      1,
				IndustryID:      1,
				Location:        "singapore",
				Remote:          true,
				SalaryCurrency:  "SGD",
				MinSalary:       1000,
				MaxSalary:       2000,
				SkillID:         []uint64{1, 2},
				CreatedAt:       &testTime,
				UpdatedAt:       &testTime,
				StartAt:         &testTime,
				ExpireAt:        &testTime,
			},
			{
				ID:              2,
				CompanyID:       2,
				HRContactID:     2,
				HiringManagerID: 2,
				JobPlatformID:   2,
				Title:           "product manager",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 5,
				EmploymentType:  "full-time",
				FunctionID:      1,
				IndustryID:      1,
				Location:        "singapore",
				Remote:          true,
				MinSalary:       3000,
				MaxSalary:       4000,
				SkillID:         []uint64{1, 2},
				CreatedAt:       &testTime,
				UpdatedAt:       &testTime,
				StartAt:         &testTime,
				ExpireAt:        &testTime,
			},
		},
	}

	expect := &pb.JobPlatform{
		Id:      1,
		Name:    "indeed",
		BaseUrl: "https://indeed.com",
		JobPosts: []*pb.JobPost{
			{
				Id:              1,
				CompanyId:       1,
				HrContactId:     1,
				HiringManagerId: 1,
				JobPlatformId:   1,
				Title:           "software engineer",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 2,
				EmploymentType:  "full-time",
				FunctionId:      1,
				IndustryId:      1,
				Location:        "singapore",
				Remote:          true,
				SalaryCurrency:  "SGD",
				MinSalary:       1000,
				MaxSalary:       2000,
				SkillId:         []uint64{1, 2},
				CreatedAt:       testPbTime,
				UpdatedAt:       testPbTime,
				StartAt:         testPbTime,
				ExpireAt:        testPbTime,
			},
			{
				Id:              2,
				CompanyId:       2,
				HrContactId:     2,
				HiringManagerId: 2,
				JobPlatformId:   2,
				Title:           "product manager",
				Description:     "description",
				SeniorityLevel:  "entry",
				YearsExperience: 5,
				EmploymentType:  "full-time",
				FunctionId:      1,
				IndustryId:      1,
				Location:        "singapore",
				Remote:          true,
				MinSalary:       3000,
				MaxSalary:       4000,
				SkillId:         []uint64{1, 2},
				CreatedAt:       testPbTime,
				UpdatedAt:       testPbTime,
				StartAt:         testPbTime,
				ExpireAt:        testPbTime,
			},
		},
	}

	got := input.ToProto()
	require.EqualValues(t, expect, got)
}
