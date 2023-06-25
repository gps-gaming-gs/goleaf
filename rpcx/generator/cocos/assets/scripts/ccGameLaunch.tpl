import { _decorator, Component, error, JsonAsset, Node, resources, TextAsset } from 'cc';
import { ResManager } from './services/ResManager';
import { TimerManager } from './services/TimerManager';
import { UIManager } from './services/UIManager';
import { EventManager } from './services/EventManager';
import { GameApp } from './bootstrap/GameApp';
import { NetManager } from './services/net/NetManager';
import { ProtoManager } from './services/net/ProtoManager';
import { NetEventDispatcher } from './services/net/NetEventDispatcher';

const { ccclass, property } = _decorator;

@ccclass('GameLaunch')
export class GameLaunch extends Component {

    @property({
        displayName: "NetMode"
    })
    private isNetMode: boolean = false;

    @property({
        visible(this: GameLaunch) {
            // 只有當showProperty為true時才顯示該屬性
            return this.isNetMode;
        },
        displayName: "Host"
    })
    private apiHost: string = "ws://127.0.0.1:9001";

    @property({
        type: TextAsset,
        visible(this: GameLaunch) {
            // 只有當showProperty為true時才顯示該屬性
            return this.isNetMode;
        }
    })
    private protoFile: TextAsset | null = null;

    protected onLoad(): void {
        resources.load('config/app', (err: any, res: JsonAsset) => {
            if (err) {
                error(err.message || err);
                return;
            }
            this.node.addComponent(EventManager);
            this.node.addComponent(TimerManager);
            this.node.addComponent(ResManager);
            this.node.addComponent(UIManager);
            this.node.addComponent(GameApp);

            if (this.isNetMode) {
                this.node.addComponent(ProtoManager).init(this.protoFile);
                this.node.addComponent(NetManager).init(this.apiHost);
                this.node.addComponent(NetEventDispatcher).init();
            }

            GameApp.Instance.enterGame(res.json);
        })
    }
}

