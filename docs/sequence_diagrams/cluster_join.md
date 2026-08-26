mermaid
---
config:
  theme: redux-color
  look: neo
---
sequenceDiagram
    participant Client
    box WorkerNode
      participant WorkerNode
      participant ContainerGroup
    end
    box MasterNode
      participant MasterNode
      participant KV-Store
      participant Load-Balancer
    end

    WorkerNode->>MasterNode: Request to join cluster at /join endpoint
    MasterNode->>KV-Store: Appends new node
    MasterNode-->>WorkerNode: Gets container config info and returns it to worker
