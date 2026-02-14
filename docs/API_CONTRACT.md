# RoboHub API Contracts Specification

This document defines the complete API contracts required for integrating repositories, packages, scenarios, and datasets into the RoboHub platform.

---

## 1. REPOSITORIES API

### 1.1 List Repositories
**Endpoint:** `GET /v1/repositories`

**Query Parameters:**
```typescript
{
  limit?: number           // Default: 20, Max: 100
  offset?: number          // Default: 0
  search?: string          // Search by name
  provider?: string        // Filter: "github" | "gitlab" | "bitbucket"
  status?: string          // Filter: "synced" | "syncing" | "needs_attention" | "error"
  autoSync?: boolean       // Filter by auto-sync enabled
}
```

**Response:**
```typescript
{
  repositories: Repository[]
  total: number
  limit: number
  offset: number
  _links?: {
    next?: string
    prev?: string
    self: string
  }
}
```

**Repository Schema:**
```typescript
interface Repository {
  id: string                          // Unique repository ID (UUID)
  name: string                        // Format: "org/repo"
  provider: "github" | "gitlab" | "bitbucket"
  url: string                         // Full repository URL
  description?: string
  defaultBranch: string
  visibility: "public" | "private"
  
  // Sync Information
  lastSynced: string                  // ISO 8601 datetime
  syncStatus: "synced" | "syncing" | "needs_attention" | "error"
  autoSync: boolean
  
  // Latest Commit Info
  latestCommit: {
    hash: string                      // Git commit SHA
    message: string
    author: string
    date: string                      // ISO 8601 datetime
    url: string
  }
  
  // Webhook Configuration
  webhookStatus: "active" | "inactive" | "error"
  webhookId?: string
  
  // Metadata
  tags: string[]
  packageCount: number
  owner: {
    id: string
    name: string
    avatarUrl?: string
  }
  
  // Timestamps
  createdAt: string                   // ISO 8601 datetime
  updatedAt: string                   // ISO 8601 datetime
}
```

---

### 1.2 Get Repository Details
**Endpoint:** `GET /v1/repositories/:id`

**Response:** Single `Repository` object

---

### 1.3 Connect Repository
**Endpoint:** `POST /v1/repositories`

**Request Body:**
```typescript
{
  provider: "github" | "gitlab" | "bitbucket"
  repoId: string              // Provider's repository ID
  repoName: string            // Format: "org/repo"
  url: string
  defaultBranch: string
  autoSync: boolean           // Auto-sync on webhook events
  webhookSecret?: string      // For webhook verification
}
```

**Response:** Single `Repository` object

---

### 1.4 Update Repository Settings
**Endpoint:** `PATCH /v1/repositories/:id`

**Request Body:**
```typescript
{
  autoSync?: boolean
  defaultBranch?: string
  tags?: string[]
}
```

**Response:** Single `Repository` object

---

### 1.5 Manual Sync Repository
**Endpoint:** `POST /v1/repositories/:id/sync`

**Request Body:**
```typescript
{
  branch?: string            // Specific branch to sync (defaults to default branch)
}
```

**Response:**
```typescript
{
  status: "syncing"
  syncId: string            // Identifier for this sync operation
  message: string
  estimatedDuration?: number // Seconds
}
```

---

### 1.6 Get Repository Sync Status
**Endpoint:** `GET /v1/repositories/:id/sync-status`

**Response:**
```typescript
{
  syncId: string
  status: "queued" | "in_progress" | "completed" | "failed"
  progress: number          // 0-100 percentage
  startedAt: string
  completedAt?: string
  errorMessage?: string
  packagesFound: number
  packagesAdded: number
  packagesUpdated: number
}
```

---

### 1.7 Webhook Receiver
**Endpoint:** `POST /v1/webhooks/repository`

**Headers:**
```
X-Provider: "github" | "gitlab" | "bitbucket"
X-Webhook-Signature: string  // HMAC-SHA256 of payload with webhook secret
X-Delivery-ID: string
```

