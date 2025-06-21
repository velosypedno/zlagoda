#!/bin/bash

# 🌸 Zlagoda API Endpoint Testing Script with Kawaii Validation 🌸
# Proper unit testing with response validation and cute emojis!

# Configuration
BASE_URL="http://localhost:8080/api"
CONTENT_TYPE="Content-Type: application/json"

# Kawaii colors and emojis
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PINK='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Test counters
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Helper functions
print_header() {
    echo -e "\n${CYAN}🌟 === $1 === 🌟${NC}"
}

print_test() {
    echo -e "${BLUE}🧪 Testing: $1${NC}"
}

assert_success() {
    local test_name="$1"
    local response="$2"
    local expected_status="$3"
    local actual_status="$4"

    ((TOTAL_TESTS++))
    if [ "$actual_status" -eq "$expected_status" ]; then
        echo -e "${GREEN}✨ PASS: $test_name (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧${NC}"
        ((PASSED_TESTS++))
        return 0
    else
        echo -e "${RED}💥 FAIL: $test_name (╥﹏╥) Expected: $expected_status, Got: $actual_status${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
        return 1
    fi
}

assert_contains() {
    local test_name="$1"
    local response="$2"
    local expected_field="$3"

    ((TOTAL_TESTS++))
    if echo "$response" | grep -q "$expected_field"; then
        echo -e "${GREEN}🎀 PASS: $test_name contains '$expected_field' (◕‿◕)♡${NC}"
        ((PASSED_TESTS++))
        return 0
    else
        echo -e "${RED}🌧️  FAIL: $test_name missing '$expected_field' (´；ω；')${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
        return 1
    fi
}

assert_not_contains() {
    local test_name="$1"
    local response="$2"
    local unexpected_field="$3"

    ((TOTAL_TESTS++))
    if ! echo "$response" | grep -q "$unexpected_field"; then
        echo -e "${GREEN}🎈 PASS: $test_name doesn't contain '$unexpected_field' (＾◡＾)${NC}"
        ((PASSED_TESTS++))
        return 0
    else
        echo -e "${RED}💔 FAIL: $test_name unexpectedly contains '$unexpected_field' (╯︵╰)${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
        return 1
    fi
}

validate_json() {
    local test_name="$1"
    local response="$2"

    ((TOTAL_TESTS++))
    if echo "$response" | python3 -m json.tool >/dev/null 2>&1; then
        echo -e "${GREEN}🍰 PASS: $test_name is valid JSON (´∀｀)♡${NC}"
        ((PASSED_TESTS++))
        return 0
    else
        echo -e "${RED}🔥 FAIL: $test_name is not valid JSON (ಥ_ಥ)${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
        return 1
    fi
}

make_request() {
    local method="$1"
    local endpoint="$2"
    local data="$3"

    if [ "$method" = "GET" ] || [ "$method" = "DELETE" ]; then
        curl -s -w "\n%{http_code}" -X "$method" "$BASE_URL/$endpoint"
    else
        curl -s -w "\n%{http_code}" -X "$method" "$BASE_URL/$endpoint" -H "$CONTENT_TYPE" -d "$data"
    fi
}

parse_response() {
    local full_response="$1"
    echo "$full_response" | head -n -1
}

parse_status() {
    local full_response="$1"
    echo "$full_response" | tail -n 1
}

# Test API health
test_api_health() {
    print_header "API Health Check"

    print_test "API server availability"
    local full_response=$(curl -s -w "\n%{http_code}" "$BASE_URL/../")
    local status=$(parse_status "$full_response")

    if [ "$status" -eq 200 ] || [ "$status" -eq 404 ]; then
        echo -e "${GREEN}🌈 API server is running! ヽ(´▽｀)/${NC}"
        return 0
    else
        echo -e "${RED}💀 API server is down! (ノಠ益ಠ)ノ${NC}"
        return 1
    fi
}

