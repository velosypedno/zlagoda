-- Zlagoda Comprehensive Sample Data Generation
-- Designed specifically for individual SQL queries testing
-- This script ensures all Vlad, Arthur, and Oleksii queries return meaningful results

BEGIN;

-- Clear existing data in dependency order
DELETE FROM sale;
DELETE FROM receipt;
DELETE FROM store_product;
DELETE FROM customer_card;
DELETE FROM employee;
DELETE FROM product;
DELETE FROM category;

-- Reset sequences
ALTER SEQUENCE category_category_id_seq RESTART WITH 1;
ALTER SEQUENCE product_product_id_seq RESTART WITH 1;

-- ===== CATEGORIES =====
-- Create 10 diverse categories for comprehensive testing
INSERT INTO category (category_name) VALUES
('Electronics'),           -- 1
('Clothing'),             -- 2
('Food & Beverages'),     -- 3
('Home & Garden'),        -- 4
('Books & Media'),        -- 5
('Sports & Outdoors'),    -- 6
('Health & Beauty'),      -- 7
('Toys & Games'),         -- 8
('Automotive'),           -- 9
('Office Supplies');      -- 10

-- ===== PRODUCTS =====
-- 5 products per category (50 total) for comprehensive coverage
INSERT INTO product (product_name, characteristics, category_id) VALUES
-- Electronics (Category 1) - Will be top sellers for Vlad1
('Samsung Galaxy S23', 'Latest smartphone with 256GB storage, 5G ready', 1),
('Sony Headphones WH-1000XM4', 'Noise-canceling wireless headphones', 1),
('MacBook Pro 14-inch', 'M3 chip, 16GB RAM, 512GB SSD', 1),
('Nintendo Switch OLED', 'Gaming console with 7-inch OLED display', 1),
('iPad Air 5th Gen', 'Tablet with M1 chip, 10.9-inch display', 1),

-- Clothing (Category 2)
('Nike Air Max 270', 'Running shoes with air cushioning', 2),
('Levi''s 501 Jeans', 'Classic straight fit denim jeans', 2),
('Adidas Hoodie', 'Cotton blend pullover hoodie', 2),
('Ray-Ban Sunglasses', 'Classic aviator style sunglasses', 2),
('North Face Jacket', 'Waterproof outdoor jacket', 2),

-- Food & Beverages (Category 3)
('Starbucks Coffee Beans', 'Premium arabica coffee beans 1kg', 3),
('Lindt Dark Chocolate', 'Swiss dark chocolate 70% cocoa', 3),
('Evian Water 6-pack', 'Natural spring water bottles', 3),
('Organic Honey', 'Raw wildflower honey 500g', 3),
('Green Tea Selection', 'Premium loose leaf tea variety pack', 3),

-- Home & Garden (Category 4)
('Philips LED Lamp', 'Energy efficient smart LED bulb', 4),
('Garden Hose 25ft', 'Heavy duty flexible garden hose', 4),
('Memory Foam Pillow', 'Ergonomic sleep pillow with bamboo cover', 4),
('Ceramic Plant Pot', 'Decorative indoor plant container', 4),
('Kitchen Knife Set', 'Professional chef knife collection', 4),

-- Books & Media (Category 5)
('The Great Gatsby', 'Classic American literature novel', 5),
('Programming Guide', 'Complete guide to modern programming', 5),
('Marvel Movies Collection', 'Blu-ray box set of Marvel films', 5),
('Cookbook Essentials', 'International cuisine recipe collection', 5),
('History Encyclopedia', 'Comprehensive world history reference', 5),

-- Sports & Outdoors (Category 6)
('Yoga Mat Premium', 'Non-slip exercise mat with alignment lines', 6),
('Tennis Racket Pro', 'Professional grade tennis racket', 6),
('Camping Tent 4-person', 'Waterproof family camping tent', 6),
('Bicycle Helmet', 'Safety helmet with ventilation system', 6),
('Hiking Backpack', 'Lightweight trekking backpack 40L', 6),

-- Health & Beauty (Category 7)
('Vitamin C Serum', 'Anti-aging facial serum with hyaluronic acid', 7),
('Moisturizing Cream', 'Daily face moisturizer for all skin types', 7),
('Electric Toothbrush', 'Sonic cleaning with multiple brush heads', 7),
('Essential Oils Set', 'Aromatherapy essential oils collection', 7),
('Protein Powder', 'Whey protein supplement vanilla flavor', 7),

-- Toys & Games (Category 8)
('LEGO Creator Set', 'Building blocks construction set 500 pieces', 8),
('Board Game Strategy', 'Family strategy board game for 2-6 players', 8),
('Puzzle 1000 pieces', 'Scenic landscape jigsaw puzzle', 8),
('Remote Control Car', 'Electric RC car with rechargeable battery', 8),
('Educational Robot', 'Programming robot for kids STEM learning', 8),

