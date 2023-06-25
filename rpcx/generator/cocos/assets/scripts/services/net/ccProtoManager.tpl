import { Component, TextAsset } from 'cc';
import protobuf from '../../lib/protobuf/protobuf.js';

export class ProtoManager extends Component {
    private static instance: ProtoManager | null = null;
    private pb: any = null;

    public static get Instance(): ProtoManager {
        return ProtoManager.instance!;
    }

    protected onLoad(): void {
        if (ProtoManager.instance) {
            this.destroy();
            return;
        }
        ProtoManager.instance = this;
    }

    public init(pbText: TextAsset): ProtoManager {
        const pbTextData = pbText.text;
        this.pb = protobuf.parse(pbTextData);
        return this;
    }

    public serializeMsg(msgName: string, msgBody: any): Uint8Array {
        const rs = this.pb.root.lookupType(msgName);
        const msg = rs.create(msgBody);
        const buf = rs.encode(msg).finish();
        return buf;
    }

    public deserializeMsg(msgName: string, msgBuf: Uint8Array): Object {
        const rs = this.pb.root.lookupType(msgName);
        const msg = rs.decode(msgBuf);
        return msg;
    }
}