# Category endpoints testing
test_categories() {
    print_header "Category Endpoints"

    # Test 1: Create category
    print_test "Create category"
    local full_response=$(make_request "POST" "categories" '{"name":"Kawaii Test Category"}')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")

    assert_success "Category creation" "$response" 201 "$status"
    validate_json "Category creation response" "$response"
    assert_contains "Category creation ID" "$response" '"id":'

    # Extract category ID for further tests
    local category_id=$(echo "$response" | grep -o '"id":[0-9]*' | cut -d':' -f2)

    # Test 2: List categories
    print_test "List categories"
    full_response=$(make_request "GET" "categories")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Category listing" "$response" 200 "$status"
    validate_json "Category listing response" "$response"
    assert_contains "Category in list" "$response" '"name":"Kawaii Test Category"'

    # Test 3: Get single category
    if [ -n "$category_id" ]; then
        print_test "Get category by ID"
        full_response=$(make_request "GET" "categories/$category_id")
        response=$(parse_response "$full_response")
        status=$(parse_status "$full_response")

        assert_success "Category retrieval" "$response" 200 "$status"
        validate_json "Category retrieval response" "$response"
        assert_contains "Category ID match" "$response" "\"id\":$category_id"
    fi

    # Test 4: Update category
    if [ -n "$category_id" ]; then
        print_test "Update category"
        full_response=$(make_request "PATCH" "categories/$category_id" '{"name":"Updated Kawaii Category"}')
        response=$(parse_response "$full_response")
        status=$(parse_status "$full_response")

        assert_success "Category update" "$response" 200 "$status"
        assert_contains "Update success message" "$response" '"message"'
    fi
}

# Product endpoints testing
test_products() {
    print_header "Product Endpoints"

    # Test 1: Create product
    print_test "Create product"
    local full_response=$(make_request "POST" "products" '{
        "category_id": 1,
        "name": "Kawaii Test Product",
        "characteristics": "Super kawaii product with magical properties ✨"
    }')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")

    assert_success "Product creation" "$response" 201 "$status"
    validate_json "Product creation response" "$response"
    assert_contains "Product creation ID" "$response" '"id":'

    # Extract product ID
    local product_id=$(echo "$response" | grep -o '"id":[0-9]*' | cut -d':' -f2)

    # Test 2: List products
    print_test "List products"
    full_response=$(make_request "GET" "products")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Product listing" "$response" 200 "$status"
    validate_json "Product listing response" "$response"
    assert_contains "Product in list" "$response" '"name":"Kawaii Test Product"'

    # Test 3: Search products
    print_test "Search products by name"
    full_response=$(make_request "GET" "products/search?name=Kawaii")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Product search" "$response" 200 "$status"
    validate_json "Product search response" "$response"

    # Test 4: Get products by category
    print_test "Get products by category"
    full_response=$(make_request "GET" "products/by-category/1")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Products by category" "$response" 200 "$status"
    validate_json "Products by category response" "$response"
}

# Store Product endpoints testing
test_store_products() {
    print_header "Store Product Endpoints"

    # Test 1: Create store product
    print_test "Create store product"
    local full_response=$(make_request "POST" "store-products" '{
        "product_id": 1,
        "selling_price": 29.99,
        "products_number": 50,
        "promotional_product": false
    }')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")

    assert_success "Store product creation" "$response" 201 "$status"
    validate_json "Store product creation response" "$response"
    assert_contains "Store product UPC" "$response" '"upc":'

    # Extract UPC for further tests
    local upc=$(echo "$response" | grep -o '"upc":"[^"]*"' | cut -d'"' -f4)

    # Test 2: List store products
    print_test "List store products"
    full_response=$(make_request "GET" "store-products")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Store product listing" "$response" 200 "$status"
    validate_json "Store product listing response" "$response"

    # Test 3: List store products with details
    print_test "List store products with details"
    full_response=$(make_request "GET" "store-products/details")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Store products with details" "$response" 200 "$status"
    validate_json "Store products with details response" "$response"
    assert_contains "Product details" "$response" '"product_name"'
    assert_contains "Category details" "$response" '"category_name"'

    # Test 4: Get promotional products
    print_test "Get promotional products"
    full_response=$(make_request "GET" "store-products/promotional")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Promotional products" "$response" 200 "$status"
    validate_json "Promotional products response" "$response"
}

# Employee endpoints testing
test_employees() {
    print_header "Employee Endpoints"

    # Test 1: Create employee
    print_test "Create employee"
    local full_response=$(make_request "POST" "employees" '{
        "empl_surname": "Kawaii",
        "empl_name": "Sakura",
        "empl_patronymic": "Cherry",
        "empl_role": "cashier",
        "salary": 30000.00,
        "date_of_birth": "1995-03-15",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Kawaii City",
        "street": "Sakura Street",
        "zip_code": "12345"
    }')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")

    assert_success "Employee creation" "$response" 201 "$status"
    validate_json "Employee creation response" "$response"
    assert_contains "Employee creation ID" "$response" '"id":'

    # Extract employee ID
    local employee_id=$(echo "$response" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)

    # Test 2: List employees
    print_test "List employees"
    full_response=$(make_request "GET" "employees")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Employee listing" "$response" 200 "$status"
    validate_json "Employee listing response" "$response"
    assert_contains "Employee in list" "$response" '"empl_surname":"Kawaii"'

    # Test 3: Get employee by ID
    if [ -n "$employee_id" ]; then
        print_test "Get employee by ID"
        full_response=$(make_request "GET" "employees/$employee_id")
        response=$(parse_response "$full_response")
        status=$(parse_status "$full_response")

        assert_success "Employee retrieval" "$response" 200 "$status"
        validate_json "Employee retrieval response" "$response"
        assert_contains "Employee ID match" "$response" "\"employee_id\":\"$employee_id\""
    fi
}

