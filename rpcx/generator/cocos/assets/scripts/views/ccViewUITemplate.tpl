import { _decorator, Component, Label, Node } from 'cc';
import { BaseUIView } from './BaseUIView';
import { EventManager } from '../services/EventManager';
// import { LoginUI } from './UIEvent';

const { ccclass, property} = _decorator;

@ccclass('{{.CtrlName}}View')
export class {{.CtrlName}}View extends BaseUIView {

    // public version: Label | null = null;

    protected onLoad(): void {
        super.onLoad();
        // this.version = this.view["version"].getComponent(Label);

        // EventManager.Instance.addEventListener(LoginUI.VERSION, this, this.updateVersion)
    }

    // protected updateVersion(eventName: string, param: any): void {
    //    this.version.string = param;
    // }
}