**Request Body:** Provider-specific webhook payload (push, PR, release events)

**Response:**
```typescript
{
  success: boolean
  message: string
  syncId?: string
}
```

---

### 1.8 Repository Activity
**Endpoint:** `GET /v1/repositories/:id/activity`

**Query Parameters:**
```typescript
{
  limit?: number
  offset?: number
  types?: string[]  // "connected" | "sync_completed" | "package_created" | "tags_updated" | "settings_changed"
}
```

**Response:**
```typescript
{
  activities: RepoActivity[]
  total: number
}

interface RepoActivity {
  id: string
  type: "connected" | "sync_completed" | "package_created" | "tags_updated" | "settings_changed"
  message: string
  timestamp: string        // ISO 8601
  metadata?: Record<string, any>
}
```

---

## 2. PACKAGES API

### 2.1 List Packages
**Endpoint:** `GET /v1/packages`

**Query Parameters:**
```typescript
{
  limit?: number                      // Default: 20
  offset?: number                     // Default: 0
  search?: string                     // Free text search
  type?: string[]                     // Multi-select: "planner" | "perception" | "control" | "sensors" | "simulation" | "infrastructure" | "other"
  repoId?: string                     // Filter by repository
  status?: string                     // "pass" | "fail" | "pending"
  sort?: string                       // "updated" | "name" | "validations" | "popular"
  tag?: string[]                      // Multi-select tags
}
```

**Response:**
```typescript
{
  packages: Package[]
  total: number
  limit: number
  offset: number
}
```

**Package Schema:**
```typescript
interface Package {
  id: string                          // UUID
  name: string                        // Package name (lowercase, no spaces)
  displayName: string                 // Human-readable name
  description: string
  documentation?: string              // Markdown content or URL
  
  // Repository Information
  repoId: string
  repoName: string                    // org/repo format
  path: string                        // Path within repo
  
  // Type Classification
  types: Array<"planner" | "perception" | "control" | "sensors" | "simulation" | "infrastructure" | "other">
  
  // Version Information
  latestVersion: string               // Semantic version
  versions: string[]                  // All available versions
  
  // Metadata
  tags: string[]
  keywords: string[]
  
  // Validation Status
  validationStatus: {
    lastValidated: string             // ISO 8601
    status: "pass" | "fail" | "pending"
    passRate: number                  // 0-100 percentage
  }
  
  // Relationships
  linkedScenariosCount: number
  linkedDatasetsCount: number
  usedInCollectionsCount: number
  
  // Owner Information
  owner: {
    id: string
    name: string
    avatarUrl?: string
  }
  
  // Last Run
  lastRun?: {
    status: "pass" | "fail" | "pending"
    runAt: string
    scenarioId: string
  }
  
  // License & Dependencies
  license?: string
  dependencies?: Array<{
    name: string
    version: string
  }>
  
  // Timestamps
  createdAt: string
  updatedAt: string
}
```

---

### 2.2 Get Package Details
**Endpoint:** `GET /v1/packages/:id`

**Response:**
```typescript
{
  ...Package
  documentation: {
    overview: {
      description: string
      capabilities: string[]
      dependencies: string[]
      architectureDiagramUrl?: string
    }
    usage: {
      installSnippet: string
      launchExample: string
      commonParameters: Array<{
        name: string
        type: string
        description: string
        default?: string
      }>
      typicalFailureModes: string[]
    }
    compatibility: {
      rosDistros: string[]
      operatingSystems: string[]
      simulators: string[]
      incompatibleCombos: Array<{
        combo: string
        reason: string
      }>
    }
    limitations: {
      nonGoals: string[]
      knownIssues: string[]
      safetyDisclaimer?: string
    }
  }
  relatedPackages: Package[]
  relatedScenarios: Scenario[]
}
```

---

### 2.3 List Package Versions
**Endpoint:** `GET /v1/packages/:id/versions`

**Query Parameters:**
```typescript
{
  limit?: number
  offset?: number
}
```

