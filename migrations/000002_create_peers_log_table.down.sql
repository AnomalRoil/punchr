BEGIN;
DROP TRIGGER IF EXISTS on_peer_update ON peers;
DROP FUNCTION IF EXISTS insert_peer_log;
DROP TABLE IF EXISTS peer_logs;
END;
