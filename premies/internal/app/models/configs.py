from dataclasses import dataclass


@dataclass
class DatabaseConfig:
    host: str
    port: int
    worker: str
    password: str
    name: str


@dataclass
class GrpcConfig:
    host: str
    port: int
    max_workers: int


@dataclass
class Config:
    database: DatabaseConfig
    grpc: GrpcConfig
