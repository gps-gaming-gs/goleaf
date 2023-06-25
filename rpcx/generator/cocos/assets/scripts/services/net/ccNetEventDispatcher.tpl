import { Component, TextAsset } from 'cc';
import { EventManager } from '../EventManager';
import { NetEvent } from './NetEvent';
import { ProtoManager } from './ProtoManager';
import { NetManager } from './NetManager';

export class NetEventDispatcher extends Component {
    private static instance: NetEventDispatcher | null = null;

    public static get Instance(): NetEventDispatcher {
        return NetEventDispatcher.instance!;
    }

    protected onLoad(): void {
        if (NetEventDispatcher.instance) {
            this.destroy();
            return;
        }
        NetEventDispatcher.instance = this;
    }

    public init(): void {
        EventManager.Instance.addEventListener(NetEvent.Message, this, this.onRecvMessage);
    }

    private onRecvMessage(eventName: string, bufData: ArrayBuffer): void {
        const uint8Buf = new Uint8Array(bufData);
        const cmdID = this.Uint8ArrayToInt(uint8Buf.slice(0,2));
        const msgBuf = uint8Buf.slice(2);

        const msgBody = ProtoManager.Instance.deserializeMsg(CmdType[cmdID], msgBuf);

        EventManager.Instance.emit(CmdType[cmdID], msgBody);
    }

    /**
     * 服務器請求
     * @param servType 服務類型
     * @param cmdType 命令類型
     * @param msgBody
     */
    public send(cmdType: number, msgBody: any): void {
        const msgBuf = ProtoManager.Instance.serializeMsg(CmdType[cmdType], msgBody);

        let _buffer = new Uint8Array(msgBuf.length+2);
        let tagBinary = this.IntToUint8Array(cmdType, 16);
        let tagUnit8 = new Uint8Array(tagBinary);

        _buffer.set(tagUnit8,0);
        _buffer.set(msgBuf.subarray(0,msgBuf.length),2);

        NetManager.Instance.send(_buffer);
    }

    private IntToUint8Array(num: number, Bits: number): number[] {
        const binaryStr: string = num.toString(2);
        const resArry: number[] = Array.from(binaryStr, Number);

        if (Bits && resArry.length < Bits) {
            resArry.unshift(...Array(Bits - resArry.length).fill(0));
        }

        const xresArry: number[] = [];
        for (let j = 0; j < Bits; j += 8) {
            xresArry.push(parseInt(resArry.slice(j, j + 8).join(""), 2));
        }

        return xresArry;
    }

    private Uint8ArrayToInt(uint8Ary: Uint8Array): number {
        let retInt: number = 0;
        const length = uint8Ary.length;
        for (let i = 0; i < length; i++) {
            retInt |= uint8Ary[i] << (8 * (length - i - 1));
        }

        return retInt;
    }
}
