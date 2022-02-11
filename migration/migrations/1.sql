CREATE TABLE IF NOT EXISTS events (
    client_time DateTime,
    device_id UUID,
    device_os String,
    session String,
    sequence Int64,
    event String,
    param_int Int16,
    param_str String,
    server_time DateTime,
    ip IPv4
) ENGINE = MergeTree() PRIMARY KEY sequence;