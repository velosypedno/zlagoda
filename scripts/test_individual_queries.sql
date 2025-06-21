-- Zlagoda Individual Queries Validation Script
-- This script tests all individual SQL queries with the generated sample data
-- Run this after generating sample data to verify all queries work correctly

\echo '=== ZLAGODA INDIVIDUAL QUERIES VALIDATION ==='
\echo 'Testing all Vlad, Arthur, and Oleksii queries with sample data'
\echo ''

-- ===== VLAD1 QUERY TEST =====
\echo '--- VLAD1: Most sold products in Electronics category (last 3 months) ---'
\echo 'Expected: 5 results with Samsung Galaxy S23 as top seller'

SELECT
    c.category_id,
    c.category_name,
    p.product_id,
    p.product_name,
    COUNT(s.upc) AS total_sales,
    SUM(s.product_number) AS total_units_sold,
    SUM(s.product_number * s.selling_price) AS total_revenue
FROM
    category c
JOIN
    product p ON c.category_id = p.category_id
JOIN
    store_product sp ON p.product_id = sp.product_id
JOIN
    sale s ON sp.upc = s.upc
JOIN
    receipt r ON s.receipt_number = r.receipt_number
WHERE
    r.print_date BETWEEN CURRENT_DATE - (3 * INTERVAL '1 month') AND CURRENT_DATE
    AND c.category_id = 1  -- Electronics
GROUP BY
    c.category_id, c.category_name, p.product_id, p.product_name
HAVING
    SUM(s.product_number) >= ALL (
        SELECT SUM(s2.product_number)
        FROM
            product p2
        JOIN
            store_product sp2 ON p2.product_id = sp2.product_id
        JOIN
            sale s2 ON sp2.upc = s2.upc
        JOIN
            receipt r2 ON s2.receipt_number = r2.receipt_number
        WHERE
            p2.category_id = c.category_id
            AND r2.print_date BETWEEN CURRENT_DATE - (3 * INTERVAL '1 month') AND CURRENT_DATE
        GROUP BY
            p2.product_id
    )
ORDER BY
    total_units_sold DESC, total_revenue DESC
LIMIT 5;

\echo ''

-- ===== VLAD2 QUERY TEST =====
\echo '--- VLAD2: Employees who never sold promotional products ---'
\echo 'Expected: CSH0000003 (Kevin Miller) and CSH0000005 (Daniel Moore)'

SELECT
    e.employee_id,
    e.empl_surname,
    e.empl_name
FROM
    employee AS e
WHERE
    NOT EXISTS (
        SELECT 1
        FROM
            receipt AS r
        WHERE
            r.employee_id = e.employee_id
            AND EXISTS (
                SELECT 1
                FROM
                    sale AS s
                    INNER JOIN store_product AS sp
                        ON s.upc = sp.upc
                WHERE
                    s.receipt_number = r.receipt_number
                    AND sp.promotional_product = TRUE
            )
    )
LIMIT 5;

\echo ''

-- ===== ARTHUR1 QUERY TEST =====
\echo '--- ARTHUR1: Category sales statistics (2024 full year) ---'
\echo 'Expected: All 10 categories with revenue data, Electronics likely leading'

SELECT
    c.category_name,
    SUM(s.product_number) as units_sold,
    SUM(s.selling_price * s.product_number) as revenue
FROM sale s
    JOIN receipt r ON s.receipt_number = r.receipt_number
    JOIN store_product sp ON s.upc = sp.upc
    JOIN product p ON sp.product_id = p.product_id
    JOIN category c ON p.category_id = c.category_id
WHERE r.print_date BETWEEN '2024-01-01'::date AND '2024-12-31'::date
GROUP BY c.category_name
ORDER BY revenue DESC, c.category_name ASC
LIMIT 5;

\echo ''

-- ===== ARTHUR2 QUERY TEST =====
\echo '--- ARTHUR2: Unsold non-promotional products with stock ---'
\echo 'Expected: 10 products that have never been sold (UPCs ending in 099)'

SELECT
    sp.upc,
    p.product_name,
    sp.products_number,
    c.category_name
FROM              store_product sp
INNER JOIN        product       p  ON p.product_id  = sp.product_id
INNER JOIN        category      c  ON c.category_id = p.category_id

WHERE NOT EXISTS (
          SELECT 1
          FROM   sale s
          WHERE  s.upc = sp.upc
      )

  AND NOT EXISTS (
          SELECT 1
          FROM   store_product sp2
          WHERE  sp2.upc = sp.upc
            AND (sp2.upc_prom IS NOT NULL
                 OR sp2.promotional_product = TRUE)
      )

  AND sp.products_number > 0
ORDER BY
    c.category_name,
    p.product_name
LIMIT 5;

\echo ''

-- ===== OLEKSII1 QUERY TEST =====
\echo '--- OLEKSII1: Cashiers who served customers with >15% discount ---'
\echo 'Expected: 5+ cashiers with high-discount customer metrics'

SELECT
	e.employee_id,
	e.empl_surname,
	e.empl_name,
	COUNT(DISTINCT cc.card_number) as high_discount_customers,
	COUNT(DISTINCT r.receipt_number) as total_receipts_high_discount,
	SUM(r.sum_total) as total_revenue_high_discount,
	AVG(r.sum_total) as avg_receipt_amount,
	AVG(cc.percent) as avg_customer_discount
