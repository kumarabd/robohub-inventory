# Visual Design Guide for v0

Quick reference for the visual design and layout of the RoboHub Inventory Dashboard.

## Color Palette

```css
/* Primary Colors */
--primary-blue: #3B82F6;
--primary-cyan: #06B6D4;

/* Status Colors */
--success: #10B981;
--warning: #F59E0B;
--error: #EF4444;

/* Dark Theme */
--bg-dark: #0F172A;
--bg-card: #1E293B;
--bg-hover: #334155;
--border: #334155;
--text-primary: #F1F5F9;
--text-secondary: #94A3B8;
```

## Layout Structure

```
┌─────────────────────────────────────────────────────────┐
│                      Top Bar                            │
│  🤖 RoboHub Inventory        Search...    👤 User      │
├──────┬──────────────────────────────────────────────────┤
│      │                                                   │
│  📁  │  Repositories                    [+ New Repo]   │
│  📊  │  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━  │
│  📦  │                                                   │
│  🗺️  │  ┌─────────────┐  ┌─────────────┐  ┌──────────┐│
│  🖥️  │  │   📁 Repo 1 │  │   📁 Repo 2 │  │ 📁 Repo 3││
│      │  │             │  │             │  │          ││
│ Nav  │  │  Git | Go   │  │  SVN | C++  │  │ Git | Py ││
│      │  │  [ros] [ai] │  │  [robotics] │  │ [ml]     ││
│      │  │             │  │             │  │          ││
│      │  │  ✏️  🗑️       │  │  ✏️  🗑️       │  │ ✏️  🗑️    ││
│      │  └─────────────┘  └─────────────┘  └──────────┘│
│      │                                                   │
│      │  ┌─────────────┐  ┌─────────────┐  ┌──────────┐│
│      │  │   📁 Repo 4 │  │   📁 Repo 5 │  │ 📁 Repo 6││
│      │  └─────────────┘  └─────────────┘  └──────────┘│
│      │                                                   │
│      │  ◄ 1 2 3 4 ►                                    │
└──────┴──────────────────────────────────────────────────┘
```

## Component Designs

### 1. Sidebar Navigation

```
┌─────────────┐
│ 🤖 RoboHub  │
├─────────────┤
│ 📁 Repos    │ ← Selected (blue bg)
│ 📊 Datasets │
│ 📦 Packages │
│ 🗺️ Scenarios│
│ 🖥️ Simulator│
├─────────────┤
│ ⚙️ Settings │
└─────────────┘
```

### 2. Resource Card

```
┌──────────────────────────┐
│ 📁 ros-navigation        │ ← Icon + Name
│ ━━━━━━━━━━━━━━━━━━━━━━━ │
│ Type: Git | Language: C++│ ← Metadata
│                          │
│ ROS 2 Navigation Stack   │ ← Description
│ with advanced path...    │
│                          │
│ [ros2] [nav] [autonomy]  │ ← Tags
│                          │
│ 📅 Feb 14, 2024          │ ← Date
│                          │
│           ✏️ Edit  🗑️ Delete│ ← Actions
└──────────────────────────┘
```

### 3. Create/Edit Form Modal

```
┌─────────────────────────────────┐
│  Create New Repository      ✕   │
├─────────────────────────────────┤
│                                 │
│  Name *                         │
│  ┌─────────────────────────┐   │
│  │ ros-navigation           │   │
│  └─────────────────────────┘   │
│                                 │
│  URL *                          │
│  ┌─────────────────────────┐   │
│  │ https://github.com/...   │   │
│  └─────────────────────────┘   │
│                                 │
│  Type *                         │
│  ┌─────────────────────────┐   │
│  │ Git ▼                    │   │
│  └─────────────────────────┘   │
│                                 │
│  Language                       │
│  ┌─────────────────────────┐   │
│  │ C++ ▼                    │   │
│  └─────────────────────────┘   │
│                                 │
│  Description                    │
│  ┌─────────────────────────┐   │
│  │                          │   │
│  │                          │   │
│  └─────────────────────────┘   │
│                                 │
│  Tags                           │
│  ┌─────────────────────────┐   │
│  │ [ros2] [nav] + Add tag   │   │
│  └─────────────────────────┘   │
│                                 │
│         [Cancel]  [Create]      │
└─────────────────────────────────┘
```

### 4. Detail View Modal

```
┌─────────────────────────────────┐
│  Repository Details         ✕   │
├─────────────────────────────────┤
│                                 │
│  📁 ros-navigation              │
│                                 │
│  ━━━━━━━━━━━━━━━━━━━━━━━━━━━━ │
│                                 │
│  Name:        ros-navigation    │
│  Type:        Git               │
│  Language:    C++               │
│  URL:         https://github... │
│                                 │
│  Description:                   │
│  ROS 2 Navigation Stack with   │
│  advanced path planning...      │
│                                 │
│  Tags:                          │
│  [ros2] [navigation] [autonomy] │
│                                 │
│  Created:     Feb 14, 2024      │
│  Updated:     Feb 14, 2024      │
│                                 │
│         [Edit]  [Delete]        │
└─────────────────────────────────┘
```

