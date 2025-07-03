#!/bin/bash

BASE_URL="http://localhost:8080/api/v1"
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTE1MDQ5NTUsImlhdCI6MTc1MTUwNDA1NSwicm9sZSI6InBhdGllbnQiLCJzdWIiOiJlYWI4ODZlMi0zOWU2LTQ2NWEtYmQzYS0yZDkzNjI4OWZiNTMifQ.w2UNpd3vh46dmwIka9x-wpsPfLPuaDyR33UDwzV2-hs"

echo "🧪 Testing Group Assignment Logic"
echo "=================================="

# Function to register a user
register_user() {
    local email="test$1@example.com"
    local response=$(curl -s -X POST "$BASE_URL/auth/register" \
        -H "Content-Type: application/json" \
        -d "{
            \"first_name\": \"Test$1\",
            \"last_name\": \"User$1\",
            \"email\": \"$email\",
            \"password\": \"password123\"
        }")
    echo "User $1 registration: $response"
}

# Register 5 users rapidly to test race condition
echo "Registering 5 users rapidly..."
for i in {1..5}; do
    register_user $i &
done

# Wait for all registrations to complete
wait

echo ""
echo "✅ Group assignment test completed!"
echo "Check the database to see if groups are balanced." 