# Chaos Engineering Special Academy

A comprehensive learning repository for **Chaos Engineering** practices using **ChaosToolkit** and **Toxiproxy** to build resilient distributed systems.

## 🎯 Overview

This repository serves as a hands-on learning platform for **Chaos Engineering** - the practice of intentionally introducing controlled failures to test and improve system resilience. The project demonstrates how to systematically inject network faults and measure system behavior under stress conditions.

### What is Chaos Engineering?

Chaos Engineering is a methodology for testing distributed systems by deliberately introducing controlled failures to:
- Discover hidden weaknesses in production systems
- Improve observability and monitoring capabilities  
- Ensure systems remain functional when components fail
- Foster a culture of reliability-focused engineering

## 📁 Repository Structure

```
spe-chaos-eng/
├── chaostoolkit/           # ChaosToolkit experiments and configuration
│   ├── experiments/        # Organized chaos experiments
│   │   ├── 01-jsonplaceholder/  # Basic API chaos testing
│   │   └── 02-microchaos/       # Microservices chaos testing
│   ├── docker-compose.yml  # ChaosToolkit + Toxiproxy setup
│   └── Dockerfile          # ChaosToolkit container configuration
├── microchaos/             # Sample microservices application
│   ├── order-api/          # Order management service (Go)
│   ├── notifier/           # Notification service (Go)
│   ├── mysql/              # Database initialization scripts
│   ├── docker-compose.yml  # Microservices stack
│   └── seed.sh             # Toxiproxy proxy configuration
├── assesment/              # Learning assessments and materials
└── README.md               # This file
```

### 📂 Detailed Folder Structure

#### `chaostoolkit/` - Chaos Engineering Experiments
Contains all ChaosToolkit experiments organized by complexity and scope:

- **`experiments/01-jsonplaceholder/`** - Beginner-friendly experiments
  - Basic API testing with JSONPlaceholder
  - Introduction to Toxiproxy integration
  - Various network fault injection scenarios (latency, bandwidth, timeout, etc.)

- **`experiments/02-microchaos/`** - Advanced microservices experiments
  - Real-world microservices chaos testing
  - Database connection fault injection
  - Comprehensive analysis and reporting

#### `microchaos/` - Sample Application Stack
A complete microservices application for chaos testing:

- **`order-api/`** - Order management service written in Go
- **`notifier/`** - Event-driven notification service
- **`mysql/`** - Database schema and initialization
- **Supporting infrastructure** - Docker Compose, Toxiproxy configuration

#### `assesment/` - Learning Materials
Assessment materials and learning resources for the Special Academy program.

## 🛠️ Prerequisites

Before starting with chaos engineering experiments, ensure you have:

- **Docker** - Container runtime for isolated environments
- **Docker Compose** - Multi-container orchestration
- **Git** - Version control (to clone this repository)

### Installation Links
- [Docker Installation Guide](https://docs.docker.com/get-docker/)
- [Docker Compose Installation Guide](https://docs.docker.com/compose/install/)

## 🚀 Quick Start

1. **Clone the repository**
   ```bash
   git clone https://gitlab.spesolution.net/data/rnd/special-academy-chaos-engineering.git
   cd special-academy-chaos-engineering
   ```

2. **Set up ChaosToolkit environment**
   ```bash
   cd chaostoolkit
   docker-compose up -d
   ```

3. **Set up microservices stack**
   ```bash
   cd ../microchaos
   docker-compose up -d
   ```

4. **Run your first experiment**
   ```bash
   cd ../chaostoolkit
   chaos run experiments/01-jsonplaceholder/01.01-simple-hit-dummy-api.json
   ```

## 📚 Learning Path

### Phase 1: Basic API Chaos Testing
Start with `experiments/01-jsonplaceholder/` to learn:
- ChaosToolkit fundamentals
- Toxiproxy integration
- Basic network fault injection
- Experiment documentation and reporting

### Phase 2: Microservices Chaos Testing  
Progress to `experiments/02-microchaos/` for:
- Real-world application testing
- Database resilience testing
- Advanced fault scenarios
- Performance impact analysis

### Phase 3: Custom Experiments
Build your own experiments based on:
- Application-specific requirements
- Production failure scenarios
- Team-specific resilience goals

## 🔬 Experiment Types

The repository includes experiments for various failure scenarios:

| Failure Type | Description | Impact Assessment |
|--------------|-------------|-------------------|
| **Latency** | Network delay simulation | Response time degradation |
| **Bandwidth** | Connection speed limitation | Throughput reduction |
| **Timeout** | Connection timeout simulation | Request failures |
| **Packet Loss** | Data transmission failures | Partial data corruption |
| **Connection Issues** | Slow close, connection drops | Resource exhaustion |
| **Rate Limiting** | Request throttling | Service availability |

## 📊 Expected Outcomes

After completing experiments in this repository, you will:

- ✅ Understand chaos engineering principles and methodology
- ✅ Design and execute safe chaos experiments
- ✅ Build shared knowledge about system resilience
- ✅ Improve reliability across your projects
- ✅ Implement proper monitoring and alerting
- ✅ Apply fault tolerance patterns (retry, circuit breaker, etc.)

## 📖 Documentation

Detailed documentation is available in each experiment folder:

- **[ChaosToolkit Setup Guide](./chaostoolkit/README.md)** - Complete environment setup
- **[Microchaos Project Guide](./microchaos/README.md)** - Application architecture and testing
- **[JSONPlaceholder Experiments](./chaostoolkit/experiments/01-jsonplaceholder/README.md)** - Basic chaos testing
- **[Microservices Experiments](./chaostoolkit/experiments/02-microchaos/README.md)** - Advanced scenarios

## 🤝 Contributing

Contributions to improve experiments, documentation, or add new scenarios are welcome. Please:

1. Fork the repository
2. Create a feature branch
3. Add your experiments or improvements
4. Submit a pull request with clear documentation

---

**Ready to make your systems more resilient?** Start with the basic experiments and gradually work your way up to complex microservices chaos testing. Remember: the goal is not to break systems, but to make them stronger through controlled failure testing.