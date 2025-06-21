#!/bin/bash

# 🌸✨ Zlagoda API Kawaii Testing Script ✨🌸
# =============================================================================
# This super kawaii script tests all API endpoints with love and care! (◕‿◕)♡
# It creates test data, performs cute operations, and cleans up like a good
# kawaii assistant while respecting database relationships! ヽ(´▽｀)/
#
# Usage: ./test_endpoints.sh
# Requirements: curl, bash, and lots of kawaii energy! ✨
# =============================================================================

# Kawaii colors for output! 🌈
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PINK='\033[0;35m'
CYAN='\033[0;36m'
PURPLE='\033[0;95m'
NC='\033[0m' # No Color

# Base URL
BASE_URL="http://localhost:8080/api"

# Test counters
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Test results storage
declare -a FAILED_TEST_DETAILS=()

# Function to print kawaii colored output
print_status() {
    local status=$1
    local message=$2
    if [ "$status" = "PASS" ]; then
        echo -e "${GREEN}✨ PASS: $message (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧${NC}"
        ((PASSED_TESTS++))
        ((TOTAL_TESTS++))
    elif [ "$status" = "FAIL" ]; then
        echo -e "${RED}💥 FAIL: $message (╥﹏╥)${NC}"
        ((FAILED_TESTS++))
        ((TOTAL_TESTS++))
        FAILED_TEST_DETAILS+=("$message")
    elif [ "$status" = "INFO" ]; then
        echo -e "${CYAN}🌟 $message ♪(´▽｀)♪${NC}"
    fi
}

# Function to test HTTP endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local expected_status=$4
    local test_name=$5

    # Prepare headers
    local headers=("-H" "Content-Type: application/json")
    if [ -n "$JWT_TOKEN" ]; then
        headers+=("-H" "Authorization: Bearer $JWT_TOKEN")
    fi

    if [ -n "$data" ]; then
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" -X "$method" \
            "${headers[@]}" \
            -d "$data" \
            "$BASE_URL$endpoint")
    else
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" -X "$method" \
            "${headers[@]}" \
            "$BASE_URL$endpoint")
    fi

    http_code=$(echo "$response" | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
    body=$(echo "$response" | sed -e 's/HTTPSTATUS:.*//')

    if [ "$http_code" -eq "$expected_status" ]; then
        print_status "PASS" "$test_name (HTTP $http_code)"
        echo "$body" # Return response body for further processing
    else
        print_status "FAIL" "$test_name (Expected HTTP $expected_status, got HTTP $http_code)"
        echo "Response: $body" >&2
    fi
}

# Function to extract ID from JSON response
extract_id() {
    echo "$1" | grep -o '"id":[0-9]*' | grep -o '[0-9]*' | head -1
}

# Function to extract field from JSON response
extract_field() {
    local json=$1
    local field=$2
    echo "$json" | grep -o "\"$field\":\"[^\"]*\"" | sed "s/\"$field\":\"\(.*\)\"/\1/"
}

