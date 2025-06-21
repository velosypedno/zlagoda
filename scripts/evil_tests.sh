#!/bin/bash

# 😈 Evil Kawaii Tests - Designed to Break Your API 😈
# These tests are meant to FAIL if your API is secure!
# If they PASS, you have security holes! (╯°□°）╯︵ ┻━┻

BASE_URL="http://localhost:8080/api"
CONTENT_TYPE="Content-Type: application/json"

# Evil colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m'

TOTAL_EVIL_TESTS=0
SECURITY_HOLES=0
PROPERLY_BLOCKED=0

print_evil_header() {
    echo -e "\n${PURPLE}👹 === $1 === 👹${NC}"
}

print_evil_test() {
    echo -e "${CYAN}🔥 Evil Test: $1${NC}"
    ((TOTAL_EVIL_TESTS++))
}

# If the test PASSES (status 200/201), it's a SECURITY HOLE!
# If the test FAILS properly (status 400+), it's GOOD!
assert_blocked() {
    local test_name="$1"
    local response="$2"
    local status="$3"

    if [ "$status" -ge 400 ]; then
        echo -e "${GREEN}😇 GOOD: $test_name was properly blocked! ^_^${NC}"
        ((PROPERLY_BLOCKED++))
    else
        echo -e "${RED}💀 SECURITY HOLE: $test_name succeeded when it should fail! Status: $status${NC}"
        echo -e "${YELLOW}📝 Response: $response${NC}"
        ((SECURITY_HOLES++))
    fi
}

make_evil_request() {
    local method="$1"
    local endpoint="$2"
    local data="$3"

    if [ "$method" = "GET" ] || [ "$method" = "DELETE" ]; then
        curl -s -w "\n%{http_code}" -X "$method" "$BASE_URL/$endpoint" 2>/dev/null
    else
        curl -s -w "\n%{http_code}" -X "$method" "$BASE_URL/$endpoint" -H "$CONTENT_TYPE" -d "$data" 2>/dev/null
    fi
}

parse_response() {
    echo "$1" | head -n -1
}

parse_status() {
    echo "$1" | tail -n 1
}

# SQL Injection attempts
test_sql_injection() {
    print_evil_header "SQL Injection Attacks"

    # Test 1: SQL injection in category name
    print_evil_test "SQL injection in category name"
    local full_response=$(make_evil_request "POST" "categories" "{\"name\":\"'; DROP TABLE categories; --\"}")
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")
    assert_blocked "SQL injection category" "$response" "$status"

    # Test 2: SQL injection in product search
    print_evil_test "SQL injection in product search"
    full_response=$(make_evil_request "GET" "products/search?name=%27%20OR%20%271%27=%271")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "SQL injection search" "$response" "$status"

    # Test 3: SQL injection in employee surname
    print_evil_test "SQL injection in employee creation"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Smith\"; DROP TABLE employees; --",
        "empl_name": "Evil",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Evil City",
        "street": "Evil Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "SQL injection employee" "$response" "$status"
}

