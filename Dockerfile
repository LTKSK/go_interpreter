# ===================================
# build stage
FROM golang:1.15.2 
#AS builder

WORKDIR /go/src/go_interpreter
COPY . .


# ===================================
# bundle stage
#JFROM scratch

#JCOPY --from=builder /

