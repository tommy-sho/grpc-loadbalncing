## Ingress health check
This directory is for health checking for Ingress resource.

### GKE Ingress
Ingress on GKE, use `/` path for health checking and if `/` could not return `200 OK` status, Ingress evaluate that pod is not health and don't flow traffic to applicable pod.