# Zlagoda Sample Data Generation

This document explains how to generate comprehensive sample data for the Zlagoda store management system to test the individual SQL queries.

## 🎯 Overview

The sample data generation creates realistic test data specifically designed to ensure all individual queries return meaningful results:

- **Vlad's Queries**: Electronics sales analysis with clear top products and employees without promotional sales
- **Arthur's Queries**: Complete category statistics and carefully placed unsold inventory  
- **Oleksii's Queries**: VIP customers buying from all categories and cashiers serving high-discount customers

## 📋 Prerequisites

1. **Docker & Docker Compose** - The database should be running in containers
2. **PostgreSQL Database** - The schema should already be created
3. **Environment Setup** - `.env` file configured (optional)
4. **Database Schema** - All tables must exist before running sample data generation

## 🚀 Quick Start

### Method 1: Using the Shell Script (Recommended)

```bash
# Make the script executable (if not already)
chmod +x generate_sample_data.sh

# Generate sample data with confirmation prompt
./generate_sample_data.sh

# Generate sample data without confirmation
./generate_sample_data.sh --force
```

### Method 2: Direct SQL Execution

```bash
# Start the database container
docker compose up -d postgres-zlagoda

# Execute the SQL script directly
docker compose exec -T postgres-zlagoda psql -U postgres -d weather < generate_sample_data.sql
```

### Method 3: Generate with Validation

```bash
# Generate data and immediately validate all individual queries
./generate_sample_data.sh --validate

# Force generation with validation (no confirmation)
./generate_sample_data.sh --force --validate
```

## 📊 Generated Data Overview

### Database Statistics
- **10 Categories**: Electronics, Clothing, Food & Beverages, Home & Garden, Books & Media, Sports & Outdoors, Health & Beauty, Toys & Games, Automotive, Office Supplies
- **50 Products**: 5 products per category with realistic details
- **10 Employees**: 3 Managers + 7 Cashiers with login credentials
- **18 Customer Cards**: Strategic discount levels (3-45%) including 3 VIP customers
- **60+ Store Products**: Regular products, promotional variants, and strategically unsold items
- **40+ Receipts**: Recent receipts for comprehensive query testing
- **120+ Sales Records**: Precisely designed transaction data for individual query optimization

### Key Data Features

#### For Individual Queries Testing:
1. **Vlad1 (Most Sold Products)**:
   - Electronics category (ID: 1) has concentrated sales with clear hierarchy
   - Samsung Galaxy S23 is #1 most sold (15+ units across multiple receipts)
   - Sony Headphones is #2 most sold (12+ units)
   - Nintendo Switch, iPad Air, and MacBook Pro follow in order
   - All sales within recent months for time-based analysis

2. **Vlad2 (Employees Without Promotional Sales)**:
   - `CSH0000003` (Kevin Miller) and `CSH0000005` (Daniel Moore) never sold promotional items
   - Other cashiers (`CSH0000001`, `CSH0000002`, `CSH0000004`, `CSH0000006`, `CSH0000007`) have promotional sales
   - Clear separation ensures reliable query results

3. **Arthur1 (Category Sales Statistics)**:
   - All 10 categories have substantial sales data with varying revenue levels
   - Electronics leads in revenue, followed by other categories
   - Date range 2024-01-01 to 2024-12-31 covers all data

4. **Arthur2 (Unsold Products)**:
   - 10 products with UPCs ending in "099" have stock but zero sales
   - One product per category for balanced representation
   - All are non-promotional with positive inventory

5. **Oleksii1 (High Discount Customers)**:
   - 5+ cashiers serve customers with >15% discount
   - Discount tiers: VIP (30-45%), High (15-25%), Medium (8-12%), Low (3-7%)
   - Comprehensive metrics: customer count, receipts, revenue, averages

6. **Oleksii2 (Customers from All Categories)**:
   - 3 VIP customers (`CRD0000000001`, `CRD0000000002`, `CRD0000000003`) bought from all 10 categories
   - Each customer has 10 receipts (one per category) within the last month
   - Recent purchase dates from last 30 days with full category coverage

## 🧪 Testing Individual Queries

After generating sample data, you can test in multiple ways:

### Automatic Validation

Run the built-in validation to test all queries at once:
```bash
# Run validation queries directly
docker compose exec -T postgres-zlagoda psql -U postgres -d weather < test_individual_queries.sql

# Or generate data with automatic validation
./generate_sample_data.sh --validate
```

### Manual Testing Parameters

Test individual queries with these specific parameters:

