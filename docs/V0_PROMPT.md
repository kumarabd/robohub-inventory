# v0 Prompt: RoboHub Inventory Management Dashboard

Create a modern, responsive dashboard for managing robotics platform inventory including Repositories, Datasets, Packages, Scenarios, and Simulators.

## Design Requirements

Build a clean, professional dashboard with:

1. **Navigation**: Sidebar with icons for each resource type (Repositories, Datasets, Packages, Scenarios, Simulators)
2. **Modern UI**: Use shadcn/ui components with a robotics/tech theme
3. **Color Scheme**: Dark mode with blue/cyan accents, similar to tech dashboards
4. **Layout**: 
   - Left sidebar navigation (collapsible)
   - Main content area with header
   - Responsive grid for cards/lists
5. **Features**:
   - Search and filter for each resource type
   - Create/Edit/Delete operations
   - List view with cards showing key information
   - Detail view modal for each item
   - Tag filtering and management
   - Pagination for lists

## Components Needed

### 1. Dashboard Layout
- Sidebar with navigation links and icons
- Top bar with breadcrumbs and user info
- Main content area

### 2. Resource List Pages (for each type)
- Grid of cards showing resources
- Search bar and filters
- "Create New" button
- Sort options (by date, name)

### 3. Resource Cards
- Show key information
- Action buttons (Edit, Delete, View Details)
- Tag badges
- Type indicators

### 4. Create/Edit Forms
- Modal dialogs with form fields
- Validation
- Tag input with autocomplete
- Submit/Cancel buttons

### 5. Detail View Modal
- Full information display
- Metadata (created_at, updated_at)
- Actions (Edit, Delete)

## API Integration

Base URL: `http://localhost:8080/api/v1`

### Endpoints Structure
All resources follow RESTful patterns:
- `GET /api/v1/{resource}?limit=10&offset=0` - List with pagination
- `GET /api/v1/{resource}/{id}` - Get by ID
- `POST /api/v1/{resource}` - Create new
- `PUT /api/v1/{resource}/{id}` - Update existing
- `DELETE /api/v1/{resource}/{id}` - Delete

### Resource Types
- `repositories` - Code repositories
- `datasets` - Training/sensor data
- `packages` - Software packages
- `scenarios` - Test scenarios
- `simulators` - Simulation environments

## Data Schemas

### 1. Repository
```typescript
interface Repository {
  id?: number;
  name: string;
  url: string;
  type: string;              // "git", "svn", etc.
  description?: string;
  language?: string;         // "Go", "Python", "C++", etc.
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}

// Type options: git, svn, mercurial
// Language options: Go, Python, C++, JavaScript, Rust, Java
```

### 2. Dataset
```typescript
interface Dataset {
  id?: number;
  name: string;
  description?: string;
  type: string;              // "sensor", "image", "lidar", "training"
  size?: number;             // Size in bytes
  location?: string;         // Storage location/URL
  format?: string;           // "rosbag", "hdf5", "json", "csv"
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}

// Type options: sensor, image, lidar, training, validation, test
// Format options: rosbag, hdf5, json, csv, parquet, bag
```

### 3. Package
```typescript
interface Package {
  id?: number;
  name: string;
  version: string;           // Semantic version (e.g., "1.0.0")
  description?: string;
  type: string;              // "ros", "python", "cpp"
  repository?: string;       // Git URL
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}

// Type options: ros, ros2, python, cpp, docker, apt
```

### 4. Scenario
```typescript
interface Scenario {
  id?: number;
  name: string;
  description?: string;
  type: string;              // "simulation", "real_world"
  config?: string;           // JSON configuration string
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}

// Type options: simulation, real_world, hybrid, unit_test, integration_test
```

### 5. Simulator
```typescript
interface Simulator {
  id?: number;
  name: string;
  description?: string;
  type: string;              // "gazebo", "unity", "custom"
  version?: string;
  config?: string;           // JSON configuration string
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}

// Type options: gazebo, unity, unreal, webots, carla, custom
```

## UI Features

### Resource Cards
Each card should display:
- **Icon** based on type
- **Name** (bold, prominent)
- **Type badge** (colored pill)
- **Description** (truncated if long)
- **Tags** (colorful badges)
- **Actions** (Edit, Delete icons)
- **Timestamp** (created date)

### Search & Filter
- Search by name
- Filter by type (dropdown)
- Filter by tags (multi-select)
- Clear filters button

### Forms
Include these fields for each resource:

**Repository Form:**
- Name* (text input)
- URL* (text input with validation)
- Type* (select: git, svn, mercurial)
- Description (textarea)
- Language (select: Go, Python, C++, JavaScript, Rust, Java)
- Tags (tag input)