**Response:**
```typescript
{
  versions: PackageVersion[]
  total: number
}

interface PackageVersion {
  id: string                          // Version ID (UUID)
  packageId: string
  version: string                     // Semantic version (e.g., "1.2.3")
  commitHash: string
  changelog: string                   // Markdown content
  
  validationSummary: {
    totalRuns: number
    passRate: number                  // 0-100 percentage
    lastRun: string                   // ISO 8601
  }
  
  releaseNotes?: string               // Markdown content
  prerelease?: boolean
  
  createdAt: string
  updatedAt: string
  isCurrent: boolean
}
```

---

### 2.4 Get Package Version Details
**Endpoint:** `GET /v1/packages/:id/versions/:versionId`

**Response:**
```typescript
{
  ...PackageVersion
  validationDetails: {
    totalRuns: number
    passed: number
    failed: number
    pending: number
    
    runsByScenario: Array<{
      scenarioId: string
      scenarioName: string
      passed: number
      failed: number
      lastRunStatus: "pass" | "fail" | "pending"
      lastRunAt: string
    }>
  }
  
  artifacts?: {
    buildUrl?: string
    sourceUrl?: string
    releaseUrl?: string
  }
}
```

---

### 2.5 Search Packages
**Endpoint:** `GET /v1/packages/search`

**Query Parameters:**
```typescript
{
  q: string                           // Search query (required)
  limit?: number
  offset?: number
  filters?: {
    types?: string[]
    minValidationRate?: number
    hasDocumentation?: boolean
  }
}
```

**Response:**
```typescript
{
  results: Package[]
  total: number
  executionTimeMs: number
}
```

---

## 3. SCENARIOS API

### 3.1 List Scenarios
**Endpoint:** `GET /v1/scenarios`

**Query Parameters:**
```typescript
{
  limit?: number
  offset?: number
  search?: string
  category?: Array<"navigation" | "perception" | "localization" | "planning">
  difficulty?: Array<"easy" | "medium" | "hard">
  maintainedBy?: Array<"RoboHub" | "Community" | "Partner">
  verified?: boolean
  simulator?: string
  sort?: "newest" | "popular" | "updated"
}
```

**Response:**
```typescript
{
  scenarios: Scenario[]
  total: number
  limit: number
  offset: number
}

interface Scenario {
  id: string                          // UUID
  name: string
  slug?: string                       // URL-friendly identifier
  description: string
  detailedDescription?: string        // Markdown
  
  // Classification
  category: "navigation" | "perception" | "localization" | "planning"
  difficulty: "easy" | "medium" | "hard"
  maintainedBy: "RoboHub" | "Community" | "Partner"
  verified: boolean
  
  // Content
  whatItTests: string[]               // Bullet points
  whyItMatters: string
  realWorldAnalogs: string[]
  domain: "indoor" | "outdoor" | "urban" | "warehouse" | "mixed"
  
  // Compatibility
  supportedSimulators: string[]       // e.g., ["Gazebo", "CARLA", "AirSim"]
  
  // Related Data
  recommendedDatasets: string[]       // Dataset IDs
  requiredInputs: Array<{
    name: string
    type: string
    description: string
  }>
  
  // Metrics
  successCriteria: Array<{
    name: string
    description: string
    threshold: string
    unit: string
  }>
  passDefinition: string
  
  // Statistics
  weeklyRunCount: number
  monthlyRunCount: number
  usedByPackagesCount: number
  usedByStacksCount: number
  averagePassRate: number             // 0-100 percentage
  
  // Metadata
  tags: string[]
  owner: {
    id: string
    name: string
    avatarUrl?: string
  }
  
  // Timestamps
  createdAt: string
  updatedAt: string
  version: string
}
```

---

### 3.2 Get Scenario Details
**Endpoint:** `GET /v1/scenarios/:id`