### 5. Search & Filter Bar

```
┌──────────────────────────────────────────┐
│  🔍 Search...            Type: All ▼     │
│                          Tags: All ▼     │
│                          [Clear Filters] │
└──────────────────────────────────────────┘
```

### 6. Tag Badge Styles

```
[ros2]         ← Blue (#3B82F6)
[navigation]   ← Cyan (#06B6D4)
[ai]           ← Purple (#8B5CF6)
[robotics]     ← Green (#10B981)
[simulation]   ← Yellow (#F59E0B)
```

## Icon Mapping

| Resource  | Icon Name (lucide-react) | Color |
|-----------|-------------------------|-------|
| Repository| `GitBranch`             | Blue  |
| Dataset   | `Database`              | Cyan  |
| Package   | `Package`               | Purple|
| Scenario  | `Map`                   | Green |
| Simulator | `Cpu`                   | Orange|

## Typography Scale

```
H1 (Page Title):    2xl (1.5rem) - font-bold
H2 (Section):       xl (1.25rem) - font-semibold
H3 (Card Title):    lg (1.125rem) - font-semibold
Body:               base (1rem) - font-normal
Small:              sm (0.875rem) - font-normal
Tiny:               xs (0.75rem) - font-normal
```

## Responsive Breakpoints

```
Mobile:     < 640px  (1 column)
Tablet:     640-1024px (2 columns)
Desktop:    > 1024px (3-4 columns)
```

## Card Grid Layout

```css
/* Desktop: 3 columns */
.grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1.5rem;
}

/* Tablet: 2 columns */
@media (max-width: 1024px) {
  .grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

/* Mobile: 1 column */
@media (max-width: 640px) {
  .grid {
    grid-template-columns: 1fr;
  }
}
```

## Hover States

```
Card:       Scale 1.02 + shadow-lg
Button:     Brightness 110%
Link:       Underline + color shift
```

## Loading States

```
┌──────────────────────────┐
│                          │
│    ⟳ Loading...          │ ← Spinner
│                          │
└──────────────────────────┘

Or skeleton cards:
┌──────────────────────────┐
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓           │
│ ▓▓▓▓▓▓ ▓▓▓▓              │
│                          │
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓     │
│ ▓▓▓▓▓▓▓▓▓▓               │
└──────────────────────────┘
```

## Empty States

```
┌──────────────────────────┐
│                          │
│         📁               │
│                          │
│   No repositories yet    │
│                          │
│   [Create First Repo]    │
│                          │
└──────────────────────────┘
```

## Toast Notifications

```
✓ Success: Repository created successfully
⚠ Warning: Connection slow
✕ Error: Failed to delete package
ℹ Info: Changes saved
```

## Button Styles

```
Primary:    Blue bg, white text, shadow
Secondary:  Transparent, border, text
Danger:     Red bg, white text
Ghost:      Transparent, no border
```

## Form Validation

```
Valid:      Green border, ✓ icon
Invalid:    Red border, ✕ icon + error message
Required:   Label with red asterisk *
```

## Accessibility Features

- Focus ring: 2px solid blue
- Skip to main content link
- ARIA labels on all interactive elements
- Keyboard navigation (Tab, Enter, Esc)
- High contrast mode support

## Animation Timings

```
Fast:       150ms (hover effects)
Normal:     200ms (transitions)
Slow:       300ms (page transitions)
Easing:     cubic-bezier(0.4, 0, 0.2, 1)
```

## Shadow System

```
sm:   0 1px 2px 0 rgb(0 0 0 / 0.05)
md:   0 4px 6px -1px rgb(0 0 0 / 0.1)
lg:   0 10px 15px -3px rgb(0 0 0 / 0.1)
xl:   0 20px 25px -5px rgb(0 0 0 / 0.1)
```

## Best Practices for v0

1. **Start simple**: Build one resource type first (Repositories)
2. **Use shadcn/ui**: Dialog, Button, Input, Badge, Card components
3. **Mobile-first**: Design for mobile, enhance for desktop
4. **Consistent spacing**: Use 4px/8px grid system
5. **Loading states**: Always show feedback during async operations
6. **Error boundaries**: Graceful error handling everywhere
7. **Optimistic UI**: Update UI immediately, revert on error
8. **Keyboard shortcuts**: Add power user features (Cmd+K for search)

---

Copy this entire document along with V0_PROMPT.md and API_SCHEMA.md to v0.dev for the best results!
