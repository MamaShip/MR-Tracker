FROM alpine:3.15.5
ARG TARGETOS
ARG TARGETARCH
COPY build/MR-Tracker_${TARGETOS}_${TARGETARCH} /usr/local/bin/MR-Tracker
RUN chmod +x /usr/local/bin/MR-Tracker