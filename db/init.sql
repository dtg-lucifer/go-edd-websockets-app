SELECT 'CREATE DATABASE edd_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'edd_db') \gexec