**Response:**
```typescript
{
  ...Scenario
  documentation: {
    overview: {
      description: string
      whyItExists: string
      realWorldAnalogs: string[]
      domain: string
    }
    successCriteria: {
      metrics: Array<{
        name: string
        description: string
        threshold: string
        unit: string
      }>
      passDefinition: string
    }
    usage: {
      whenToRun: string[]
      requiredInputs: Array<{
        name: string
        type: string
        description: string
      }>
      typicalDuration: string
      estimatedCostRange?: string
      setup: string                  // Markdown
    }
    limitations: {
      assumptions: string[]
      edgeCasesNotCovered: string[]
      environmentalConstraints: string[]
    }
  }
  
  // Related Entities
  compatiblePackages: Package[]
  recommendedDatasets: Dataset[]
  recentRuns: ScenarioRunSummary[]
}
```

---

### 3.3 Run Scenario
**Endpoint:** `POST /v1/scenarios/:id/runs`

**Request Body:**
```typescript
{
  packageId: string
  packageVersion: string
  datasetId?: string
  parameters?: Record<string, any>
  tags?: string[]
  
  // Optional simulation config
  simulationConfig?: {
    simulator: string
    maxDuration: number               // Seconds
    randomSeed?: number
  }
}
```

**Response:**
```typescript
{
  runId: string                       // UUID
  scenarioId: string
  packageId: string
  status: "queued"
  createdAt: string
  estimatedDuration?: number
  _links: {
    status: string                    // Link to poll for status
    logs: string
  }
}
```

---

### 3.4 Get Scenario Run Status
**Endpoint:** `GET /v1/scenarios/:id/runs/:runId`

**Response:**
```typescript
{
  id: string
  scenarioId: string
  packageId: string
  packageVersion: string
  datasetId?: string
  
  status: "queued" | "running" | "passed" | "failed"
  progress: number                    // 0-100 percentage
  
  startedAt?: string
  completedAt?: string
  duration?: number                   // Seconds
  
  result?: {
    metrics: Array<{
      name: string
      value: number
      threshold: string
      unit: string
      passed: boolean
    }>
    logs: string                      // URL to downloadable logs
    artifacts: Array<{
      name: string
      url: string
      sizeBytes: number
    }>
    errorMessage?: string
    stackTrace?: string
  }
}
```

---

### 3.5 Scenario Run History
**Endpoint:** `GET /v1/scenarios/:id/run-history`

**Query Parameters:**
```typescript
{
  limit?: number
  offset?: number
  packageId?: string
  status?: string
  dateRange?: {
    start: string                     // ISO 8601
    end: string                       // ISO 8601
  }
}
```

**Response:**
```typescript
{
  runs: Array<{
    id: string
    packageId: string
    packageName: string
    packageVersion: string
    status: "pass" | "fail" | "pending"
    passRate?: number
    runAt: string
    duration?: number
    shortOutcome: string
  }>
  total: number
  stats: {
    passRate: number
    avgDuration: number
    totalRuns: number
  }
}
```

---

### 3.6 List Scenario Categories
**Endpoint:** `GET /v1/scenarios/categories`

**Response:**
```typescript
{
  categories: Array<{
    name: "navigation" | "perception" | "localization" | "planning"
    displayName: string
    description: string
    count: number                     // Number of scenarios in category
    icon?: string
  }>
}
```

---

## 4. DATASETS API

### 4.1 List Datasets
**Endpoint:** `GET /v1/datasets`

**Query Parameters:**
```typescript
{
  limit?: number
  offset?: number
  search?: string
  type?: Array<"autonomous-driving" | "robotics" | "indoor-mapping" | "synthetic">
  modality?: Array<"camera" | "lidar" | "radar" | "imu" | "gps" | "multimodal">
  format?: Array<"rosbag2" | "bag" | "parquet" | "custom">
  license?: Array<"MIT" | "Apache-2.0" | "CC-BY" | "CC-BY-NC" | "proprietary">
  sizeFilter?: "small" | "medium" | "large"
  visibility?: "public" | "private"
  sort?: "newest" | "size" | "popularity" | "name"
  owner?: string
}
```

