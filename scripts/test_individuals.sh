#!/bin/bash

# 🌟 Zlagoda Individual Endpoints Testing Script 🌟
# Tests all individual query endpoints

# Configuration
BASE_URL="http://localhost:8080/api"
CONTENT_TYPE="Content-Type: application/json"
AUTH_TOKEN=""

# Colors for output
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

print_header() {
    echo -e "\n${CYAN}🌟 === $1 === 🌟${NC}"
}

print_test() {
    echo -e "${BLUE}🧪 Testing: $1${NC}"
}

test_endpoint() {
    local name="$1"
    local endpoint="$2"
    local expected_fields="$3"

    print_test "$name"
    ((TOTAL_TESTS++))

    echo -e "${YELLOW}📡 GET $BASE_URL/$endpoint${NC}"

    local response
    if [ -n "$AUTH_TOKEN" ]; then
        response=$(curl -s -H "Authorization: Bearer $AUTH_TOKEN" "$BASE_URL/$endpoint")
    else
        response=$(curl -s "$BASE_URL/$endpoint")
    fi
    local exit_code=$?

    if [ $exit_code -ne 0 ]; then
        echo -e "${RED}❌ FAIL: curl failed with exit code $exit_code${NC}"
        ((FAILED_TESTS++))
        return 1
    fi

    # Check if response contains error
    if echo "$response" | grep -q '"error"'; then
        echo -e "${RED}❌ FAIL: API returned error${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
        return 1
    fi

    # Check if response is valid JSON
    if ! echo "$response" | python3 -m json.tool >/dev/null 2>&1; then
        echo -e "${RED}❌ FAIL: Invalid JSON response${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
        return 1
    fi

    # Check if response contains expected structure
    if echo "$response" | grep -q '"description"' && echo "$response" | grep -q '"results"'; then
        echo -e "${GREEN}✅ PASS: Valid response structure${NC}"

        # Pretty print the response
        echo -e "${CYAN}📋 Response:${NC}"
        echo "$response" | python3 -m json.tool | head -20

        # Check if results array exists and show count
        local result_count=$(echo "$response" | python3 -c "
import json, sys
try:
    data = json.load(sys.stdin)
    results = data.get('results', [])
    print(len(results))
except:
    print('0')
" 2>/dev/null)

        echo -e "${PINK}📊 Results count: $result_count${NC}"
        ((PASSED_TESTS++))
        return 0
    else
        echo -e "${RED}❌ FAIL: Missing expected fields (description, results)${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
        return 1
    fi
}

wait_for_api() {
    print_header "Waiting for API to be ready"
    local max_attempts=30
    local attempt=1

    while [ $attempt -le $max_attempts ]; do
        echo -e "${YELLOW}⏳ Attempt $attempt/$max_attempts...${NC}"

        if curl -s "$BASE_URL/../" >/dev/null 2>&1; then
            echo -e "${GREEN}🎉 API is ready!${NC}"
            return 0
        fi

        sleep 2
        ((attempt++))
    done

    echo -e "${RED}💀 API failed to start after $max_attempts attempts${NC}"
    return 1
}

get_auth_token() {
    print_header "Getting Authentication Token"

    # First create a test employee with auth
    local employee_data='{
        "login": "testuser",
        "password": "testpass123",
        "surname": "Test",
        "name": "User",
        "role": "Manager",
        "salary": 50000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Test City",
        "street": "Test Street",
        "zip_code": "12345"
    }'

    echo -e "${YELLOW}🔐 Creating test user...${NC}"
    local register_response=$(curl -s -H "$CONTENT_TYPE" -d "$employee_data" "$BASE_URL/register")

    if echo "$register_response" | grep -q '"token"'; then
        AUTH_TOKEN=$(echo "$register_response" | python3 -c "
import json, sys
try:
    data = json.load(sys.stdin)
    print(data.get('token', ''))
except:
    print('')
" 2>/dev/null)
        echo -e "${GREEN}✅ User created and token obtained${NC}"
        return 0
    fi

    # If registration fails, try login with existing user
    echo -e "${YELLOW}🔐 Trying to login with existing test user...${NC}"
    local login_data='{"login": "testuser", "password": "testpass123"}'
    local login_response=$(curl -s -H "$CONTENT_TYPE" -d "$login_data" "$BASE_URL/login")

    if echo "$login_response" | grep -q '"token"'; then
        AUTH_TOKEN=$(echo "$login_response" | python3 -c "
import json, sys
try:
    data = json.load(sys.stdin)
    print(data.get('token', ''))
except:
    print('')
" 2>/dev/null)
        echo -e "${GREEN}✅ Login successful, token obtained${NC}"
        return 0
    fi

    echo -e "${RED}❌ Failed to obtain authentication token${NC}"
    echo -e "${YELLOW}📝 Register response: $register_response${NC}"
    echo -e "${YELLOW}📝 Login response: $login_response${NC}"
    return 1
}

test_individual_endpoints() {
    print_header "Testing Individual Query Endpoints"

    # Test Vlad1 - Most sold product in a category within a time period
    test_endpoint "Vlad1 - Most sold products in category" "vlad1?category_id=1&months=3"

    # Test Vlad2 - Employees who never sold promotional products
    test_endpoint "Vlad2 - Employees without promo sales" "vlad2"

    # Test Arthur1 - Category sales statistics within date range
    test_endpoint "Arthur1 - Category sales stats" "arthur1?start_date=2024-01-01&end_date=2024-12-31"

    # Test Arthur2 - Products never sold and not promotional
    test_endpoint "Arthur2 - Unsold non-promotional products" "arthur2"

    # Test Oleksii1 - Cashiers who served high discount customers
    test_endpoint "Oleksii1 - Cashiers with high discount customers" "oleksii1?discount_threshold=15"

    # Test Oleksii2 - Customers who bought from all categories
    test_endpoint "Oleksii2 - Customers from all categories" "oleksii2"
}

test_parameter_validation() {
    print_header "Testing Parameter Validation"

    print_test "Vlad1 without required category_id"
    ((TOTAL_TESTS++))
    local response
    if [ -n "$AUTH_TOKEN" ]; then
        response=$(curl -s -H "Authorization: Bearer $AUTH_TOKEN" "$BASE_URL/vlad1")
    else
        response=$(curl -s "$BASE_URL/vlad1")
    fi
    if echo "$response" | grep -q '"error".*category_id.*required'; then
        echo -e "${GREEN}✅ PASS: Properly rejects missing category_id${NC}"
        ((PASSED_TESTS++))
    else
        echo -e "${RED}❌ FAIL: Should reject missing category_id${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
    fi

    print_test "Arthur1 without required dates"
    ((TOTAL_TESTS++))
    if [ -n "$AUTH_TOKEN" ]; then
        response=$(curl -s -H "Authorization: Bearer $AUTH_TOKEN" "$BASE_URL/arthur1")
    else
        response=$(curl -s "$BASE_URL/arthur1")
    fi
    if echo "$response" | grep -q '"error".*date.*required'; then
        echo -e "${GREEN}✅ PASS: Properly rejects missing dates${NC}"
        ((PASSED_TESTS++))
    else
        echo -e "${RED}❌ FAIL: Should reject missing dates${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
    fi

    print_test "Vlad1 with invalid category_id"
    ((TOTAL_TESTS++))
    if [ -n "$AUTH_TOKEN" ]; then
        response=$(curl -s -H "Authorization: Bearer $AUTH_TOKEN" "$BASE_URL/vlad1?category_id=invalid")
    else
        response=$(curl -s "$BASE_URL/vlad1?category_id=invalid")
    fi
    if echo "$response" | grep -q '"error".*Invalid.*category_id'; then
        echo -e "${GREEN}✅ PASS: Properly rejects invalid category_id${NC}"
        ((PASSED_TESTS++))
    else
        echo -e "${RED}❌ FAIL: Should reject invalid category_id${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((FAILED_TESTS++))
    fi
}

show_summary() {
    print_header "Test Results Summary"
    echo -e "${CYAN}📊 Total Tests: $TOTAL_TESTS${NC}"
    echo -e "${GREEN}✅ Passed: $PASSED_TESTS${NC}"
    echo -e "${RED}❌ Failed: $FAILED_TESTS${NC}"

    if [ $FAILED_TESTS -eq 0 ]; then
        echo -e "${GREEN}🎉 All individual endpoints are working perfectly! 🎉${NC}"
        echo -e "${PINK}🌟 Your individual queries implementation is complete! 🌟${NC}"
    elif [ $FAILED_TESTS -lt $((TOTAL_TESTS / 4)) ]; then
        echo -e "${YELLOW}⚠️  Most tests passed with some minor issues${NC}"
    else
        echo -e "${RED}😰 Several tests failed - check the implementation${NC}"
    fi

    local success_rate=$((PASSED_TESTS * 100 / TOTAL_TESTS))
    echo -e "${BLUE}📈 Success Rate: $success_rate%${NC}"
}

cleanup() {
    print_header "Cleanup"
    echo -e "${YELLOW}🧹 Stopping docker compose...${NC}"
    docker compose down
}

main() {
    echo -e "${PINK}🌟 Starting Individual Endpoints Testing 🌟${NC}"
    echo -e "${CYAN}🏠 Base URL: $BASE_URL${NC}"
    echo -e "${BLUE}⏰ Timestamp: $(date)${NC}"

    # Check if python3 is available
    if ! command -v python3 &> /dev/null; then
        echo -e "${YELLOW}🐍 Python3 not found. JSON validation will be limited${NC}"
    fi

    # Start docker compose
    print_header "Starting Services"
    echo -e "${YELLOW}🚀 Starting docker compose...${NC}"
    docker compose up -d

    # Wait for API to be ready
    if ! wait_for_api; then
        echo -e "${RED}💀 Cannot proceed without API. Exiting.${NC}"
        cleanup
        exit 1
    fi

    # Get authentication token
    if ! get_auth_token; then
        echo -e "${RED}💀 Cannot proceed without authentication token. Exiting.${NC}"
        cleanup
        exit 1
    fi

    # Run tests
    test_individual_endpoints
    test_parameter_validation

    # Show results
    show_summary

    # Cleanup
    cleanup

    if [ $FAILED_TESTS -eq 0 ]; then
        exit 0
    else
        exit 1
    fi
}

# Handle Ctrl+C gracefully
trap cleanup EXIT

# Show help
if [[ "$1" == "--help" || "$1" == "-h" ]]; then
    echo -e "${PINK}🌟 Individual Endpoints Testing Script Help 🌟${NC}"
    echo ""
    echo -e "${CYAN}Usage: $0${NC}"
    echo ""
    echo -e "${BLUE}This script tests all individual query endpoints:${NC}"
    echo -e "${GREEN}  🔍 Vlad1    - Most sold products in category${NC}"
    echo -e "${GREEN}  👤 Vlad2    - Employees without promotional sales${NC}"
    echo -e "${GREEN}  📊 Arthur1  - Category sales statistics${NC}"
    echo -e "${GREEN}  📦 Arthur2  - Unsold non-promotional products${NC}"
    echo -e "${GREEN}  💳 Oleksii1 - Cashiers with high discount customers${NC}"
    echo -e "${GREEN}  🛒 Oleksii2 - Customers from all categories${NC}"
    echo ""
    echo -e "${YELLOW}The script will:${NC}"
    echo -e "${CYAN}  1. Start docker compose services${NC}"
    echo -e "${CYAN}  2. Wait for API to be ready${NC}"
    echo -e "${CYAN}  3. Test all individual endpoints${NC}"
    echo -e "${CYAN}  4. Validate parameter handling${NC}"
    echo -e "${CYAN}  5. Show detailed results${NC}"
    echo -e "${CYAN}  6. Clean up services${NC}"
    echo ""
    echo -e "${PINK}Happy testing! 🎉${NC}"
    exit 0
fi

# Run main function
main "$@"
