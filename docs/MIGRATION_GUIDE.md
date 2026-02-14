# Automatic Database Migrations with GORM

This service uses GORM's AutoMigrate feature to automatically handle database schema changes.

## How It Works

When the service starts, it automatically:

1. **Detects breaking changes** - Checks if the existing schema is incompatible (e.g., old integer IDs vs new UUIDs)
2. **Drops tables if needed** - If breaking changes detected, drops and recreates all tables
3. **Runs GORM AutoMigrate** - Applies schema changes automatically
4. **Loads seed data** - If tables are empty, populates with sample data

## Environment Variables

Control migration behavior with these environment variables:

### `FORCE_DROP_TABLES`

Force drop all tables and recreate them.

```bash
FORCE_DROP_TABLES=true go run cmd/main.go
```

**Use cases:**
- Testing with fresh data
- Applying breaking schema changes
- Resetting development environment

**⚠️ Warning:** This will delete ALL data in the database!

### `LOAD_SEED_DATA`

Force load seed data even if tables aren't empty.

```bash
LOAD_SEED_DATA=true go run cmd/main.go
```

**Use cases:**
- Populating test/development environment
- Adding sample data for demos
- Quick setup for new developers

## Automatic Detection

The system automatically detects when to drop tables by checking:

1. If `FORCE_DROP_TABLES=true` is set
2. If tables exist with old schema (integer IDs instead of UUIDs)

Example check:
```sql
SELECT COUNT(*) 
FROM information_schema.columns 
WHERE table_name = 'repositories' 
AND column_name = 'id' 
AND data_type IN ('integer', 'bigint')
```

If this returns > 0, the system knows the old schema exists and will drop/recreate.

## Migration Process

### Step 1: Connection
```go
db, err := database.Connect(cfg)
```

### Step 2: Detection
```go
if shouldDropTables(db) {
    // Detected old schema or FORCE_DROP_TABLES=true
    dropAllTables(db)
}
```

### Step 3: Enable UUID Extension
```sql
CREATE EXTENSION IF NOT EXISTS "pgcrypto"
```

### Step 4: GORM AutoMigrate
```go
db.AutoMigrate(
    &repository.Repository{},
    &pkg.Package{},
    &scenario.Scenario{},
    &dataset.Dataset{},
    &simulator.Simulator{},
)
```

### Step 5: Load Seed Data (if empty)
```go
if shouldLoadSeedData(db) {
    loadSeedData(db)
}
```

## Seed Data

The system includes realistic sample data:

- **2 Repositories**: ros-planning/navigation2, ros-perception/perception_pcl
- **3 Packages**: nav2_planner, nav2_controller, pcl_ros
- **3 Scenarios**: warehouse navigation, urban driving, object detection
- **3 Datasets**: warehouse data, CARLA urban, indoor objects
- **3 Simulators**: Gazebo, CARLA, Unity

## Usage Examples

### Normal Development Start

```bash
# First time - creates tables and loads seed data
go run cmd/main.go

# Subsequent starts - just connects, no changes
go run cmd/main.go
```

### Fresh Start with Data

```bash
# Drop everything and reload
FORCE_DROP_TABLES=true LOAD_SEED_DATA=true go run cmd/main.go
```

### Migration from Old Schema

```bash
# Automatically detected and handled
go run cmd/main.go
# Output: "Detected old schema with integer IDs - will drop and recreate tables"
```

### Production Deployment

```bash
# Never force drop in production!
# LOAD_SEED_DATA should also be false
go run cmd/main.go
```

For production, you might want to:
1. Backup data first
2. Run migration in maintenance window
3. Verify data after migration

## GORM AutoMigrate Features

GORM's AutoMigrate will:

- ✅ Create tables that don't exist
- ✅ Add missing columns
- ✅ Add missing indexes
- ✅ Create foreign key constraints
- ✅ Add check constraints

GORM's AutoMigrate **will NOT**:

