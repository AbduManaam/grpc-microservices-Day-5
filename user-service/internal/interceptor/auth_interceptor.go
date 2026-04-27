package interceptor

import (
    "context"
    "strings"

    "github.com/golang-jwt/jwt/v5"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
)

var secret = []byte("super-secret-key")

func AuthInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {

    // Skip auth for public RPCs
    if strings.Contains(info.FullMethod, "Login") ||
        strings.Contains(info.FullMethod, "Register") {
        return handler(ctx, req)
    }

    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, status.Error(codes.Unauthenticated, "missing metadata")
    }

    values := md.Get("authorization")
    if len(values) == 0 {
        return nil, status.Error(codes.Unauthenticated, "missing token")
    }

    tokenStr := strings.TrimPrefix(values[0], "Bearer ")

    _, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return secret, nil
    })
    if err != nil {
        return nil, status.Error(codes.Unauthenticated, "invalid token")
    }

    return handler(ctx, req)
}