# Customer Card endpoints testing
test_customer_cards() {
    print_header "Customer Card Endpoints"

    # Test 1: Create customer card
    print_test "Create customer card"
    local full_response=$(make_request "POST" "customer-cards" '{
        "cust_surname": "Anime",
        "cust_name": "Waifu",
        "cust_patronymic": "Chan",
        "phone_number": "+380987654321",
        "city": "Anime City",
        "street": "Waifu Avenue",
        "zip_code": "54321",
        "percent": 10
    }')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")

    assert_success "Customer card creation" "$response" 201 "$status"
    validate_json "Customer card creation response" "$response"
    assert_contains "Customer card ID" "$response" '"id":'

    # Extract card ID
    local card_id=$(echo "$response" | grep -o '"id":"[^"]*"' | cut -d'"' -f4)

    # Test 2: List customer cards
    print_test "List customer cards"
    full_response=$(make_request "GET" "customer-cards")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Customer card listing" "$response" 200 "$status"
    validate_json "Customer card listing response" "$response"
    assert_contains "Customer card in list" "$response" '"cust_surname":"Anime"'

    # Test 3: Get customer card by number
    if [ -n "$card_id" ]; then
        print_test "Get customer card by number"
        full_response=$(make_request "GET" "customer-cards/$card_id")
        response=$(parse_response "$full_response")
        status=$(parse_status "$full_response")

        assert_success "Customer card retrieval" "$response" 200 "$status"
        validate_json "Customer card retrieval response" "$response"
        assert_contains "Card number match" "$response" "\"card_number\":\"$card_id\""
    fi
}

# Receipt endpoints testing
test_receipts() {
    print_header "Receipt Endpoints"

    # Get existing employee for receipt creation
    local employees_response=$(make_request "GET" "employees")
    local employees=$(parse_response "$employees_response")
    local employee_id=$(echo "$employees" | grep -o '"employee_id":"[^"]*"' | head -1 | cut -d'"' -f4)

    if [ -n "$employee_id" ]; then
        # Test 1: Create receipt
        print_test "Create receipt"
        local full_response=$(make_request "POST" "receipts" "{
            \"employee_id\": \"$employee_id\",
            \"print_date\": \"2024-01-15\",
            \"sum_total\": 99.99
        }")
        local response=$(parse_response "$full_response")
        local status=$(parse_status "$full_response")

        assert_success "Receipt creation" "$response" 201 "$status"
        validate_json "Receipt creation response" "$response"
        assert_contains "Receipt creation ID" "$response" '"id":'
    fi

    # Test 2: List receipts
    print_test "List receipts"
    full_response=$(make_request "GET" "receipts")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Receipt listing" "$response" 200 "$status"
    validate_json "Receipt listing response" "$response"
}

# Sales endpoints testing
test_sales() {
    print_header "Sales Endpoints"

    # Test 1: List sales
    print_test "List sales"
    local full_response=$(make_request "GET" "sales")
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")

    assert_success "Sales listing" "$response" 200 "$status"
    validate_json "Sales listing response" "$response"

    # Test 2: List sales with details
    print_test "List sales with details"
    full_response=$(make_request "GET" "sales/details")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Sales with details" "$response" 200 "$status"
    validate_json "Sales with details response" "$response"

    # Test 3: Get top selling products
    print_test "Get top selling products"
    full_response=$(make_request "GET" "sales/top-products")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Top selling products" "$response" 200 "$status"
    validate_json "Top selling products response" "$response"
}

