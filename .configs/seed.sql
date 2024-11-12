-- Switch to the trinity database
\c trinitydb

-- Enable the UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the Users table
CREATE TABLE IF NOT EXISTS Users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the Campaigns table
CREATE TABLE IF NOT EXISTS Campaigns (
    campaign_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    discount_rate DECIMAL(5, 2) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    max_redemptions INT NOT NULL
);

-- Create the Vouchers table
CREATE TABLE IF NOT EXISTS Vouchers (
    voucher_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    campaign_id UUID REFERENCES Campaigns(campaign_id),
    user_id UUID REFERENCES Users(user_id),
    is_valid BOOLEAN DEFAULT TRUE,
    expiry_date DATE,
    redeemed_at TIMESTAMP,
    UNIQUE (user_id, campaign_id)
);

-- Create the Transactions table
CREATE TABLE IF NOT EXISTS Transactions (
    transaction_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES Users(user_id),
    voucher_id UUID REFERENCES Vouchers(voucher_id),
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    amount DECIMAL(10, 2) NOT NULL,
    final_amount DECIMAL(10, 2) NOT NULL
);

-- Create the Notifications table
CREATE TABLE IF NOT EXISTS Notifications (
    notification_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES Users(user_id),
    campaign_id UUID REFERENCES Campaigns(campaign_id),
    message TEXT,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the CampaignAnalytics table
CREATE TABLE IF NOT EXISTS CampaignAnalytics (
    campaign_id UUID PRIMARY KEY REFERENCES Campaigns(campaign_id),
    issued_vouchers INT,
    redeemed_vouchers INT,
    conversion_rate DECIMAL(5, 2),
    total_discount_amount DECIMAL(10, 2)
);

-- Insert sample data into Users
INSERT INTO Users (user_id, email) VALUES
(uuid_generate_v4(), 'alice@example.com'),
(uuid_generate_v4(), 'bob@example.com'),
(uuid_generate_v4(), 'charlie@example.com');

-- Insert sample data into Campaigns
INSERT INTO Campaigns (campaign_id, name, discount_rate, start_date, end_date, max_redemptions) VALUES
(uuid_generate_v4(), 'New Year Sale', 20.00, '2024-01-01', '2024-01-31', 1000),
(uuid_generate_v4(), 'Spring Discount', 15.00, '2024-03-01', '2024-03-31', 500),
(uuid_generate_v4(), 'Summer Special', 10.00, '2024-06-01', '2024-06-30', 300);

-- Insert sample data into Vouchers (assuming user and campaign IDs exist)
INSERT INTO Vouchers (voucher_id, campaign_id, user_id, is_valid, expiry_date) VALUES
(uuid_generate_v4(), (SELECT campaign_id FROM Campaigns WHERE name='New Year Sale'), (SELECT user_id FROM Users WHERE email='alice@example.com'), TRUE, '2024-01-31'),
(uuid_generate_v4(), (SELECT campaign_id FROM Campaigns WHERE name='Spring Discount'), (SELECT user_id FROM Users WHERE email='bob@example.com'), TRUE, '2024-03-31');

-- Insert sample data into Transactions
INSERT INTO Transactions (transaction_id, user_id, voucher_id, amount, final_amount) VALUES
(uuid_generate_v4(), (SELECT user_id FROM Users WHERE email='alice@example.com'), (SELECT voucher_id FROM Vouchers WHERE user_id=(SELECT user_id FROM Users WHERE email='alice@example.com') AND is_valid=TRUE), 100.00, 80.00),
(uuid_generate_v4(), (SELECT user_id FROM Users WHERE email='bob@example.com'), (SELECT voucher_id FROM Vouchers WHERE user_id=(SELECT user_id FROM Users WHERE email='bob@example.com') AND is_valid=TRUE), 50.00, 42.50);

-- Insert sample data into Notifications
INSERT INTO Notifications (notification_id, user_id, campaign_id, message) VALUES
(uuid_generate_v4(), (SELECT user_id FROM Users WHERE email='alice@example.com'), (SELECT campaign_id FROM Campaigns WHERE name='New Year Sale'), 'Your voucher is ready for New Year Sale!'),
(uuid_generate_v4(), (SELECT user_id FROM Users WHERE email='bob@example.com'), (SELECT campaign_id FROM Campaigns WHERE name='Spring Discount'), 'Spring Discount is live now!');

-- Insert sample data into CampaignAnalytics
INSERT INTO CampaignAnalytics (campaign_id, issued_vouchers, redeemed_vouchers, conversion_rate, total_discount_amount) VALUES
((SELECT campaign_id FROM Campaigns WHERE name='New Year Sale'), 500, 250, 50.00, 5000.00),
((SELECT campaign_id FROM Campaigns WHERE name='Spring Discount'), 300, 150, 50.00, 2250.00);
