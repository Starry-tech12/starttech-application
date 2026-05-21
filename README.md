# StartTech Application (Frontend + Backend CI/CD)

## 🧠 Overview

This repository contains the full application:

- React Frontend (S3 + CloudFront)
- Golang Backend (Docker + EC2 + ASG)
- CI/CD pipelines using GitHub Actions

---

# 🏗️ Architecture

Frontend:
- React app
- Hosted on S3
- Distributed via CloudFront

Backend:
- Golang API
- Dockerized
- Runs on EC2 Auto Scaling Group
- Behind Application Load Balancer

Database:
- MongoDB Atlas

Cache:
- Redis (ElastiCache)

---

# 🚀 CI/CD PIPELINES

## Frontend Pipeline

Triggered on push to `main`:

- Install dependencies
- Run tests
- Build React app
- Security scan (npm audit)
- Deploy to S3
- Invalidate CloudFront cache

---

## Backend Pipeline

Triggered on push to `main`:

- Run Go tests
- Build binary
- Build Docker image
- Push to ECR
- Deploy to EC2
- Run health checks

---

# ⚙️ Deployment Flow

1. Code pushed to GitHub
2. GitHub Actions runs CI/CD
3. Infrastructure already provisioned via Terraform
4. Backend deployed via Docker
5. Frontend deployed to S3
6. CloudFront serves users globally

---

# 🧪 Health Check

Backend endpoint:
GET /health → OK


---

# 🔐 Security

- Secrets stored in GitHub Secrets
- IAM roles for EC2 access
- No credentials in code
- Vulnerability scanning enabled

---

# 📊 Monitoring

- CloudWatch Logs for backend
- EC2 metrics
- ALB metrics
- Deployment logs via GitHub Actions

---

# 🧾 Scripts

- deploy-frontend.sh
- deploy-backend.sh
- health-check.sh
- rollback.sh

---

# 📌 Notes

Ensure the following before deployment:

- AWS credentials configured
- S3 bucket exists
- CloudFront distribution ID set
- EC2 security groups allow port 8080