import { _decorator, Component, log, Node } from 'cc';
import { MainGameAudioManager } from '../services/audio/MainGameAudioManager';
import { ResManager } from '../services/ResManager';
import { MainGameRes } from '../GameConfig';
import { BGMAudioManager } from '../services/audio/BGMAudioManager';
import { UIManager } from '../services/UIManager';
import { EventManager } from '../services/EventManager';
import { NetEvent } from '../services/net/NetEvent';

export class GameApp extends Component {

    private static instance: GameApp | null = null;

    public static get Instance(): GameApp {
        return GameApp.instance!;
    }

    protected onLoad(): void {
        if(!GameApp.instance) {
            GameApp.instance = this;
        } else {
            this.destroy();
            return;
        }
        EventManager.Instance.addEventListener(NetEvent.Connecting, this, this.onWebSocketConnecting);
        EventManager.Instance.addEventListener(NetEvent.Connected, this, this.onWebSocketConnected);
    }

    private onWebSocketConnecting(eventName: string, msgBody: any): void {
        console.log("WS: ", eventName, msgBody);
    }

    private onWebSocketConnected(eventName: string, msgBody: any): void {
        console.log("WS: ", eventName, msgBody);
    }

    /**
     * 進入遊戲
     */
    public enterGame(config: object): void {
        ResManager.Instance.preloadResPkg(MainGameRes, (now: any, total: any) => {
            console.log(`now: ${now}`, `total: ${total}`);
        }, () => {
            this.node.addComponent(BGMAudioManager);
            this.node.addComponent(MainGameAudioManager);
            this.loadingScene();
        })
    }

    /**
     * 載入場景
     */
    protected loadingScene() {
        UIManager.Instance.showUIView("Background");
        UIManager.Instance.showUIView("LoginUI");
        // 實例化相關物件
        log("loadingScene.....");
    }

    /**
     * 登入畫面
     */
    protected loadingLogin() {

    }
}