**Response:**
```typescript
{
  datasets: Dataset[]
  total: number
  limit: number
  offset: number
}

interface Dataset {
  id: string                          // UUID
  name: string
  slug?: string
  description: string
  detailedDescription?: string        // Markdown
  
  // Classification
  type: "autonomous-driving" | "robotics" | "indoor-mapping" | "synthetic"
  modality: "camera" | "lidar" | "radar" | "imu" | "gps" | "multimodal"
  format: "rosbag2" | "bag" | "parquet" | "custom"
  license: "MIT" | "Apache-2.0" | "CC-BY" | "CC-BY-NC" | "proprietary"
  
  // Content
  tags: string[]
  whatsInside: string[]               // Bullet points
  usageNotes?: string
  
  // Data Information
  sizeGB: number
  samplesCount: number
  sequencesCount?: number
  duration?: number                   // Seconds
  
  // Compatibility
  supportedScenarios?: string[]       // Scenario IDs
  roboticsPlatforms?: string[]
  
  // Metadata
  source: "uploaded" | "external_link" | "partner"
  ownerType: "user" | "organization"
  ownerId: string
  ownerName: string
  visibility: "public" | "private"
  
  // Preview
  previewAssets?: {
    thumbnailUrl?: string
    sampleFrames: string[]            // URLs to preview images
    videoPreview?: string             // Video URL
  }
  
  // Schema Information
  schema?: {
    topics: Array<{
      name: string
      messageType: string
      frequency: string
      description: string
    }>
    dataSplits?: Array<{
      name: string
      percentage: number
      description: string
    }>
  }
  
  // Statistics
  downloadCount: number
  usedInRuns: number
  avgRating?: number
  ratingCount?: number
  
  // Timestamps
  createdAt: string
  updatedAt: string
}
```

---

### 4.2 Get Dataset Details
**Endpoint:** `GET /v1/datasets/:id`

**Response:**
```typescript
{
  ...Dataset
  documentation: {
    overview: {
      description: string
      collectionMethod: "synthetic" | "recorded" | "mixed"
      environmentCoverage: string[]
      recordingDetails?: {
        platform: string
        duration: string
        locations: string[]
      }
    }
    schema: {
      topics: Array<{
        name: string
        messageType: string
        frequency: string
        description: string
        sampleMessage?: object
      }>
      dataSplits?: Array<{
        name: string
        percentage: number
        description: string
      }>
    }
    usage: {
      recommendedScenarios: string[]
      typicalUseCases: string[]
      performanceConsiderations: string[]
      tutorialUrl?: string
    }
    limitations: {
      biases: string[]
      missingConditions: string[]
      privacyEthicsNotes?: string[]
    }
  }
  
  // Download Information
  downloadUrl?: string
  directStreamUrl?: string
  checksums?: {
    md5: string
    sha256: string
  }
  
  // Related
  relatedDatasets: Dataset[]
  usedByScenarios: Array<{
    scenarioId: string
    scenarioName: string
    runCount: number
  }>
}
```

---

### 4.3 Upload Dataset
**Endpoint:** `POST /v1/datasets`

**Request Body (multipart/form-data):**
```typescript
{
  file: File                          // The dataset file
  
  // Metadata
  name: string
  description: string
  type: "autonomous-driving" | "robotics" | "indoor-mapping" | "synthetic"
  modality: "camera" | "lidar" | "radar" | "imu" | "gps" | "multimodal"
  format: "rosbag2" | "bag" | "parquet" | "custom"
  license: "MIT" | "Apache-2.0" | "CC-BY" | "CC-BY-NC" | "proprietary"
  visibility: "public" | "private"
  
  tags?: string[]
  documentation?: string             // Markdown as text
  
  // Optional metadata
  roboticsPlatforms?: string[]
  relatedScenarios?: string[]
}
```

