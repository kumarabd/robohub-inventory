# RoboHub Inventory API Schema Reference

Complete API schema documentation for frontend integration.

## Base URL

```
http://localhost:8080/api/v1
```

For production, replace with your deployed service URL or use the Docker image:
```
http://your-domain.com/api/v1
```

## Common Headers

```http
Content-Type: application/json
Accept: application/json
```

## Response Format

All API responses follow this structure:

### Success Response (Single Item)
```json
{
  "id": 1,
  "name": "example",
  "created_at": "2024-02-14T10:30:00Z",
  "updated_at": "2024-02-14T10:30:00Z",
  ...
}
```

### Success Response (List)
```json
{
  "items": [...],
  "total": 42,
  "limit": 10,
  "offset": 0
}
```

### Error Response
```json
{
  "error": "Error message description",
  "code": "ERROR_CODE"
}
```

---

## 1. Repositories

### Schema

```typescript
interface Repository {
  id: number;                // Auto-generated
  name: string;              // Required, unique
  url: string;               // Required
  type: string;              // Required (git, svn, mercurial)
  description?: string;      // Optional
  language?: string;         // Optional (Go, Python, C++, JavaScript, Rust, Java)
  tags?: string[];           // Optional array
  created_at: string;        // ISO 8601 timestamp
  updated_at: string;        // ISO 8601 timestamp
}
```

### Endpoints

#### List Repositories
```http
GET /api/v1/repositories?limit=10&offset=0
```

**Query Parameters:**
- `limit` (optional): Number of items to return (default: 10)
- `offset` (optional): Number of items to skip (default: 0)

**Response Example:**
```json
{
  "items": [
    {
      "id": 1,
      "name": "ros-navigation",
      "url": "https://github.com/ros-planning/navigation2",
      "type": "git",
      "description": "ROS 2 Navigation Stack",
      "language": "C++",
      "tags": ["ros2", "navigation", "autonomous"],
      "created_at": "2024-02-14T10:30:00Z",
      "updated_at": "2024-02-14T10:30:00Z"
    }
  ],
  "total": 1
}
```

#### Get Repository by ID
```http
GET /api/v1/repositories/{id}
```

**Response Example:**
```json
{
  "id": 1,
  "name": "ros-navigation",
  "url": "https://github.com/ros-planning/navigation2",
  "type": "git",
  "description": "ROS 2 Navigation Stack",
  "language": "C++",
  "tags": ["ros2", "navigation", "autonomous"],
  "created_at": "2024-02-14T10:30:00Z",
  "updated_at": "2024-02-14T10:30:00Z"
}
```

#### Create Repository
```http
POST /api/v1/repositories
```

**Request Body:**
```json
{
  "name": "ros-navigation",
  "url": "https://github.com/ros-planning/navigation2",
  "type": "git",
  "description": "ROS 2 Navigation Stack",
  "language": "C++",
  "tags": ["ros2", "navigation", "autonomous"]
}
```

**Response:** Same as Get Repository by ID

#### Update Repository
```http
PUT /api/v1/repositories/{id}
```

**Request Body:** Same as Create Repository

**Response:** Updated repository object

#### Delete Repository
```http
DELETE /api/v1/repositories/{id}
```

**Response:** 204 No Content

---

## 2. Datasets

### Schema

```typescript
interface Dataset {
  id: number;                // Auto-generated
  name: string;              // Required, unique
  description?: string;      // Optional
  type: string;              // Required (sensor, image, lidar, training, validation, test)
  size?: number;             // Optional, size in bytes
  location?: string;         // Optional, storage URL
  format?: string;           // Optional (rosbag, hdf5, json, csv, parquet)
  tags?: string[];           // Optional array
  created_at: string;        // ISO 8601 timestamp
  updated_at: string;        // ISO 8601 timestamp
}
```

### Endpoints

#### List Datasets
```http
GET /api/v1/datasets?limit=10&offset=0
```

**Response Example:**
```json
{
  "items": [
    {
      "id": 1,
      "name": "autonomous-driving-v1",
      "description": "Autonomous driving sensor data from urban environments",
      "type": "sensor",
      "size": 10737418240,
      "location": "s3://robohub-data/autonomous-driving-v1",
      "format": "rosbag",
      "tags": ["lidar", "camera", "urban", "autonomous-driving"],
      "created_at": "2024-02-14T10:30:00Z",
      "updated_at": "2024-02-14T10:30:00Z"
    }
  ],
  "total": 1
}
```

#### Get Dataset by ID
```http
GET /api/v1/datasets/{id}
```

#### Create Dataset
```http
POST /api/v1/datasets
```

**Request Body:**
```json
{
  "name": "autonomous-driving-v1",
  "description": "Autonomous driving sensor data from urban environments",
  "type": "sensor",
  "size": 10737418240,
  "location": "s3://robohub-data/autonomous-driving-v1",
  "format": "rosbag",
  "tags": ["lidar", "camera", "urban", "autonomous-driving"]
}
```

