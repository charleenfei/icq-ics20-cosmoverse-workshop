/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage =
  "charleenfei.icqics20cosmoverseworkshop.eightball";

export interface MsgFeelingLucky {
  creator: string;
  offering: string;
}

export interface MsgFeelingLuckyResponse {}

const baseMsgFeelingLucky: object = { creator: "", offering: "" };

export const MsgFeelingLucky = {
  encode(message: MsgFeelingLucky, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.offering !== "") {
      writer.uint32(18).string(message.offering);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFeelingLucky {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFeelingLucky } as MsgFeelingLucky;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.offering = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFeelingLucky {
    const message = { ...baseMsgFeelingLucky } as MsgFeelingLucky;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.offering !== undefined && object.offering !== null) {
      message.offering = String(object.offering);
    } else {
      message.offering = "";
    }
    return message;
  },

  toJSON(message: MsgFeelingLucky): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.offering !== undefined && (obj.offering = message.offering);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFeelingLucky>): MsgFeelingLucky {
    const message = { ...baseMsgFeelingLucky } as MsgFeelingLucky;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.offering !== undefined && object.offering !== null) {
      message.offering = object.offering;
    } else {
      message.offering = "";
    }
    return message;
  },
};

const baseMsgFeelingLuckyResponse: object = {};

export const MsgFeelingLuckyResponse = {
  encode(_: MsgFeelingLuckyResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFeelingLuckyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgFeelingLuckyResponse,
    } as MsgFeelingLuckyResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgFeelingLuckyResponse {
    const message = {
      ...baseMsgFeelingLuckyResponse,
    } as MsgFeelingLuckyResponse;
    return message;
  },

  toJSON(_: MsgFeelingLuckyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgFeelingLuckyResponse>
  ): MsgFeelingLuckyResponse {
    const message = {
      ...baseMsgFeelingLuckyResponse,
    } as MsgFeelingLuckyResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FeelingLucky(request: MsgFeelingLucky): Promise<MsgFeelingLuckyResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  FeelingLucky(request: MsgFeelingLucky): Promise<MsgFeelingLuckyResponse> {
    const data = MsgFeelingLucky.encode(request).finish();
    const promise = this.rpc.request(
      "charleenfei.icqics20cosmoverseworkshop.eightball.Msg",
      "FeelingLucky",
      data
    );
    return promise.then((data) =>
      MsgFeelingLuckyResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
