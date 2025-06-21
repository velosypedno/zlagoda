#!/bin/bash

# Zlagoda Sample Data Generation Script
# This script generates comprehensive sample data for testing individual queries

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_header() {
    echo -e "\n${BLUE}🌟 === $1 === 🌟${NC}"
}

# Check if docker compose is available
check_docker_compose() {
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed or not in PATH"
        exit 1
    fi

    if ! docker compose version &> /dev/null; then
        print_error "Docker Compose is not available"
        exit 1
    fi
}

# Check if containers are running
check_containers() {
    print_info "Checking if database container is running..."

    if ! docker compose ps | grep -q "postgres-zlagoda.*running"; then
        print_warning "Database container is not running. Starting services..."
        docker compose up -d postgres-zlagoda

        print_info "Waiting for database to be ready..."
        sleep 10

        # Wait for database to be ready
        local retries=30
        while [ $retries -gt 0 ]; do
            if docker compose exec postgres-zlagoda pg_isready -U postgres &> /dev/null; then
                print_success "Database is ready!"
                break
            fi
            print_info "Waiting for database... ($retries attempts remaining)"
            sleep 2
            ((retries--))
        done

        if [ $retries -eq 0 ]; then
            print_error "Database failed to start within timeout"
            exit 1
        fi
    else
        print_success "Database container is running"
    fi
}

# Load environment variables
load_env() {
    if [ -f .env ]; then
        print_info "Loading environment variables from .env file..."
        export $(grep -v '^#' .env | xargs)
    else
        print_warning ".env file not found, using default values"
        export DB_HOST=${DB_HOST:-postgres-zlagoda}
        export DB_PORT=${DB_PORT:-5432}
        export DB_NAME=${DB_NAME:-weather}
        export DB_USER=${DB_USER:-postgres}
        export DB_PASSWORD=${DB_PASSWORD:-password}
    fi
}

# Run the SQL script
run_sql_script() {
    local sql_file="generate_sample_data.sql"

    if [ ! -f "$sql_file" ]; then
        print_error "SQL script '$sql_file' not found!"
        print_info "Make sure you're running this script from the zlagoda directory"
        exit 1
    fi

    print_info "Executing sample data generation script..."
    print_warning "This will clear existing data and generate new sample data"

    # Ask for confirmation unless --force flag is used
    if [[ "$1" != "--force" ]]; then
        read -p "Are you sure you want to continue? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            print_info "Operation cancelled"
            exit 0
        fi
    fi

    print_info "Running SQL script via Docker..."

    # Execute the SQL script using docker compose
    if docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" < "$sql_file"; then
        print_success "Sample data generated successfully!"
    else
        print_error "Failed to execute SQL script"
        exit 1
    fi
}

# Show usage information
show_usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Generate sample data for Zlagoda database"
    echo ""
    echo "Options:"
    echo "  --force      Skip confirmation prompt"
    echo "  --validate   Run validation queries after data generation"
    echo "  --help       Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0                    # Generate data with confirmation"
    echo "  $0 --force            # Generate data without confirmation"
    echo "  $0 --validate         # Generate data and run validation"
    echo "  $0 --force --validate # Generate data and validate without confirmation"
}

# Verify data generation
verify_data() {
    print_header "Verifying Generated Data"

    print_info "Checking data counts..."

    # Check categories
    local categories=$(docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM category;")
    print_info "Categories: $(echo $categories | tr -d ' ')"

    # Check products
    local products=$(docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM product;")
    print_info "Products: $(echo $products | tr -d ' ')"

    # Check employees
    local employees=$(docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM employee;")
    print_info "Employees: $(echo $employees | tr -d ' ')"

    # Check customer cards
    local cards=$(docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM customer_card;")
    print_info "Customer Cards: $(echo $cards | tr -d ' ')"

    # Check store products
    local store_products=$(docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM store_product;")
    print_info "Store Products: $(echo $store_products | tr -d ' ')"

    # Check receipts
    local receipts=$(docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM receipt;")
    print_info "Receipts: $(echo $receipts | tr -d ' ')"

    # Check sales
    local sales=$(docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM sale;")
    print_info "Sales: $(echo $sales | tr -d ' ')"

    print_success "Data verification completed!"
}

# Run validation queries to test individual query functionality
validate_queries() {
    print_header "Validating Individual Queries"

    local validation_file="test_individual_queries.sql"

    if [ ! -f "$validation_file" ]; then
        print_error "Validation script '$validation_file' not found!"
        print_info "Make sure you're running this script from the zlagoda directory"
        return 1
    fi

    print_info "Running individual query validation tests..."
    print_warning "This will test all Vlad, Arthur, and Oleksii queries"

    # Execute the validation script using docker compose
    if docker compose exec -T postgres-zlagoda psql -U "$DB_USER" -d "$DB_NAME" -f /dev/stdin < "$validation_file"; then
        print_success "Query validation completed successfully!"
        print_info "Check the output above to verify all queries return expected results"
    else
        print_error "Failed to execute validation queries"
        return 1
    fi
}

# Show sample queries for testing
show_sample_queries() {
    print_header "Sample Individual Queries to Test"

    echo -e "${YELLOW}You can now test the individual queries:${NC}"
    echo ""
    echo -e "${BLUE}Vlad1 (Most sold products in category):${NC}"
    echo "  Category ID: 1 (Electronics), Months: 3"
    echo ""
    echo -e "${BLUE}Vlad2 (Employees without promotional sales):${NC}"
    echo "  No parameters needed"
    echo ""
    echo -e "${BLUE}Arthur1 (Category sales statistics):${NC}"
    echo "  Start Date: 2024-01-01, End Date: 2024-12-31"
    echo ""
    echo -e "${BLUE}Arthur2 (Unsold non-promotional products):${NC}"
    echo "  No parameters needed"
    echo ""
    echo -e "${BLUE}Oleksii1 (Cashiers with high discount customers):${NC}"
    echo "  Discount Threshold: 15%"
    echo ""
    echo -e "${BLUE}Oleksii2 (Customers from all categories):${NC}"
    echo "  No parameters needed"
    echo ""
    print_info "Access the frontend at http://localhost:3000 to test these queries!"
}

# Main execution
main() {
    print_header "Zlagoda Sample Data Generator"

    # Parse command line arguments
    FORCE_MODE=false
    VALIDATE_MODE=false

    while [[ $# -gt 0 ]]; do
        case $1 in
            --help|-h)
                show_usage
                exit 0
                ;;
            --force)
                FORCE_MODE=true
                shift
                ;;
            --validate)
                VALIDATE_MODE=true
                shift
                ;;
            "")
                break
                ;;
            *)
                print_error "Unknown option: $1"
                show_usage
                exit 1
                ;;
        esac
    done

    # Check prerequisites
    check_docker_compose

    # Load environment variables
    load_env

    # Check and start containers if needed
    check_containers

    # Run the SQL script
    if [ "$FORCE_MODE" = true ]; then
        run_sql_script --force
    else
        run_sql_script
    fi

    # Verify the generated data
    verify_data

    # Run validation queries if requested
    if [ "$VALIDATE_MODE" = true ]; then
        validate_queries
    fi

    # Show sample queries
    show_sample_queries

    print_success "Sample data generation completed successfully! 🎉"
}

# Run main function with all arguments
main "$@"