FROM employee e
JOIN receipt r ON e.employee_id = r.employee_id
JOIN customer_card cc ON r.card_number = cc.card_number
WHERE cc.percent > 15  -- 15% threshold
	AND e.empl_role = 'Cashier'
GROUP BY e.employee_id, e.empl_surname, e.empl_name
HAVING COUNT(DISTINCT cc.card_number) > 0
ORDER BY high_discount_customers DESC,
		total_revenue_high_discount DESC,
		e.empl_surname ASC
LIMIT 5;

\echo ''

-- ===== OLEKSII2 QUERY TEST =====
\echo '--- OLEKSII2: Customers who bought from all categories in last month ---'
\echo 'Expected: 3 VIP customers (CRD0000000001, CRD0000000002, CRD0000000003)'

SELECT
    cc.card_number,
    cc.cust_surname,
    cc.cust_name,
    cc.phone_number
FROM customer_card cc
WHERE NOT EXISTS (
    SELECT 1
    FROM category c
    WHERE NOT EXISTS (
        SELECT 1
        FROM sale s
        JOIN receipt       r  ON r.receipt_number = s.receipt_number
        JOIN store_product sp ON sp.upc           = s.upc
        JOIN product       p  ON p.product_id     = sp.product_id
        WHERE r.card_number = cc.card_number
          AND r.print_date  >= CURRENT_DATE - INTERVAL '1 month'
          AND p.category_id = c.category_id
    )
)
LIMIT 5;

\echo ''

-- ===== ADDITIONAL VALIDATION QUERIES =====
\echo '--- DATA VALIDATION SUMMARY ---'

\echo 'Total counts verification:'
SELECT 'Categories' as table_name, COUNT(*) as count FROM category
UNION ALL
SELECT 'Products', COUNT(*) FROM product
UNION ALL
SELECT 'Employees', COUNT(*) FROM employee
UNION ALL
SELECT 'Customer Cards', COUNT(*) FROM customer_card
UNION ALL
SELECT 'Store Products', COUNT(*) FROM store_product
UNION ALL
SELECT 'Receipts', COUNT(*) FROM receipt
UNION ALL
SELECT 'Sales', COUNT(*) FROM sale;

\echo ''
\echo 'VIP customers category coverage verification:'
SELECT
    cc.card_number,
    cc.cust_surname,
    COUNT(DISTINCT p.category_id) as categories_purchased
FROM customer_card cc
JOIN receipt r ON cc.card_number = r.card_number
JOIN sale s ON r.receipt_number = s.receipt_number
JOIN store_product sp ON s.upc = sp.upc
JOIN product p ON sp.product_id = p.product_id
WHERE r.print_date >= CURRENT_DATE - INTERVAL '1 month'
  AND cc.card_number IN ('CRD0000000001', 'CRD0000000002', 'CRD0000000003')
GROUP BY cc.card_number, cc.cust_surname
ORDER BY cc.card_number;

\echo ''
\echo 'Employees without promotional sales verification:'
SELECT
    e.employee_id,
    e.empl_surname,
    'No promotional sales' as status
FROM employee e
WHERE e.empl_role = 'Cashier'
AND NOT EXISTS (
    SELECT 1 FROM receipt r
    JOIN sale s ON r.receipt_number = s.receipt_number
    JOIN store_product sp ON s.upc = sp.upc
    WHERE r.employee_id = e.employee_id
    AND sp.promotional_product = TRUE
)
UNION ALL
SELECT
    e.employee_id,
    e.empl_surname,
    'Has promotional sales' as status
FROM employee e
WHERE e.empl_role = 'Cashier'
AND EXISTS (
    SELECT 1 FROM receipt r
    JOIN sale s ON r.receipt_number = s.receipt_number
    JOIN store_product sp ON s.upc = sp.upc
    WHERE r.employee_id = e.employee_id
    AND sp.promotional_product = TRUE
)
ORDER BY status, employee_id;

\echo ''
\echo 'Electronics category sales volume (for Vlad1):'
SELECT
    p.product_name,
    SUM(s.product_number) as total_units_sold,
    COUNT(s.receipt_number) as number_of_sales,
    SUM(s.product_number * s.selling_price) as total_revenue
FROM sale s
JOIN store_product sp ON s.upc = sp.upc
JOIN product p ON sp.product_id = p.product_id
WHERE p.category_id = 1  -- Electronics
GROUP BY p.product_id, p.product_name
ORDER BY total_units_sold DESC;

\echo ''
\echo '=== VALIDATION COMPLETE ==='
\echo 'All individual queries have been tested with sample data'
\echo 'Expected results:'
\echo '- Vlad1: Top Electronics products with Samsung Galaxy S23 leading'
\echo '- Vlad2: CSH0000003 and CSH0000005 (employees without promotional sales)'
\echo '- Arthur1: All 10 categories with revenue data'
\echo '- Arthur2: 10 unsold products (UPCs ending in 099)'
\echo '- Oleksii1: 5+ cashiers serving high-discount customers'
\echo '- Oleksii2: 3 VIP customers buying from all categories'
