FROM alpine:latest

EXTERN { \
    \"Config\": { \
        \"Image\": \"alpine:latest\", \
        \"Cmd\": [\"sh\", \"-c\", \"echo fizzle5 | tee /tmp/bazz\"] \
    }, \
    \"HostConfig\": { \
        \"Mounts\": [{ \
            \"Source\": \"/tmp\", \
            \"Target\": \"/tmp\", \
            \"Type\": \"bind\" \
        }] \
    } \
}
