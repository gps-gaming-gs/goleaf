import { _decorator, Component, Label, Node } from 'cc';
import { BaseController } from './BaseController';
import { {{.CtrlName}}View } from '../views/{{.CtrlName}}View';
import { EventManager } from '../services/EventManager';
// import { LoginUI } from '../views/UIEvent';

const { ccclass, property } = _decorator;

@ccclass('{{.CtrlName}}Controller')
export class LoginUIController extends BaseController {

    private bindView : {{.CtrlName}}IView | null = null;

    protected onLoad(): void {
        this.bindView = this.node.getComponent({{.CtrlName}}View);
    }
}