- ❌ Drop columns (keeps old columns even if removed from struct)
- ❌ Modify column types (e.g., can't change varchar(50) to varchar(100))
- ❌ Rename columns
- ❌ Drop tables

This is why we have the `dropAllTables` function for breaking changes.

## Table Drop Order

Tables are dropped in reverse dependency order to handle foreign keys:

1. simulators
2. datasets
3. scenarios
4. packages
5. repositories

## Schema Detection Logic

The system intelligently detects when schema is incompatible:

```go
func shouldDropTables(db *gorm.DB) bool {
    // Check environment variable
    if os.Getenv("FORCE_DROP_TABLES") == "true" {
        return true
    }

    // Check for old schema (integer IDs)
    var count int64
    db.Raw(`
        SELECT COUNT(*) 
        FROM information_schema.columns 
        WHERE table_name = 'repositories' 
        AND column_name = 'id' 
        AND data_type IN ('integer', 'bigint')
    `).Scan(&count)
    
    return count > 0
}
```

## Logs

The migration system provides detailed logging:

```
Running database migrations...
Detected old schema with integer IDs - will drop and recreate tables
Dropped table: simulators
Dropped table: datasets
Dropped table: scenarios
Dropped table: packages
Dropped table: repositories
Migrations completed successfully
Loading seed data...
Seed data loaded successfully: 2 repos, 3 packages, 3 scenarios, 3 datasets, 3 simulators
Database connected and migrated successfully
```

## Docker Compose Usage

In `docker-compose.yml`:

```yaml
services:
  app:
    environment:
      # For development - fresh start every time
      - FORCE_DROP_TABLES=true
      - LOAD_SEED_DATA=true
      
      # For production - no drops, no seed data
      # - FORCE_DROP_TABLES=false
      # - LOAD_SEED_DATA=false
```

## Testing

### Test with fresh database

```bash
# Start postgres
docker run -d --name test-postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=robohub_inventory \
  -p 5432:5432 \
  postgres:15-alpine

# Run with auto-migration
FORCE_DROP_TABLES=true LOAD_SEED_DATA=true go run cmd/main.go

# Verify data
curl http://localhost:8080/api/v1/repositories
```

### Test schema upgrade

```bash
# 1. Start with old schema (simulated)
# 2. Change entity definitions
# 3. Restart application
go run cmd/main.go
# System automatically detects and migrates
```

## Troubleshooting

### "Permission denied" on pgcrypto

**Solution**: Ensure postgres user has SUPERUSER or extension creation privileges.

```sql
ALTER USER postgres CREATEDB;
GRANT ALL PRIVILEGES ON DATABASE robohub_inventory TO postgres;
```

### Tables not dropping

**Solution**: Check for active connections or lingering transactions.

```sql
-- Check active connections
SELECT * FROM pg_stat_activity WHERE datname = 'robohub_inventory';

-- Terminate connections (if needed)
SELECT pg_terminate_backend(pid) FROM pg_stat_activity 
WHERE datname = 'robohub_inventory' AND pid <> pg_backend_pid();
```

### Seed data not loading

**Solution**: Check logs for specific errors. Common issues:
- Foreign key violations (order matters)
- UUID generation issues (pgcrypto not enabled)
- Duplicate data (name uniqueness)

### GORM AutoMigrate failing

**Solution**: Drop tables manually and let system recreate:

```sql
DROP TABLE IF EXISTS simulators, datasets, scenarios, packages, repositories CASCADE;
```

Then restart the application.

## Benefits

1. **Zero-configuration**: Works out of the box
2. **Automatic detection**: Knows when to drop/recreate
3. **Safe defaults**: Won't drop in production unless explicitly told
4. **Fast development**: Quick iterations with `FORCE_DROP_TABLES=true`
5. **Seed data**: Always have test data available
6. **No SQL files**: Everything in Go code

## Migration to Production

When deploying to production:

1. **Backup first**:
```bash
pg_dump -U postgres robohub_inventory > backup_$(date +%Y%m%d).sql
```

2. **Deploy with migrations disabled** (first time):
```bash
# Let it fail on first run to see what would happen
# Then set FORCE_DROP_TABLES=true if needed
```

3. **Verify**:
```bash
psql -U postgres -d robohub_inventory -c "\dt"
psql -U postgres -d robohub_inventory -c "SELECT COUNT(*) FROM repositories"
```

4. **Monitor logs** for any migration errors

## Comparison with SQL Migrations

### Old Way (SQL files)
```bash
psql -f migrations/001_schema.sql
psql -f migrations/002_seed.sql
```

**Pros**: Full control, explicit  
**Cons**: Manual tracking, error-prone, separate from code

### New Way (GORM AutoMigrate)
```go
db, _ := database.Connect(cfg)
// Done! Tables created and migrated automatically
```

**Pros**: Automatic, code-driven, always in sync  
**Cons**: Less control over exact SQL, can't rename columns easily

## When to Use Manual Migrations

Consider manual SQL migrations for:
- Complex data transformations
- Column renames (GORM can't do this)
- Production with zero-downtime requirements
- Multi-step migrations with data preservation

For development and most cases, GORM AutoMigrate is perfect!

---

**No SQL migration files needed - everything happens automatically!** 🎉
