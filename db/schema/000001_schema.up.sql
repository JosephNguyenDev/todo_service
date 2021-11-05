BEGIN;

CREATE FUNCTION update_updated_at_column() RETURNS trigger
        LANGUAGE plpgsql
        AS $$
    BEGIN
        NEW.updated_at = NOW();
        RETURN NEW;
    END;
$$;

CREATE TABLE IF NOT EXISTS todo_list (
    id SERIAL PRIMARY KEY,
    todo_list_name VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- update updated_at trigger
CREATE TRIGGER todo_list_updated_at_modtime BEFORE UPDATE ON todo_list FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    todo_name VARCHAR(255),
    todo_list_name VARCHAR(255) REFERENCES todo_list (todo_list_name) ON UPDATE CASCADE ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    content TEXT NOT NULL,
    done BOOLEAN NOT NULL
);

-- update updated_at trigger
CREATE TRIGGER todo_updated_at_modtime BEFORE UPDATE ON todo FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

COMMIT;