#### Update Dataset
```http
PUT /api/v1/datasets/{id}
```

#### Delete Dataset
```http
DELETE /api/v1/datasets/{id}
```

---

## 3. Packages

### Schema

```typescript
interface Package {
  id: number;                // Auto-generated
  name: string;              // Required, unique
  version: string;           // Required, semantic version
  description?: string;      // Optional
  type: string;              // Required (ros, ros2, python, cpp, docker, apt)
  repository?: string;       // Optional, Git URL
  tags?: string[];           // Optional array
  created_at: string;        // ISO 8601 timestamp
  updated_at: string;        // ISO 8601 timestamp
}
```

### Endpoints

#### List Packages
```http
GET /api/v1/packages?limit=10&offset=0
```

**Response Example:**
```json
{
  "items": [
    {
      "id": 1,
      "name": "ros-navigation",
      "version": "1.0.0",
      "description": "ROS Navigation Stack",
      "type": "ros2",
      "repository": "https://github.com/ros-planning/navigation2",
      "tags": ["ros2", "navigation", "robotics"],
      "created_at": "2024-02-14T10:30:00Z",
      "updated_at": "2024-02-14T10:30:00Z"
    }
  ],
  "total": 1
}
```

#### Get Package by ID
```http
GET /api/v1/packages/{id}
```

#### Create Package
```http
POST /api/v1/packages
```

**Request Body:**
```json
{
  "name": "ros-navigation",
  "version": "1.0.0",
  "description": "ROS Navigation Stack",
  "type": "ros2",
  "repository": "https://github.com/ros-planning/navigation2",
  "tags": ["ros2", "navigation", "robotics"]
}
```

#### Update Package
```http
PUT /api/v1/packages/{id}
```

#### Delete Package
```http
DELETE /api/v1/packages/{id}
```

---

## 4. Scenarios

### Schema

```typescript
interface Scenario {
  id: number;                // Auto-generated
  name: string;              // Required, unique
  description?: string;      // Optional
  type: string;              // Required (simulation, real_world, hybrid, unit_test, integration_test)
  config?: string;           // Optional, JSON string
  tags?: string[];           // Optional array
  created_at: string;        // ISO 8601 timestamp
  updated_at: string;        // ISO 8601 timestamp
}
```

### Config Field Examples

The `config` field is a JSON string that can contain scenario-specific configuration:

```json
{
  "environment": "urban",
  "weather": "clear",
  "traffic_density": "medium",
  "duration": 300,
  "start_position": {"x": 0, "y": 0, "z": 0},
  "goal_position": {"x": 100, "y": 50, "z": 0}
}
```

### Endpoints

#### List Scenarios
```http
GET /api/v1/scenarios?limit=10&offset=0
```

**Response Example:**
```json
{
  "items": [
    {
      "id": 1,
      "name": "urban-navigation",
      "description": "Navigate through urban environment with obstacles",
      "type": "simulation",
      "config": "{\"environment\":\"urban\",\"weather\":\"clear\",\"duration\":300}",
      "tags": ["navigation", "urban", "simulation"],
      "created_at": "2024-02-14T10:30:00Z",
      "updated_at": "2024-02-14T10:30:00Z"
    }
  ],
  "total": 1
}
```

#### Get Scenario by ID
```http
GET /api/v1/scenarios/{id}
```

#### Create Scenario
```http
POST /api/v1/scenarios
```

**Request Body:**
```json
{
  "name": "urban-navigation",
  "description": "Navigate through urban environment with obstacles",
  "type": "simulation",
  "config": "{\"environment\":\"urban\",\"weather\":\"clear\",\"duration\":300}",
  "tags": ["navigation", "urban", "simulation"]
}
```

#### Update Scenario
```http
PUT /api/v1/scenarios/{id}
```

#### Delete Scenario
```http
DELETE /api/v1/scenarios/{id}
```

---

## 5. Simulators

### Schema

```typescript
interface Simulator {
  id: number;                // Auto-generated
  name: string;              // Required, unique
  description?: string;      // Optional
  type: string;              // Required (gazebo, unity, unreal, webots, carla, custom)
  version?: string;          // Optional
  config?: string;           // Optional, JSON string
  tags?: string[];           // Optional array
  created_at: string;        // ISO 8601 timestamp
  updated_at: string;        // ISO 8601 timestamp
}
```

### Config Field Examples

The `config` field is a JSON string for simulator-specific settings:

```json
{
  "physics_engine": "ODE",
  "render_mode": "headless",
  "real_time_factor": 1.0,
  "max_step_size": 0.001,
  "plugins": ["ros_control", "gazebo_ros"]
}
```

### Endpoints

#### List Simulators
```http
GET /api/v1/simulators?limit=10&offset=0
```

