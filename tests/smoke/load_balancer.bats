#!/usr/bin/env bats

certificate_name="l0-$LAYER0_INSTANCE-api"

@test "create environment test" {
    l0 environment create test
}

@test "loadbalancer create test loadbalancer1" {
    l0 loadbalancer create test loadbalancer1
}

@test "loadbalancer list" {
    l0 loadbalancer list
}

@test "loadbalancer get loadbalancer1" {
    l0 loadbalancer get loadbalancer1
}

@test "loadbalancer addport loadbalancer1 8000:8000/http" {
    l0 loadbalancer addport loadbalancer1 8000:8000/http
}

@test "loadbalancer dropport loadbalancer1 8000" {
    l0 loadbalancer dropport loadbalancer1 8000
}

@test "loadbalancer delete loadbalancer1" {
    l0 loadbalancer delete loadbalancer1
}

@test "loadbalancer create --port 80:80/http --port 443:443/https --private --certificate $certificate_name loadbalancer2" {
    l0 loadbalancer create --port 80:80/http --port 443:443/https --private --certificate $certificate_name test loadbalancer2
}

@test "loadbalancer list" {
    l0 loadbalancer list
}

@test "loadbalancer get loadbalancer2" {
    l0 loadbalancer get loadbalancer2
}

@test "loadbalancer get l\*" {
    l0 loadbalancer get l\*
}

@test "loadbalancer delete loadbalancer2" {
    l0 loadbalancer delete loadbalancer2
}

@test "loadbalancer create --healthcheck-target TCP:80 --healthcheck-interval 30 --healthcheck-timeout 5 --healthcheck-healthy-threshold 2 --healthcheck-unhealthy-threshold 2 loadbalancer3" {
    l0 loadbalancer create --healthcheck-target TCP:80 --healthcheck-interval 30 --healthcheck-timeout 5 --healthcheck-healthy-threshold 2 --healthcheck-unhealthy-threshold 2 test loadbalancer3
}

@test "loadbalancer healthcheck loadbalancer3" {
    l0 loadbalancer healthcheck loadbalancer3
}

@test "loadbalancer healthcheck --healthcheck-target TCP:88 --healthcheck-interval 45 --healthcheck-timeout 10 --healthy-threshold 5 --unhealthy-threshold 3 loadbalancer3" {
    l0 loadbalancer healthcheck --healthcheck-target TCP:88 --healthcheck-interval 45 --healthcheck-timeout 10 --healthy-threshold 5 --unhealthy-threshold 3 loadbalancer3
}

@test "loadbalancer delete loadbalancer3" {
    l0 loadbalancer delete loadbalancer3
}

@test "loadbalancer create --port 80:80/http test loadbalancer4" {
    l0 loadbalancer create --port 80:80/http test loadbalancer4
}

@test "loadbalancer list" {
    l0 loadbalancer list
}

@test "loadbalancer get loadbalancer4" {
    l0 loadbalancer get loadbalancer4
}

# this deletes the remaining service(s), load balancer(s), and task(s)
@test "environment delete test" {
    l0 environment delete test
}
