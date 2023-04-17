package dbschema

import (
	"context"
)

func SelectMeetingAll(ctx context.Context, db DB, limit int, offset int) ([]Meeting, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, video_id, title, thumbnail, started_at, ended_at, description ` +
		`FROM emoine.meeting ` +
		`LIMIT ? OFFSET ?`
	// run
	logf(sqlstr, limit, offset)

	rows, err := db.QueryContext(ctx, sqlstr, limit, offset)
	if err != nil {
		return nil, logerror(err)
	}
	m := []Meeting{}
	for rows.Next() {
		var m1 Meeting
		err = rows.Scan(
			&m1.ID,
			&m1.VideoID,
			&m1.Title,
			&m1.Thumbnail,
			&m1.StartedAt,
			&m1.EndedAt,
			&m1.Description,
		)
		if err != nil {
			return nil, logerror(err)
		}
		m = append(m, m1)
	}
	return m, nil
}

func CountMeeting(ctx context.Context, db DB) (int, error) {
	// query
	const sqlstr = `SELECT ` +
		`COUNT(*) ` +
		`FROM emoine.meeting`
	// run
	logf(sqlstr)

	var count int
	err := db.QueryRowContext(ctx, sqlstr).Scan(&count)
	if err != nil {
		return 0, logerror(err)
	}
	return count, nil
}
