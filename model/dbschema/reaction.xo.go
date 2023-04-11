package dbschema

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Reaction represents a row from 'emoine.reaction'.
type Reaction struct {
	ID        UUID      `json:"id"`         // id
	UserID    string    `json:"user_id"`    // user_id
	MeetingID UUID      `json:"meeting_id"` // meeting_id
	StampID   UUID      `json:"stamp_id"`   // stamp_id
	CreatedAt time.Time `json:"created_at"` // created_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Reaction exists in the database.
func (r *Reaction) Exists() bool {
	return r._exists
}

// Deleted returns true when the Reaction has been marked for deletion from
// the database.
func (r *Reaction) Deleted() bool {
	return r._deleted
}

// Insert inserts the Reaction to the database.
func (r *Reaction) Insert(ctx context.Context, db DB) error {
	switch {
	case r._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case r._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO emoine.reaction (` +
		`id, user_id, meeting_id, stamp_id, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, r.ID, r.UserID, r.MeetingID, r.StampID, r.CreatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, r.ID, r.UserID, r.MeetingID, r.StampID, r.CreatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	r._exists = true
	return nil
}

// Update updates a Reaction in the database.
func (r *Reaction) Update(ctx context.Context, db DB) error {
	switch {
	case !r._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case r._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE emoine.reaction SET ` +
		`user_id = ?, meeting_id = ?, stamp_id = ?, created_at = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, r.UserID, r.MeetingID, r.StampID, r.CreatedAt, r.ID)
	if _, err := db.ExecContext(ctx, sqlstr, r.UserID, r.MeetingID, r.StampID, r.CreatedAt, r.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Reaction to the database.
func (r *Reaction) Save(ctx context.Context, db DB) error {
	if r.Exists() {
		return r.Update(ctx, db)
	}
	return r.Insert(ctx, db)
}

// Upsert performs an upsert for Reaction.
func (r *Reaction) Upsert(ctx context.Context, db DB) error {
	switch {
	case r._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO emoine.reaction (` +
		`id, user_id, meeting_id, stamp_id, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id = VALUES(id), user_id = VALUES(user_id), meeting_id = VALUES(meeting_id), stamp_id = VALUES(stamp_id), created_at = VALUES(created_at)`
	// run
	logf(sqlstr, r.ID, r.UserID, r.MeetingID, r.StampID, r.CreatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, r.ID, r.UserID, r.MeetingID, r.StampID, r.CreatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	r._exists = true
	return nil
}

// Delete deletes the Reaction from the database.
func (r *Reaction) Delete(ctx context.Context, db DB) error {
	switch {
	case !r._exists: // doesn't exist
		return nil
	case r._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM emoine.reaction ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, r.ID)
	if _, err := db.ExecContext(ctx, sqlstr, r.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	r._deleted = true
	return nil
}

// ReactionByID retrieves a row from 'emoine.reaction' as a Reaction.
//
// Generated from index 'reaction_id_pkey'.
func ReactionByID(ctx context.Context, db DB, id UUID) (*Reaction, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, meeting_id, stamp_id, created_at ` +
		`FROM emoine.reaction ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	r := Reaction{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&r.ID, &r.UserID, &r.MeetingID, &r.StampID, &r.CreatedAt); err != nil {
		return nil, logerror(err)
	}
	return &r, nil
}