import { keyvault, pingpong, redis, sql } from "./pingpong";

export const services = [
  {
    id: "frontend",
    name: "Frontend v3",
    description: "React + Vite",
    endpoint: null,
  },
  {
    id: "backend",
    name: "Backend",
    description: "Golang",
    endpoint: pingpong,
  },
  {
    id: "sql",
    name: "Postgres SQL",
    description: "Database connection",
    endpoint: sql,
  },
  {
    id: "redis",
    name: "Redis Cache",
    description: "Cache connection",
    endpoint: redis,
  },
  {
    id: "keyvault",
    name: "Azure Key Vault - Example API Key",
    description: "Secrets connection",
    endpoint: keyvault,
  },
];