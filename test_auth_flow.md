# Firebase Auth Flow Testing Guide

## Overview
The new auth flow allows frontend to exchange Firebase tokens for custom JWT tokens with refresh capabilities.

## Flow Steps

### 1. Frontend → Backend: Verify Firebase Token
**Endpoint:** `POST /api/verify-token`

**Request:**
```json
{
  "token": "firebase-id-token-from-frontend"
}
```

**Response:**
```json
{
  "access_token": "custom-jwt-access-token",
  "refresh_token": "custom-jwt-refresh-token", 
  "expires_in": 900,
  "user": {
    "id": 1,
    "firebase_uid": "firebase-user-id",
    "email": "user@example.com",
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

### 2. Frontend → Backend: Use Access Token
**Headers:** `Authorization: Bearer {access_token}`

All protected routes now accept either:
- Custom JWT access tokens (preferred)
- Firebase ID tokens (for backward compatibility)

### 3. Frontend → Backend: Refresh Token
**Endpoint:** `POST /api/refresh-token`

**Request:**
```json
{
  "refresh_token": "refresh-token-from-step-1"
}
```

**Response:**
```json
{
  "access_token": "new-custom-jwt-access-token",
  "refresh_token": "new-custom-jwt-refresh-token",
  "expires_in": 900
}
```

## Token Expiration
- **Access Token:** 15 minutes
- **Refresh Token:** 7 days

## Environment Setup
Set `JWT_SECRET` environment variable:
```bash
export JWT_SECRET="your-super-secret-jwt-key-change-this-in-production"
```

## Testing with curl

### 1. Verify Firebase Token
```bash
curl -X POST http://localhost:8080/api/verify-token \
  -H "Content-Type: application/json" \
  -d '{"token": "your-firebase-id-token"}'
```

### 2. Use Access Token
```bash
curl -X GET http://localhost:8080/api/users \
  -H "Authorization: Bearer your-access-token"
```

### 3. Refresh Token
```bash
curl -X POST http://localhost:8080/api/refresh-token \
  -H "Content-Type: application/json" \
  -d '{"refresh_token": "your-refresh-token"}'
```