# Function to login and get JWT token
login_and_get_token() {
    print_status "INFO" "Logging in to get authentication token! (◕‿◕)♡"

    # First register a test user
    register_response=$(curl -s -w "HTTPSTATUS:%{http_code}" -X "POST" \
        -H "Content-Type: application/json" \
        -d '{
            "login": "kawaii_test_user",
            "password": "kawaii123",
            "surname": "Test",
            "name": "Kawaii",
            "patronymic": "Chan",
            "role": "Manager",
            "salary": 50000.00,
            "date_of_birth": "1990-01-01",
            "date_of_start": "2020-01-01",
            "phone_number": "+380999999999",
            "city": "Test City",
            "street": "Test Street 1",
            "zip_code": "12345"
        }' \
        "http://localhost:8080/api/register")

    register_http_code=$(echo "$register_response" | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
    register_body=$(echo "$register_response" | sed -e 's/HTTPSTATUS:.*//')

    if [ "$register_http_code" -eq 201 ]; then
        JWT_TOKEN=$(echo "$register_body" | grep -o '"token":"[^"]*"' | sed 's/"token":"\(.*\)"/\1/')
        print_status "PASS" "User registered and token obtained! ✨"
    else
        # If registration fails, try login with existing user
        print_status "INFO" "Registration failed, trying login with existing user..."

        login_response=$(curl -s -w "HTTPSTATUS:%{http_code}" -X "POST" \
            -H "Content-Type: application/json" \
            -d '{
                "login": "kawaii_test_user",
                "password": "kawaii123"
            }' \
            "http://localhost:8080/api/login")

        login_http_code=$(echo "$login_response" | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
        login_body=$(echo "$login_response" | sed -e 's/HTTPSTATUS:.*//')

        if [ "$login_http_code" -eq 200 ]; then
            JWT_TOKEN=$(echo "$login_body" | grep -o '"token":"[^"]*"' | sed 's/"token":"\(.*\)"/\1/')
            print_status "PASS" "Login successful and token obtained! ✨"
        else
            print_status "FAIL" "Failed to login and get token! Tests will fail (╥﹏╥)"
            echo "Register response: $register_body" >&2
            echo "Login response: $login_body" >&2
        fi
    fi
}
# Store test IDs for cleanup
CATEGORY_ID=""
EMPLOYEE_ID=""
CUSTOMER_CARD_NUMBER=""
PRODUCT_ID=""
STORE_PRODUCT_UPC=""
RECEIPT_NUMBER=""
RECEIPT_NUMBER_2=""

# Authentication token
JWT_TOKEN=""

# Start kawaii banner
echo -e "${PINK}🌸✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨🌸${NC}"
echo -e "${PINK}✨                                                      ✨${NC}"
echo -e "${PINK}✨           🌸 KAWAII API TESTING BEGINS! 🌸           ✨${NC}"
echo -e "${PINK}✨                    (◕‿◕)♡                           ✨${NC}"
echo -e "${PINK}✨                                                      ✨${NC}"
echo -e "${PINK}🌸✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨🌸${NC}"
echo

# Login first to get authentication token
login_and_get_token
echo

echo -e "${CYAN}🏷️ === KAWAII CATEGORY ENDPOINTS === 🏷️${NC}"

# Test Category Creation
print_status "INFO" "Testing Category endpoints with magical powers! ✨"
category_response=$(test_endpoint "POST" "/categories" '{"name":"🌸 Kawaii Test Category 🌸"}' 201 "Create Magical Category")
CATEGORY_ID=$(extract_id "$category_response")

# Test Get Categories List
test_endpoint "GET" "/categories" "" 200 "List All Cute Categories"

# Test Get Category by ID
if [ -n "$CATEGORY_ID" ]; then
    test_endpoint "GET" "/categories/$CATEGORY_ID" "" 200 "Get Category by Kawaii ID"

    # Test Update Category
    test_endpoint "PATCH" "/categories/$CATEGORY_ID" '{"name":"🌸 Updated Kawaii Category 🌸"}' 200 "Update Category with Love"
fi

echo

echo -e "${PINK}💳 === KAWAII CUSTOMER CARD ENDPOINTS === 💳${NC}"

# Test Customer Card Creation
print_status "INFO" "Testing Customer Card endpoints with care! (｡◕‿◕｡)"
customer_card_response=$(test_endpoint "POST" "/customer-cards" '{
    "cust_surname":"Kawaii",
    "cust_name":"Sakura",
    "cust_patronymic":"Chan",
    "phone_number":"+380123456789",
    "city":"🌸 Kawaii City 🌸",
    "street":"Rainbow Street 123",
    "zip_code":"12345",
    "percent":5
}' 201 "Create Kawaii Customer Card")
CUSTOMER_CARD_NUMBER=$(extract_field "$customer_card_response" "id")

# Test Get Customer Cards List
test_endpoint "GET" "/customer-cards" "" 200 "List All Precious Customer Cards"

# Test Get Customer Card by Number
if [ -n "$CUSTOMER_CARD_NUMBER" ]; then
    test_endpoint "GET" "/customer-cards/$CUSTOMER_CARD_NUMBER" "" 200 "Get Customer Card by Magic Number"

    # Test Update Customer Card
    test_endpoint "PATCH" "/customer-cards/$CUSTOMER_CARD_NUMBER" '{"percent":10}' 200 "Update Card with More Love"
fi

echo

echo -e "${BLUE}👥 === KAWAII EMPLOYEE ENDPOINTS === 👥${NC}"

# Test Employee Creation
print_status "INFO" "Testing Employee endpoints with friendship! ٩(◕‿◕)۶"
employee_response=$(test_endpoint "POST" "/employees" '{
    "empl_surname":"Sakura",
    "empl_name":"Kawaii",
    "empl_patronymic":"Chan",
    "empl_role":"cashier",
    "salary":25000.50,
    "date_of_birth":"1995-03-15",
    "date_of_start":"2020-01-10",
    "phone_number":"+380987654321",
    "city":"🌸 Kawaii City 🌸",
    "street":"Rainbow Street 123",
    "zip_code":"12345"
}' 201 "Hire Kawaii Employee")
EMPLOYEE_ID=$(extract_field "$employee_response" "id")

# Test Get Employees List
test_endpoint "GET" "/employees" "" 200 "List All Amazing Employees"

# Test Get Employee by ID
if [ -n "$EMPLOYEE_ID" ]; then
    test_endpoint "GET" "/employees/$EMPLOYEE_ID" "" 200 "Get Employee by Magical ID"

    # Test Update Employee
    test_endpoint "PATCH" "/employees/$EMPLOYEE_ID" '{"salary":26000.00}' 200 "Give Employee a Kawaii Raise"
fi

echo

echo -e "${PURPLE}📦 === KAWAII PRODUCT ENDPOINTS === 📦${NC}"

# Test Product Creation (requires category)
print_status "INFO" "Testing Product endpoints with sparkles! ✨"
if [ -n "$CATEGORY_ID" ]; then
    product_response=$(test_endpoint "POST" "/products" "{
        \"category_id\":$CATEGORY_ID,
        \"name\":\"🌸 Kawaii Test Product 🌸\",
        \"characteristics\":\"Super kawaii product with magical properties ✨\"
    }" 201 "Create Magical Product")
    PRODUCT_ID=$(extract_id "$product_response")

    # Test Get Products List
    test_endpoint "GET" "/products" "" 200 "List All Wonderful Products"

    # Test Get Product by ID
    if [ -n "$PRODUCT_ID" ]; then
        test_endpoint "GET" "/products/$PRODUCT_ID" "" 200 "Get Product by Sparkly ID"

        # Test Update Product
        test_endpoint "PATCH" "/products/$PRODUCT_ID" '{"name":"🌸 Super Kawaii Product 🌸"}' 200 "Make Product Even More Kawaii"

        # Test Get Products by Category
        test_endpoint "GET" "/products/by-category/$CATEGORY_ID" "" 200 "Get Products by Kawaii Category"
    fi

    # Test Search Products by Name
    test_endpoint "GET" "/products/search?name=Kawaii" "" 200 "Search for Kawaii Products"
fi

echo

echo -e "${YELLOW}🏪 === KAWAII STORE PRODUCT ENDPOINTS === 🏪${NC}"

# Test Store Product Creation (requires product)
print_status "INFO" "Testing Store Product endpoints with joy! ヽ(´▽｀)/"
if [ -n "$PRODUCT_ID" ]; then
    store_product_response=$(test_endpoint "POST" "/store-products" "{
        \"product_id\":$PRODUCT_ID,
        \"selling_price\":29.99,
        \"products_number\":100,
        \"promotional_product\":false
    }" 201 "Add Product to Kawaii Store")
    STORE_PRODUCT_UPC=$(extract_field "$store_product_response" "upc")

    # Test Get Store Products List
    test_endpoint "GET" "/store-products" "" 200 "List All Store Treasures"

    # Test Get Store Products with Details
    test_endpoint "GET" "/store-products/details" "" 200 "Get Store Products with Kawaii Details"

    if [ -n "$STORE_PRODUCT_UPC" ]; then
        # Test Get Store Product by UPC
        test_endpoint "GET" "/store-products/$STORE_PRODUCT_UPC" "" 200 "Get Store Product by Magic UPC"

        # Test Update Store Product
        test_endpoint "PATCH" "/store-products/$STORE_PRODUCT_UPC" '{"selling_price":34.99}' 200 "Update Product Price with Love"

        # Test Stock Check
        test_endpoint "GET" "/store-products/$STORE_PRODUCT_UPC/stock-check?quantity=10" "" 200 "Check Kawaii Stock Availability"

        # Test Quantity Update
        test_endpoint "PATCH" "/store-products/$STORE_PRODUCT_UPC/quantity" '{"quantity_change":50}' 200 "Add More Kawaii Products"

        # Test Delivery Update
        test_endpoint "PATCH" "/store-products/$STORE_PRODUCT_UPC/delivery" '{"quantity_change":25,"new_price":39.99}' 200 "Receive Kawaii Delivery"
    fi

    # Test Get Store Products by Product ID
    test_endpoint "GET" "/store-products/by-product/$PRODUCT_ID" "" 200 "Get Store Products by Magical Product ID"

    # Test Get Store Products by Category
    test_endpoint "GET" "/store-products/by-category/$CATEGORY_ID" "" 200 "Get Store Products by Kawaii Category"

    # Test Search Store Products by Name
    test_endpoint "GET" "/store-products/search?name=Kawaii" "" 200 "Search for Kawaii Store Products"

    # Test Get Promotional Products
    test_endpoint "GET" "/store-products/promotional" "" 200 "Get Special Promotional Kawaii Products"
fi

echo

echo -e "${GREEN}🧾 === KAWAII RECEIPT ENDPOINTS === 🧾${NC}"

# Test Receipt Creation (requires employee)
print_status "INFO" "Testing Receipt endpoints with happiness! (◡‿◡)♡"
if [ -n "$EMPLOYEE_ID" ]; then
    receipt_response=$(test_endpoint "POST" "/receipts" "{
        \"employee_id\":\"$EMPLOYEE_ID\",
        \"print_date\":\"2024-01-15\",
        \"sum_total\":159.99
    }" 201 "Create Kawaii Receipt")
    RECEIPT_NUMBER=$(extract_field "$receipt_response" "id")

    # Test Get Receipts List
    test_endpoint "GET" "/receipts" "" 200 "List All Beautiful Receipts"

    if [ -n "$RECEIPT_NUMBER" ]; then
        # Test Get Receipt by Number
        test_endpoint "GET" "/receipts/$RECEIPT_NUMBER" "" 200 "Get Receipt by Magic Number"

        # Test Update Receipt
        test_endpoint "PATCH" "/receipts/$RECEIPT_NUMBER" '{"sum_total":179.99}' 200 "Update Receipt with More Love"

        # Test Receipt Total
        test_endpoint "GET" "/receipts/$RECEIPT_NUMBER/total" "" 200 "Calculate Kawaii Receipt Total"
    fi

    # Test Complete Receipt Creation
    if [ -n "$EMPLOYEE_ID" ] && [ -n "$STORE_PRODUCT_UPC" ] && [ -n "$CUSTOMER_CARD_NUMBER" ]; then
        complete_receipt_response=$(test_endpoint "POST" "/receipts/complete" "{
            \"employee_id\":\"$EMPLOYEE_ID\",
            \"card_number\":\"$CUSTOMER_CARD_NUMBER\",
            \"print_date\":\"2024-01-16\",
            \"items\":[{
                \"upc\":\"$STORE_PRODUCT_UPC\",
                \"product_number\":2,
                \"selling_price\":39.99
            }]
        }" 201 "Create Complete Kawaii Receipt with Items")
        RECEIPT_NUMBER_2=$(extract_field "$complete_receipt_response" "id")
    fi
