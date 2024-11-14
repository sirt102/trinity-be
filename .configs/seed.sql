-- Switch to the trinity database
\c trinitydb

-- Enable the UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the roles table
CREATE TABLE IF NOT EXISTS roles (
    role_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) UNIQUE NOT NULL,
    admin_permission BOOLEAN DEFAULT FALSE
);

-- Create the subscriptions table
CREATE TABLE IF NOT EXISTS subscriptions (
    subscription_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    level INT NOT NULL,
    description TEXT
);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id UUID REFERENCES roles(role_id)
);

-- Create the campaigns table
CREATE TABLE IF NOT EXISTS campaigns (
    campaign_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    max_redemptions INT NOT NULL,
    available INT NOT NULL
);

-- Create the vouchers table
CREATE TABLE IF NOT EXISTS vouchers (
    voucher_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    campaign_id UUID REFERENCES campaigns(campaign_id),
    discount_rate DECIMAL(5, 2) DEFAULT 0.00,
    expiry_date DATE NOT NULL,
    is_valid BOOLEAN DEFAULT TRUE
);

-- Create the user_vouchers table to link vouchers with specific users
CREATE TABLE IF NOT EXISTS user_vouchers (
    user_voucher_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    voucher_id UUID REFERENCES vouchers(voucher_id),
    user_id UUID REFERENCES users(user_id),
    redeemed_at TIMESTAMP,
    UNIQUE (user_id, voucher_id)
);

-- Create the transactions table
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(user_id),
    voucher_id UUID REFERENCES vouchers(voucher_id),
    subscription_id UUID REFERENCES subscriptions(subscription_id),
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    amount DECIMAL(10, 2) NOT NULL,
    final_amount DECIMAL(10, 2) NOT NULL,
    status_payment VARCHAR(20) NOT NULL
);

-- Create the user_subscriptions table to track user subscriptions over time
CREATE TABLE IF NOT EXISTS user_subscriptions (
    user_subscription_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    subscription_id UUID REFERENCES subscriptions(subscription_id) ON DELETE CASCADE,
    subscribed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_date TIMESTAMP
);

-- Insert sample data into roles
INSERT INTO roles (role_id, name, admin_permission) VALUES
(uuid_generate_v4(), 'admin', TRUE),
(uuid_generate_v4(), 'user', FALSE);

-- Insert sample data into subscriptions
INSERT INTO subscriptions (subscription_id, name, price, level, description) VALUES
(uuid_generate_v4(), 'bronze', 60.00, 0, 'Bronze plan with basic access'),
(uuid_generate_v4(), 'silver', 100.00, 1, 'Silver plan with additional features');

-- Insert sample data into users
INSERT INTO users (user_id, user_name, email, role_id) VALUES
(uuid_generate_v4(), 'alice', 'alice@example.com', (SELECT role_id FROM roles WHERE name = 'admin')),
(uuid_generate_v4(), 'bob', 'bob@example.com', (SELECT role_id FROM roles WHERE name = 'user')),
(uuid_generate_v4(), 'charlie', 'charlie@example.com', (SELECT role_id FROM roles WHERE name = 'user'));

-- Insert sample data into campaigns
INSERT INTO campaigns (campaign_id, name, start_date, end_date, max_redemptions, available) VALUES
(uuid_generate_v4(), 'end_of_year_sale', '2024-11-11', '2024-12-12', 100, 100),
(uuid_generate_v4(), 'new_member', '2024-01-01', '2025-01-31', 200, 200);

-- Insert sample data into vouchers (general voucher information)
INSERT INTO vouchers (voucher_id, campaign_id, discount_rate, expiry_date) VALUES
(uuid_generate_v4(), (SELECT campaign_id FROM campaigns WHERE name='new_member'), 0.20, '2025-01-31'),
(uuid_generate_v4(), (SELECT campaign_id FROM campaigns WHERE name='end_of_year_sale'), 0.30, '2024-12-12');

-- Insert sample data into user_vouchers to assign vouchers to users
INSERT INTO user_vouchers (user_voucher_id, voucher_id, user_id) VALUES
(uuid_generate_v4(), (SELECT voucher_id FROM vouchers WHERE discount_rate=0.20), (SELECT user_id FROM users WHERE email='alice@example.com')),
(uuid_generate_v4(), (SELECT voucher_id FROM vouchers WHERE discount_rate=0.30), (SELECT user_id FROM users WHERE email='bob@example.com'));

-- Insert sample data into transactions with updated amounts and correct voucher application
-- INSERT INTO transactions (transaction_id, user_id, voucher_id, subscription_id, transaction_date, amount, final_amount, status_payment) VALUES
-- (uuid_generate_v4(), (SELECT user_id FROM users WHERE email='alice@example.com'), 
--     (SELECT voucher_id FROM vouchers WHERE discount_rate=0.20), 
--     (SELECT subscription_id FROM subscriptions WHERE name='silver'), CURRENT_TIMESTAMP, 100.00, 0, 'pending'),
-- (uuid_generate_v4(), (SELECT user_id FROM users WHERE email='bob@example.com'), 
--     (SELECT voucher_id FROM vouchers WHERE discount_rate=0.30), 
--     (SELECT subscription_id FROM subscriptions WHERE name='silver'), CURRENT_TIMESTAMP, 100.00, 0, 'pending');
