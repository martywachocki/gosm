CREATE TABLE services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    protocol TEXT NOT NULL,
    host TEXT NOT NULL,
    port TEXT,
    uptime_start INTEGER
);