BEGIN;

DROP TABLE IF EXISTS todo;
DROP TABLE IF EXISTS todo_list;
DROP FUNCTION IF EXISTS update_updated_at_column;
DROP TRIGGER IF EXISTS todo_updated_at_modtime ON todo;
DROP TRIGGER IF EXISTS todo_list_updated_at_modtime ON todo_list;

COMMIT;