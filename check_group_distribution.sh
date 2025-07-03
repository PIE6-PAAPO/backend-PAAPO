#!/bin/bash

echo "📊 Checking Group Distribution"
echo "=============================="

# Connect to the database and count users in each group
echo "Querying database for group distribution..."

# Get the database container name
DB_CONTAINER=$(docker-compose ps -q db)

if [ -z "$DB_CONTAINER" ]; then
    echo "❌ Database container not found"
    exit 1
fi

# Execute SQL query to count users in each group
echo "Group 1 (Test Group - is_test_group = true):"
docker exec $DB_CONTAINER psql -U postgres -d paapo -c "SELECT COUNT(*) as test_group_count FROM users WHERE is_test_group = true;"

echo ""
echo "Group 2 (Control Group - is_test_group = false):"
docker exec $DB_CONTAINER psql -U postgres -d paapo -c "SELECT COUNT(*) as control_group_count FROM users WHERE is_test_group = false;"

echo ""
echo "Total users:"
docker exec $DB_CONTAINER psql -U postgres -d paapo -c "SELECT COUNT(*) as total_users FROM users;"

echo ""
echo "Detailed breakdown:"
docker exec $DB_CONTAINER psql -U postgres -d paapo -c "SELECT is_test_group, COUNT(*) as count FROM users GROUP BY is_test_group ORDER BY is_test_group;"

echo ""
echo "✅ Group distribution check completed!" 