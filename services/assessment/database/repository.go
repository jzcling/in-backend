package database

import (
	"context"
	"fmt"
	"strings"

	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pkg/errors"

	"in-backend/services/assessment/interfaces"
	"in-backend/services/assessment/models"
)

// Repository implements the assessment Repository interface
type repository struct {
	DB *pg.DB
}

// NewRepository declares a new Repository that implements assessment Repository
func NewRepository(db *pg.DB) interfaces.Repository {
	return &repository{
		DB: db,
	}
}

/* --------------- Assessment --------------- */

// CreateAssessment creates a new Assessment
func (r *repository) CreateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error) {
	if m == nil {
		return nil, errors.New("Input parameter assessment is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation(relQuestions).
		Relation(relAttempts).
		Returning("*").
		Where("name = ?", m.Name).
		OnConflict("DO NOTHING").
		SelectOrInsert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert assessment %v", m)
		return nil, err
	}

	return m, nil
}

// GetAllAssessments returns all Assessments
func (r *repository) GetAllAssessments(ctx context.Context, f models.AssessmentFilters, role *string, cid *uint64) ([]*models.Assessment, error) {
	var m []*models.Assessment
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("a.id in (?)", pg.In(f.ID))
	}
	if f.Name != "" {
		q = q.Where("lower(a.name) like ?", "%"+strings.ToLower(f.Name)+"%")
	}
	if len(f.Difficulty) > 0 {
		q = q.Where("a.difficulty in (?)", pg.In(f.Difficulty))
	}
	if len(f.Type) > 0 {
		q = q.Where("a.type in (?)", pg.In(f.Type))
	}
	if f.CandidateID > 0 {
		q = q.Where("as.candidate_id = ?", f.CandidateID)
	}
	if len(f.Status) > 0 {
		q = q.Where("as.status in (?)", pg.In(f.Status))
	}
	if f.MinScore > 0 {
		q = q.Where("as.score >= ?", f.MinScore)
	}

	if *role == "Admin" {
		q = q.Relation(relAttempts)
	}
	if *role == "Owner" {
		q = q.Relation(relAttempts, func(q *orm.Query) (*orm.Query, error) {
			return q.Where("aa.candidate_id = ?", *cid), nil
		})
	}
	err := q.Relation(relQuestions).
		Returning("*").
		Select()
	return m, err
}

// GetAssessmentByID returns a Assessment by ID
func (r *repository) GetAssessmentByID(ctx context.Context, id uint64, role *string, cid *uint64) (*models.Assessment, error) {
	m := models.Assessment{ID: id}
	q := r.DB.WithContext(ctx).Model(&m).
		Where("id = ?", id).
		Relation(relQuestions)
	if *role == "Admin" {
		q = q.Relation(relAttempts)
	} else {
		q = q.Relation(relAttempts, func(q *orm.Query) (*orm.Query, error) {
			return q.Where("aa.candidate_id = ?", *cid).Order("aa.started_at desc"), nil
		})
	}
	err := q.Returning("*").First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateAssessment updates a Assessment
func (r *repository) UpdateAssessment(ctx context.Context, m *models.Assessment) (*models.Assessment, error) {
	if m == nil {
		return nil, errors.New("Assessment is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relQuestions).
		Relation(relAttempts).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update assessment with id %v", m.ID))
	}

	return m, nil
}

// DeleteAssessment deletes a Assessment by ID
func (r *repository) DeleteAssessment(ctx context.Context, id uint64) error {
	m := &models.Assessment{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete assessment with id %v", id))
	}
	return nil
}

/* --------------- Assessment Attempt --------------- */

// CreateAssessmentAttempt creates a new AssessmentAttempt
func (r *repository) CreateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	if m == nil {
		return nil, errors.New("Input parameter assessment attempt is nil")
	}

	tx, err := r.DB.BeginContext(ctx)
	defer tx.Close()

	_, err = tx.Model(m).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert assessment attempt %v", m)
		tx.Rollback()
		return nil, err
	}

	var aaqSlice []models.AttemptQuestion
	for _, q := range m.Questions {
		aaq := models.AttemptQuestion{
			AttemptID:   m.ID,
			QuestionID:  q.ID,
			CandidateID: m.CandidateID,
			Selection:   -1,
			Score:       -1,
		}
		aaqSlice = append(aaqSlice, aaq)
	}

	_, err = tx.Model(&aaqSlice).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert questions for assessment attempt")
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return m, nil
}

