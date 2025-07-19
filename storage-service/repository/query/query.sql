-- name: HealthCheck :one
select (current_timestamp - pg_postmaster_start_time())::interval;