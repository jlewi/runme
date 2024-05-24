/* eslint-disable */
// @generated by protobuf-ts 2.9.4 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/runner/v2alpha1/runner.proto" (package "runme.runner.v2alpha1", syntax proto3)
// tslint:disable
// @ts-nocheck
/* eslint-disable */
// @generated by protobuf-ts 2.9.4 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/runner/v2alpha1/runner.proto" (package "runme.runner.v2alpha1", syntax proto3)
// tslint:disable
// @ts-nocheck
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { UInt32Value } from "../../../google/protobuf/wrappers_pb";
import { ProgramConfig } from "./config_pb";
/**
 * @generated from protobuf enum runme.runner.v2alpha1.ResolveProgramRequest.Mode
 */
export var ResolveProgramRequest_Mode;
(function (ResolveProgramRequest_Mode) {
    /**
     * unspecified is auto (default) which prompts for all
     * unresolved environment variables.
     * Subsequent runs will likely resolve via the session.
     *
     * @generated from protobuf enum value: MODE_UNSPECIFIED = 0;
     */
    ResolveProgramRequest_Mode[ResolveProgramRequest_Mode["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * prompt always means to prompt for all environment variables.
     *
     * @generated from protobuf enum value: MODE_PROMPT_ALL = 1;
     */
    ResolveProgramRequest_Mode[ResolveProgramRequest_Mode["PROMPT_ALL"] = 1] = "PROMPT_ALL";
    /**
     * skip means to not prompt for any environment variables.
     * All variables will be marked as resolved.
     *
     * @generated from protobuf enum value: MODE_SKIP_ALL = 2;
     */
    ResolveProgramRequest_Mode[ResolveProgramRequest_Mode["SKIP_ALL"] = 2] = "SKIP_ALL";
})(ResolveProgramRequest_Mode || (ResolveProgramRequest_Mode = {}));
/**
 * @generated from protobuf enum runme.runner.v2alpha1.ResolveProgramResponse.Status
 */
export var ResolveProgramResponse_Status;
(function (ResolveProgramResponse_Status) {
    /**
     * unspecified is the default value and it means unresolved.
     *
     * @generated from protobuf enum value: STATUS_UNSPECIFIED = 0;
     */
    ResolveProgramResponse_Status[ResolveProgramResponse_Status["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * resolved means that the variable is resolved.
     *
     * @generated from protobuf enum value: STATUS_RESOLVED = 1;
     */
    ResolveProgramResponse_Status[ResolveProgramResponse_Status["RESOLVED"] = 1] = "RESOLVED";
    /**
     * unresolved with message means that the variable is unresolved
     * but it contains a message. E.g. FOO=this is message.
     *
     * @generated from protobuf enum value: STATUS_UNRESOLVED_WITH_MESSAGE = 2;
     */
    ResolveProgramResponse_Status[ResolveProgramResponse_Status["UNRESOLVED_WITH_MESSAGE"] = 2] = "UNRESOLVED_WITH_MESSAGE";
    /**
     * unresolved with placeholder means that the variable is unresolved
     * but it contains a placeholder. E.g. FOO="this is placeholder".
     *
     * @generated from protobuf enum value: STATUS_UNRESOLVED_WITH_PLACEHOLDER = 3;
     */
    ResolveProgramResponse_Status[ResolveProgramResponse_Status["UNRESOLVED_WITH_PLACEHOLDER"] = 3] = "UNRESOLVED_WITH_PLACEHOLDER";
    /**
     * unresolved with secret means that the variable is unresolved
     * and it requires treatment as a secret.
     *
     * @generated from protobuf enum value: STATUS_UNRESOLVED_WITH_SECRET = 4;
     */
    ResolveProgramResponse_Status[ResolveProgramResponse_Status["UNRESOLVED_WITH_SECRET"] = 4] = "UNRESOLVED_WITH_SECRET";
})(ResolveProgramResponse_Status || (ResolveProgramResponse_Status = {}));
/**
 * @generated from protobuf enum runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot.Status
 */
export var MonitorEnvStoreResponseSnapshot_Status;
(function (MonitorEnvStoreResponseSnapshot_Status) {
    /**
     * @generated from protobuf enum value: STATUS_UNSPECIFIED = 0;
     */
    MonitorEnvStoreResponseSnapshot_Status[MonitorEnvStoreResponseSnapshot_Status["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * @generated from protobuf enum value: STATUS_LITERAL = 1;
     */
    MonitorEnvStoreResponseSnapshot_Status[MonitorEnvStoreResponseSnapshot_Status["LITERAL"] = 1] = "LITERAL";
    /**
     * @generated from protobuf enum value: STATUS_HIDDEN = 2;
     */
    MonitorEnvStoreResponseSnapshot_Status[MonitorEnvStoreResponseSnapshot_Status["HIDDEN"] = 2] = "HIDDEN";
    /**
     * @generated from protobuf enum value: STATUS_MASKED = 3;
     */
    MonitorEnvStoreResponseSnapshot_Status[MonitorEnvStoreResponseSnapshot_Status["MASKED"] = 3] = "MASKED";
})(MonitorEnvStoreResponseSnapshot_Status || (MonitorEnvStoreResponseSnapshot_Status = {}));
/**
 * env store implementation
 *
 * @generated from protobuf enum runme.runner.v2alpha1.SessionEnvStoreType
 */
export var SessionEnvStoreType;
(function (SessionEnvStoreType) {
    /**
     * uses default env store
     *
     * @generated from protobuf enum value: SESSION_ENV_STORE_TYPE_UNSPECIFIED = 0;
     */
    SessionEnvStoreType[SessionEnvStoreType["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * uses owl store
     *
     * @generated from protobuf enum value: SESSION_ENV_STORE_TYPE_OWL = 1;
     */
    SessionEnvStoreType[SessionEnvStoreType["OWL"] = 1] = "OWL";
})(SessionEnvStoreType || (SessionEnvStoreType = {}));
/**
 * @generated from protobuf enum runme.runner.v2alpha1.ExecuteStop
 */
export var ExecuteStop;
(function (ExecuteStop) {
    /**
     * @generated from protobuf enum value: EXECUTE_STOP_UNSPECIFIED = 0;
     */
    ExecuteStop[ExecuteStop["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * @generated from protobuf enum value: EXECUTE_STOP_INTERRUPT = 1;
     */
    ExecuteStop[ExecuteStop["INTERRUPT"] = 1] = "INTERRUPT";
    /**
     * @generated from protobuf enum value: EXECUTE_STOP_KILL = 2;
     */
    ExecuteStop[ExecuteStop["KILL"] = 2] = "KILL";
})(ExecuteStop || (ExecuteStop = {}));
/**
 * SessionStrategy determines a session selection in
 * an initial execute request.
 *
 * @generated from protobuf enum runme.runner.v2alpha1.SessionStrategy
 */
export var SessionStrategy;
(function (SessionStrategy) {
    /**
     * Uses the session_id field to determine the session.
     * If none is present, a new session is created.
     *
     * @generated from protobuf enum value: SESSION_STRATEGY_UNSPECIFIED = 0;
     */
    SessionStrategy[SessionStrategy["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * Uses the most recent session on the server.
     * If there is none, a new one is created.
     *
     * @generated from protobuf enum value: SESSION_STRATEGY_MOST_RECENT = 1;
     */
    SessionStrategy[SessionStrategy["MOST_RECENT"] = 1] = "MOST_RECENT";
})(SessionStrategy || (SessionStrategy = {}));
/**
 * @generated from protobuf enum runme.runner.v2alpha1.MonitorEnvStoreType
 */
export var MonitorEnvStoreType;
(function (MonitorEnvStoreType) {
    /**
     * @generated from protobuf enum value: MONITOR_ENV_STORE_TYPE_UNSPECIFIED = 0;
     */
    MonitorEnvStoreType[MonitorEnvStoreType["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * possible expansion to have a "timeline" view
     * MONITOR_ENV_STORE_TYPE_TIMELINE = 2;
     *
     * @generated from protobuf enum value: MONITOR_ENV_STORE_TYPE_SNAPSHOT = 1;
     */
    MonitorEnvStoreType[MonitorEnvStoreType["SNAPSHOT"] = 1] = "SNAPSHOT";
})(MonitorEnvStoreType || (MonitorEnvStoreType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class Project$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.Project", [
            { no: 1, name: "root", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "env_load_order", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.Project
 */
export const Project = new Project$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Session$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.Session", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "env", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.Session
 */
export const Session = new Session$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateSessionRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.CreateSessionRequest", [
            { no: 1, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 2, name: "env", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "project", kind: "message", T: () => Project },
            { no: 4, name: "env_store_type", kind: "enum", T: () => ["runme.runner.v2alpha1.SessionEnvStoreType", SessionEnvStoreType, "SESSION_ENV_STORE_TYPE_"] }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.CreateSessionRequest
 */
export const CreateSessionRequest = new CreateSessionRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateSessionResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.CreateSessionResponse", [
            { no: 1, name: "session", kind: "message", T: () => Session }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.CreateSessionResponse
 */
export const CreateSessionResponse = new CreateSessionResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSessionRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.GetSessionRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.GetSessionRequest
 */
export const GetSessionRequest = new GetSessionRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSessionResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.GetSessionResponse", [
            { no: 1, name: "session", kind: "message", T: () => Session }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.GetSessionResponse
 */
export const GetSessionResponse = new GetSessionResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListSessionsRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ListSessionsRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ListSessionsRequest
 */
export const ListSessionsRequest = new ListSessionsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListSessionsResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ListSessionsResponse", [
            { no: 1, name: "sessions", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Session }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ListSessionsResponse
 */
export const ListSessionsResponse = new ListSessionsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateSessionRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.UpdateSessionRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 3, name: "env", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "project", kind: "message", T: () => Project }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.UpdateSessionRequest
 */
export const UpdateSessionRequest = new UpdateSessionRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateSessionResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.UpdateSessionResponse", [
            { no: 1, name: "session", kind: "message", T: () => Session }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.UpdateSessionResponse
 */
export const UpdateSessionResponse = new UpdateSessionResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteSessionRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.DeleteSessionRequest", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.DeleteSessionRequest
 */
export const DeleteSessionRequest = new DeleteSessionRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteSessionResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.DeleteSessionResponse", []);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.DeleteSessionResponse
 */
export const DeleteSessionResponse = new DeleteSessionResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Winsize$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.Winsize", [
            { no: 1, name: "rows", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 2, name: "cols", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 3, name: "x", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 4, name: "y", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.Winsize
 */
export const Winsize = new Winsize$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ExecuteRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ExecuteRequest", [
            { no: 1, name: "config", kind: "message", T: () => ProgramConfig },
            { no: 8, name: "input_data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 9, name: "stop", kind: "enum", T: () => ["runme.runner.v2alpha1.ExecuteStop", ExecuteStop, "EXECUTE_STOP_"] },
            { no: 10, name: "winsize", kind: "message", T: () => Winsize },
            { no: 20, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 21, name: "session_strategy", kind: "enum", T: () => ["runme.runner.v2alpha1.SessionStrategy", SessionStrategy, "SESSION_STRATEGY_"] },
            { no: 22, name: "project", kind: "message", T: () => Project },
            { no: 23, name: "store_stdout_in_env", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ExecuteRequest
 */
export const ExecuteRequest = new ExecuteRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ExecuteResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ExecuteResponse", [
            { no: 1, name: "exit_code", kind: "message", T: () => UInt32Value },
            { no: 2, name: "stdout_data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 3, name: "stderr_data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 4, name: "pid", kind: "message", T: () => UInt32Value },
            { no: 5, name: "mime_type", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ExecuteResponse
 */
export const ExecuteResponse = new ExecuteResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResolveProgramCommandList$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ResolveProgramCommandList", [
            { no: 1, name: "lines", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ResolveProgramCommandList
 */
export const ResolveProgramCommandList = new ResolveProgramCommandList$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResolveProgramRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ResolveProgramRequest", [
            { no: 1, name: "commands", kind: "message", oneof: "source", T: () => ResolveProgramCommandList },
            { no: 2, name: "script", kind: "scalar", oneof: "source", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "mode", kind: "enum", T: () => ["runme.runner.v2alpha1.ResolveProgramRequest.Mode", ResolveProgramRequest_Mode, "MODE_"] },
            { no: 4, name: "env", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "session_strategy", kind: "enum", T: () => ["runme.runner.v2alpha1.SessionStrategy", SessionStrategy, "SESSION_STRATEGY_"] },
            { no: 7, name: "project", kind: "message", T: () => Project }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ResolveProgramRequest
 */
export const ResolveProgramRequest = new ResolveProgramRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResolveProgramResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ResolveProgramResponse", [
            { no: 1, name: "script", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "commands", kind: "message", T: () => ResolveProgramCommandList },
            { no: 3, name: "vars", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => ResolveProgramResponse_VarResult }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ResolveProgramResponse
 */
export const ResolveProgramResponse = new ResolveProgramResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResolveProgramResponse_VarResult$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ResolveProgramResponse.VarResult", [
            { no: 1, name: "status", kind: "enum", T: () => ["runme.runner.v2alpha1.ResolveProgramResponse.Status", ResolveProgramResponse_Status, "STATUS_"] },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "original_value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "resolved_value", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ResolveProgramResponse.VarResult
 */
export const ResolveProgramResponse_VarResult = new ResolveProgramResponse_VarResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MonitorEnvStoreRequest$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.MonitorEnvStoreRequest", [
            { no: 1, name: "session", kind: "message", T: () => Session }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.MonitorEnvStoreRequest
 */
export const MonitorEnvStoreRequest = new MonitorEnvStoreRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MonitorEnvStoreResponseSnapshot$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot", [
            { no: 1, name: "envs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => MonitorEnvStoreResponseSnapshot_SnapshotEnv }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot
 */
export const MonitorEnvStoreResponseSnapshot = new MonitorEnvStoreResponseSnapshot$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MonitorEnvStoreResponseSnapshot_SnapshotEnv$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot.SnapshotEnv", [
            { no: 1, name: "status", kind: "enum", T: () => ["runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot.Status", MonitorEnvStoreResponseSnapshot_Status, "STATUS_"] },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "spec", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "origin", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "original_value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "resolved_value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "create_time", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "update_time", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "errors", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => MonitorEnvStoreResponseSnapshot_Error }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot.SnapshotEnv
 */
export const MonitorEnvStoreResponseSnapshot_SnapshotEnv = new MonitorEnvStoreResponseSnapshot_SnapshotEnv$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MonitorEnvStoreResponseSnapshot_Error$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot.Error", [
            { no: 1, name: "code", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 2, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.MonitorEnvStoreResponseSnapshot.Error
 */
export const MonitorEnvStoreResponseSnapshot_Error = new MonitorEnvStoreResponseSnapshot_Error$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MonitorEnvStoreResponse$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.MonitorEnvStoreResponse", [
            { no: 1, name: "type", kind: "enum", T: () => ["runme.runner.v2alpha1.MonitorEnvStoreType", MonitorEnvStoreType, "MONITOR_ENV_STORE_TYPE_"] },
            { no: 2, name: "snapshot", kind: "message", oneof: "data", T: () => MonitorEnvStoreResponseSnapshot }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.MonitorEnvStoreResponse
 */
export const MonitorEnvStoreResponse = new MonitorEnvStoreResponse$Type();
/**
 * @generated ServiceType for protobuf service runme.runner.v2alpha1.RunnerService
 */
export const RunnerService = new ServiceType("runme.runner.v2alpha1.RunnerService", [
    { name: "CreateSession", options: {}, I: CreateSessionRequest, O: CreateSessionResponse },
    { name: "GetSession", options: {}, I: GetSessionRequest, O: GetSessionResponse },
    { name: "ListSessions", options: {}, I: ListSessionsRequest, O: ListSessionsResponse },
    { name: "UpdateSession", options: {}, I: UpdateSessionRequest, O: UpdateSessionResponse },
    { name: "DeleteSession", options: {}, I: DeleteSessionRequest, O: DeleteSessionResponse },
    { name: "MonitorEnvStore", serverStreaming: true, options: {}, I: MonitorEnvStoreRequest, O: MonitorEnvStoreResponse },
    { name: "Execute", serverStreaming: true, clientStreaming: true, options: {}, I: ExecuteRequest, O: ExecuteResponse },
    { name: "ResolveProgram", options: {}, I: ResolveProgramRequest, O: ResolveProgramResponse }
]);