// GetAssessmentAttemptByID returns a AssessmentAttempt by ID
func (r *repository) GetAssessmentAttemptByID(ctx context.Context, id uint64) (*models.AssessmentAttempt, error) {
	m := models.AssessmentAttempt{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where("aa.id = ?", id).
		Relation(relAssessment).
		Relation(relQuestions).
		Relation(relQuestionAttempts).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// GetLatestAssessmentAttemptByCandidate returns the latest AssessmentAttempt by a Candidate
func (r *repository) GetLatestAssessmentAttemptByCandidate(ctx context.Context, cid uint64) (*models.AssessmentAttempt, error) {
	m := models.AssessmentAttempt{}
	err := r.DB.WithContext(ctx).Model(&m).
		Where("aa.candidate_id = ?", cid).
		Order("aa.started_at desc").
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateAssessmentAttempt updates a AssessmentAttempt
func (r *repository) UpdateAssessmentAttempt(ctx context.Context, m *models.AssessmentAttempt) (*models.AssessmentAttempt, error) {
	if m == nil {
		return nil, errors.New("AssessmentAttempt is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Returning("*").
		Relation(relAssessment).
		Relation(relQuestions).
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update assessment attempt with id %v", m.ID))
	}

	return m, nil
}

// DeleteAssessmentAttempt deletes a AssessmentAttempt by ID
func (r *repository) DeleteAssessmentAttempt(ctx context.Context, id uint64) error {
	m := &models.AssessmentAttempt{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete assessment attempt with id %v", id))
	}
	return nil
}

/* --------------- Question --------------- */

// CreateQuestion creates a new Question
func (r *repository) CreateQuestion(ctx context.Context, m *models.Question) (*models.Question, error) {
	if m == nil {
		return nil, errors.New("Input parameter question is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Relation(relTags).
		Relation(relAssessments).
		Relation(relAttempts).
		Returning("*").
		Insert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert question %v", m)
		return nil, err
	}

	return m, nil
}

// BulkCreateQuestion creates a new Question
func (r *repository) BulkCreateQuestion(ctx context.Context, m []*models.Question) ([]*models.Question, error) {
	empty := []*models.Question{}
	if len(m) == 0 {
		return empty, errors.New("Input parameter questions is empty")
	}

	// get all tag names first before creating questions
	// this is because the tags are lost after inserting questions with go-pg
	tags := make(map[int]([]string))
	for i, q := range m {
		var names []string
		for _, t := range q.Tags {
			names = append(names, t.Name)
		}
		tags[i] = names
	}

	tx, err := r.DB.BeginContext(ctx)
	defer tx.Close()

	_, err = tx.Model(&m).
		Returning("*").
		Insert()
	if err != nil {
		tx.Rollback()
		err = errors.Wrapf(err, "Failed to insert questions %v", m)
		return empty, err
	}

	// get all assessment questions for bulk insert later
	var assessmentQuestions []*models.AssessmentQuestion
	// get all question tags for bulk insert later
	var questionTags []*models.QuestionTag
	// get all question IDs to pull data to return later
	var qid []uint64
	for i, q := range m {
		aq := &models.AssessmentQuestion{
			AssessmentID: q.Assessments[0].ID,
			QuestionID:   q.ID,
		}
		assessmentQuestions = append(assessmentQuestions, aq)

		if len(tags[i]) > 0 {
			for _, n := range tags[i] {
				// create tag if it doesn't exist
				tag := &models.Tag{Name: n}
				_, err = tx.Model(tag).
					Where("name = ?", tag.Name).
					OnConflict("DO NOTHING").
					Returning("*").
					SelectOrInsert()
				if err != nil {
					tx.Rollback()
					err = errors.Wrapf(err, "Failed to insert tag %v", tag)
					return empty, err
				}

				qt := &models.QuestionTag{
					QuestionID: q.ID,
					TagID:      tag.ID,
				}
				questionTags = append(questionTags, qt)
			}
		}

		qid = append(qid, q.ID)
	}

	if len(assessmentQuestions) > 0 {
		_, err = tx.Model(&assessmentQuestions).
			Returning("*").
			Insert()
		if err != nil {
			tx.Rollback()
			err = errors.Wrapf(err, "Failed to insert assessment questions %v", assessmentQuestions)
			return empty, err
		}
	}

	if len(questionTags) > 0 {
		_, err = tx.Model(&questionTags).
			Returning("*").
			Insert()
		if err != nil {
			tx.Rollback()
			err = errors.Wrapf(err, "Failed to insert question tags %v", questionTags)
			return empty, err
		}
	}

	if err := tx.Commit(); err != nil {
		return empty, err
	}

	f := models.QuestionFilters{ID: qid}
	questions, err := r.GetAllQuestions(ctx, f)
	if err != nil {
		return m, nil
	}

	return questions, nil
}

// GetAllQuestions returns all Questions
func (r *repository) GetAllQuestions(ctx context.Context, f models.QuestionFilters) ([]*models.Question, error) {
	var m []*models.Question
	q := r.DB.WithContext(ctx).Model(&m)
	if len(f.ID) > 0 {
		q = q.Where("q.id in (?)", pg.In(f.ID))
	}
	if len(f.Tags) > 0 {
		q = q.Where("t.name in (?)", pg.In(f.Tags))
	}
	err := q.Relation(relTags).
		Relation(relAssessments).
		Relation(relAttempts).
		Returning("*").
		Select()
	return m, err
}

// GetQuestionByID returns a Question by ID
func (r *repository) GetQuestionByID(ctx context.Context, id uint64) (*models.Question, error) {
	m := models.Question{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where("id = ?", id).
		Relation(relTags).
		Relation(relAssessments).
		Relation(relAttempts).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateQuestion updates a Question
func (r *repository) UpdateQuestion(ctx context.Context, m *models.Question) (*models.Question, error) {
	if m == nil {
		return nil, errors.New("Question is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Relation(relTags).
		Relation(relAssessments).
		Relation(relAttempts).
		Returning("*").
		Update()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot update question with id %v", m.ID))
	}

	return m, nil
}

// DeleteQuestion deletes a Question by ID
func (r *repository) DeleteQuestion(ctx context.Context, id uint64) error {
	m := &models.Question{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete question with id %v", id))
	}
	return nil
}

/* --------------- Tag --------------- */

// CreateTag creates a new Tag
func (r *repository) CreateTag(ctx context.Context, m *models.Tag) (*models.Tag, error) {
	if m == nil {
		return nil, errors.New("Input parameter tag is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).
		Where("name = ?", m.Name).
		OnConflict("DO NOTHING").
		Returning("*").
		SelectOrInsert()
	if err != nil {
		err = errors.Wrapf(err, "Failed to insert tag %v", m)
		return nil, err
	}

	return m, nil
}

// DeleteTag deletes a Tag by ID
func (r *repository) DeleteTag(ctx context.Context, id uint64) error {
	m := &models.Tag{ID: id}
	_, err := r.DB.WithContext(ctx).Model(m).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot delete tag with id %v", id))
	}
	return nil
}

/* --------------- Attempt Question --------------- */

// GetAttemptQuestionByID returns a AttemptQuestion by ID
func (r *repository) GetAttemptQuestionByID(ctx context.Context, id uint64) (*models.AttemptQuestion, error) {
	m := models.AttemptQuestion{ID: id}
	err := r.DB.WithContext(ctx).Model(&m).
		Where("id = ?", id).
		Returning("*").
		First()
	//pg returns error when no rows in the result set
	if err == pg.ErrNoRows {
		return nil, nil
	}
	return &m, err
}

// UpdateAttemptQuestion updates a Attempt Question
func (r *repository) UpdateAttemptQuestion(ctx context.Context, m *models.AttemptQuestion) (*models.AttemptQuestion, error) {
	if m == nil {
		return nil, errors.New("AttemptQuestion is nil")
	}

	_, err := r.DB.WithContext(ctx).Model(m).WherePK().
		Returning("*").
		Update()
	if err != nil {
		err = errors.Wrapf(err, "Failed to update attempt question %v", m)
		return nil, err
	}

	return m, nil
}
