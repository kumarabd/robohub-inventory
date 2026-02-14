# Quick Start: Using the Updated Schema

## 🚀 Apply Database Changes

```bash
# 1. Backup existing data (if any)
pg_dump -U postgres robohub_inventory > backup_$(date +%Y%m%d).sql

# 2. Apply schema migration
psql -U postgres -d robohub_inventory -f migrations/001_update_schema_to_contract.sql

# 3. Load sample data
psql -U postgres -d robohub_inventory -f migrations/002_seed_data.sql

# 4. Verify
psql -U postgres -d robohub_inventory -c "SELECT COUNT(*) FROM repositories;"
```

## 📊 Example API Requests

### Get All Repositories

```bash
curl http://localhost:8080/api/v1/repositories
```

**Response:**
```json
[
  {
    "id": "a1b2c3d4-e5f6-...",
    "name": "ros-planning/navigation2",
    "provider": "github",
    "url": "https://github.com/ros-planning/navigation2",
    "defaultBranch": "main",
    "syncStatus": "synced",
    "latestCommit": {
      "hash": "a1b2c3d4e5f6",
      "message": "Add new planner plugin",
      "author": "John Doe",
      "date": "2024-02-14T10:30:00Z"
    },
    "owner": {
      "id": "user-001",
      "name": "ros-planning"
    },
    "packageCount": 5,
    "tags": ["ros2", "navigation", "autonomous"],
    "createdAt": "2024-02-14T10:00:00Z"
  }
]
```

### Create a Package

```bash
curl -X POST http://localhost:8080/api/v1/packages \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my_planner",
    "displayName": "My Custom Planner",
    "description": "Custom path planning algorithm",
    "repoId": "a1b2c3d4-e5f6-...",
    "types": ["planner", "navigation"],
    "latestVersion": "1.0.0",
    "tags": ["custom", "planning"],
    "owner": {
      "id": "user-123",
      "name": "My Team"
    }
  }'
```

### Create a Scenario

```bash
curl -X POST http://localhost:8080/api/v1/scenarios \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Test Scenario",
    "description": "Test autonomous navigation",
    "category": "navigation",
    "difficulty": "medium",
    "maintainedBy": "Community",
    "verified": false,
    "whatItTests": ["Path planning", "Obstacle avoidance"],
    "domain": "indoor",
    "supportedSimulators": ["Gazebo"],
    "tags": ["navigation", "test"],
    "owner": {
      "id": "user-123",
      "name": "My Team"
    }
  }'
```

### Create a Dataset

```bash
curl -X POST http://localhost:8080/api/v1/datasets \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my_dataset",
    "description": "Custom navigation dataset",
    "type": "robotics",
    "modality": "multimodal",
    "format": "rosbag2",
    "license": "MIT",
    "sizeGB": 5.5,
    "samplesCount": 1000,
    "source": "uploaded",
    "ownerType": "user",
    "ownerId": "user-123",
    "ownerName": "My Team",
    "visibility": "public",
    "tags": ["navigation", "indoor"]
  }'
```

## 🔍 Query with UUID

All resources now use UUIDs. Save the `id` from create responses:

```bash
# Get by UUID
curl http://localhost:8080/api/v1/packages/a1b2c3d4-e5f6-7890-abcd-ef1234567890

# Update by UUID
curl -X PUT http://localhost:8080/api/v1/packages/a1b2c3d4-e5f6-... \
  -H "Content-Type: application/json" \
  -d '{"name": "updated_name", ...}'

# Delete by UUID
curl -X DELETE http://localhost:8080/api/v1/packages/a1b2c3d4-e5f6-...
```

## 📝 Key Schema Differences

### Before → After

#### IDs
```javascript
// Before
{ "id": 1 }

// After
{ "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890" }
```

#### Timestamps
```javascript
// Before
{ "created_at": "...", "updated_at": "..." }

// After
{ "createdAt": "...", "updatedAt": "..." }
```

#### Package Types
```javascript
// Before
{ "type": "ros" }

// After
{ "types": ["planner", "navigation"] }
```

#### Repository
```javascript
// Before
{
  "id": 1,
  "name": "repo",
  "type": "git",
  "language": "Go"
}

// After
{
  "id": "uuid...",
  "name": "org/repo",
  "provider": "github",
  "syncStatus": "synced",
  "latestCommit": {...},
  "owner": {...}
}
```

## 🎯 Frontend Integration Checklist

- [ ] Update ID handling from `number` to `string`
- [ ] Update TypeScript interfaces with new field names
- [ ] Handle UUID format in URL parameters
- [ ] Update timestamp field names (camelCase)
- [ ] Handle array fields (types, tags, etc.)
- [ ] Parse JSONB fields (owner, latestCommit, etc.)
- [ ] Update forms for new required fields
- [ ] Test all CRUD operations with new schema

## 🔧 Common Operations

### Get Repository with Packages

```sql
SELECT 
  r.*,
  json_agg(p.*) as packages
FROM repositories r
LEFT JOIN packages p ON p.repo_id = r.id
WHERE r.id = 'uuid...'
GROUP BY r.id;
```

### Find Packages by Type

```sql
SELECT * FROM packages 
WHERE 'planner' = ANY(types);
```

### Search Scenarios by Category

```sql
SELECT * FROM scenarios 
WHERE category = 'navigation' 
AND difficulty = 'easy';
```

### Get Dataset with Schema

```sql
SELECT 
  name,
  schema->'topics' as topics,
  schema->'dataSplits' as data_splits
FROM datasets 
WHERE id = 'uuid...';
```

## 🐛 Troubleshooting

### "Invalid ID" Error
- **Cause**: Sending integer instead of UUID
- **Fix**: Use UUID string format

### "Column does not exist" Error
- **Cause**: Using old column names
- **Fix**: Use camelCase field names

### JSONB Query Issues
- **Fix**: Use `->` for JSON object, `->>` for text

```sql
-- Get nested value
SELECT owner->>'name' FROM packages;

-- Query array
SELECT * FROM packages WHERE 'planner' = ANY(types);
```

## 📚 Documentation

- **Full Schema**: See `SCHEMA_UPDATE_SUMMARY.md`
- **Migrations**: See `migrations/README.md`
- **API Reference**: See `API_SCHEMA.md`
- **Contract**: See `API_CONTRACT.md`

---

**Ready to Go!** Your service is now fully compliant with the API_CONTRACT.md specification. 🎉