fi

echo

echo -e "${RED}💰 === KAWAII SALE ENDPOINTS === 💰${NC}"

# Test Sale Creation (requires UPC and receipt)
print_status "INFO" "Testing Sale endpoints with excitement! ✧*｡ヾ(｡>﹏<｡)ﾉ✧*｡"
if [ -n "$STORE_PRODUCT_UPC" ] && [ -n "$RECEIPT_NUMBER" ]; then
    test_endpoint "POST" "/sales" "{
        \"upc\":\"$STORE_PRODUCT_UPC\",
        \"receipt_number\":\"$RECEIPT_NUMBER\",
        \"product_number\":3,
        \"selling_price\":39.99
    }" 201 "Create Kawaii Sale"

    # Test Get Sales List
    test_endpoint "GET" "/sales" "" 200 "List All Amazing Sales"

    # Test Get Sales with Details
    test_endpoint "GET" "/sales/details" "" 200 "Get Sales with Kawaii Details"

    # Test Get Sale by Key
    test_endpoint "GET" "/sales/$STORE_PRODUCT_UPC/$RECEIPT_NUMBER" "" 200 "Get Sale by Magic Key"

    # Test Get Sales by Receipt
    test_endpoint "GET" "/sales/by-receipt/$RECEIPT_NUMBER" "" 200 "Get Sales by Kawaii Receipt"

    # Test Get Sales by Receipt with Details
    test_endpoint "GET" "/sales/by-receipt/$RECEIPT_NUMBER/details" "" 200 "Get Sales by Receipt with Details"

    # Test Get Sales by UPC
    test_endpoint "GET" "/sales/by-upc/$STORE_PRODUCT_UPC" "" 200 "Get Sales by Magic UPC"

    # Test Update Sale
    test_endpoint "PATCH" "/sales/$STORE_PRODUCT_UPC/$RECEIPT_NUMBER" '{"product_number":5}' 200 "Update Sale with More Kawaii"

    # Test Top Selling Products
    test_endpoint "GET" "/sales/top-products?limit=5" "" 200 "Get Top Selling Kawaii Products"

    # Test Sales Stats by Product
    if [ -n "$PRODUCT_ID" ]; then
        test_endpoint "GET" "/sales/stats/product/$PRODUCT_ID?start_date=2024-01-01&end_date=2024-12-31" "" 200 "Get Kawaii Sales Statistics"
    fi