# Check endpoints testing
test_checks() {
    print_header "Check Endpoints"

    # Get existing data for check creation
    local employees_response=$(make_request "GET" "employees")
    local employees=$(parse_response "$employees_response")
    local employee_id=$(echo "$employees" | grep -o '"employee_id":"[^"]*"' | head -1 | cut -d'"' -f4)

    local cards_response=$(make_request "GET" "customer-cards")
    local cards=$(parse_response "$cards_response")
    local card_number=$(echo "$cards" | grep -o '"card_number":"[^"]*"' | head -1 | cut -d'"' -f4)

    local store_products_response=$(make_request "GET" "store-products")
    local store_products=$(parse_response "$store_products_response")
    local upc=$(echo "$store_products" | grep -o '"upc":"[^"]*"' | head -1 | cut -d'"' -f4)

    if [ -n "$employee_id" ] && [ -n "$upc" ]; then
        # Test 1: Create check
        print_test "Create check (complex transaction)"
        local check_data="{
            \"employee_id\": \"$employee_id\",
            \"card_number\": \"$card_number\",
            \"print_date\": \"2024-01-15\",
            \"items\": [
                {
                    \"upc\": \"$upc\",
                    \"product_number\": 3,
                    \"selling_price\": 29.99
                }
            ]
        }"

        local full_response=$(make_request "POST" "checks" "$check_data")
        local response=$(parse_response "$full_response")
        local status=$(parse_status "$full_response")

        assert_success "Check creation" "$response" 201 "$status"
        validate_json "Check creation response" "$response"
        assert_contains "Receipt number" "$response" '"receipt_number":'
        assert_contains "Total sum" "$response" '"total_sum":'
        assert_contains "VAT" "$response" '"vat":'
        assert_contains "Print date" "$response" '"print_date":'

        # Validate calculation logic
        local total_sum=$(echo "$response" | grep -o '"total_sum":[0-9.]*' | cut -d':' -f2)
        ((TOTAL_TESTS++))
        if [ "$total_sum" = "89.97" ]; then
            echo -e "${GREEN}🧮 PASS: Total calculation correct (29.99 * 3 = 89.97) (◕‿◕)♡${NC}"
            ((PASSED_TESTS++))
        else
            echo -e "${RED}💸 FAIL: Total calculation wrong. Expected: 89.97, Got: $total_sum (╥﹏╥)${NC}"
            ((FAILED_TESTS++))
        fi
    else
        echo -e "${YELLOW}⚠️  Skipping check creation - missing required data (employee_id: $employee_id, upc: $upc) (´･ω･')${NC}"
    fi
}

# Test error handling
test_error_handling() {
    print_header "Error Handling"

    # Test 1: Invalid category creation
    print_test "Invalid category creation (empty name)"
    local full_response=$(make_request "POST" "categories" '{"name":""}')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")

    assert_success "Invalid category rejection" "$response" 400 "$status"
    assert_contains "Error message present" "$response" '"error"'

    # Test 2: Non-existent resource
    print_test "Get non-existent category"
    full_response=$(make_request "GET" "categories/99999")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Non-existent category" "$response" 404 "$status"
    assert_contains "Not found error" "$response" '"error"'

    # Test 3: Invalid employee phone number
    print_test "Invalid employee phone number"
    full_response=$(make_request "POST" "employees" '{
        "empl_surname": "Test",
        "empl_name": "User",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "invalid",
        "city": "Test City",
        "street": "Test Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")

    assert_success "Invalid phone rejection" "$response" 400 "$status"
    assert_contains "Phone validation error" "$response" '"error"'
}

# Cleanup function
cleanup_test_data() {
    print_header "Cleanup"
    echo -e "${PINK}🧹 Test cleanup is disabled to preserve data (◕‿◕)${NC}"
    echo -e "${YELLOW}💭 You can manually delete test data if needed (´∀')${NC}"
}

