CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  image VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  description TEXT,
  rating INT NOT NULL,
  num_reviews INT NOT NULL DEFAULT 0,
  price DECIMAL(10,2) NOT NULL,
  count_in_stock INT NOT NULL,
  time_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  time_now TIMESTAMP
);