**Response Example:**
```json
{
  "items": [
    {
      "id": 1,
      "name": "gazebo-classic",
      "description": "Gazebo Classic simulation environment",
      "type": "gazebo",
      "version": "11.12.0",
      "config": "{\"physics_engine\":\"ODE\",\"render_mode\":\"headless\"}",
      "tags": ["gazebo", "ros", "simulation"],
      "created_at": "2024-02-14T10:30:00Z",
      "updated_at": "2024-02-14T10:30:00Z"
    }
  ],
  "total": 1
}
```

#### Get Simulator by ID
```http
GET /api/v1/simulators/{id}
```

#### Create Simulator
```http
POST /api/v1/simulators
```

**Request Body:**
```json
{
  "name": "gazebo-classic",
  "description": "Gazebo Classic simulation environment",
  "type": "gazebo",
  "version": "11.12.0",
  "config": "{\"physics_engine\":\"ODE\",\"render_mode\":\"headless\"}",
  "tags": ["gazebo", "ros", "simulation"]
}
```

#### Update Simulator
```http
PUT /api/v1/simulators/{id}
```

#### Delete Simulator
```http
DELETE /api/v1/simulators/{id}
```

---

## Type Enumerations

### Repository Types
- `git`
- `svn`
- `mercurial`

### Repository Languages
- `Go`
- `Python`
- `C++`
- `JavaScript`
- `TypeScript`
- `Rust`
- `Java`

### Dataset Types
- `sensor`
- `image`
- `lidar`
- `training`
- `validation`
- `test`

### Dataset Formats
- `rosbag`
- `hdf5`
- `json`
- `csv`
- `parquet`
- `bag`

### Package Types
- `ros`
- `ros2`
- `python`
- `cpp`
- `docker`
- `apt`

### Scenario Types
- `simulation`
- `real_world`
- `hybrid`
- `unit_test`
- `integration_test`

### Simulator Types
- `gazebo`
- `unity`
- `unreal`
- `webots`
- `carla`
- `custom`

---

## Error Codes

| Code | Status | Description |
|------|--------|-------------|
| 200 | OK | Request successful |
| 201 | Created | Resource created successfully |
| 204 | No Content | Resource deleted successfully |
| 400 | Bad Request | Invalid request body or parameters |
| 404 | Not Found | Resource not found |
| 409 | Conflict | Resource with name already exists |
| 500 | Internal Server Error | Server error |

---

## CORS Configuration

The service should be configured to allow CORS requests from your frontend:

```go
// Add to your server configuration
AllowOrigins: ["http://localhost:3000", "https://your-frontend-domain.com"]
AllowMethods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"]
AllowHeaders: ["Content-Type", "Authorization"]
```

---

## Testing with cURL

### Create a Repository
```bash
curl -X POST http://localhost:8080/api/v1/repositories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test-repo",
    "url": "https://github.com/user/repo",
    "type": "git",
    "language": "Go",
    "tags": ["test"]
  }'
```

### List Packages
```bash
curl http://localhost:8080/api/v1/packages?limit=5&offset=0
```

### Update a Dataset
```bash
curl -X PUT http://localhost:8080/api/v1/datasets/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "updated-dataset",
    "type": "sensor",
    "format": "rosbag"
  }'
```

### Delete a Scenario
```bash
curl -X DELETE http://localhost:8080/api/v1/scenarios/1
```

---

## Frontend Integration Example (React + TypeScript)

```typescript
// api/client.ts
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

export async function fetchRepositories(limit = 10, offset = 0) {
  const response = await fetch(`${API_BASE_URL}/repositories?limit=${limit}&offset=${offset}`);
  if (!response.ok) throw new Error('Failed to fetch repositories');
  return response.json();
}

export async function createPackage(data: Partial<Package>) {
  const response = await fetch(`${API_BASE_URL}/packages`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  });
  if (!response.ok) throw new Error('Failed to create package');
  return response.json();
}

export async function updateDataset(id: number, data: Partial<Dataset>) {
  const response = await fetch(`${API_BASE_URL}/datasets/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  });
  if (!response.ok) throw new Error('Failed to update dataset');
  return response.json();
}

export async function deleteScenario(id: number) {
  const response = await fetch(`${API_BASE_URL}/scenarios/${id}`, {
    method: 'DELETE'
  });
  if (!response.ok) throw new Error('Failed to delete scenario');
}
```

---

## Health Check

```http
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "timestamp": "2024-02-14T10:30:00Z"
}
```

---

## Notes for Frontend Developers

1. **Tags**: Always send tags as an array of strings, even if empty: `[]`
2. **Timestamps**: All timestamps are in ISO 8601 format (UTC)
3. **Size**: Dataset size is in bytes (convert to MB/GB in UI)
4. **Config**: Config fields for Scenarios and Simulators are JSON strings - parse before using
5. **Pagination**: Total count is returned with list responses for pagination
6. **Unique Names**: Names must be unique across each resource type
7. **ID**: ID field is auto-generated and read-only

---

For more information, see the main [README.md](./README.md) and [V0_PROMPT.md](./V0_PROMPT.md).
