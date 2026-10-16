
-- Migration: 001_init.sql
-- Project: NebengKampus

-- Enable geospatial support
CREATE EXTENSION IF NOT EXISTS postgis;


-- 1. USERS
CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX users_email_unique_idx
    ON users (LOWER(email));


-- 2. TRIPS
CREATE TABLE trips (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    driver_id BIGINT NOT NULL
        REFERENCES users(id) ON DELETE RESTRICT,

    origin_name VARCHAR(255) NOT NULL,
    origin_location GEOGRAPHY(POINT, 4326) NOT NULL,

    destination_name VARCHAR(255) NOT NULL,
    destination_location GEOGRAPHY(POINT, 4326) NOT NULL,

    departure_time TIMESTAMPTZ NOT NULL,

    total_seats INTEGER NOT NULL
        CHECK (total_seats > 0),

    available_seats INTEGER NOT NULL,

    cost_per_seat NUMERIC(12, 2) NOT NULL
        CHECK (cost_per_seat >= 0),

    status VARCHAR(20) NOT NULL DEFAULT 'open'
        CHECK (status IN (
            'open', 'full', 'cancelled', 'completed'
        )),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT trips_available_seats_check
        CHECK (
            available_seats >= 0
            AND available_seats <= total_seats
        )
);

CREATE INDEX trips_driver_id_idx
    ON trips (driver_id);

CREATE INDEX trips_departure_time_idx
    ON trips (departure_time);

CREATE INDEX trips_origin_location_idx
    ON trips USING GIST (origin_location);

CREATE INDEX trips_destination_location_idx
    ON trips USING GIST (destination_location);


-- 3. BOOKINGS
CREATE TABLE bookings (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    trip_id BIGINT NOT NULL
        REFERENCES trips(id) ON DELETE RESTRICT,

    passenger_id BIGINT NOT NULL
        REFERENCES users(id) ON DELETE RESTRICT,

    seats_booked INTEGER NOT NULL
        CHECK (seats_booked > 0),

    status VARCHAR(20) NOT NULL DEFAULT 'pending'
        CHECK (status IN (
            'pending', 'confirmed', 'cancelled'
        )),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT bookings_one_per_passenger_per_trip
        UNIQUE (trip_id, passenger_id)
);

CREATE INDEX bookings_passenger_id_idx
    ON bookings (passenger_id);
