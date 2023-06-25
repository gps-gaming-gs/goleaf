import { _decorator, Component, Node } from 'cc';
export class EventManager extends Component {

    private static instance: EventManager | null = null;

    private eventMap: Record<string, any> = {};

    protected onLoad(): void {
        if (EventManager.instance === null) {
            EventManager.instance = this;
        } else {
            this.destroy();
            return;
        }
    }

    public static get Instance(): EventManager {
        return EventManager.instance!;
    }

    /**
     * 加入監聽
     */
    public addEventListener(eventName: string, caller: any, func: Function): EventManager {
        if (!this.eventMap[eventName]) {
            this.eventMap[eventName] = [];
        }

        let eventQueue = this.eventMap[eventName];
        eventQueue.push({
            caller: caller,
            func: func,
        });

        return this;
    }

    /**
     * 發送
     * @param eventName
     * @param param
     * @returns
     */
    public emit(eventName: string, eventData: any): EventManager {
        if (!this.eventMap[eventName]) {
            return;
        }

        let eventQueue = this.eventMap[eventName];
        for (let i = 0; i < eventQueue.length; i++) {
            const obj = eventQueue[i];
            obj.func.call(obj.caller, eventName, eventData);
        }

        return this;
    }

    /**
     * 移除監聽
     */
    public removeListener(eventName: string, caller: any, func: Function): EventManager {
        if (!this.eventMap[eventName]) {
            return;
        }

        let eventQueue = this.eventMap[eventName];
        for (let i = 0; i < eventQueue.length; i++) {
            const obj = eventQueue[i];
            if (obj.caller == caller && obj.func == func) {
                eventQueue.splice(i, 1);
                break;
            }
        }

        if (eventQueue.length <=0 ) {
            this.eventMap[eventName] = null;
        }

        return this;
    }
}

