import { _decorator, Component, find, instantiate, Node, Prefab } from 'cc';
import { ResManager } from './ResManager';

export class UIManager extends Component {
    private static instance: UIManager | null = null;
    private canvas: Node | null = null;
    private uiMap: Record<string, Node | null> = {};

    protected onLoad(): void {
        if (UIManager.instance === null) {
            UIManager.instance = this;
        } else {
            this.destroy();
            return;
        }

        this.canvas = find("Canvas");
    }

    public static get Instance(): UIManager {
        return UIManager.instance!;
    }

    public showUIPrefab(uiPrefab: Prefab, parent?: Node): UIManager {
        const defaultParent = this.canvas;
        const uiView = instantiate(uiPrefab);
        parent = parent ?? defaultParent;
        parent?.addChild(uiView);
        this.uiMap[uiPrefab.data.name] = uiView;
        // 掛載View
        try {
            uiView.addComponent(uiPrefab.data.name + "View");
        } catch (err) {
            console.warn("warn adding view component:", err);
        }

        // 掛載Controller
        try {
            uiView.addComponent(uiPrefab.data.name + "Controller");
        } catch (err) {
            console.warn("warn adding controller component:", err);
        }

        return this;
    }

    public showUIView(viewName: string, parent?: Node): UIManager {
        const defaultParent = this.canvas;
        const uiPrefab = ResManager.Instance.getAsset("GUI", "UIPrefab/" + viewName);

        if (!uiPrefab) {
            console.error("UI prefab not found:", viewName);
            return;
        }

        const uiView = instantiate(uiPrefab);
        parent = parent ?? defaultParent;
        parent?.addChild(uiView);
        this.uiMap[viewName] = uiView;
        // 掛載View
        try {
            uiView.addComponent(viewName + "View");
        } catch (err) {
            console.warn("Warn adding view component:", err);
        }

        // 掛載Controller
        try {
            uiView.addComponent(viewName + "Controller");
        } catch (err) {
            console.warn("Warn adding controller component:", err);
        }

        return this;
    }

    public removeUIView(viewName: string): UIManager {
        if (this.uiMap[viewName]) {
            this.uiMap[viewName]?.destroy();
            delete this.uiMap[viewName];
        }
        return this;
    }

    public clearAll(): UIManager {
        for (const key in this.uiMap) {
            if (this.uiMap[key]) {
                this.uiMap[key]?.destroy();
                delete this.uiMap[key];
            }
        }
        return this;
    }
}
