# 🚀 Quick v0 Prompt (Copy & Paste This)

Build a modern dashboard for RoboHub Inventory Management with dark theme using Next.js, TypeScript, and shadcn/ui.

## What to Build

A responsive inventory management system for robotics with 5 resource types:
1. **Repositories** (code repos)
2. **Datasets** (training data)
3. **Packages** (software packages)
4. **Scenarios** (test scenarios)
5. **Simulators** (simulation environments)

## Core Features

- ✅ Sidebar navigation with icons for each resource type
- ✅ List view with cards in a grid (3 columns desktop, 2 tablet, 1 mobile)
- ✅ Search and filter by type/tags
- ✅ Create/Edit modals with forms
- ✅ Delete with confirmation
- ✅ Tag system with colored badges
- ✅ Pagination (10 items per page)
- ✅ Dark theme with blue/cyan accents

## API Endpoints

Base URL: `http://localhost:8080/api/v1`

All resources follow REST patterns:
```
GET    /{resource}?limit=10&offset=0  → List
GET    /{resource}/{id}                → Get
POST   /{resource}                     → Create
PUT    /{resource}/{id}                → Update
DELETE /{resource}/{id}                → Delete
```

Resources: `repositories`, `datasets`, `packages`, `scenarios`, `simulators`

## Data Schemas (TypeScript)

### Repository
```typescript
{
  id?: number;
  name: string;              // Required, unique
  url: string;               // Required
  type: string;              // git | svn | mercurial
  description?: string;
  language?: string;         // Go | Python | C++ | JavaScript
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}
```

### Dataset
```typescript
{
  id?: number;
  name: string;              // Required, unique
  description?: string;
  type: string;              // sensor | image | lidar | training
  size?: number;             // Bytes
  location?: string;         // URL
  format?: string;           // rosbag | hdf5 | json | csv
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}
```

### Package
```typescript
{
  id?: number;
  name: string;              // Required, unique
  version: string;           // Required (e.g., "1.0.0")
  description?: string;
  type: string;              // ros | ros2 | python | cpp
  repository?: string;       // Git URL
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}
```

### Scenario
```typescript
{
  id?: number;
  name: string;              // Required, unique
  description?: string;
  type: string;              // simulation | real_world | hybrid
  config?: string;           // JSON string
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}
```

### Simulator
```typescript
{
  id?: number;
  name: string;              // Required, unique
  description?: string;
  type: string;              // gazebo | unity | unreal | carla
  version?: string;
  config?: string;           // JSON string
  tags?: string[];
  created_at?: string;
  updated_at?: string;
}
```

## UI Components

### Sidebar (lucide-react icons)
- GitBranch → Repositories
- Database → Datasets
- Package → Packages
- Map → Scenarios
- Cpu → Simulators

### Resource Card
```
┌──────────────────────┐
│ 📁 Resource Name     │
│ Type: Git | Lang: Go │
│ Description text...  │
│ [tag1] [tag2] [tag3] │
│ 📅 Feb 14, 2024      │
│         ✏️ Edit  🗑️ Delete │
└──────────────────────┘
```

### Create/Edit Modal
- Form with all fields
- Type selectors (dropdowns)
- Tag input with badges
- Validation (required fields marked with *)
- Cancel/Submit buttons

## Color Scheme
```css
Primary: #3B82F6 (blue)
Secondary: #06B6D4 (cyan)
Success: #10B981 (green)
Error: #EF4444 (red)
Background: #0F172A (dark)
Card: #1E293B (dark gray)
```

## Implementation Notes

1. Use `fetch` API for HTTP requests
2. Show loading spinners during API calls
3. Toast notifications for success/errors
4. Empty states with "Create First..." button
5. Debounced search (300ms)
6. Optimistic UI updates
7. Error boundaries

## Example API Call

```typescript
// Fetch repositories
const response = await fetch('http://localhost:8080/api/v1/repositories?limit=10&offset=0');
const data = await response.json();
// Returns: { items: Repository[], total: number }

// Create package
await fetch('http://localhost:8080/api/v1/packages', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: "ros-navigation",
    version: "1.0.0",
    type: "ros2",
    tags: ["navigation", "ros2"]
  })
});
```

## Start With
Build the Repositories page first with:
- List view (cards in grid)
- Search bar
- Create modal
- Edit modal
- Delete confirmation

Then replicate for other resources with their specific fields.

---

**Use shadcn/ui components**: Button, Dialog, Input, Card, Badge, Select, Textarea
**Dark theme with modern look** similar to Vercel/Linear dashboards