fi

echo

echo -e "${PURPLE}❌ === KAWAII ERROR HANDLING TESTS === ❌${NC}"

print_status "INFO" "Testing error handling with patience! (´･ω･')"

# Test invalid endpoints
test_endpoint "GET" "/invalid-endpoint" "" 404 "Try Invalid Endpoint (Expected Failure)"

# Test invalid IDs
test_endpoint "GET" "/categories/99999" "" 404 "Get Non-existent Category (Expected Failure)"
test_endpoint "GET" "/products/99999" "" 404 "Get Non-existent Product (Expected Failure)"
test_endpoint "GET" "/employees/INV" "" 400 "Get Employee with Invalid ID Format (Expected Failure)"
test_endpoint "GET" "/customer-cards/INVALID" "" 400 "Get Customer Card with Invalid Number (Expected Failure)"

# Test invalid data formats
test_endpoint "POST" "/categories" '{"invalid":"data"}' 400 "Create Category with Invalid Data (Expected Failure)"
test_endpoint "POST" "/categories" '{"name":""}' 400 "Create Category with Empty Name (Expected Failure)"
test_endpoint "POST" "/employees" '{"empl_name":"Test"}' 400 "Create Employee with Missing Fields (Expected Failure)"

# Test boundary conditions
test_endpoint "POST" "/products" '{"category_id":-1,"name":"Test","characteristics":"Test"}' 400 "Create Product with Negative Category ID (Expected Failure)"

