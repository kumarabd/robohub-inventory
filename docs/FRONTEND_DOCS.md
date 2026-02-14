# Frontend Integration Documentation

This directory contains complete documentation for building a frontend UI for the RoboHub Inventory Service.

## 📚 Documentation Files

### For v0.dev Users

1. **[V0_QUICK_PROMPT.md](./V0_QUICK_PROMPT.md)** ⭐ **START HERE**
   - Quick copy-paste prompt for v0.dev
   - Concise with all essential information
   - Best for getting started quickly

2. **[V0_PROMPT.md](./V0_PROMPT.md)**
   - Comprehensive prompt with detailed requirements
   - Complete UI/UX specifications
   - Implementation guidelines and best practices

3. **[DESIGN_GUIDE.md](./DESIGN_GUIDE.md)**
   - Visual design system
   - Component mockups (ASCII art)
   - Color palette and typography
   - Responsive layouts

### For Developers

4. **[API_SCHEMA.md](./API_SCHEMA.md)**
   - Complete API reference
   - Request/response examples
   - TypeScript interfaces
   - cURL examples for testing
   - Error codes and handling

## 🚀 Quick Start

### Option 1: Use v0.dev (Fastest)

1. Copy the entire content of **[V0_QUICK_PROMPT.md](./V0_QUICK_PROMPT.md)**
2. Paste into [v0.dev](https://v0.dev)
3. Let AI generate the UI
4. Copy generated code to your project
5. Update API base URL to: `http://localhost:8080/api/v1`

### Option 2: Build Manually

1. Read **[API_SCHEMA.md](./API_SCHEMA.md)** for API details
2. Follow **[DESIGN_GUIDE.md](./DESIGN_GUIDE.md)** for styling
3. Use **[V0_PROMPT.md](./V0_PROMPT.md)** for requirements
4. Start with one resource type (Repositories)
5. Replicate for other resources

## 🏗️ Tech Stack Recommendations

### Recommended
- **Framework**: Next.js 14+ (App Router)
- **Language**: TypeScript
- **UI Library**: shadcn/ui
- **Styling**: Tailwind CSS
- **Icons**: lucide-react
- **Data Fetching**: SWR or TanStack Query

### Alternative
- React + Vite
- Vue 3 + Nuxt
- Svelte + SvelteKit

## 📋 Features to Implement

### Core Features (MVP)
- ✅ List view for all 5 resource types
- ✅ Create new resources
- ✅ Edit existing resources
- ✅ Delete resources (with confirmation)
- ✅ Search functionality
- ✅ Filter by type and tags
- ✅ Pagination

### Nice-to-Have Features
- ⭐ Dashboard home with statistics
- 🔍 Advanced search with multiple filters
- 📊 Data visualization (charts)
- 📥 Import/Export (JSON)
- ⌨️ Keyboard shortcuts
- 🌓 Dark/Light mode toggle
- 📱 Mobile app (React Native)

## 🎨 Design System

### Colors
```
Primary:    #3B82F6 (Blue)
Secondary:  #06B6D4 (Cyan)
Success:    #10B981 (Green)
Warning:    #F59E0B (Yellow)
Error:      #EF4444 (Red)
Background: #0F172A (Dark Blue)
Card:       #1E293B (Dark Gray)
```

### Typography
- Font Family: Inter or System UI
- Headings: 600-700 weight
- Body: 400 weight

### Spacing
- Base unit: 4px
- Grid gap: 24px (1.5rem)
- Card padding: 24px (1.5rem)

## 🔌 API Connection

### Local Development
```typescript
const API_BASE_URL = 'http://localhost:8080/api/v1';
```

### Production
```typescript
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'https://your-api.com/api/v1';
```

### Docker Container
If using the published Docker image:
```bash
docker run -p 8080:8080 \
  -e DB_HOST=your-db \
  -e DB_PASSWORD=your-password \
  ghcr.io/kumarabd/robohub-inventory:latest
```

## 📖 Resource Types Overview

| Resource | Icon | Description | Example |
|----------|------|-------------|---------|
| Repository | 📁 | Code repositories | Git/SVN repos |
| Dataset | 📊 | Training/sensor data | ROS bags, images |
| Package | 📦 | Software packages | ROS, Python, C++ |
| Scenario | 🗺️ | Test scenarios | Simulation tests |
| Simulator | 🖥️ | Simulation environments | Gazebo, Unity |

## 🔍 Example Queries

### List Repositories
```bash
curl http://localhost:8080/api/v1/repositories?limit=10&offset=0
```

### Create Package
```bash
curl -X POST http://localhost:8080/api/v1/packages \
  -H "Content-Type: application/json" \
  -d '{
    "name": "ros-navigation",
    "version": "1.0.0",
    "type": "ros2",
    "tags": ["navigation", "ros2"]
  }'
```

### Update Dataset
```bash
curl -X PUT http://localhost:8080/api/v1/datasets/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "updated-name", "type": "sensor"}'
```

### Delete Scenario
```bash
curl -X DELETE http://localhost:8080/api/v1/scenarios/1
```

## 🧪 Testing the API

### Start the Service
```bash
# Using Docker
docker run -p 8080:8080 ghcr.io/kumarabd/robohub-inventory:latest

# Or locally
make run
```

### Health Check
```bash
curl http://localhost:8080/health
```

Should return:
```json
{
  "status": "healthy",
  "timestamp": "2024-02-14T10:30:00Z"
}
```

## 📱 Responsive Design

### Breakpoints
- Mobile: < 640px (1 column)
- Tablet: 640-1024px (2 columns)
- Desktop: > 1024px (3-4 columns)

### Mobile Considerations
- Collapsible sidebar
- Stacked cards (1 column)
- Touch-friendly buttons (min 44px)
- Bottom navigation option

## 🔐 Security Notes

### CORS
The API needs CORS enabled for frontend access. Add to server config:
```go
AllowOrigins: ["http://localhost:3000", "https://your-frontend.com"]
```

### Authentication (Future)
Currently no authentication. Consider adding:
- JWT tokens
- OAuth2
- API keys

## 🐛 Common Issues

### CORS Errors
**Problem**: `Access-Control-Allow-Origin` error  
**Solution**: Configure CORS on the backend or use a proxy

### 404 Not Found
**Problem**: API endpoint not found  
**Solution**: Check API base URL and endpoint paths

### Network Error
**Problem**: Cannot connect to API  
**Solution**: Ensure backend service is running on port 8080

## 📚 Additional Resources

- [Backend README](./README.md) - Service documentation
- [GitHub Actions Workflows](./.github/workflows/README.md) - CI/CD documentation
- [Quick Start Guide](./.github/QUICKSTART.md) - Setup instructions

## 🤝 Contributing

When building the frontend:
1. Follow the API schema exactly
2. Handle all error cases
3. Show loading states
4. Add proper TypeScript types
5. Write clean, maintainable code
6. Test with real API calls

## 📞 Support

For questions or issues:
1. Check [API_SCHEMA.md](./API_SCHEMA.md) for API details
2. Review [V0_PROMPT.md](./V0_PROMPT.md) for requirements
3. See [DESIGN_GUIDE.md](./DESIGN_GUIDE.md) for UI guidance
4. Open an issue in the repository

---

**Ready to build?** Start with [V0_QUICK_PROMPT.md](./V0_QUICK_PROMPT.md)! 🚀