### Frontend Testing (http://localhost:3000)

1. **Navigate to Individual Query Pages**:
   - `/individuals/vlad` - Vlad's queries
   - `/individuals/arthur` - Arthur's queries  
   - `/individuals/oleksii` - Oleksii's queries

2. **Recommended Test Parameters**:

   **Vlad1 (Most Sold Products)**: 
   - Category: Electronics (ID: 1)
   - Months: 3
   - Expected: Samsung Galaxy S23 as top seller

   **Vlad2 (Employees Without Promotional Sales)**:
   - No parameters needed
   - Expected: CSH0000003 and CSH0000005

   **Arthur1 (Category Sales Statistics)**:
   - Start Date: 2024-01-01
   - End Date: 2024-12-31
   - Expected: All 10 categories with revenue data

   **Arthur2 (Unsold Products)**:
   - No parameters needed
   - Expected: 10 products with UPCs ending in "099"

   **Oleksii1 (High Discount Customers)**:
   - Discount Threshold: 15%
   - Expected: 5+ cashiers with detailed metrics

   **Oleksii2 (All Categories)**:
   - No parameters needed
   - Expected: 3 VIP customers

### Backend API Testing

```bash
# Get authentication token first
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"login":"manager1","password":"password"}'

# Test individual queries (replace TOKEN with actual token)
# Test Vlad1 - Should return Samsung Galaxy S23 as #1, Sony Headphones as #2
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:8080/api/vlad1?category_id=1&months=3"

# Test Vlad2 - Should return CSH0000003 (Kevin Miller) and CSH0000005 (Daniel Moore)
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:8080/api/vlad2"

# Test Arthur1 - Should return all 10 categories with Electronics leading in revenue
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:8080/api/arthur1?start_date=2024-01-01&end_date=2024-12-31"

# Test Arthur2 - Should return exactly 10 unsold products (UPCs ending in 099)
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:8080/api/arthur2"

# Test Oleksii1 - Should return 5+ cashiers with comprehensive high-discount metrics
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:8080/api/oleksii1?discount_threshold=15"

# Test Oleksii2 - Should return exactly 3 VIP customers with full category coverage
curl -H "Authorization: Bearer TOKEN" \
  "http://localhost:8080/api/oleksii2"
```

## 🔍 Data Structure Details

### Employee Accounts
```
Managers:
- manager1/password (MGR0000001) - Alice Johnson
- manager2/password (MGR0000002) - Robert Smith
- manager3/password (MGR0000003) - Emma Williams

Cashiers:
- cashier1/password (CSH0000001) - Michael Brown
- cashier2/password (CSH0000002) - Sarah Davis
- cashier3/password (CSH0000003) - Kevin Miller *no promotional sales
- cashier4/password (CSH0000004) - Jessica Wilson
- cashier5/password (CSH0000005) - Daniel Moore *no promotional sales
- cashier6/password (CSH0000006) - Lisa Taylor
- cashier7/password (CSH0000007) - Mark Anderson
```

### Customer Card Types
- **Very High Discount (30-45%)**: VIP customers (CRD0000000001-003)
- **High Discount (15-25%)**: Premium loyalty customers (CRD0000000004-008)
- **Medium Discount (8-12%)**: Regular loyalty customers (CRD0000000009-012)
- **Low Discount (3-7%)**: Basic program members (CRD0000000013-018)

### Product Categories
1. Electronics (phones, laptops, headphones)
2. Clothing (shoes, jeans, hoodies)
3. Food & Beverages (coffee, chocolate, water)
4. Home & Garden (lamps, hoses, pillows)
5. Books & Media (novels, guides, DVDs)
6. Sports & Outdoors (yoga mats, tennis rackets)
7. Health & Beauty (serums, creams, toothbrushes)
8. Toys & Games (LEGO, board games, puzzles)
9. Automotive (motor oil, phone mounts)
10. Office Supplies (mice, notebooks, desks)

### Time Periods
- **Recent Sales**: November 2024 - January 2025 (for current analysis)
- **Historical Data**: June 2024 - October 2024 (for time-based queries)
- **Peak Seasons**: Holiday periods with concentrated Electronics sales
- **VIP Activity**: Recent month purchases across all categories for Oleksii2

## 🛠 Troubleshooting

### Common Issues

1. **Database Connection Failed**:
   ```bash
   # Check if containers are running
   docker compose ps
   
   # Start database if needed
   docker compose up -d postgres-zlagoda
   ```

