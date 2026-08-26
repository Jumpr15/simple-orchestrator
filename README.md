# **Documentation**

## **Design Philosophy**
- The program should be simple to run, and the components of the system and their purposes should be simple to understand
- The program should be opinionated in order to ensure simplicity while also providing a usable service (inflexible and minimal extensibility)

## **System Design**

### **Master Node**:
- Handles external traffic from clients and re-directs to available container instances (via reverse proxy)
- Handles nodes joining the cluster
- Stores cluster config data, cluster-level node data and node-level container data
- Performs heartbeats/container and node state queries on worker nodes

### **Worker Node**:
## **Diagrams**
Context Diagram
<img width="1600" height="1008" alt="Context Diagram (Current) (1)" src="https://github.com/user-attachments/assets/1cceb90a-d03c-42a5-900a-619be08e2053" />

Master App Diagram
<img width="3482" height="1936" alt="Master Node App Diagram (Current) (1)" src="https://github.com/user-attachments/assets/a76f9817-5a99-4ba6-bcb8-c415060dcb9c" />

Worker App Diagram
<img width="2752" height="2416" alt="Worker Node App Diagram (Current) (2)" src="https://github.com/user-attachments/assets/9a066cd2-151f-463b-8dbb-dc2d8308ecd0" />

