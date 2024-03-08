CREATE TABLE IF NOT EXISTS job (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    version INT NOT NULL, -- Optimistic locking
    name VARCHAR(255), -- Name of the job
    status VARCHAR(255) NULL, -- Status of the job
    requester_id INT,
    destination_id INT,
    equipment_id INT
);
```