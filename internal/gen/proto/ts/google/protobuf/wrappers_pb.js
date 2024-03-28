/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "google/protobuf/wrappers.proto" (package "google.protobuf", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
//
// Wrappers for primitive (non-message) types. These types are useful
// for embedding primitives in the `google.protobuf.Any` type and for places
// where we need to distinguish between the absence of a primitive
// typed field and its default value.
//
// These wrappers have no meaningful use within repeated fields as they lack
// the ability to detect presence on individual elements.
// These wrappers have no meaningful use within a map or a oneof since
// individual entries of a map or fields of a oneof can already detect presence.
//
/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "google/protobuf/wrappers.proto" (package "google.protobuf", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
//
// Wrappers for primitive (non-message) types. These types are useful
// for embedding primitives in the `google.protobuf.Any` type and for places
// where we need to distinguish between the absence of a primitive
// typed field and its default value.
//
// These wrappers have no meaningful use within repeated fields as they lack
// the ability to detect presence on individual elements.
// These wrappers have no meaningful use within a map or a oneof since
// individual entries of a map or fields of a oneof can already detect presence.
//
import { ScalarType } from "@protobuf-ts/runtime";
import { LongType } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
// @generated message type with reflection information, may provide speed optimized methods
class DoubleValue$Type extends MessageType {
    constructor() {
        super("google.protobuf.DoubleValue", [
            { no: 1, name: "value", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
    /**
     * Encode `DoubleValue` to JSON number.
     */
    internalJsonWrite(message, options) {
        return this.refJsonWriter.scalar(2, message.value, "value", false, true);
    }
    /**
     * Decode `DoubleValue` from JSON number.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, 1, undefined, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.DoubleValue
 */
export const DoubleValue = new DoubleValue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FloatValue$Type extends MessageType {
    constructor() {
        super("google.protobuf.FloatValue", [
            { no: 1, name: "value", kind: "scalar", T: 2 /*ScalarType.FLOAT*/ }
        ]);
    }
    /**
     * Encode `FloatValue` to JSON number.
     */
    internalJsonWrite(message, options) {
        return this.refJsonWriter.scalar(1, message.value, "value", false, true);
    }
    /**
     * Decode `FloatValue` from JSON number.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, 1, undefined, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.FloatValue
 */
export const FloatValue = new FloatValue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Int64Value$Type extends MessageType {
    constructor() {
        super("google.protobuf.Int64Value", [
            { no: 1, name: "value", kind: "scalar", T: 3 /*ScalarType.INT64*/ }
        ]);
    }
    /**
     * Encode `Int64Value` to JSON string.
     */
    internalJsonWrite(message, options) {
        return this.refJsonWriter.scalar(ScalarType.INT64, message.value, "value", false, true);
    }
    /**
     * Decode `Int64Value` from JSON string.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, ScalarType.INT64, LongType.STRING, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.Int64Value
 */
export const Int64Value = new Int64Value$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UInt64Value$Type extends MessageType {
    constructor() {
        super("google.protobuf.UInt64Value", [
            { no: 1, name: "value", kind: "scalar", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
    /**
     * Encode `UInt64Value` to JSON string.
     */
    internalJsonWrite(message, options) {
        return this.refJsonWriter.scalar(ScalarType.UINT64, message.value, "value", false, true);
    }
    /**
     * Decode `UInt64Value` from JSON string.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, ScalarType.UINT64, LongType.STRING, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.UInt64Value
 */
export const UInt64Value = new UInt64Value$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Int32Value$Type extends MessageType {
    constructor() {
        super("google.protobuf.Int32Value", [
            { no: 1, name: "value", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    /**
     * Encode `Int32Value` to JSON string.
     */
    internalJsonWrite(message, options) {
        return this.refJsonWriter.scalar(5, message.value, "value", false, true);
    }
    /**
     * Decode `Int32Value` from JSON string.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, 5, undefined, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.Int32Value
 */
export const Int32Value = new Int32Value$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UInt32Value$Type extends MessageType {
    constructor() {
        super("google.protobuf.UInt32Value", [
            { no: 1, name: "value", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
    /**
     * Encode `UInt32Value` to JSON string.
     */
    internalJsonWrite(message, options) {
        return this.refJsonWriter.scalar(13, message.value, "value", false, true);
    }
    /**
     * Decode `UInt32Value` from JSON string.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, 13, undefined, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.UInt32Value
 */
export const UInt32Value = new UInt32Value$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BoolValue$Type extends MessageType {
    constructor() {
        super("google.protobuf.BoolValue", [
            { no: 1, name: "value", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    /**
     * Encode `BoolValue` to JSON bool.
     */
    internalJsonWrite(message, options) {
        return message.value;
    }
    /**
     * Decode `BoolValue` from JSON bool.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, 8, undefined, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.BoolValue
 */
export const BoolValue = new BoolValue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class StringValue$Type extends MessageType {
    constructor() {
        super("google.protobuf.StringValue", [
            { no: 1, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    /**
     * Encode `StringValue` to JSON string.
     */
    internalJsonWrite(message, options) {
        return message.value;
    }
    /**
     * Decode `StringValue` from JSON string.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, 9, undefined, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.StringValue
 */
export const StringValue = new StringValue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class BytesValue$Type extends MessageType {
    constructor() {
        super("google.protobuf.BytesValue", [
            { no: 1, name: "value", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    /**
     * Encode `BytesValue` to JSON string.
     */
    internalJsonWrite(message, options) {
        return this.refJsonWriter.scalar(12, message.value, "value", false, true);
    }
    /**
     * Decode `BytesValue` from JSON string.
     */
    internalJsonRead(json, options, target) {
        if (!target)
            target = this.create();
        target.value = this.refJsonReader.scalar(json, 12, undefined, "value");
        return target;
    }
}
/**
 * @generated MessageType for protobuf message google.protobuf.BytesValue
 */
export const BytesValue = new BytesValue$Type();
