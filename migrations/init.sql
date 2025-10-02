CREATE TABLE IF NOT EXISTS fhr_records (
    id SERIAL PRIMARY KEY,
    time_recorded TIMESTAMP WITH TIME ZONE NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    is_retrieved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (time_recorded)
);

CREATE TABLE IF NOT EXISTS uc_records (
    id SERIAL PRIMARY KEY,
    time_recorded TIMESTAMP WITH TIME ZONE NOT NULL,
    value DOUBLE PRECISION NOT NULL,
    is_retrieved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (time_recorded)
);

CREATE INDEX idx_fhr_records_not_retrieved ON fhr_records(time_recorded) WHERE NOT is_retrieved;
CREATE INDEX idx_uc_records_not_retrieved ON uc_records(time_recorded) WHERE NOT is_retrieved;