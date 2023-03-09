package postgres

import (
	example "app-module/internal/app/example/domain"
	"app-module/pkg/errors"
	"app-module/pkg/postgres"
	"database/sql"

	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  *example.Instance
		want   string
		err    error
		expect error
	}{
		{
			input:  &example.Instance{Test: "tester"},
			want:   "tester",
			err:    nil,
			expect: nil,
		},
		{
			input:  &example.Instance{Test: "builder"},
			want:   "builder",
			err:    fmt.Errorf("example unique"),
			expect: fmt.Errorf("example unique"),
		},
	}

	for i := range tests {
		if tests[i].err != nil {
			mock.ExpectExec("INSERT INTO example").WithArgs(tests[i].want).WillReturnError(tests[i].err)
		} else {
			mock.ExpectExec("INSERT INTO example").WithArgs(tests[i].want).WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err := repo.Create(tests[i].input)

		assert.Equal(t, tests[i].expect, errors.Cause(err), "TestCase # %d", i+1)
	}
}

func TestGetOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input        int
		want         int
		result       *sqlmock.Rows
		err          error
		expectResult *example.Instance
		expectErr    error
	}{
		{
			input:  1,
			want:   1,
			result: sqlmock.NewRows([]string{"name"}).AddRow("tester"),
			err:    nil,
			expectResult: &example.Instance{
				ID:   1,
				Test: "tester",
			},
			expectErr: nil,
		},
		{
			input:        2,
			want:         2,
			result:       nil,
			err:          sql.ErrNoRows,
			expectResult: nil,
			expectErr:    sql.ErrNoRows,
		},
		{
			input:        0,
			want:         0,
			result:       nil,
			err:          fmt.Errorf("some error"),
			expectResult: nil,
			expectErr:    fmt.Errorf("some error"),
		},
	}

	for i := range tests {
		if tests[i].err != nil {
			mock.ExpectQuery("SELECT name FROM example").WithArgs(tests[i].want).WillReturnError(tests[i].err).RowsWillBeClosed()
		} else {
			mock.ExpectQuery("SELECT name FROM example").WithArgs(tests[i].want).WillReturnRows(tests[i].result)
		}

		result, err := repo.GetOne(tests[i].input)

		assert.Equal(t, tests[i].expectErr, errors.Cause(err), "error TestCase # %d", i+1)

		assert.Equal(t, tests[i].expectResult, result, "result TestCase # %d", i+1)
	}
}

func TestGetMany(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input        []int
		want         []int
		result       *sqlmock.Rows
		err          error
		expectResult []*example.Instance
		expectErr    error
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
			result: sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "tester"),
			err:    nil,
			expectResult: []*example.Instance{
				{
					ID:   1,
					Test: "tester",
				},
			},
			expectErr: nil,
		},
		{
			input:        nil,
			want:         nil,
			result:       nil,
			err:          fmt.Errorf("some error"),
			expectResult: nil,
			expectErr:    fmt.Errorf("some error"),
		},
	}

	for i := range tests {
		if tests[i].err != nil {
			mock.ExpectQuery("SELECT id, name FROM example").WithArgs(pq.Array(tests[i].want)).WillReturnError(tests[i].err)
		} else {
			mock.ExpectQuery("SELECT id, name FROM example").WithArgs(pq.Array(tests[i].want)).WillReturnRows(tests[i].result)
		}

		result, err := repo.GetMany(tests[i].input)

		assert.Equal(t, tests[i].expectErr, errors.Cause(err), "error TestCase # %d", i+1)

		assert.Equal(t, tests[i].expectResult, result, "result TestCase # %d", i+1)
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  example.Instance
		want   example.Instance
		err    error
		expect error
	}{
		{
			input:  example.Instance{ID: 1, Test: "tester"},
			want:   example.Instance{ID: 1, Test: "tester"},
			err:    nil,
			expect: nil,
		},
		{
			input:  example.Instance{ID: 2, Test: "tester"},
			want:   example.Instance{ID: 2, Test: "tester"},
			err:    fmt.Errorf("example unique"),
			expect: fmt.Errorf("example unique"),
		},
	}

	for i := range tests {
		if tests[i].err != nil {
			mock.ExpectExec("UPDATE example").WithArgs(tests[i].want.ID, tests[i].want.Test).WillReturnError(tests[i].err)
		} else {
			mock.ExpectExec("UPDATE example").WithArgs(tests[i].want.ID, tests[i].want.Test).WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err := repo.Update(&tests[i].input)

		assert.Equal(t, tests[i].expect, errors.Cause(err), "TestCase # %d", i+1)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := New(nil, &postgres.Postgres{
		DB: db,
	})

	tests := []struct {
		input  int
		want   int
		err    error
		expect error
	}{
		{
			input:  1,
			want:   1,
			err:    nil,
			expect: nil,
		},
		{
			input:  0,
			want:   0,
			err:    fmt.Errorf("example unique"),
			expect: fmt.Errorf("example unique"),
		},
	}

	for i := range tests {
		if tests[i].err != nil {
			mock.ExpectExec("DELETE FROM example").WithArgs(tests[i].want).WillReturnError(tests[i].err)
		} else {
			mock.ExpectExec("DELETE FROM example").WithArgs(tests[i].want).WillReturnResult(sqlmock.NewResult(1, 1))
		}

		err := repo.Delete(tests[i].input)

		assert.Equal(t, tests[i].expect, errors.Cause(err), "TestCase # %d", i+1)
	}
}
