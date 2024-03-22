# nvidia-smi-exporter
![Docker Pulls](https://img.shields.io/docker/pulls/petrgrishin/nvidia-smi-exporter?style=for-the-badge&color=orange&link=https%3A%2F%2Fhub.docker.com%2Fr%2Fpetrgrishin%2Fnvidia-smi-exporter)

Nvidia SMI metrics exporter for Prometheus

## Start daemon
```
docker run -d --net="host" petrgrishin/nvidia-smi-exporter
```

## Metrics
```
nvidia_fan_speed{gpu="0", name="GeForce GTX 1080 Ti"} 69
nvidia_temperature_gpu{gpu="0", name="GeForce GTX 1080 Ti"} 61
nvidia_clocks_gr{gpu="0", name="GeForce GTX 1080 Ti"} 1860
nvidia_clocks_sm{gpu="0", name="GeForce GTX 1080 Ti"} 1860
nvidia_clocks_mem{gpu="0", name="GeForce GTX 1080 Ti"} 5005
nvidia_power_draw{gpu="0", name="GeForce GTX 1080 Ti"} 262.52
nvidia_utilization_gpu{gpu="0", name="GeForce GTX 1080 Ti"} 100
nvidia_utilization_memory{gpu="0", name="GeForce GTX 1080 Ti"} 80
nvidia_memory_total{gpu="0", name="GeForce GTX 1080 Ti"} 11172
nvidia_memory_free{gpu="0", name="GeForce GTX 1080 Ti"} 10587
nvidia_memory_used{gpu="0", name="GeForce GTX 1080 Ti"} 585
```
