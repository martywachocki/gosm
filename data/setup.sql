CREATE TABLE services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    protocol TEXT NOT NULL,
    host TEXT NOT NULL,
    port TEXT
);

INSERT INTO services (name, protocol, host) VALUES('Detroit Web Solutions Website', 'https', 'https://detroitws.com/');
INSERT INTO services (name, protocol, host) VALUES('Better Health Cafe Website', 'http', 'http://betterhealthcafe.com/');
INSERT INTO services (name, protocol, host) VALUES('Propel FSM API', 'https', 'https://api.propelfsm.com/');