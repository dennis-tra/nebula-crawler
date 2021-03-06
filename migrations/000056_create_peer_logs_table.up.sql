BEGIN;

CREATE TABLE peer_logs
(
    id         INT GENERATED ALWAYS AS IDENTITY,
    peer_id    INT         NOT NULL,
    field      TEXT        NOT NULL,
    old        TEXT        NOT NULL,
    new        TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_peer_logs_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION insert_peer_log()
    RETURNS TRIGGER AS
$$
BEGIN
    IF OLD.agent_version_id != NEW.agent_version_id THEN
        INSERT INTO peer_logs (peer_id, field, old, new, created_at)
        VALUES (NEW.id, 'agent_version_id', OLD.agent_version_id, NEW.agent_version_id, NOW());
        RETURN NEW;
    END IF;

    IF OLD.protocols_set_id != NEW.protocols_set_id THEN
        INSERT INTO peer_logs (peer_id, field, old, new, created_at)
        VALUES (NEW.id, 'protocols_set_id', OLD.protocols_set_id, NEW.protocols_set_id, NOW());
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER on_peer_update
    BEFORE UPDATE
    ON peers
    FOR EACH ROW
EXECUTE PROCEDURE insert_peer_log();

COMMIT;
