// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "in-backend/services/assessment/pb"
)

// AssessmentServiceClient is an autogenerated mock type for the AssessmentServiceClient type
type AssessmentServiceClient struct {
	mock.Mock
}

// BulkCreateQuestion provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) BulkCreateQuestion(ctx context.Context, in *pb.BulkCreateQuestionRequest, opts ...grpc.CallOption) (*pb.BulkCreateQuestionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.BulkCreateQuestionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.BulkCreateQuestionRequest, ...grpc.CallOption) *pb.BulkCreateQuestionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.BulkCreateQuestionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.BulkCreateQuestionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAssessment provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) CreateAssessment(ctx context.Context, in *pb.CreateAssessmentRequest, opts ...grpc.CallOption) (*pb.Assessment, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.Assessment
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateAssessmentRequest, ...grpc.CallOption) *pb.Assessment); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Assessment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateAssessmentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAssessmentAttempt provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) CreateAssessmentAttempt(ctx context.Context, in *pb.CreateAssessmentAttemptRequest, opts ...grpc.CallOption) (*pb.AssessmentAttempt, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AssessmentAttempt
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateAssessmentAttemptRequest, ...grpc.CallOption) *pb.AssessmentAttempt); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AssessmentAttempt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateAssessmentAttemptRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQuestion provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) CreateQuestion(ctx context.Context, in *pb.CreateQuestionRequest, opts ...grpc.CallOption) (*pb.Question, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.Question
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateQuestionRequest, ...grpc.CallOption) *pb.Question); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateQuestionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTag provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) CreateTag(ctx context.Context, in *pb.CreateTagRequest, opts ...grpc.CallOption) (*pb.Tag, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.Tag
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateTagRequest, ...grpc.CallOption) *pb.Tag); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateTagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAssessment provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) DeleteAssessment(ctx context.Context, in *pb.DeleteAssessmentRequest, opts ...grpc.CallOption) (*pb.DeleteAssessmentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.DeleteAssessmentResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteAssessmentRequest, ...grpc.CallOption) *pb.DeleteAssessmentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteAssessmentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteAssessmentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAssessmentAttempt provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) DeleteAssessmentAttempt(ctx context.Context, in *pb.DeleteAssessmentAttemptRequest, opts ...grpc.CallOption) (*pb.DeleteAssessmentAttemptResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.DeleteAssessmentAttemptResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteAssessmentAttemptRequest, ...grpc.CallOption) *pb.DeleteAssessmentAttemptResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteAssessmentAttemptResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteAssessmentAttemptRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteQuestion provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) DeleteQuestion(ctx context.Context, in *pb.DeleteQuestionRequest, opts ...grpc.CallOption) (*pb.DeleteQuestionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.DeleteQuestionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteQuestionRequest, ...grpc.CallOption) *pb.DeleteQuestionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteQuestionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteQuestionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTag provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) DeleteTag(ctx context.Context, in *pb.DeleteTagRequest, opts ...grpc.CallOption) (*pb.DeleteTagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.DeleteTagResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeleteTagRequest, ...grpc.CallOption) *pb.DeleteTagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeleteTagResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeleteTagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAssessments provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) GetAllAssessments(ctx context.Context, in *pb.GetAllAssessmentsRequest, opts ...grpc.CallOption) (*pb.GetAllAssessmentsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.GetAllAssessmentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllAssessmentsRequest, ...grpc.CallOption) *pb.GetAllAssessmentsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllAssessmentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllAssessmentsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllQuestions provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) GetAllQuestions(ctx context.Context, in *pb.GetAllQuestionsRequest, opts ...grpc.CallOption) (*pb.GetAllQuestionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.GetAllQuestionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAllQuestionsRequest, ...grpc.CallOption) *pb.GetAllQuestionsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetAllQuestionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAllQuestionsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAssessmentAttemptByID provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) GetAssessmentAttemptByID(ctx context.Context, in *pb.GetAssessmentAttemptByIDRequest, opts ...grpc.CallOption) (*pb.AssessmentAttempt, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AssessmentAttempt
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAssessmentAttemptByIDRequest, ...grpc.CallOption) *pb.AssessmentAttempt); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AssessmentAttempt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAssessmentAttemptByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAssessmentByID provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) GetAssessmentByID(ctx context.Context, in *pb.GetAssessmentByIDRequest, opts ...grpc.CallOption) (*pb.Assessment, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.Assessment
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAssessmentByIDRequest, ...grpc.CallOption) *pb.Assessment); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Assessment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAssessmentByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuestionByID provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) GetQuestionByID(ctx context.Context, in *pb.GetQuestionByIDRequest, opts ...grpc.CallOption) (*pb.Question, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.Question
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetQuestionByIDRequest, ...grpc.CallOption) *pb.Question); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetQuestionByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalGetAssessmentAttemptByID provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) LocalGetAssessmentAttemptByID(ctx context.Context, in *pb.GetAssessmentAttemptByIDRequest, opts ...grpc.CallOption) (*pb.AssessmentAttempt, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AssessmentAttempt
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetAssessmentAttemptByIDRequest, ...grpc.CallOption) *pb.AssessmentAttempt); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AssessmentAttempt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetAssessmentAttemptByIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalUpdateAssessmentAttempt provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) LocalUpdateAssessmentAttempt(ctx context.Context, in *pb.UpdateAssessmentAttemptRequest, opts ...grpc.CallOption) (*pb.AssessmentAttempt, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AssessmentAttempt
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateAssessmentAttemptRequest, ...grpc.CallOption) *pb.AssessmentAttempt); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AssessmentAttempt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateAssessmentAttemptRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAssessment provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) UpdateAssessment(ctx context.Context, in *pb.UpdateAssessmentRequest, opts ...grpc.CallOption) (*pb.Assessment, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.Assessment
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateAssessmentRequest, ...grpc.CallOption) *pb.Assessment); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Assessment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateAssessmentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAssessmentAttempt provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) UpdateAssessmentAttempt(ctx context.Context, in *pb.UpdateAssessmentAttemptRequest, opts ...grpc.CallOption) (*pb.AssessmentAttempt, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AssessmentAttempt
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateAssessmentAttemptRequest, ...grpc.CallOption) *pb.AssessmentAttempt); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AssessmentAttempt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateAssessmentAttemptRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAttemptQuestion provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) UpdateAttemptQuestion(ctx context.Context, in *pb.UpdateAttemptQuestionRequest, opts ...grpc.CallOption) (*pb.AttemptQuestion, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AttemptQuestion
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateAttemptQuestionRequest, ...grpc.CallOption) *pb.AttemptQuestion); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AttemptQuestion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateAttemptQuestionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateQuestion provides a mock function with given fields: ctx, in, opts
func (_m *AssessmentServiceClient) UpdateQuestion(ctx context.Context, in *pb.UpdateQuestionRequest, opts ...grpc.CallOption) (*pb.Question, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.Question
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdateQuestionRequest, ...grpc.CallOption) *pb.Question); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdateQuestionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
