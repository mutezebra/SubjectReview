namespace go base

struct Base {
    1: optional i32 Code (api.body="code");
    2: optional string Msg (api.body="msg");
}

struct BaseResp {
    1: optional Base Base (api.body="base");
}

struct PingReq {}

struct PingResp {
}

service PingService {
    PingResp Ping(1: PingReq req) (api.post="/ping")
}
