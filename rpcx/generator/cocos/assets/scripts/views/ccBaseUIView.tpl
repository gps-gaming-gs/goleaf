import { _decorator, Component, Node, Label, Button } from 'cc';

export class BaseUIView extends Component {

  protected view: Record<string, Node> = {};

  protected onLoad(): void {
    this.loadAllNodeInView(this.node, "");
  }

  /**
   * 取得視圖內的節點
   * @param root
   * @param path
   */
  protected loadAllNodeInView(root: Node, path: string): void {
    root.children.forEach((child) => {
      const fullPath = path + child.name;
      this.view[fullPath] = child;
      this.loadAllNodeInView(child, fullPath);
    });
  }

  /**
   * 按鈕監聽
   */
  public addButtonListener(viewName: string, caller: any, func: Function) {
    const viewNode = this.view[viewName];
    if (!viewNode) return;

    const button = viewNode.getComponent(Button);
    if (!button) return;

    button.node.on(Button.EventType.CLICK, func, caller);
  }

  // 其他UI事件
}