-- Automotive (Category 9)
('Motor Oil 5W-30', 'Synthetic motor oil 4L container', 9),
('Car Phone Mount', 'Dashboard smartphone holder with suction', 9),
('Tire Pressure Gauge', 'Digital tire pressure monitoring tool', 9),
('Car Air Freshener', 'Long-lasting vanilla scent air freshener', 9),
('Jump Starter Kit', 'Portable car battery jump starter', 9),

-- Office Supplies (Category 10)
('Wireless Mouse', 'Ergonomic wireless computer mouse', 10),
('Notebook Set', 'Lined notebooks pack of 5', 10),
('Desk Organizer', 'Bamboo desktop storage organizer', 10),
('Printer Paper A4', 'High quality copy paper 500 sheets', 10),
('Standing Desk', 'Adjustable height standing desk converter', 10);

-- ===== EMPLOYEES =====
-- Create managers and cashiers with login credentials
INSERT INTO employee (employee_id, login, hashed_password, empl_surname, empl_name, empl_patronymic, empl_role, salary, date_of_birth, date_of_start, phone_number, city, street, zip_code) VALUES
-- Managers
('MGR0000001', 'manager1', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Johnson', 'Alice', 'Marie', 'Manager', 45000.00, '1985-03-15', '2020-01-01', '+380501234567', 'Kyiv', 'Independence Square 1', '01001'),
('MGR0000002', 'manager2', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Smith', 'Robert', 'James', 'Manager', 47000.00, '1982-07-22', '2019-06-01', '+380501234568', 'Kyiv', 'Khreshchatyk 2', '01002'),
('MGR0000003', 'manager3', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Williams', 'Emma', 'Grace', 'Manager', 46000.00, '1987-11-10', '2021-03-01', '+380501234569', 'Kyiv', 'Maidan Nezalezhnosti 3', '01003'),

-- Cashiers (some will never sell promotional products for Vlad2)
('CSH0000001', 'cashier1', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Brown', 'Michael', 'David', 'Cashier', 28000.00, '1995-01-12', '2022-01-15', '+380501234570', 'Kyiv', 'Prorizna 10', '01010'),
('CSH0000002', 'cashier2', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Davis', 'Sarah', 'Lynn', 'Cashier', 29000.00, '1993-05-20', '2022-02-01', '+380501234571', 'Kyiv', 'Velyka Vasylkivska 15', '01015'),
('CSH0000003', 'cashier3', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Miller', 'Kevin', 'Paul', 'Cashier', 27500.00, '1996-09-08', '2022-03-01', '+380501234572', 'Kyiv', 'Shevchenko Blvd 20', '01020'),
('CSH0000004', 'cashier4', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Wilson', 'Jessica', 'Rose', 'Cashier', 28500.00, '1994-12-03', '2022-04-01', '+380501234573', 'Kyiv', 'Bohdana Khmelnytskoho 25', '01025'),
('CSH0000005', 'cashier5', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Moore', 'Daniel', 'Chris', 'Cashier', 27000.00, '1997-04-15', '2022-05-01', '+380501234574', 'Kyiv', 'Taras Shevchenko 30', '01030'),
('CSH0000006', 'cashier6', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Taylor', 'Lisa', 'Anne', 'Cashier', 28200.00, '1995-08-27', '2022-06-01', '+380501234575', 'Kyiv', 'Volodymyrska 35', '01035'),
('CSH0000007', 'cashier7', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Anderson', 'Mark', 'Thomas', 'Cashier', 29200.00, '1992-06-18', '2022-07-01', '+380501234576', 'Kyiv', 'Horodetskyi 40', '01040');

-- ===== CUSTOMER CARDS =====
-- Create diverse discount levels for comprehensive testing
INSERT INTO customer_card (card_number, cust_surname, cust_name, cust_patronymic, phone_number, city, street, zip_code, percent) VALUES
-- VIP customers with very high discounts (30-45%) - for Oleksii2 (will buy from all categories)
('CRD0000000001', 'Ivanov', 'Volodymyr', 'Petrovych', '+380671111111', 'Kyiv', 'Hrushevskoho 1', '01001', 45),
('CRD0000000002', 'Petrenko', 'Oksana', 'Ivanivna', '+380672222222', 'Kyiv', 'Sichovykh Striltsiv 2', '01002', 35),
('CRD0000000003', 'Kovalenko', 'Andriy', 'Mykolayovych', '+380673333333', 'Kyiv', 'Lesi Ukrainky 3', '01003', 30),

-- High discount customers (15-25%) - for Oleksii1 testing
('CRD0000000004', 'Shevchenko', 'Mariya', 'Oleksandrivna', '+380674444444', 'Kyiv', 'Bohdana Khmelnytskoho 4', '01004', 25),
('CRD0000000005', 'Bondarenko', 'Igor', 'Sergiyovych', '+380675555555', 'Kyiv', 'Chervonoarmiyska 5', '01005', 20),
('CRD0000000006', 'Lysenko', 'Tetyana', 'Viktorivna', '+380676666666', 'Kyiv', 'Pushkinska 6', '01006', 18),
('CRD0000000007', 'Moroz', 'Viktor', 'Dmytrovych', '+380677777777', 'Kyiv', 'Tolstoho 7', '01007', 16),
('CRD0000000008', 'Tkachenko', 'Nataliya', 'Romanivna', '+380678888888', 'Kyiv', 'Antonovycha 8', '01008', 15),

-- Medium discount customers (8-12%)
('CRD0000000009', 'Savchenko', 'Oleksiy', 'Volodymyrovych', '+380679999999', 'Kyiv', 'Laboratorna 9', '01009', 12),
('CRD0000000010', 'Kravchenko', 'Yuliya', 'Anatoliyivna', '+380670000000', 'Kyiv', 'Pylypa Orlyka 10', '01010', 10),
('CRD0000000011', 'Nazarenko', 'Roman', 'Bogdanovych', '+380671111110', 'Kyiv', 'Mykhaylivska 11', '01011', 9),
('CRD0000000012', 'Rudenko', 'Svitlana', 'Yaroslavivna', '+380672222220', 'Kyiv', 'Striletska 12', '01012', 8),

-- Low discount customers (3-7%)
('CRD0000000013', 'Marchenko', 'Dmitro', 'Vasylovych', '+380673333330', 'Kyiv', 'Bankova 13', '01013', 7),
('CRD0000000014', 'Kuznetsova', 'Anna', 'Mikhaylivna', '+380674444440', 'Kyiv', 'Institutska 14', '01014', 5),
('CRD0000000015', 'Popov', 'Sergiy', 'Oleksandrovych', '+380675555550', 'Kyiv', 'Sofiyivska 15', '01015', 4),
('CRD0000000016', 'Melnyk', 'Kateryna', 'Petrivna', '+380676666660', 'Kyiv', 'Mykolayivska 16', '01016', 3),
('CRD0000000017', 'Koval', 'Mykola', 'Ivanovych', '+380677777770', 'Kyiv', 'Desyatynna 17', '01017', 6),
('CRD0000000018', 'Polishchuk', 'Yelena', 'Stanislavivna', '+380678888880', 'Kyiv', 'Yaroslaviv Val 18', '01018', 5);

-- ===== STORE PRODUCTS =====
-- Create comprehensive store products with both promotional and regular items
INSERT INTO store_product (upc, upc_prom, product_id, selling_price, products_number, promotional_product) VALUES
-- Electronics - Will be top sellers (for Vlad1)
('100000000001', NULL, 1, 899.99, 25, FALSE),  -- Samsung Galaxy S23
('100000000002', NULL, 2, 349.99, 30, FALSE),  -- Sony Headphones
('100000000003', NULL, 3, 1999.99, 8, FALSE),  -- MacBook Pro
('100000000004', NULL, 4, 349.99, 20, FALSE),  -- Nintendo Switch
('100000000005', NULL, 5, 649.99, 15, FALSE),  -- iPad Air

-- Electronics Promotional products (80% of regular price)
('100000000006', '100000000001', 1, 719.99, 25, TRUE),   -- Samsung promo
('100000000007', '100000000002', 2, 279.99, 30, TRUE),   -- Sony promo

-- Clothing
('200000000001', NULL, 6, 129.99, 40, FALSE),  -- Nike Air Max
('200000000002', NULL, 7, 79.99, 35, FALSE),   -- Levi's Jeans
('200000000003', NULL, 8, 49.99, 50, FALSE),   -- Adidas Hoodie
('200000000004', NULL, 9, 159.99, 25, FALSE),  -- Ray-Ban
('200000000005', NULL, 10, 299.99, 15, FALSE), -- North Face

-- Food & Beverages
('300000000001', NULL, 11, 24.99, 60, FALSE),  -- Coffee Beans
('300000000002', NULL, 12, 12.99, 80, FALSE),  -- Chocolate
('300000000003', NULL, 13, 8.99, 100, FALSE),  -- Water
('300000000004', NULL, 14, 18.99, 45, FALSE),  -- Honey
('300000000005', NULL, 15, 22.99, 30, FALSE),  -- Tea

-- Home & Garden
('400000000001', NULL, 16, 19.99, 70, FALSE),  -- LED Lamp
('400000000002', NULL, 17, 39.99, 25, FALSE),  -- Garden Hose
('400000000003', NULL, 18, 59.99, 35, FALSE),  -- Pillow
('400000000004', NULL, 19, 29.99, 40, FALSE),  -- Plant Pot
('400000000005', NULL, 20, 149.99, 20, FALSE), -- Knife Set

-- Books & Media
('500000000001', NULL, 21, 14.99, 50, FALSE),  -- Great Gatsby
('500000000002', NULL, 22, 39.99, 30, FALSE),  -- Programming Guide
('500000000003', NULL, 23, 99.99, 15, FALSE),  -- Marvel Collection
('500000000004', NULL, 24, 29.99, 25, FALSE),  -- Cookbook
('500000000005', NULL, 25, 49.99, 20, FALSE),  -- Encyclopedia

-- Sports & Outdoors
('600000000001', NULL, 26, 49.99, 40, FALSE),  -- Yoga Mat
('600000000002', NULL, 27, 199.99, 15, FALSE), -- Tennis Racket
('600000000003', NULL, 28, 299.99, 10, FALSE), -- Camping Tent
('600000000004', NULL, 29, 79.99, 25, FALSE),  -- Bicycle Helmet
('600000000005', NULL, 30, 149.99, 18, FALSE), -- Hiking Backpack

-- Health & Beauty
('700000000001', NULL, 31, 29.99, 60, FALSE),  -- Vitamin C Serum
('700000000002', NULL, 32, 24.99, 55, FALSE),  -- Moisturizer
('700000000003', NULL, 33, 89.99, 30, FALSE),  -- Electric Toothbrush
('700000000004', NULL, 34, 39.99, 35, FALSE),  -- Essential Oils
('700000000005', NULL, 35, 49.99, 40, FALSE),  -- Protein Powder

-- Toys & Games
('800000000001', NULL, 36, 79.99, 30, FALSE),  -- LEGO Set
('800000000002', NULL, 37, 44.99, 25, FALSE),  -- Board Game
('800000000003', NULL, 38, 19.99, 45, FALSE),  -- Puzzle
('800000000004', NULL, 39, 129.99, 20, FALSE), -- RC Car
('800000000005', NULL, 40, 199.99, 15, FALSE), -- Educational Robot

-- Automotive
('900000000001', NULL, 41, 34.99, 50, FALSE),  -- Motor Oil
('900000000002', NULL, 42, 24.99, 40, FALSE),  -- Phone Mount
('900000000003', NULL, 43, 19.99, 35, FALSE),  -- Pressure Gauge
('900000000004', NULL, 44, 6.99, 100, FALSE),  -- Air Freshener
('900000000005', NULL, 45, 89.99, 25, FALSE),  -- Jump Starter

-- Office Supplies
('101000000001', NULL, 46, 29.99, 60, FALSE),  -- Wireless Mouse
('101000000002', NULL, 47, 15.99, 80, FALSE),  -- Notebooks
('101000000003', NULL, 48, 49.99, 30, FALSE),  -- Desk Organizer
('101000000004', NULL, 49, 12.99, 100, FALSE), -- Printer Paper
('101000000005', NULL, 50, 299.99, 10, FALSE), -- Standing Desk

-- Products that will NEVER be sold (for Arthur2) - some from each category with stock but no sales
('100000000099', NULL, 3, 1999.99, 5, FALSE),  -- MacBook (unsold)
('200000000099', NULL, 10, 299.99, 8, FALSE),  -- North Face (unsold)
('300000000099', NULL, 15, 22.99, 12, FALSE),  -- Tea (unsold)
('400000000099', NULL, 20, 149.99, 6, FALSE),  -- Knife Set (unsold)
('500000000099', NULL, 25, 49.99, 10, FALSE),  -- Encyclopedia (unsold)
('600000000099', NULL, 30, 149.99, 7, FALSE),  -- Hiking Backpack (unsold)
('700000000099', NULL, 35, 49.99, 9, FALSE),   -- Protein Powder (unsold)
('800000000099', NULL, 40, 199.99, 4, FALSE),  -- Educational Robot (unsold)
('900000000099', NULL, 45, 89.99, 11, FALSE),  -- Jump Starter (unsold)
('101000000099', NULL, 50, 299.99, 3, FALSE);  -- Standing Desk (unsold)

-- ===== RECEIPTS =====
-- Create receipts across different time periods and employees
-- Recent receipts (last month) for Oleksii2 testing
INSERT INTO receipt (receipt_number, employee_id, card_number, print_date, sum_total, vat) VALUES
-- VIP customer receipts (recent month) - these customers will buy from ALL categories
('RCP0000001', 'CSH0000001', 'CRD0000000001', CURRENT_DATE - INTERVAL '5 days', 1249.98, 249.996),    -- Electronics
('RCP0000002', 'CSH0000002', 'CRD0000000001', CURRENT_DATE - INTERVAL '8 days', 209.98, 41.996),     -- Clothing
('RCP0000003', 'CSH0000003', 'CRD0000000001', CURRENT_DATE - INTERVAL '12 days', 37.98, 7.596),      -- Food
('RCP0000004', 'CSH0000004', 'CRD0000000001', CURRENT_DATE - INTERVAL '15 days', 59.98, 11.996),     -- Home
('RCP0000005', 'CSH0000005', 'CRD0000000001', CURRENT_DATE - INTERVAL '18 days', 44.98, 8.996),      -- Books
('RCP0000006', 'CSH0000006', 'CRD0000000001', CURRENT_DATE - INTERVAL '20 days', 249.98, 49.996),    -- Sports
('RCP0000007', 'CSH0000007', 'CRD0000000001', CURRENT_DATE - INTERVAL '22 days', 54.98, 10.996),     -- Beauty
('RCP0000008', 'CSH0000001', 'CRD0000000001', CURRENT_DATE - INTERVAL '25 days', 79.99, 15.998),     -- Toys
('RCP0000009', 'CSH0000002', 'CRD0000000001', CURRENT_DATE - INTERVAL '27 days', 34.99, 6.998),      -- Auto
('RCP0000010', 'CSH0000003', 'CRD0000000001', CURRENT_DATE - INTERVAL '29 days', 29.99, 5.998),      -- Office

-- Second VIP customer (all categories)
('RCP0000011', 'CSH0000004', 'CRD0000000002', CURRENT_DATE - INTERVAL '3 days', 699.99, 139.998),    -- Electronics
('RCP0000012', 'CSH0000005', 'CRD0000000002', CURRENT_DATE - INTERVAL '6 days', 129.99, 25.998),     -- Clothing
('RCP0000013', 'CSH0000006', 'CRD0000000002', CURRENT_DATE - INTERVAL '9 days', 24.99, 4.998),       -- Food
('RCP0000014', 'CSH0000007', 'CRD0000000002', CURRENT_DATE - INTERVAL '11 days', 39.99, 7.998),      -- Home
('RCP0000015', 'CSH0000001', 'CRD0000000002', CURRENT_DATE - INTERVAL '14 days', 14.99, 2.998),      -- Books
('RCP0000016', 'CSH0000002', 'CRD0000000002', CURRENT_DATE - INTERVAL '17 days', 49.99, 9.998),      -- Sports
('RCP0000017', 'CSH0000003', 'CRD0000000002', CURRENT_DATE - INTERVAL '19 days', 29.99, 5.998),      -- Beauty
('RCP0000018', 'CSH0000004', 'CRD0000000002', CURRENT_DATE - INTERVAL '23 days', 44.99, 8.998),      -- Toys
('RCP0000019', 'CSH0000005', 'CRD0000000002', CURRENT_DATE - INTERVAL '26 days', 24.99, 4.998),      -- Auto
('RCP0000020', 'CSH0000006', 'CRD0000000002', CURRENT_DATE - INTERVAL '28 days', 15.99, 3.198),      -- Office

-- Third VIP customer (all categories)
('RCP0000021', 'CSH0000007', 'CRD0000000003', CURRENT_DATE - INTERVAL '4 days', 349.99, 69.998),     -- Electronics
('RCP0000022', 'CSH0000001', 'CRD0000000003', CURRENT_DATE - INTERVAL '7 days', 79.99, 15.998),      -- Clothing
('RCP0000023', 'CSH0000002', 'CRD0000000003', CURRENT_DATE - INTERVAL '10 days', 12.99, 2.598),      -- Food
('RCP0000024', 'CSH0000003', 'CRD0000000003', CURRENT_DATE - INTERVAL '13 days', 19.99, 3.998),      -- Home
('RCP0000025', 'CSH0000004', 'CRD0000000003', CURRENT_DATE - INTERVAL '16 days', 39.99, 7.998),      -- Books
('RCP0000026', 'CSH0000005', 'CRD0000000003', CURRENT_DATE - INTERVAL '21 days', 199.99, 39.998),    -- Sports
('RCP0000027', 'CSH0000006', 'CRD0000000003', CURRENT_DATE - INTERVAL '24 days', 89.99, 17.998),     -- Beauty
('RCP0000028', 'CSH0000007', 'CRD0000000003', CURRENT_DATE - INTERVAL '26 days', 19.99, 3.998),      -- Toys
('RCP0000029', 'CSH0000001', 'CRD0000000003', CURRENT_DATE - INTERVAL '28 days', 19.99, 3.998),      -- Auto
('RCP0000030', 'CSH0000002', 'CRD0000000003', CURRENT_DATE - INTERVAL '30 days', 49.99, 9.998),      -- Office

-- High discount customer receipts for Oleksii1 testing
('RCP0000031', 'CSH0000001', 'CRD0000000004', CURRENT_DATE - INTERVAL '1 days', 899.99, 179.998),    -- High discount customer
('RCP0000032', 'CSH0000002', 'CRD0000000005', CURRENT_DATE - INTERVAL '2 days', 349.99, 69.998),     -- High discount customer
('RCP0000033', 'CSH0000003', 'CRD0000000006', CURRENT_DATE - INTERVAL '3 days', 279.99, 55.998),     -- High discount customer (CSH0000003 - no promo)
('RCP0000034', 'CSH0000004', 'CRD0000000007', CURRENT_DATE - INTERVAL '4 days', 159.99, 31.998),     -- High discount customer
('RCP0000035', 'CSH0000005', 'CRD0000000008', CURRENT_DATE - INTERVAL '5 days', 129.99, 25.998),     -- High discount customer (CSH0000005 - no promo)

-- Additional receipts for volume testing
('RCP0000036', 'CSH0000006', 'CRD0000000009', CURRENT_DATE - INTERVAL '1 month', 199.99, 39.998),
('RCP0000037', 'CSH0000007', 'CRD0000000010', CURRENT_DATE - INTERVAL '2 months', 299.99, 59.998),
('RCP0000038', 'CSH0000001', 'CRD0000000011', CURRENT_DATE - INTERVAL '2 months', 149.99, 29.998),
('RCP0000039', 'CSH0000002', 'CRD0000000012', CURRENT_DATE - INTERVAL '3 months', 89.99, 17.998),
('RCP0000040', 'CSH0000004', 'CRD0000000013', CURRENT_DATE - INTERVAL '3 months', 179.99, 35.998),

-- Historical 2024 receipts for Arthur1 testing (category sales statistics)
('RCP0000041', 'CSH0000001', 'CRD0000000009', '2024-06-15', 299.99, 59.998),
('RCP0000042', 'CSH0000002', 'CRD0000000010', '2024-07-20', 179.99, 35.998),
('RCP0000043', 'CSH0000003', 'CRD0000000011', '2024-08-10', 249.99, 49.998),
('RCP0000044', 'CSH0000004', 'CRD0000000012', '2024-09-05', 129.99, 25.998),
('RCP0000045', 'CSH0000005', 'CRD0000000013', '2024-10-12', 199.99, 39.998);

-- ===== SALES =====
-- Comprehensive sales data designed for individual query testing

-- VIP Customer 1 (CRD0000000001) - Buys from ALL 10 categories in recent month
INSERT INTO sale (upc, receipt_number, product_number, selling_price) VALUES
-- Electronics
('100000000001', 'RCP0000001', 1, 899.99),    -- Samsung Galaxy S23
('100000000002', 'RCP0000001', 1, 349.99),    -- Sony Headphones
-- Clothing
('200000000001', 'RCP0000002', 1, 129.99),    -- Nike shoes
('200000000002', 'RCP0000002', 1, 79.99),     -- Levi's jeans
-- Food & Beverages
('300000000001', 'RCP0000003', 1, 24.99),     -- Coffee beans
('300000000002', 'RCP0000003', 1, 12.99),     -- Chocolate
-- Home & Garden
('400000000001', 'RCP0000004', 2, 19.99),     -- LED Lamp
('400000000002', 'RCP0000004', 1, 39.99),     -- Garden hose
-- Books & Media
('500000000001', 'RCP0000005', 2, 14.99),     -- Great Gatsby
('500000000002', 'RCP0000005', 1, 39.99),     -- Programming guide
-- Sports & Outdoors
('600000000001', 'RCP0000006', 3, 49.99),     -- Yoga mat
('600000000002', 'RCP0000006', 1, 199.99),    -- Tennis racket
-- Health & Beauty
('700000000001', 'RCP0000007', 1, 29.99),     -- Vitamin C serum
('700000000002', 'RCP0000007', 1, 24.99),     -- Moisturizer
-- Toys & Games
('800000000001', 'RCP0000008', 1, 79.99),     -- LEGO set
-- Automotive
('900000000001', 'RCP0000009', 1, 34.99),     -- Motor oil
-- Office Supplies
('101000000001', 'RCP0000010', 1, 29.99),     -- Wireless mouse

-- VIP Customer 2 (CRD0000000002) - Buys from ALL 10 categories
('100000000004', 'RCP0000011', 2, 349.99),    -- Nintendo Switch
('200000000003', 'RCP0000012', 1, 49.99),     -- Hoodie
('300000000003', 'RCP0000013', 1, 8.99),      -- Water
('400000000003', 'RCP0000014', 1, 59.99),     -- Pillow
('500000000003', 'RCP0000015', 1, 99.99),     -- Marvel collection
('600000000003', 'RCP0000016', 1, 299.99),    -- Camping tent
('700000000003', 'RCP0000017', 1, 89.99),     -- Electric toothbrush
('800000000002', 'RCP0000018', 1, 44.99),     -- Board game
('900000000002', 'RCP0000019', 1, 24.99),     -- Phone mount
('101000000002', 'RCP0000020', 1, 15.99),     -- Notebooks

-- VIP Customer 3 (CRD0000000003) - Buys from ALL 10 categories
('100000000005', 'RCP0000021', 1, 649.99),    -- iPad Air
('200000000004', 'RCP0000022', 1, 159.99),    -- Ray-Ban
('300000000004', 'RCP0000023', 1, 18.99),     -- Honey
('400000000004', 'RCP0000024', 1, 29.99),     -- Plant pot
('500000000004', 'RCP0000025', 1, 29.99),     -- Cookbook
('600000000004', 'RCP0000026', 1, 79.99),     -- Bicycle helmet
('700000000004', 'RCP0000027', 1, 39.99),     -- Essential oils
('800000000003', 'RCP0000028', 1, 19.99),     -- Puzzle
('900000000003', 'RCP0000029', 1, 19.99),     -- Pressure gauge
('101000000003', 'RCP0000030', 1, 49.99),     -- Desk organizer

-- HIGH VOLUME ELECTRONICS SALES for Vlad1 (most sold products in Electronics category)
-- Samsung Galaxy S23 - Will be #1 most sold
('100000000001', 'RCP0000031', 2, 899.99),    -- High discount customer
('100000000001', 'RCP0000032', 1, 899.99),
('100000000001', 'RCP0000036', 1, 899.99),
('100000000001', 'RCP0000037', 2, 899.99),
('100000000001', 'RCP0000038', 1, 899.99),
('100000000001', 'RCP0000039', 1, 899.99),

-- Sony Headphones - Will be #2 most sold in Electronics
('100000000002', 'RCP0000032', 2, 349.99),
('100000000002', 'RCP0000034', 1, 349.99),
('100000000002', 'RCP0000036', 1, 349.99),
('100000000002', 'RCP0000037', 1, 349.99),
('100000000002', 'RCP0000038', 2, 349.99),

-- Nintendo Switch - Will be #3 most sold in Electronics
('100000000004', 'RCP0000033', 1, 349.99),
('100000000004', 'RCP0000035', 1, 349.99),
('100000000004', 'RCP0000039', 2, 349.99),
('100000000004', 'RCP0000040', 1, 349.99),

-- iPad Air - #4 in Electronics
('100000000005', 'RCP0000034', 1, 649.99),
('100000000005', 'RCP0000040', 1, 649.99),

-- MacBook Pro - #5 in Electronics
('100000000003', 'RCP0000031', 1, 1999.99),

-- PROMOTIONAL PRODUCT SALES (these will be sold by cashiers CSH0000001, CSH0000002, CSH0000004, CSH0000006, CSH0000007)
-- CSH0000003 and CSH0000005 will NEVER sell promotional products (for Vlad2)
('100000000006', 'RCP0000031', 1, 719.99),    -- Samsung promo by CSH0000001
('100000000007', 'RCP0000032', 1, 279.99),    -- Sony promo by CSH0000002
('100000000006', 'RCP0000034', 1, 719.99),    -- Samsung promo by CSH0000004
('100000000007', 'RCP0000036', 1, 279.99),    -- Sony promo by CSH0000006
('100000000006', 'RCP0000037', 1, 719.99),    -- Samsung promo by CSH0000007

-- SALES FOR HIGH DISCOUNT CUSTOMERS (for Oleksii1 - cashiers serving customers with >15% discount)
-- These cashiers serve customers with high discounts: CSH0000001, CSH0000002, CSH0000004, CSH0000005
('200000000001', 'RCP0000031', 1, 129.99),    -- CSH0000001 serves CRD0000000004 (25% discount)
('200000000002', 'RCP0000032', 1, 79.99),     -- CSH0000002 serves CRD0000000005 (20% discount)
('300000000001', 'RCP0000034', 2, 24.99),     -- CSH0000004 serves CRD0000000007 (16% discount)
('300000000002', 'RCP0000035', 3, 12.99),     -- CSH0000005 serves CRD0000000008 (15% discount)

-- COMPREHENSIVE CATEGORY SALES for Arthur1 (all 10 categories with substantial sales)
-- Additional sales across all categories to ensure good revenue data
('200000000003', 'RCP0000033', 2, 49.99),     -- Clothing
('200000000005', 'RCP0000035', 1, 299.99),    -- Clothing - North Face
('300000000005', 'RCP0000036', 1, 22.99),     -- Food - Tea
('400000000005', 'RCP0000037', 1, 149.99),    -- Home - Knife set
('500000000005', 'RCP0000038', 1, 49.99),     -- Books - Encyclopedia
('600000000005', 'RCP0000039', 1, 149.99),    -- Sports - Hiking backpack
('700000000005', 'RCP0000040', 2, 49.99),     -- Beauty - Protein powder
('800000000004', 'RCP0000036', 1, 129.99),    -- Toys - RC Car
('800000000005', 'RCP0000037', 1, 199.99),    -- Toys - Educational robot
('900000000004', 'RCP0000038', 5, 6.99),      -- Auto - Air freshener
('900000000005', 'RCP0000039', 1, 89.99),     -- Auto - Jump starter
('101000000004', 'RCP0000040', 10, 12.99),    -- Office - Printer paper
('101000000005', 'RCP0000031', 1, 299.99),    -- Office - Standing desk

-- CSH0000003 and CSH0000005 ONLY REGULAR PRODUCTS (never promotional) - for Vlad2
-- CSH0000003 additional sales (RCP0000033)
('200000000004', 'RCP0000033', 1, 159.99),    -- Ray-Ban (non-promotional)
('700000000003', 'RCP0000033', 1, 89.99),     -- Electric toothbrush (non-promotional)

-- CSH0000005 additional sales (RCP0000035)
('600000000001', 'RCP0000035', 1, 49.99),     -- Yoga mat (non-promotional)
('900000000001', 'RCP0000035', 1, 34.99),     -- Motor oil (non-promotional)

-- HISTORICAL 2024 SALES for Arthur1 testing (all categories coverage)
-- RCP0000041 - Electronics and Clothing
('100000000001', 'RCP0000041', 1, 899.99),    -- Samsung Galaxy S23 - Electronics
('200000000001', 'RCP0000041', 1, 129.99),    -- Nike shoes - Clothing

-- RCP0000042 - Food & Beverages and Home & Garden
('300000000001', 'RCP0000042', 2, 24.99),     -- Coffee beans - Food
('400000000001', 'RCP0000042', 3, 19.99),     -- LED Lamp - Home

-- RCP0000043 - Books & Media and Sports & Outdoors
('500000000001', 'RCP0000043', 1, 14.99),     -- Great Gatsby - Books
('600000000001', 'RCP0000043', 2, 49.99),     -- Yoga mat - Sports

-- RCP0000044 - Health & Beauty and Toys & Games
('700000000001', 'RCP0000044', 1, 29.99),     -- Vitamin C serum - Beauty
('800000000001', 'RCP0000044', 1, 79.99),     -- LEGO set - Toys

-- RCP0000045 - Automotive and Office Supplies
('900000000001', 'RCP0000045', 2, 34.99),     -- Motor oil - Automotive
('101000000001', 'RCP0000045', 3, 29.99);     -- Wireless mouse - Office

-- NOTE: Products with UPCs ending in 099 will NEVER be sold (for Arthur2)
-- These are the unsold products: 100000000099, 200000000099, etc.

COMMIT;

-- ===== VALIDATION SUMMARY =====
-- Run this to verify the data works for all individual queries:

-- Vlad1: Electronics (category_id=1) products with sales in last 3 months - should return 5 products
-- SELECT p.product_name, SUM(s.product_number) as total_units
-- FROM sale s JOIN store_product sp ON s.upc = sp.upc
-- JOIN product p ON sp.product_id = p.product_id
-- WHERE p.category_id = 1
-- GROUP BY p.product_id, p.product_name
-- ORDER BY total_units DESC LIMIT 5;

-- Vlad2: Employees who never sold promotional products - should return CSH0000003, CSH0000005
-- SELECT e.employee_id, e.empl_surname FROM employee e WHERE e.empl_role = 'Cashier'
-- AND NOT EXISTS (SELECT 1 FROM receipt r JOIN sale s ON r.receipt_number = s.receipt_number
-- JOIN store_product sp ON s.upc = sp.upc WHERE r.employee_id = e.employee_id
-- AND sp.promotional_product = TRUE);

-- Arthur1: Category sales statistics - should return all 10 categories with revenue data
-- SELECT c.category_name, SUM(s.product_number * s.selling_price) as revenue
-- FROM sale s JOIN store_product sp ON s.upc = sp.upc
-- JOIN product p ON sp.product_id = p.product_id
-- JOIN category c ON p.category_id = c.category_id
-- GROUP BY c.category_name ORDER BY revenue DESC;

-- Arthur2: Unsold non-promotional products - should return 10 products (UPCs ending in 099)
-- SELECT sp.upc, p.product_name FROM store_product sp
-- JOIN product p ON sp.product_id = p.product_id
-- WHERE NOT EXISTS (SELECT 1 FROM sale s WHERE s.upc = sp.upc)
-- AND sp.promotional_product = FALSE AND sp.products_number > 0;

-- Oleksii1: Cashiers serving high discount customers (>15%) - should return 5+ cashiers
-- SELECT e.employee_id, COUNT(DISTINCT cc.card_number) as high_discount_customers
-- FROM employee e JOIN receipt r ON e.employee_id = r.employee_id
-- JOIN customer_card cc ON r.card_number = cc.card_number
-- WHERE cc.percent > 15 AND e.empl_role = 'Cashier'
-- GROUP BY e.employee_id;

-- Oleksii2: Customers who bought from all categories in last month - should return 3 VIP customers
-- SELECT cc.card_number, cc.cust_surname FROM customer_card cc
-- WHERE NOT EXISTS (SELECT 1 FROM category c WHERE NOT EXISTS
-- (SELECT 1 FROM sale s JOIN receipt r ON r.receipt_number = s.receipt_number
-- JOIN store_product sp ON sp.upc = s.upc JOIN product p ON p.product_id = sp.product_id
-- WHERE r.card_number = cc.card_number AND r.print_date >= CURRENT_DATE - INTERVAL '1 month'
-- AND p.category_id = c.category_id));
