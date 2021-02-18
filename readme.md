# bussindex 

bussindex = swedish for Bus factor  
*The bus factor is a measurement of the risk resulting from information and capabilities not being shared among team members, derived from the phrase "in case they get hit by a bus."* - Wikipedia

Simple idea with fun and powerful tech. In other words - a Demo. <span style="color:magenta">**(work-in-progress)** </span>

This demo application includes:  

- React hooks + Redux (RTK)  
- Golang
- Neo4j  
- Kubernetes
- [Skaffold](https://skaffold.dev/)


## You will need
* Docker and Kubernetes (e.g Docker-Desktop with kubernetes enabled, see settings.)  
* [Skaffold](https://skaffold.dev/)
* [Helm](https://helm.sh/docs/intro/install/)
* Go (for the backend development)
* Node (for the frontend development)


## Init (one time only)
1) `` helm repo add equinor-charts https://equinor.github.io/helm-charts/charts/ ``   
2) `` helm repo update ``  

## Development environment
start the backend with ``skaffold dev`` (in the root folder)  
start the frontend with ``npm run dev``  (in ./web)  

Webpack will restart (and refresh) your webb app when you save a change in the web application code and Skaffold will restart (build images and deploy to your k8s cluster) your API when you save a 
change in the api code.

When the system is running you can inspect your Neo4j database by forwarding the ports ``kubectl port-forward <name of neo4j pod> 7474:7474 7687:7687`` and navigate to http://localhost:7474 and login (see credentials in ./api/cmd/main.go, hard coded for now..).