# Main execution with kawaii summary
main() {
    echo -e "${PINK}🌸✨ Starting Kawaii API Testing ✨🌸${NC}"
    echo -e "${CYAN}🏠 Base URL: $BASE_URL${NC}"
    echo -e "${BLUE}⏰ Timestamp: $(date)${NC}"

    # Check if python3 is available for JSON validation
    if ! command -v python3 &> /dev/null; then
        echo -e "${YELLOW}🐍 Python3 not found. JSON validation will be skipped (´･ω･')${NC}"
    fi

    # Test API health first
    if ! test_api_health; then
        echo -e "${RED}💀 API is not available. Exiting with sadness (╥﹏╥)${NC}"
        exit 1
    fi

    # Run tests based on arguments or run all
    if [ $# -eq 0 ]; then
        print_header "Running All Kawaii Tests"
        test_categories
        test_products
        test_store_products
        test_employees
        test_customer_cards
        test_receipts
        test_sales
        test_checks
        test_error_handling
    else
        for arg in "$@"; do
            case $arg in
                categories)
                    test_categories
                    ;;
                products)
                    test_products
                    ;;
                store-products)
                    test_store_products
                    ;;
                employees)
                    test_employees
                    ;;
                customer-cards)
                    test_customer_cards
                    ;;
                receipts)
                    test_receipts
                    ;;
                sales)
                    test_sales
                    ;;
                checks)
                    test_checks
                    ;;
                errors)
                    test_error_handling
                    ;;
                cleanup)
                    cleanup_test_data
                    ;;
                *)
                    echo -e "${RED}❌ Unknown test category: $arg (｡•́︿•̀｡)${NC}"
                    echo -e "${CYAN}📝 Available: categories, products, store-products, employees, customer-cards, receipts, sales, checks, errors, cleanup${NC}"
                    ;;
            esac
        done
    fi

    # Kawaii test summary
    print_header "Kawaii Test Results"
    echo -e "${CYAN}📊 Total Tests: $TOTAL_TESTS${NC}"
    echo -e "${GREEN}✨ Passed: $PASSED_TESTS${NC}"
    echo -e "${RED}💥 Failed: $FAILED_TESTS${NC}"

    if [ $FAILED_TESTS -eq 0 ]; then
        echo -e "${GREEN}🎉 All tests passed! Your API is absolutely kawaii! (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧${NC}"
        echo -e "${PINK}🌸 Perfect score! You're amazing! (´∀｀)♡${NC}"
    elif [ $FAILED_TESTS -lt $((TOTAL_TESTS / 4)) ]; then
        echo -e "${YELLOW}🌟 Most tests passed! Just a few tiny issues to fix! (◕‿◕)${NC}"
        echo -e "${CYAN}💪 You're doing great! Keep it up! ٩(◕‿◕)۶${NC}"
    else
        echo -e "${RED}😰 Several tests failed, but don't give up! (｡•́︿•̀｡)${NC}"
        echo -e "${PINK}💖 Every failure is a step towards success! Fighting! ٩(◕‿◕)۶${NC}"
    fi

    # Success rate
    local success_rate=$((PASSED_TESTS * 100 / TOTAL_TESTS))
    echo -e "${BLUE}📈 Success Rate: $success_rate%${NC}"

    if [ $success_rate -ge 90 ]; then
        echo -e "${GREEN}🏆 S Rank! Absolutely Perfect! (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧${NC}"
    elif [ $success_rate -ge 80 ]; then
        echo -e "${CYAN}🥇 A Rank! Excellent work! (◕‿◕)♡${NC}"
    elif [ $success_rate -ge 70 ]; then
        echo -e "${YELLOW}🥈 B Rank! Good job! (´∀｀)${NC}"
    else
        echo -e "${PINK}🌸 Keep trying! You'll get there! (◕‿◕)${NC}"
    fi

    echo -e "${PINK}🌸 Testing completed with kawaii love! 🌸${NC}"
}

# Show kawaii help
if [[ "$1" == "--help" || "$1" == "-h" ]]; then
    echo -e "${PINK}🌸 Kawaii API Testing Script Help 🌸${NC}"
    echo ""
    echo -e "${CYAN}Usage: $0 [category1] [category2] ...${NC}"
    echo ""
    echo -e "${BLUE}Available test categories:${NC}"
    echo -e "${GREEN}  🏷️  categories      - Test category endpoints${NC}"
    echo -e "${GREEN}  📦 products        - Test product endpoints${NC}"
    echo -e "${GREEN}  🏪 store-products  - Test store product endpoints${NC}"
    echo -e "${GREEN}  👥 employees       - Test employee endpoints${NC}"
    echo -e "${GREEN}  💳 customer-cards  - Test customer card endpoints${NC}"
    echo -e "${GREEN}  🧾 receipts        - Test receipt endpoints${NC}"
    echo -e "${GREEN}  💰 sales           - Test sales endpoints${NC}"
    echo -e "${GREEN}  ✅ checks          - Test check endpoints${NC}"
    echo -e "${GREEN}  ❌ errors          - Test error handling${NC}"
    echo -e "${GREEN}  🧹 cleanup         - Clean up test data${NC}"
    echo ""
    echo -e "${YELLOW}Examples:${NC}"
    echo -e "${CYAN}  $0                    # Run all kawaii tests ✨${NC}"
    echo -e "${CYAN}  $0 categories         # Test only categories 🏷️${NC}"
    echo -e "${CYAN}  $0 products sales     # Test products and sales 📦💰${NC}"
    echo ""
    echo -e "${PINK}Have fun testing! (◕‿◕)♡${NC}"
    exit 0
fi

# Run main function with all arguments
main "$@"
