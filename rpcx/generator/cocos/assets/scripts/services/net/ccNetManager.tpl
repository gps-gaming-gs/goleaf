import { Component, _decorator } from "cc";
import { EventManager } from "../EventManager";
import { NetEvent } from "./NetEvent";

enum State {
    Disconnected = 0,
    Connecting = 1,
    Connected = 2,
}

export class NetManager extends Component {

    private static instance: NetManager | null = null;

    private host: string;

    private sock: WebSocket | null = null;

    private state: State = State.Disconnected;

    public static get Instance(): NetManager {
        return NetManager.instance!;
    }

    protected onLoad(): void {
        if (NetManager.instance) {
            this.destroy();
            return;
        }
        NetManager.instance = this;
        this.state = State.Disconnected;
    }

    protected update(dt: number): void {
        if (this.state !== State.Disconnected) return;

        this.connectToServer();
    }

    public init(host: string): NetManager {
        this.host = host;
        return this;
    }

    public closeSocket(): NetManager {
        if (this.state === State.Connected && this.sock) {
            this.sock.close();
            this.sock = null;
        }
        this.state = State.Disconnected;
        EventManager.Instance.emit(NetEvent.Disconnected, null);

        return this;
    }

    public send(dataArrayBuf: ArrayBuffer): void {
        if (this.state === State.Connected && this.sock) {
            this.sock.send(dataArrayBuf);
        }
    }

    private connectToServer(): void {
        if (this.state !== State.Disconnected) return;

        this.state = State.Connecting;
        this.sock = new WebSocket(this.host);
        this.sock.binaryType = "arraybuffer";

        this.sock.onopen = this.onOpened.bind(this);
        this.sock.onmessage = this.onRecvData.bind(this);
        this.sock.onclose = this.onSocketClose.bind(this);
        this.sock.onerror = this.onSocketErr.bind(this);

        EventManager.Instance.emit(NetEvent.Connecting, null);
    }

    private onOpened(event: Event): void {
        this.state = State.Connected;
        EventManager.Instance.emit(NetEvent.Connected, null);
    }

    private onRecvData(event: MessageEvent<ArrayBuffer>): void {
        EventManager.Instance.emit(NetEvent.Message, event.data);
    }

    private onSocketClose(event: CloseEvent): void {
        this.closeSocket();
    }

    private onSocketErr(event: Event): void {
        this.closeSocket();
    }
}
