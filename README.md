# Flower Server

## Overview

The Flower Server is the backend component of a project designed to help you keep track of the flowers you have at home and remind you when they need to be watered. This server is written in Go (Golang) and is designed to run in a Kubernetes cluster.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Deployment](#deployment)

## Features

- **Flower Inventory**: Keep a record of the flowers you have at home.
- **Kubernetes Deployment**: Ready-to-use Kubernetes deployment files included.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) is installed. If not, you can download it [here](https://golang.org/dl/).
- A Kubernetes cluster is set up and configured. You can use tools like [Minikube](https://minikube.sigs.k8s.io/docs/start/) for local development.
- Docker is installed if you plan to build Docker images for your application.

## Getting Started

To get started with Flower Server, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/vavrajosef/flower-server.git


2. Change into the project directory:

    ```bash
    cd flower-server

3. Build and run the application:

    ```bash
    go build
    ./flower-server

Access the Flower Server API at http://localhost:8080.

## Deployment

To deploy Flower Server in a Kubernetes cluster, follow these steps:

1. Navigate to the deployment folder:

    ```bash
    cd deployment

2. Apply the Kubernetes manifests:

    ```bash
    kubectl apply -f .

3. Monitor the deployment and ensure all pods are running:

    ```bash
    kubectl get pods