2. **Permission Denied on Script**:
   ```bash
   chmod +x generate_sample_data.sh
   ```

3. **Schema Not Found**:
   ```bash
   # Run migrations first
   ./migrate.sh
   ```

4. **Foreign Key Constraints**:
   - The script handles deletion order automatically
   - If issues persist, check if schema matches expected structure

### Verification Queries

Use the automated validation script:
```bash
# Run comprehensive validation
docker compose exec -T postgres-zlagoda psql -U postgres -d weather < test_individual_queries.sql
```

Or run manual verification:
```sql
-- Check data counts
SELECT 'categories' as table_name, COUNT(*) as count FROM category
UNION ALL
SELECT 'products', COUNT(*) FROM product
UNION ALL  
SELECT 'employees', COUNT(*) FROM employee
UNION ALL
SELECT 'customer_cards', COUNT(*) FROM customer_card
UNION ALL
SELECT 'store_products', COUNT(*) FROM store_product
UNION ALL
SELECT 'receipts', COUNT(*) FROM receipt
UNION ALL
SELECT 'sales', COUNT(*) FROM sale;

-- Verify VIP customers (should return 3 customers buying from all 10 categories)
SELECT cc.card_number, cc.cust_surname, COUNT(DISTINCT p.category_id) as categories_bought
FROM customer_card cc
JOIN receipt r ON cc.card_number = r.card_number  
JOIN sale s ON r.receipt_number = s.receipt_number
JOIN store_product sp ON s.upc = sp.upc
JOIN product p ON sp.product_id = p.product_id
WHERE r.print_date >= CURRENT_DATE - INTERVAL '1 month'
GROUP BY cc.card_number, cc.cust_surname
HAVING COUNT(DISTINCT p.category_id) = 10;

-- Verify employees without promotional sales (should return CSH0000003, CSH0000005)
SELECT e.employee_id, e.empl_surname 
FROM employee e 
WHERE e.empl_role = 'Cashier' 
AND NOT EXISTS (
    SELECT 1 FROM receipt r 
    JOIN sale s ON r.receipt_number = s.receipt_number 
    JOIN store_product sp ON s.upc = sp.upc 
    WHERE r.employee_id = e.employee_id 
    AND sp.promotional_product = TRUE
);
```

### Expected Query Results

After generating sample data, you should see exactly:

- **Vlad1**: Samsung Galaxy S23 (#1 with 15+ units), Sony Headphones (#2 with 12+ units), Nintendo Switch (#3), iPad Air (#4), MacBook Pro (#5)
- **Vlad2**: CSH0000003 (Kevin Miller) and CSH0000005 (Daniel Moore) - exactly 2 cashiers who never sold promotional items
- **Arthur1**: All 10 categories with revenue data - Electronics leading, followed by other categories in descending order
- **Arthur2**: Exactly 10 products with UPCs ending in "099" - one unsold product per category with positive stock
- **Oleksii1**: 5+ cashiers with detailed metrics for serving customers with >15% discount
- **Oleksii2**: Exactly 3 VIP customers (CRD0000000001, CRD0000000002, CRD0000000003) who bought from all 10 categories in the last month

### Validation Success Criteria
✅ All queries return meaningful, non-empty results
✅ Vlad1 shows clear hierarchy with Samsung Galaxy S23 leading in Electronics
✅ Vlad2 shows exactly CSH0000003 and CSH0000005 (no promotional sales)
✅ Arthur1 covers all 10 categories with substantial revenue figures
✅ Arthur2 shows exactly 10 strategically unsold products
✅ Oleksii1 shows comprehensive cashier performance metrics
✅ Oleksii2 shows 3 VIP customers with complete category coverage
✅ Data spans appropriate time periods for time-based queries
✅ Promotional vs regular product sales are clearly separated

## 🔄 Regenerating Data

To regenerate data:

```bash
# Clear and regenerate
./generate_sample_data.sh --force

# Or manually clear first
docker compose exec -T postgres-zlagoda psql -U postgres -d weather -c "
DELETE FROM sale; 
DELETE FROM receipt; 
DELETE FROM store_product; 
DELETE FROM customer_card; 
DELETE FROM employee; 
DELETE FROM product; 
DELETE FROM category;"
```

## 📞 Support

If you encounter issues:

1. Check the container logs: `docker compose logs postgres-zlagoda`
2. Verify the database schema exists
3. Ensure environment variables are correct
4. Run the verification queries above

The sample data is designed to provide comprehensive test coverage for all individual SQL queries while maintaining realistic business scenarios.