# Test malformed JSON
test_endpoint "POST" "/categories" '{"name":}' 400 "Create Category with Malformed JSON (Expected Failure)"
test_endpoint "POST" "/categories" '{"name":"Test",' 400 "Create Category with Incomplete JSON (Expected Failure)"

echo

echo -e "${CYAN}🔍 === ADDITIONAL KAWAII VALIDATION TESTS === 🔍${NC}"

print_status "INFO" "Testing additional validation with thoroughness! (｡◕‿‿◕｡)"

# Test phone number validation
test_endpoint "POST" "/employees" '{"empl_surname":"Test","empl_name":"User","empl_role":"cashier","salary":25000,"date_of_birth":"1990-01-01","date_of_start":"2020-01-01","phone_number":"+38012345678","city":"City","street":"Street","zip_code":"12345"}' 400 "Create Employee with Invalid Phone Length (Expected Failure)"

# Test date validation
test_endpoint "POST" "/employees" '{"empl_surname":"Test","empl_name":"User","empl_role":"cashier","salary":25000,"date_of_birth":"2010-01-01","date_of_start":"2020-01-01","phone_number":"+380123456789","city":"City","street":"Street","zip_code":"12345"}' 400 "Create Employee Under 18 (Expected Failure)"
test_endpoint "POST" "/receipts" '{"employee_id":"REfBl7RUh1","print_date":"2025-12-31","sum_total":100}' 400 "Create Receipt with Future Date (Expected Failure)"

