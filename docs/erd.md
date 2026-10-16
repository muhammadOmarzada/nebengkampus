
# NebengKampus — Database ERD

## Core Entities

### users
- id (PK)
- name
- email (unique)
- password_hash
- created_at

### trips
- id (PK)
- driver_id (FK -> users.id)
- origin_name
- origin_location (PostGIS, direncanakan)
- destination_name
- destination_location (PostGIS, direncanakan)
- departure_time
- total_seats
- available_seats
- cost_per_seat
- status
- created_at

### bookings
- id (PK)
- trip_id (FK -> trips.id)
- passenger_id (FK -> users.id)
- seats_booked
- status
- created_at

## Relationships

- One user can create many trips.
- One user can make many bookings.
- One trip can have many bookings.
- Each booking belongs to one trip and one passenger.

## Business Rules

- A driver cannot book their own trip.
- A booking cannot exceed available seats.
- A trip cannot accept bookings after departure.
- Cancelled bookings must follow defined seat-restoration rules.
- Seat availability must remain correct when multiple users book concurrently.

## Future Scope

- Midtrans Sandbox integration.
- Payment records and payment webhook handling.
- Wallet and ledger, if supported by the eventual payment flow.
