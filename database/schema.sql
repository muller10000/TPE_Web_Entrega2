CREATE TABLE movies (
    id SERIAL PRIMARY KEY,          -- identificador único
    title TEXT NOT NULL,            -- título de la película
    director TEXT,                  -- director de la película
    year INT,                        -- año de estreno
    genre TEXT,                     -- género
    rating NUMERIC(3,1),            -- calificación (ej: 8.5)
    created_at TIMESTAMP NOT NULL DEFAULT now()  -- fecha de creación del registro
);