# Test UPC format validation
test_endpoint "GET" "/store-products/INVALID" "" 400 "Get Store Product with Invalid UPC (Expected Failure)"
test_endpoint "POST" "/sales" '{"upc":"INVALID","receipt_number":"1234567890","product_number":1,"selling_price":10.00}' 400 "Create Sale with Invalid UPC (Expected Failure)"

echo

echo -e "${PINK}🧹 === KAWAII CLEANUP TIME === 🧹${NC}"

print_status "INFO" "Cleaning up test data with care! (◕‿◕)♡"

# Delete in proper order to handle foreign key constraints
# 1. Delete sales first (they reference store products and receipts)
if [ -n "$STORE_PRODUCT_UPC" ] && [ -n "$RECEIPT_NUMBER" ]; then
    test_endpoint "DELETE" "/sales/$STORE_PRODUCT_UPC/$RECEIPT_NUMBER" "" 200 "Delete Kawaii Sale"
fi

# 2. Delete receipts (they reference employees and customer cards)
if [ -n "$RECEIPT_NUMBER" ]; then
    test_endpoint "DELETE" "/receipts/$RECEIPT_NUMBER" "" 200 "Delete Kawaii Receipt"
fi

if [ -n "$RECEIPT_NUMBER_2" ]; then
    test_endpoint "DELETE" "/receipts/$RECEIPT_NUMBER_2" "" 200 "Delete Complete Kawaii Receipt"
fi

# 3. Delete store products (they reference products)
if [ -n "$STORE_PRODUCT_UPC" ]; then
    test_endpoint "DELETE" "/store-products/$STORE_PRODUCT_UPC" "" 200 "Remove Product from Kawaii Store"
fi

# 4. Delete products (they reference categories)
if [ -n "$PRODUCT_ID" ]; then
    test_endpoint "DELETE" "/products/$PRODUCT_ID" "" 200 "Delete Kawaii Product"
fi

# 5. Delete categories (no dependencies)
if [ -n "$CATEGORY_ID" ]; then
    test_endpoint "DELETE" "/categories/$CATEGORY_ID" "" 200 "Delete Kawaii Category"
fi

# 6. Delete employees and customer cards (no dependencies from our test data)
if [ -n "$EMPLOYEE_ID" ]; then
    test_endpoint "DELETE" "/employees/$EMPLOYEE_ID" "" 200 "Say Goodbye to Kawaii Employee"
fi

if [ -n "$CUSTOMER_CARD_NUMBER" ]; then
    test_endpoint "DELETE" "/customer-cards/$CUSTOMER_CARD_NUMBER" "" 200 "Delete Kawaii Customer Card"
fi

echo

echo -e "${PINK}🌸✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨🌸${NC}"
echo -e "${PINK}✨                                                      ✨${NC}"
echo -e "${PINK}✨              🌟 KAWAII TEST SUMMARY 🌟              ✨${NC}"
echo -e "${PINK}✨                                                      ✨${NC}"
echo -e "${PINK}🌸✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨🌸${NC}"
echo
echo -e "${CYAN}📊 Total Kawaii Tests: ${YELLOW}$TOTAL_TESTS${NC} ✨"
echo -e "${GREEN}🎉 Passed Tests: ${YELLOW}$PASSED_TESTS${NC} (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧"

