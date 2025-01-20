graph TD
    A[API REST] --> B[Authentication JWT]
    
    B --> C[Endpoints CRUD]
    
    C --> D[POST /candidates]
    C --> E[GET /candidates]
    C --> F[GET /candidates/:id]
    C --> G[PUT /candidates/:id]
    C --> H[DELETE /candidates/:id]
    
    D --> I[Validación Token]
    E --> I
    F --> I
    G --> I
    H --> I
    
    I --> J[Service Layer]
    J --> K[Repository Layer]
    K --> L[MySQL Database]
    
    subgraph "Response Codes"
    M[200 OK - Success]
    N[201 Created]
    O[400 Bad Request]
    P[401 Unauthorized]
    Q[404 Not Found]
    R[500 Server Error]
    end