**Dataset Form:**
- Name* (text input)
- Type* (select: sensor, image, lidar, training, validation, test)
- Description (textarea)
- Size (number input with unit selector: MB, GB, TB)
- Location (text input)
- Format (select: rosbag, hdf5, json, csv, parquet)
- Tags (tag input)

**Package Form:**
- Name* (text input)
- Version* (text input with pattern validation for semver)
- Type* (select: ros, ros2, python, cpp, docker, apt)
- Description (textarea)
- Repository (text input for Git URL)
- Tags (tag input)

**Scenario Form:**
- Name* (text input)
- Type* (select: simulation, real_world, hybrid, unit_test, integration_test)
- Description (textarea)
- Config (code editor textarea for JSON)
- Tags (tag input)

**Simulator Form:**
- Name* (text input)
- Type* (select: gazebo, unity, unreal, webots, carla, custom)
- Version (text input)
- Description (textarea)
- Config (code editor textarea for JSON)
- Tags (tag input)

### Icons
Use appropriate icons from lucide-react:
- Repository: `GitBranch`
- Dataset: `Database`
- Package: `Package`
- Scenario: `Map`
- Simulator: `Cpu` or `Box`

### Interactions
- **Click card** → Open detail modal
- **Click edit** → Open edit form modal
- **Click delete** → Show confirmation dialog
- **Submit form** → Show loading state, then success/error toast
- **Search** → Debounced search (300ms)
- **Filter change** → Immediate re-fetch

## Technical Details

### API Configuration
```typescript
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

// Headers for all requests
const headers = {
  'Content-Type': 'application/json',
};
```

### Error Handling
- Show toast notifications for errors
- Display user-friendly error messages
- Handle network errors gracefully
- Show loading states during API calls

### State Management
- Use React hooks (useState, useEffect)
- Consider using SWR or React Query for data fetching
- Implement optimistic updates for better UX

### Pagination
- Default: 10 items per page
- Show page numbers
- Next/Previous buttons
- Jump to page input

## Example API Calls

### Fetch Repositories
```typescript
// GET /api/v1/repositories?limit=10&offset=0
const response = await fetch(`${API_BASE_URL}/repositories?limit=10&offset=0`);
const data = await response.json();
// Returns: { items: Repository[], total: number }
```

### Create Package
```typescript
// POST /api/v1/packages
const newPackage = {
  name: "ros-navigation",
  version: "1.0.0",
  type: "ros",
  description: "ROS Navigation Stack",
  repository: "https://github.com/ros-planning/navigation2",
  tags: ["ros", "navigation", "robotics"]
};

const response = await fetch(`${API_BASE_URL}/packages`, {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify(newPackage)
});
```

### Update Dataset
```typescript
// PUT /api/v1/datasets/{id}
const updatedDataset = {
  name: "sensor-data-v2",
  type: "lidar",
  format: "rosbag",
  size: 1073741824, // 1GB in bytes
  tags: ["lidar", "outdoor", "autonomous"]
};

const response = await fetch(`${API_BASE_URL}/datasets/${id}`, {
  method: 'PUT',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify(updatedDataset)
});
```

### Delete Simulator
```typescript
// DELETE /api/v1/simulators/{id}
const response = await fetch(`${API_BASE_URL}/simulators/${id}`, {
  method: 'DELETE'
});
```

## Style Guide

### Colors
- Primary: Blue (#3B82F6)
- Secondary: Cyan (#06B6D4)
- Success: Green (#10B981)
- Warning: Yellow (#F59E0B)
- Error: Red (#EF4444)
- Background Dark: #0F172A
- Card Background: #1E293B
- Border: #334155

### Typography
- Headings: Font weight 600-700
- Body: Font weight 400
- Monospace for code/config fields

### Spacing
- Card padding: 1.5rem
- Grid gap: 1.5rem
- Form field spacing: 1rem

### Animations
- Smooth transitions (200-300ms)
- Hover effects on cards and buttons
- Loading spinners for async operations

## Additional Features (Nice to Have)

1. **Dashboard Home**: Overview with statistics and recent items
2. **Bulk Operations**: Select multiple items for batch delete
3. **Import/Export**: JSON import/export functionality
4. **Search History**: Recent searches saved
5. **Favorites**: Star items for quick access
6. **Dark/Light Mode Toggle**: Theme switcher
7. **Activity Log**: Track all changes
8. **Keyboard Shortcuts**: For power users

## Accessibility
- Proper ARIA labels
- Keyboard navigation support
- Focus indicators
- Screen reader friendly
- Color contrast compliance (WCAG AA)

---

**Note**: Start with the core functionality (list, create, edit, delete) for one resource type, then replicate for others. Use TypeScript for type safety and better developer experience.