if [ $FAILED_TESTS -gt 0 ]; then
    echo -e "${RED}💔 Failed Tests: ${YELLOW}$FAILED_TESTS${NC} (╥﹏╥)"
else
    echo -e "${GREEN}💖 Failed Tests: ${YELLOW}0${NC} ✨ Perfect! ✨"
fi

# Calculate success rate
if [ $TOTAL_TESTS -gt 0 ]; then
    SUCCESS_RATE=$((PASSED_TESTS * 100 / TOTAL_TESTS))
    echo -e "${BLUE}📈 Kawaii Success Rate: ${YELLOW}${SUCCESS_RATE}%${NC} 🌟"
fi

if [ ${#FAILED_TEST_DETAILS[@]} -gt 0 ]; then
    echo
    echo -e "${RED}😰 Failed Tests Details:${NC}"
    for detail in "${FAILED_TEST_DETAILS[@]}"; do
        echo -e "${RED}💔 $detail${NC}"
    done
    echo
    echo -e "${YELLOW}🔧 Kawaii Troubleshooting Tips:${NC}"
    echo -e "${CYAN}💡 1. Make sure the API server is running on http://localhost:8080 ✨${NC}"
    echo -e "${CYAN}💡 2. Check database connectivity and schema (◕‿◕)${NC}"
    echo -e "${CYAN}💡 3. Verify all required environment variables are set 🌟${NC}"
    echo -e "${CYAN}💡 4. Check server logs for detailed error messages ♪(´▽｀)♪${NC}"
fi

echo

echo -e "${PINK}🌸✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨🌸${NC}"
echo -e "${PINK}✨                                                      ✨${NC}"
echo -e "${PINK}✨              🎊 FINAL KAWAII RESULTS 🎊             ✨${NC}"
echo -e "${PINK}✨                                                      ✨${NC}"
echo -e "${PINK}🌸✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨✨🌸${NC}"

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}🎉✨ ALL TESTS PASSED WITH KAWAII PERFECTION! ✨🎉${NC}"
    echo -e "${PINK}🌸 Your API is absolutely kawaii and working beautifully! 🌸${NC}"
    echo -e "${CYAN}💖 Tested endpoints: Categories, Products, Store Products, Employees, Customer Cards, Receipts, Sales ♪(´▽｀)♪${NC}"
    echo -e "${BLUE}🌟 Operations tested: CRUD, validation, error handling, edge cases ✨${NC}"

    if [ $SUCCESS_RATE -eq 100 ]; then
        echo -e "${PURPLE}🏆 PERFECT KAWAII SCORE! You're absolutely amazing! (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧${NC}"
    elif [ $SUCCESS_RATE -ge 95 ]; then
        echo -e "${GREEN}🥇 EXCELLENT KAWAII WORK! Almost perfect! (◕‿◕)♡${NC}"
    elif [ $SUCCESS_RATE -ge 90 ]; then
        echo -e "${YELLOW}🥈 GREAT KAWAII JOB! Very good! ヽ(´▽｀)/${NC}"
    fi

    echo -e "${PINK}🌸 Testing completed with maximum kawaii love! 🌸${NC}"
    exit 0
else
    echo -e "${RED}😰 SOME TESTS FAILED BUT DON'T GIVE UP! 😰${NC}"
    echo -e "${YELLOW}💪 Every failure is a step towards kawaii success! Fighting! ٩(◕‿◕)۶${NC}"
    echo -e "${PINK}💖 Please review the failed tests and fix them with love! (｡◕‿◕｡)${NC}"

    if [ $SUCCESS_RATE -ge 80 ]; then
        echo -e "${CYAN}🌟 You're doing great! Just a few more fixes needed! (◕‿◕)${NC}"
    elif [ $SUCCESS_RATE -ge 60 ]; then
        echo -e "${BLUE}💭 Good progress! Keep working on it! ♪(´▽｀)♪${NC}"
    else
        echo -e "${PURPLE}🌸 Don't worry! Every expert was once a beginner! You can do it! 💪${NC}"
    fi

    echo -e "${PINK}🌸 Keep coding with kawaii spirit! 🌸${NC}"
    exit 1
fi
