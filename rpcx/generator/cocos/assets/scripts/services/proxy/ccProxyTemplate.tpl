import { EventManager } from "../EventManager";
import { NetEventDispatcher } from "../net/NetEventDispatcher";

export class {{.ProxyName}}Proxy {

    private static instance: {{.ProxyName}}Proxy | null = null;

    public static get Instance(): {{.ProxyName}}Proxy {
        if (!{{.ProxyName}}Proxy.instance) {
            {{.ProxyName}}Proxy.instance = new {{.ProxyName}}Proxy();
        }
        return {{.ProxyName}}Proxy.instance!;
    }

    constructor() {
        // EventManager.Instance.addEventListener(CmdType[CmdType.LoginResp], this, this.onServerResponse)
    }

    // private onServerResponse(eventName: string, resp: msg.LoginResp) {
    //     console.log(`resp: ${resp.message}`);
    // }

    // public login(uname:string, upwd: string) {
    //     NetEventDispatcher.Instance.send(CmdType.LoginReq, new msg.LoginReq({
    //         name: uname,
    //         password: upwd,
    //     }));
    // }
}