**Response:**
```typescript
{
  id: string
  status: "draft"
  importId: string                    // For tracking upload progress
  message: string
  _links: {
    upload: string                    // URL to track upload progress
    self: string
  }
}
```

---

### 4.4 Get Upload Progress
**Endpoint:** `GET /v1/datasets/:id/import/:importId`

**Response:**
```typescript
interface DatasetImport {
  id: string
  datasetId?: string
  
  status: "draft" | "uploading" | "processing" | "ready" | "failed"
  progressPct: number                 // 0-100
  
  currentStep: "validating" | "uploading" | "extracting" | "indexing" | "ready"
  stepProgress?: number               // 0-100 for current step
  
  uploadedBytes?: number
  totalBytes?: number
  
  estimatedTimeRemaining?: number     // Seconds
  
  errorMessage?: string
  errorDetails?: string
  
  createdAt: string
  updatedAt: string
}
```

---

### 4.5 Dataset Statistics
**Endpoint:** `GET /v1/datasets/:id/statistics`

**Response:**
```typescript
{
  downloadCount: number
  weeklyDownloads: number
  usedInScenarioRuns: number
  averageRating: number               // 0-5
  ratingCount: number
  
  usageByCategory: Array<{
    category: string
    count: number
  }>
  
  recentActivity: Array<{
    type: "download" | "used_in_run" | "rated" | "commented"
    timestamp: string
    count?: number
  }>
}
```

---

## 5. SEARCH & DISCOVERY API

### 5.1 Global Search
**Endpoint:** `GET /v1/search`

**Query Parameters:**
```typescript
{
  q: string                           // Search query (required)
  types?: Array<"package" | "scenario" | "dataset" | "repository">
  limit?: number
  offset?: number
}
```

**Response:**
```typescript
{
  results: Array<
    | { type: "package"; data: Package }
    | { type: "scenario"; data: Scenario }
    | { type: "dataset"; data: Dataset }
    | { type: "repository"; data: Repository }
  >
  total: number
  facets: {
    packages: number
    scenarios: number
    datasets: number
    repositories: number
  }
}
```

---

## 6. ERROR RESPONSES

All endpoints use standard HTTP status codes and return error responses in this format:

```typescript
{
  status: number
  error: string
  message: string
  code?: string                       // Machine-readable error code
  details?: Record<string, any>       // Additional error context
  requestId?: string                  // For debugging
}
```

**Common Error Codes:**
- `VALIDATION_ERROR` (400): Invalid input parameters
- `UNAUTHORIZED` (401): Missing or invalid authentication
- `FORBIDDEN` (403): Permission denied
- `NOT_FOUND` (404): Resource not found
- `CONFLICT` (409): Resource already exists
- `RATE_LIMITED` (429): Too many requests
- `INTERNAL_ERROR` (500): Server error
- `SERVICE_UNAVAILABLE` (503): Service temporarily down

---

## 7. AUTHENTICATION

All endpoints (except public list endpoints) require authentication via:

**Header Option 1: Bearer Token**
```
Authorization: Bearer <token>
```

**Header Option 2: API Key**
```
X-API-Key: <api_key>
```

**Header Option 3: Agent ID (for internal integrations)**
```
X-Agent-ID: <agent_id>
```

---

## 8. PAGINATION & FILTERING

### Pagination Standards
All list endpoints support:
- `limit` (default: 20, max: 100)
- `offset` (default: 0)

Responses include:
```typescript
{
  items: T[]
  total: number
  limit: number
  offset: number
  _links?: {
    next?: string
    prev?: string
    first?: string
    last?: string
    self: string
  }
}
```

### Filtering Standards
Multi-value filters use array syntax: `?types=planner&types=control`

---

## 9. WEBHOOKS

### Supported Events

**Repository Events:**
- `repository.connected`
- `repository.sync_completed`
- `repository.sync_failed`
- `package.created`
- `package.updated`

**Scenario Events:**
- `scenario.run_completed`
- `scenario.run_failed`

