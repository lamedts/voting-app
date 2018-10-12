DROP TABLE IF EXISTS votes;

CREATE TABLE IF NOT EXISTS votes (
	id				serial						PRIMARY KEY,
    voter_id        int                         NOT NULL,
    vote          VARCHAR(255),
	name			VARCHAR(255),
    gender			VARCHAR(255),
    age 			VARCHAR(255),
    nationality     VARCHAR(255),
    
    created_at				TIMESTAMP WITH TIME ZONE	NOT NULL DEFAULT now(),
	updated_at				TIMESTAMP WITH TIME ZONE	NOT NULL DEFAULT now()
);