# XSS and script injection
test_xss_injection() {
    print_evil_header "XSS and Script Injection"

    # Test 1: XSS in category name
    print_evil_test "XSS in category name"
    local full_response=$(make_evil_request "POST" "categories" '{"name":"<script>alert(\"XSS\")</script>"}')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")
    assert_blocked "XSS category" "$response" "$status"

    # Test 2: XSS in product characteristics
    print_evil_test "XSS in product characteristics"
    full_response=$(make_evil_request "POST" "products" '{
        "category_id": 1,
        "name": "Evil Product",
        "characteristics": "<img src=x onerror=alert(\"XSS\")>"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "XSS product" "$response" "$status"

    # Test 3: JavaScript in customer name
    print_evil_test "JavaScript injection in customer"
    full_response=$(make_evil_request "POST" "customer-cards" '{
        "cust_surname": "javascript:alert(\"evil\")",
        "cust_name": "Evil",
        "phone_number": "+380123456789",
        "city": "Evil City",
        "street": "Evil Street",
        "zip_code": "12345",
        "percent": 5
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "JavaScript customer" "$response" "$status"
}

# Invalid data types and formats
test_invalid_data_types() {
    print_evil_header "Invalid Data Types"

    # Test 1: String where number expected
    print_evil_test "String as category ID in product"
    local full_response=$(make_evil_request "POST" "products" '{
        "category_id": "not_a_number",
        "name": "Evil Product",
        "characteristics": "Evil characteristics"
    }')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")
    assert_blocked "String as number" "$response" "$status"

    # Test 2: Negative salary
    print_evil_test "Negative salary"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Evil",
        "empl_name": "Employee",
        "empl_role": "cashier",
        "salary": -50000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Evil City",
        "street": "Evil Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Negative salary" "$response" "$status"

    # Test 3: Invalid price (string)
    print_evil_test "String as price in store product"
    full_response=$(make_evil_request "POST" "store-products" '{
        "product_id": 1,
        "selling_price": "free",
        "products_number": 50,
        "promotional_product": false
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "String price" "$response" "$status"

    # Test 4: Boolean as string field
    print_evil_test "Boolean as employee name"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": true,
        "empl_name": false,
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Evil City",
        "street": "Evil Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Boolean as string" "$response" "$status"
}

# Boundary value attacks
test_boundary_values() {
    print_evil_header "Boundary Value Attacks"

    # Test 1: Extremely long strings
    print_evil_test "10MB string in category name"
    local long_string=$(printf 'A%.0s' {1..10485760})  # 10MB string
    full_response=$(make_evil_request "POST" "categories" "{\"name\":\"$long_string\"}")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "10MB string" "$response" "$status"

    # Test 2: Extremely large numbers
    print_evil_test "Massive salary number"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Rich",
        "empl_name": "Person",
        "empl_role": "cashier",
        "salary": 999999999999999999999999999999,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Rich City",
        "street": "Rich Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Massive number" "$response" "$status"

    # Test 3: Zero and negative quantities
    print_evil_test "Negative product quantity"
    full_response=$(make_evil_request "POST" "store-products" '{
        "product_id": 1,
        "selling_price": 10.00,
        "products_number": -1000,
        "promotional_product": false
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Negative quantity" "$response" "$status"

    # Test 4: Invalid percentage (over 100)
    print_evil_test "Over 100% customer discount"
    full_response=$(make_evil_request "POST" "customer-cards" '{
        "cust_surname": "Greedy",
        "cust_name": "Customer",
        "phone_number": "+380123456789",
        "city": "Greedy City",
        "street": "Greedy Street",
        "zip_code": "12345",
        "percent": 150
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Over 100% discount" "$response" "$status"
}

# Malformed JSON and null bytes
test_malformed_data() {
    print_evil_header "Malformed Data Attacks"

    # Test 1: Invalid JSON
    print_evil_test "Completely invalid JSON"
    local full_response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL/categories" -H "$CONTENT_TYPE" -d 'this is not json at all!' 2>/dev/null)
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")
    assert_blocked "Invalid JSON" "$response" "$status"

    # Test 2: Null bytes in strings
    print_evil_test "Null bytes in category name"
    full_response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL/categories" -H "$CONTENT_TYPE" -d $'{"name":"Evil\x00Category"}' 2>/dev/null)
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Null bytes" "$response" "$status"

    # Test 3: Unicode control characters
    print_evil_test "Unicode control characters in name"
    full_response=$(make_evil_request "POST" "categories" '{"name":"Evil\u0000\u0001\u0002Category"}')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Unicode control chars" "$response" "$status"

    # Test 4: Nested JSON depth bomb
    print_evil_test "JSON depth bomb"
    local depth_bomb='{"a":{"b":{"c":{"d":{"e":{"f":{"g":{"h":{"i":{"j":{"k":{"l":{"m":{"n":{"o":{"p":"deep"}}}}}}}}}}}}}}}'
    full_response=$(make_evil_request "POST" "categories" "{\"name\":$depth_bomb}")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "JSON depth bomb" "$response" "$status"
}

# Date and time attacks
test_date_attacks() {
    print_evil_header "Date and Time Attacks"

    # Test 1: Future birth date
    print_evil_test "Future birth date"
    local full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Time",
        "empl_name": "Traveler",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "2099-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Future City",
        "street": "Future Street",
        "zip_code": "12345"
    }')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")
    assert_blocked "Future birth date" "$response" "$status"

    # Test 2: Invalid date format
    print_evil_test "Invalid date format"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Bad",
        "empl_name": "Date",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "32/13/2024",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Bad City",
        "street": "Bad Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Invalid date format" "$response" "$status"

    # Test 3: Year 1900 (potential Y2K-style bug)
    print_evil_test "Year 1900 edge case"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Old",
        "empl_name": "Timer",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "1900-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380123456789",
        "city": "Old City",
        "street": "Old Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Year 1900" "$response" "$status"
}

# Phone number validation attacks
test_phone_attacks() {
    print_evil_header "Phone Number Validation Attacks"

    # Test 1: Non-Ukrainian phone
    print_evil_test "Non-Ukrainian phone number"
    local full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Foreign",
        "empl_name": "Person",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+1234567890",
        "city": "Foreign City",
        "street": "Foreign Street",
        "zip_code": "12345"
    }')
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")
    assert_blocked "Non-Ukrainian phone" "$response" "$status"

    # Test 2: Phone with letters
    print_evil_test "Phone number with letters"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Letter",
        "empl_name": "Phone",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380CALLME12",
        "city": "Letter City",
        "street": "Letter Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Phone with letters" "$response" "$status"

    # Test 3: Too short phone
    print_evil_test "Too short phone number"
    full_response=$(make_evil_request "POST" "employees" '{
        "empl_surname": "Short",
        "empl_name": "Phone",
        "empl_role": "cashier",
        "salary": 25000,
        "date_of_birth": "1990-01-01",
        "date_of_start": "2023-01-01",
        "phone_number": "+380",
        "city": "Short City",
        "street": "Short Street",
        "zip_code": "12345"
    }')
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Too short phone" "$response" "$status"
}

# Authorization bypassing attempts
test_authorization_bypass() {
    print_evil_header "Authorization Bypass Attempts"

    # Test 1: Access with fake admin headers
    print_evil_test "Fake admin header bypass"
    local full_response=$(curl -s -w "\n%{http_code}" -X GET "$BASE_URL/employees" -H "X-Admin: true" -H "Authorization: Bearer fake" 2>/dev/null)
    local response=$(parse_response "$full_response")
    local status=$(parse_status "$full_response")
    # This might actually succeed if there's no auth, but we're testing if extra headers cause issues
    echo -e "${CYAN}📝 Admin header test status: $status${NC}"

    # Test 2: Delete non-existent but try directory traversal
    print_evil_test "Directory traversal in delete"
    full_response=$(make_evil_request "DELETE" "categories/../../../etc/passwd")
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    assert_blocked "Directory traversal" "$response" "$status"

    # Test 3: HTTP method override
    print_evil_test "HTTP method override attack"
    full_response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL/categories" -H "$CONTENT_TYPE" -H "X-HTTP-Method-Override: DELETE" -d '{"name":"test"}' 2>/dev/null)
    response=$(parse_response "$full_response")
    status=$(parse_status "$full_response")
    echo -e "${CYAN}📝 Method override test status: $status${NC}"
}

# Main evil execution
main_evil() {
    echo -e "${PURPLE}👹💀 EVIL KAWAII SECURITY TESTS 💀👹${NC}"
    echo -e "${RED}These tests are designed to FAIL! 😈${NC}"
    echo -e "${YELLOW}If they pass, you have security holes! 🕳️${NC}"
    echo -e "${CYAN}Base URL: $BASE_URL${NC}"
    echo -e "${PURPLE}Time: $(date)${NC}"
    echo ""

    # Run all evil tests
    test_sql_injection
    test_xss_injection
    test_invalid_data_types
    test_boundary_values
    test_malformed_data
    test_date_attacks
    test_phone_attacks
    test_authorization_bypass

    # Evil summary
    print_evil_header "EVIL TEST RESULTS"
    echo -e "${CYAN}👹 Total Evil Tests: $TOTAL_EVIL_TESTS${NC}"
    echo -e "${GREEN}😇 Properly Blocked: $PROPERLY_BLOCKED${NC}"
    echo -e "${RED}💀 Security Holes Found: $SECURITY_HOLES${NC}"

    if [ $SECURITY_HOLES -eq 0 ]; then
        echo -e "${GREEN}🛡️  EXCELLENT! Your API blocked all evil attacks! ^_^${NC}"
        echo -e "${GREEN}🏆 S-Rank Security! Your code is truly kawaii and secure!${NC}"
    elif [ $SECURITY_HOLES -lt $((TOTAL_EVIL_TESTS / 4)) ]; then
        echo -e "${YELLOW}⚠️  Some security holes found, but mostly secure!${NC}"
        echo -e "${YELLOW}🔧 Fix the issues above and you'll be perfect!${NC}"
    else
        echo -e "${RED}💥 DANGER! Multiple security vulnerabilities found!${NC}"
        echo -e "${RED}🚨 Your API needs serious security hardening!${NC}"
    fi

    local security_score=$((PROPERLY_BLOCKED * 100 / TOTAL_EVIL_TESTS))
    echo -e "${PURPLE}📊 Security Score: $security_score%${NC}"

    if [ $security_score -ge 95 ]; then
        echo -e "${GREEN}🏆 SECURITY RANK: S (Absolutely Secure!)${NC}"
    elif [ $security_score -ge 85 ]; then
        echo -e "${CYAN}🥇 SECURITY RANK: A (Excellent Security!)${NC}"
    elif [ $security_score -ge 75 ]; then
        echo -e "${YELLOW}🥈 SECURITY RANK: B (Good Security!)${NC}"
    elif [ $security_score -ge 60 ]; then
        echo -e "${YELLOW}🥉 SECURITY RANK: C (Acceptable Security)${NC}"
    else
        echo -e "${RED}💀 SECURITY RANK: F (DANGEROUS!)${NC}"
    fi

    echo -e "${PURPLE}👹 Evil testing completed! 😈${NC}"
}

# Show evil help
if [[ "$1" == "--help" || "$1" == "-h" ]]; then
    echo -e "${PURPLE}👹 Evil Kawaii Security Testing Script 👹${NC}"
    echo ""
    echo -e "${RED}This script contains MALICIOUS tests designed to break your API!${NC}"
    echo -e "${YELLOW}Good APIs should REJECT these requests with 4xx/5xx status codes.${NC}"
    echo -e "${GREEN}If your API accepts these requests, you have security holes!${NC}"
    echo ""
    echo -e "${CYAN}Test categories:${NC}"
    echo -e "${RED}  💉 SQL Injection attempts${NC}"
    echo -e "${RED}  🔥 XSS and script injection${NC}"
    echo -e "${RED}  🎭 Invalid data types${NC}"
    echo -e "${RED}  📏 Boundary value attacks${NC}"
    echo -e "${RED}  💀 Malformed data${NC}"
    echo -e "${RED}  📅 Date/time attacks${NC}"
    echo -e "${RED}  📱 Phone validation attacks${NC}"
    echo -e "${RED}  🔓 Authorization bypass attempts${NC}"
    echo ""
    echo -e "${PURPLE}Usage: $0${NC}"
    echo -e "${YELLOW}Use responsibly! Only test your own APIs!${NC}"
    exit 0
fi

# Run evil tests
main_evil "$@"