**Dataset Events:**
- `dataset.uploaded`
- `dataset.published`

### Webhook Payload Format
```typescript
{
  id: string                          // Webhook event ID
  type: string                        // Event type
  timestamp: string                   // ISO 8601
  data: Record<string, any>           // Event-specific data
  
  // Retry information
  attempt: number
  maxRetries: number
}
```

---

## 10. RATE LIMITING

All endpoints are rate limited:
- **Public endpoints**: 100 requests/minute
- **Authenticated endpoints**: 1000 requests/minute
- **Webhooks**: 10,000 requests/minute

Response headers:
```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1234567890  // Unix timestamp
```

---

## 11. BATCH OPERATIONS

### Batch Get Packages
**Endpoint:** `POST /v1/packages/batch`

**Request:**
```typescript
{
  ids: string[]                       // Max 50 IDs
}
```

**Response:**
```typescript
{
  packages: Package[]
  notFound: string[]
}
```

### Batch Get Scenarios
**Endpoint:** `POST /v1/scenarios/batch`

**Request:**
```typescript
{
  ids: string[]                       // Max 50 IDs
}
```

**Response:**
```typescript
{
  scenarios: Scenario[]
  notFound: string[]
}
```

---

## 12. VERSIONING

API version is specified in the URL: `/v1/`, `/v2/`, etc.

Current version: **v1**

Deprecation policy:
- Deprecated endpoints will include header: `Deprecation: true`
- Sunset header indicates removal date: `Sunset: Wed, 31 Dec 2025 23:59:59 GMT`
- Minimum 6-month notice before removal

---

## 13. MONITORING & OBSERVABILITY

### Health Check
**Endpoint:** `GET /health`

**Response:**
```typescript
{
  status: "ok" | "degraded" | "down"
  version: string
  timestamp: string
  services: {
    database: "ok" | "degraded" | "down"
    cache: "ok" | "degraded" | "down"
    storage: "ok" | "degraded" | "down"
  }
}
```

### Status Page
Documentation for current incidents: `https://status.robohub.ai`

---

## 14. IMPLEMENTATION CHECKLIST

### Phase 1: Core Entities
- [ ] Repository API (connect, sync, list, get)
- [ ] Package API (list, get, versions)
- [ ] Scenario API (list, get, run)
- [ ] Dataset API (list, get, upload)

### Phase 2: Advanced Features
- [ ] Global search
- [ ] Webhooks
- [ ] Batch operations
- [ ] Advanced filtering

### Phase 3: Analytics
- [ ] Run statistics
- [ ] Dataset statistics
- [ ] Usage analytics
- [ ] Reporting endpoints

---

## 15. EXAMPLE USAGE

### Connect and Sync a Repository
```bash
# Step 1: Connect repository
curl -X POST https://api.robohub.ai/v1/repositories \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "provider": "github",
    "repoId": "12345",
    "repoName": "org/repo",
    "url": "https://github.com/org/repo",
    "defaultBranch": "main",
    "autoSync": true
  }'

# Step 2: Monitor sync status
curl https://api.robohub.ai/v1/repositories/{id}/sync-status \
  -H "Authorization: Bearer <token>"
```

### Run a Package Against a Scenario
```bash
curl -X POST https://api.robohub.ai/v1/scenarios/{scenarioId}/runs \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "packageId": "pkg-123",
    "packageVersion": "1.0.0",
    "datasetId": "dataset-456"
  }'
```

### Upload a Dataset
```bash
curl -X POST https://api.robohub.ai/v1/datasets \
  -H "Authorization: Bearer <token>" \
  -F "file=@dataset.bag" \
  -F "name=My Dataset" \
  -F "type=robotics" \
  -F "modality=multimodal" \
  -F "format=bag" \
  -F "license=MIT"
```

---

## Version History

- **v1.0.0** (2024): Initial API specification
  - Core CRUD operations for repositories, packages, scenarios, datasets
  - Search and discovery
  - Webhook support
  - Rate limiting and pagination

---

