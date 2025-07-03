#!/bin/bash

BASE_URL="http://localhost:8080/api/v1"
TOKEN=""

echo "🧪 Testing Training Session Endpoints"
echo "======================================"

# First, let's register a test user
echo "1. Registering test user..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Test",
    "last_name": "User",
    "email": "test@example.com",
    "password": "password123"
  }')

echo "Register response: $REGISTER_RESPONSE"

# Extract token from register response
TOKEN=$(echo $REGISTER_RESPONSE | grep -o '"access_token":"[^"]*"' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
    echo "❌ Failed to get token from registration"
    exit 1
fi

echo "✅ Got token: ${TOKEN:0:20}..."

# Test 1: Create a training session
echo ""
echo "2. Creating training session..."
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/training-sessions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "start_date": "2025-07-03T00:00:00Z",
    "start_time": "2025-07-03T10:00:00Z"
  }')

echo "Create response: $CREATE_RESPONSE"

# Test 2: Get active training session
echo ""
echo "3. Getting active training session..."
ACTIVE_RESPONSE=$(curl -s -X GET "$BASE_URL/training-sessions/active" \
  -H "Authorization: Bearer $TOKEN")

echo "Active session response: $ACTIVE_RESPONSE"

# Test 3: Get all training sessions
echo ""
echo "4. Getting all training sessions..."
ALL_RESPONSE=$(curl -s -X GET "$BASE_URL/training-sessions" \
  -H "Authorization: Bearer $TOKEN")

echo "All sessions response: $ALL_RESPONSE"

# Test 4: End the training session
echo ""
echo "5. Ending training session..."
END_RESPONSE=$(curl -s -X PUT "$BASE_URL/training-sessions/end" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "comments": "Great session! Feeling stronger."
  }')

echo "End session response: $END_RESPONSE"

# Test 5: Verify no active session
echo ""
echo "6. Verifying no active session after ending..."
NO_ACTIVE_RESPONSE=$(curl -s -X GET "$BASE_URL/training-sessions/active" \
  -H "Authorization: Bearer $TOKEN")

echo "No active session response: $NO_ACTIVE_RESPONSE"

echo ""
echo "✅ Training session tests completed!" 