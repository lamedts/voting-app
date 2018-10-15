package datastore

type VoteResult struct {
	Vote  string `db:"vote"`
	Count int32  `db:"count"`
}

type Votes struct {
	VoterID int32  `db:"voter_id"`
	Vote    string `db:"vote"`
	// Id             int       `db:"id"`
	// Name           string    `db:"name"`
	// Gender         string    `db:"vgenderote"`
	// Age            string    `db:"age"`
	// Nationality    string    `db:"nationality"`
	// LastUpdateTime time.Time `db:"updated_at" json:"-"`
	// CreationTime   time.Time `db:"created_at" json:"transaction_time"`
}
