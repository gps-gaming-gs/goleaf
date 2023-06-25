import { EventManager } from "../EventManager";
import { NetEventDispatcher } from "../net/NetEventDispatcher";

export class AuthProxy {

    private static instance: AuthProxy | null = null;

    public static get Instance(): AuthProxy {
        if (!AuthProxy.instance) {
            AuthProxy.instance = new AuthProxy();
        }
        return AuthProxy.instance!;
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

