DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_class c
                   JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
                   WHERE n.nspname = 'public' AND c.relname = 'orders' AND c.relkind = 'r') THEN
        CREATE TABLE orders (
            id SERIAL PRIMARY KEY,
            payment_method VARCHAR(255) NOT NULL,
            tax_price DECIMAL(10,2) NOT NULL,
            shipping_price DECIMAL(10,2) NOT NULL,
            total_price DECIMAL(10,2) NOT NULL,
            time_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            time_now TIMESTAMP
        );
    END IF;
END $$;
