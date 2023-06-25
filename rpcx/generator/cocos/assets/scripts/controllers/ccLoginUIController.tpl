import { _decorator, Component, Label, Node } from 'cc';
import { BaseController } from './BaseController';
import { LoginUIView } from '../views/LoginUIView';
import { EventManager } from '../services/EventManager';
import { LoginUI } from '../views/UIEvent';
import { AuthProxy } from '../services/proxy/AuthProxy';

const { ccclass, property } = _decorator;

@ccclass('LoginUIController')
export class LoginUIController extends BaseController {

    private loginView : LoginUIView | null = null;

    protected onLoad(): void {
        this.loginView = this.node.getComponent(LoginUIView);

        this.loginView.addButtonListener("btnStart", this, ()=>{
            EventManager.Instance.emit(LoginUI.VERSION, "1.000xx");
            AuthProxy.Instance.login("Xuan", "12345");
        });
    }
}

