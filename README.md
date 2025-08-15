# ToDo App's backend
A simple JWT-based event management application.

## Tech Stack
- Go 1.21  
- Gin (HTTP-framework)  
- GORM (ORM for database)  
- PostgreSQL
- JWT (`github.com/golang-jwt/jwt`)

---

## API Endpoints

| Method | Path                      | Description                     |
|--------|---------------------------|---------------------------------|
| GET    | `/ping`                   | Server's health check           |
| POST   | `/api/v1/users/register`  | Registration                    |
| POST   | `/api/v1/users/login`     | Login and JWT token             |
| POST   | `/api/v1/events/`         | Create an event                 |
| GET    | `/api/v1/events/`         | Get all events                  |
| GET    | `/api/v1/events/{id}`     | Get event by ID                 |
| PUT    | `/api/v1/events/{id}`     | Update event                    |
| DELETE | `/api/v1/events/{id}`     | Delete event                    |

---

## Testing via curl (for CMD)

### 1. Health check
```
curl -i http://localhost:8080/ping
```

### 2. Registration
```
curl -i -X POST http://localhost:8080/api/v1/users/register -H "Content-Type: application/json" -d "{\"email\":\"example@example.com\",\"password\":\"pass123\",\"name\":\"YourName\"}"
```

### 3. Login and saving token
```
curl -i -X POST http://localhost:8080/api/v1/users/login -H "Content-Type: application/json" -d "{\"email\":\"example@example.com\",\"password\":\"pass123\"}"
set TOKEN=<ваш_токен>
```

### 4. Create Event
```
curl -i -X POST http://localhost:8080/api/v1/events/ -H "Authorization: Bearer %TOKEN%" -H "Content-Type: application/json" -d "{\"name\":\"Buy milk\",\"content\":\"Remember to buy 2 liters\",\"time\":\"2025-08-15T10:00:00Z\"}"
```

### 5. Get all Events
```
curl -i http://localhost:8080/api/v1/events/ -H "Authorization: Bearer %TOKEN%"
```

### 6. Get Event by ID
```
curl -i http://localhost:8080/api/v1/events/1 -H "Authorization: Bearer %TOKEN%"
```

### 7. Update Event
```
curl -i -X PUT http://localhost:8080/api/v1/events/1 -H "Authorization: Bearer %TOKEN%" -H "Content-Type: application/json" -d "{\"name\":\"Buy almond milk\",\"content\":\"Switch to almond milk\",\"time\":\"2025-08-15T11:00:00Z\"}"
```

### 8. Delete Event
```
curl -i -X DELETE http://localhost:8080/api/v1/events/3 -H "Authorization: Bearer %TOKEN%"
```

---
## pgAdmin
 
 I also used pgAdmin (official GUI for PostgreSQL) to verify that events were created successfully.
